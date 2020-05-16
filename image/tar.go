/*

Modified by Fengkun Dong 2020

Original file form https://github.com/moby/moby/blob/master/pkg/archive/archive.go

Copyright 2013-2018 Docker, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package image

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/containerd/continuity/fs"
	"github.com/docker/docker/pkg/system"
	"golang.org/x/sys/unix"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const (
	modeISDIR  = 040000  // Directory
	modeISFIFO = 010000  // FIFO
	modeISREG  = 0100000 // Regular file
	modeISLNK  = 0120000 // Symbolic link
	modeISBLK  = 060000  // Block special file
	modeISCHR  = 020000  // Character special file
	modeISSOCK = 0140000 // Socket
)

const WhiteoutPrefix = ".wh."
const WhiteoutMetaPrefix = WhiteoutPrefix + WhiteoutPrefix
const WhiteoutOpaqueDir = WhiteoutMetaPrefix + ".opq"

const (
	// AUFSWhiteoutFormat is the default format for whiteouts
	AUFSWhiteoutFormat WhiteoutFormat = iota
	// OverlayWhiteoutFormat formats whiteout according to the overlay
	// standard.
	OverlayWhiteoutFormat
)

type WhiteoutFormat int

// IDMap contains a single entry for user namespace range remapping. An array
// of IDMap entries represents the structure that will be provided to the Linux
// kernel for creating a user namespace.
type IDMap struct {
	ContainerID int `json:"container_id"`
	HostID      int `json:"host_id"`
	Size        int `json:"size"`
}

// Identity is either a UID and GID pair or a SID (but not both)
type Identity struct {
	UID int
	GID int
	SID string
}

// IdentityMapping contains a mappings of UIDs and GIDs
type IdentityMapping struct {
	uids []IDMap
	gids []IDMap
}

type overlayWhiteoutConverter struct {
	inUserNS bool
}

type tarWhiteoutConverter interface {
	ConvertWrite(*tar.Header, string, os.FileInfo) (*tar.Header, error)
	ConvertRead(*tar.Header, string) (bool, error)
}

type tarAppender struct {
	TarWriter *tar.Writer

	// for hardlink mapping
	SeenFiles       map[uint64]string
	IdentityMapping IdentityMapping

	// For packing and unpacking whiteout files in the
	// non standard format. The whiteout files defined
	// by the AUFS standard are used as the tar whiteout
	// standard.
	WhiteoutConverter tarWhiteoutConverter
}

func Untar(tarfile, dstPath string) error {
	return unTar(tarfile, dstPath, "tar")
}

func unTar(tarFile, dstPath, fileType string) error {
	var dirs []*tar.Header
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	var tr *tar.Reader

	if fileType == "gzip" {
		gr, err := gzip.NewReader(srcFile)
		if err != nil {
			return err
		}
		defer gr.Close()
		tr = tar.NewReader(gr)
	} else {
		tr = tar.NewReader(srcFile)
	}

	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		dstFile := filepath.Join(dstPath, hdr.Name)
		err = createTarFile(dstFile, dstPath, hdr, tr, true)
		if err != nil {
			logger.WithError(err).Error("failed to untar tarball")
		}
		if hdr.Typeflag == tar.TypeDir {
			dirs = append(dirs, hdr)
		}
	}
	for _, hdr := range dirs {
		path := filepath.Join(dstPath, hdr.Name)
		if err := os.Chtimes(path, hdr.AccessTime, hdr.ModTime); err != nil {
			return err
		}
		//if err := system.Chtimes(path, hdr.AccessTime, hdr.ModTime); err != nil {
		//	return err
		//}
	}
	return nil
}

func tarDir(src string, dst string) error {
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()
	ta := &tarAppender{
		SeenFiles:       make(map[uint64]string),
		TarWriter:       tar.NewWriter(fw),
		IdentityMapping: IdentityMapping{},
	}
	defer ta.TarWriter.Close()
	ta.WhiteoutConverter = getWhiteoutConverter(0, true)

	stat, err := os.Lstat(src)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		// We can't later join a non-dir with any includes because the
		// 'walk' will error if "file/." is stat-ed and "file" is not a
		// directory. So, we must split the source path and use the
		// basename as the include.

		src, _ = SplitPathDirEntry(src)
	}

	return filepath.Walk(src, func(filePath string, f os.FileInfo, err error) error {
		relFilePath, err := filepath.Rel(src, filePath)
		if err != nil || (relFilePath == "." && f.IsDir()) {
			// Error getting relative path OR we are looking
			// at the source directory path. Skip in both situations.
			return nil
		}

		return ta.addTarFile(filePath, relFilePath)
	})
}

func createTarFile(path, extractDir string, hdr *tar.Header, reader io.Reader, inUserns bool) error {
	// hdr.Mode is in linux format, which we can use for sycalls,
	// but for os.Foo() calls we need the mode converted to os.FileMode,
	// so use hdrInfo.Mode() (they differ for e.g. setuid bits)
	hdrInfo := hdr.FileInfo()

	switch hdr.Typeflag {
	case tar.TypeDir:
		// Create directory unless it exists as a directory already.
		// In that case we just want to merge the two
		if fi, err := os.Lstat(path); !(err == nil && fi.IsDir()) {
			if err := os.Mkdir(path, hdrInfo.Mode()); err != nil {
				return err
			}
		}

	case tar.TypeReg, tar.TypeRegA:
		// Source is regular file. We use system.OpenFileSequential to use sequential
		// file access to avoid depleting the standby list on Windows.
		// On Linux, this equates to a regular os.OpenFile
		//
		// modified by unixeno
		// since islands only designed for linux, just use os.OpenFile instead
		// file, err := system.OpenFileSequential(path, os.O_CREATE|os.O_WRONLY, hdrInfo.Mode())
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, hdrInfo.Mode())
		if err != nil {
			return err
		}
		if _, err := io.Copy(file, reader); err != nil {
			file.Close()
			return err
		}
		file.Close()

	case tar.TypeBlock, tar.TypeChar:
		if inUserns { // cannot create devices in a userns
			return nil
		}
		// Handle this is an OS-specific way
		if err := handleTarTypeBlockCharFifo(hdr, path); err != nil {
			return err
		}

	case tar.TypeFifo:
		// Handle this is an OS-specific way
		if err := handleTarTypeBlockCharFifo(hdr, path); err != nil {
			return err
		}

	case tar.TypeLink:
		targetPath := filepath.Join(extractDir, hdr.Linkname)
		// check for hardlink breakout
		if !strings.HasPrefix(targetPath, extractDir) {
			return fmt.Errorf("invalid hardlink %q -> %q", targetPath, hdr.Linkname)
		}
		if err := os.Link(targetPath, path); err != nil {
			return err
		}

	case tar.TypeSymlink:
		// 	path 				-> hdr.Linkname = targetPath
		// e.g. /extractDir/path/to/symlink 	-> ../2/file	= /extractDir/path/2/file
		targetPath := filepath.Join(filepath.Dir(path), hdr.Linkname)

		// the reason we don't need to check symlinks in the path (with FollowSymlinkInScope) is because
		// that symlink would first have to be created, which would be caught earlier, at this very check:
		if !strings.HasPrefix(targetPath, extractDir) {
			return fmt.Errorf("invalid symlink %q -> %q", path, hdr.Linkname)
		}
		if err := os.Symlink(hdr.Linkname, path); err != nil {
			return err
		}

	case tar.TypeXGlobalHeader:
		logger.Debug("PAX Global Extended Headers found and ignored")
		return nil

	default:
		return fmt.Errorf("unhandled tar header type %d", hdr.Typeflag)
	}

	if err := os.Lchown(path, hdr.Uid, hdr.Gid); err != nil {
		return err
	}

	var errors []string
	for key, value := range hdr.Xattrs {
		if err := system.Lsetxattr(path, key, []byte(value), 0); err != nil {
			if err == syscall.ENOTSUP || err == syscall.EPERM {
				// We ignore errors here because not all graphdrivers support
				// xattrs *cough* old versions of AUFS *cough*. However only
				// ENOTSUP should be emitted in that case, otherwise we still
				// bail.
				// EPERM occurs if modifying xattrs is not allowed. This can
				// happen when running in userns with restrictions (ChromeOS).
				errors = append(errors, err.Error())
				continue
			}
			return err
		}

	}

	if len(errors) > 0 {
		logger.WithField("errors", errors).Warn("ignored xattrs in archive: underlying filesystem doesn't support them")
	}

	// There is no LChmod, so ignore mode for symlink. Also, this
	// must happen after chown, as that can modify the file mode
	if err := handleLChmod(hdr, path, hdrInfo); err != nil {
		return err
	}

	aTime := hdr.AccessTime
	if aTime.Before(hdr.ModTime) {
		// Last access time should never be before last modified time.
		aTime = hdr.ModTime
	}

	// system.Chtimes doesn't support a NOFOLLOW flag atm
	if hdr.Typeflag == tar.TypeLink {
		if fi, err := os.Lstat(hdr.Linkname); err == nil && (fi.Mode()&os.ModeSymlink == 0) {
			if err := system.Chtimes(path, aTime, hdr.ModTime); err != nil {
				return err
			}
		}
	} else if hdr.Typeflag != tar.TypeSymlink {
		if err := system.Chtimes(path, aTime, hdr.ModTime); err != nil {
			return err
		}
	} else {
		ts := []syscall.Timespec{timeToTimespec(aTime), timeToTimespec(hdr.ModTime)}
		//
		//if err := system.LUtimesNano(path, ts); err != nil && err != system.ErrNotSupportedPlatform {
		if err := system.LUtimesNano(path, ts); err != nil {
			return err
		}
	}
	return nil
}

// addTarFile adds to the tar archive a file from `path` as `name`
func (ta *tarAppender) addTarFile(path, name string) error {
	fi, err := os.Lstat(path)
	if err != nil {
		return err
	}

	var link string
	if fi.Mode()&os.ModeSymlink != 0 {
		var err error
		link, err = os.Readlink(path)
		if err != nil {
			return err
		}
	}

	hdr, err := FileInfoHeader(name, fi, link)
	if err != nil {
		return err
	}

	// if it's not a directory and has more than 1 link,
	// it's hard linked, so set the type flag accordingly
	if !fi.IsDir() && fi.Sys().(*syscall.Stat_t).Nlink > 1 {
		inode, err := getInodeFromStat(fi.Sys())
		if err != nil {
			return err
		}
		// a link should have a name that it links too
		// and that linked name should be first in the tar archive
		if oldpath, ok := ta.SeenFiles[inode]; ok {
			hdr.Typeflag = tar.TypeLink
			hdr.Linkname = oldpath
			hdr.Size = 0 // This Must be here for the writer math to add up!
		} else {
			ta.SeenFiles[inode] = name
		}
	}

	// check whether the file is overlayfs whiteout
	// if yes, skip re-mapping container ID mappings.
	isOverlayWhiteout := fi.Mode()&os.ModeCharDevice != 0 && hdr.Devmajor == 0 && hdr.Devminor == 0

	// handle re-mapping container ID mappings back to host ID mappings before
	// writing tar headers/files. We skip whiteout files because they were written
	// by the kernel and already have proper ownership relative to the host
	if !isOverlayWhiteout && !strings.HasPrefix(filepath.Base(hdr.Name), WhiteoutPrefix) && !ta.IdentityMapping.Empty() {
		fileIDPair, err := getFileUIDGID(fi.Sys())
		if err != nil {
			return err
		}
		hdr.Uid, hdr.Gid, err = ta.IdentityMapping.ToContainer(fileIDPair)
		if err != nil {
			return err
		}
	}

	if ta.WhiteoutConverter != nil {
		wo, err := ta.WhiteoutConverter.ConvertWrite(hdr, path, fi)
		if err != nil {
			return err
		}

		// If a new whiteout file exists, write original hdr, then
		// replace hdr with wo to be written after. Whiteouts should
		// always be written after the original. Note the original
		// hdr may have been updated to be a whiteout with returning
		// a whiteout header
		if wo != nil {
			if err := ta.TarWriter.WriteHeader(hdr); err != nil {
				return err
			}
			if hdr.Typeflag == tar.TypeReg && hdr.Size > 0 {
				return fmt.Errorf("tar: cannot use whiteout for non-empty file")
			}
			hdr = wo
		}
	}

	if err := ta.TarWriter.WriteHeader(hdr); err != nil {
		return err
	}

	if hdr.Typeflag == tar.TypeReg && hdr.Size > 0 {
		// We use system.OpenSequential to ensure we use sequential file
		// access on Windows to avoid depleting the standby list.
		// On Linux, this equates to a regular os.Open.
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(ta.TarWriter, file)
		file.Close()
		if err != nil {
			return err
		}
		err = ta.TarWriter.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

// FileInfoHeader creates a populated Header from fi.
// Compared to archive pkg this function fills in more information.
// Also, regardless of Go version, this function fills file type bits (e.g. hdr.Mode |= modeISDIR),
// which have been deleted since Go 1.9 archive/tar.
func FileInfoHeader(name string, fi os.FileInfo, link string) (*tar.Header, error) {
	hdr, err := tar.FileInfoHeader(fi, link)
	if err != nil {
		return nil, err
	}
	hdr.Format = tar.FormatPAX
	hdr.ModTime = hdr.ModTime.Truncate(time.Second)
	hdr.AccessTime = time.Time{}
	hdr.ChangeTime = time.Time{}
	hdr.Mode = fillGo18FileTypeBits(int64(os.FileMode(hdr.Mode)), fi)
	hdr.Name = canonicalTarName(name, fi.IsDir())
	if err := setHeaderForSpecialDevice(hdr, name, fi.Sys()); err != nil {
		return nil, err
	}
	return hdr, nil
}

func getInodeFromStat(stat interface{}) (inode uint64, err error) {
	s, ok := stat.(*syscall.Stat_t)

	if ok {
		inode = s.Ino
	}

	return
}

func setHeaderForSpecialDevice(hdr *tar.Header, name string, stat interface{}) (err error) {
	s, ok := stat.(*syscall.Stat_t)

	if ok {
		// Currently go does not fill in the major/minors
		if s.Mode&unix.S_IFBLK != 0 ||
			s.Mode&unix.S_IFCHR != 0 {
			hdr.Devmajor = int64(unix.Major(uint64(s.Rdev))) // nolint: unconvert
			hdr.Devminor = int64(unix.Minor(uint64(s.Rdev))) // nolint: unconvert
		}
	}

	return
}

// canonicalTarName provides a platform-independent and consistent posix-style
// path for files and directories to be archived regardless of the platform.
func canonicalTarName(name string, isDir bool) string {
	// suffix with '/' for directories
	if isDir && !strings.HasSuffix(name, "/") {
		name += "/"
	}
	return name
}

// fillGo18FileTypeBits fills type bits which have been removed on Go 1.9 archive/tar
// https://github.com/golang/go/commit/66b5a2f
func fillGo18FileTypeBits(mode int64, fi os.FileInfo) int64 {
	fm := fi.Mode()
	switch {
	case fm.IsRegular():
		mode |= modeISREG
	case fi.IsDir():
		mode |= modeISDIR
	case fm&os.ModeSymlink != 0:
		mode |= modeISLNK
	case fm&os.ModeDevice != 0:
		if fm&os.ModeCharDevice != 0 {
			mode |= modeISCHR
		} else {
			mode |= modeISBLK
		}
	case fm&os.ModeNamedPipe != 0:
		mode |= modeISFIFO
	case fm&os.ModeSocket != 0:
		mode |= modeISSOCK
	}
	return mode
}

// handleTarTypeBlockCharFifo is an OS-specific helper function used by
// createTarFile to handle the following types of header: Block; Char; Fifo
func handleTarTypeBlockCharFifo(hdr *tar.Header, path string) error {
	if RunningInUserNS() {
		// cannot create a device if running in user namespace
		return nil
	}

	mode := uint32(hdr.Mode & 07777)
	switch hdr.Typeflag {
	case tar.TypeBlock:
		mode |= unix.S_IFBLK
	case tar.TypeChar:
		mode |= unix.S_IFCHR
	case tar.TypeFifo:
		mode |= unix.S_IFIFO
	}

	return unix.Mknod(path, mode, int(unix.Mkdev(uint32(hdr.Devmajor), uint32(hdr.Devminor))))
	//return system.Mknod(path, mode, int(system.Mkdev(hdr.Devmajor, hdr.Devminor)))
}

func handleLChmod(hdr *tar.Header, path string, hdrInfo os.FileInfo) error {
	if hdr.Typeflag == tar.TypeLink {
		if fi, err := os.Lstat(hdr.Linkname); err == nil && (fi.Mode()&os.ModeSymlink == 0) {
			if err := os.Chmod(path, hdrInfo.Mode()); err != nil {
				return err
			}
		}
	} else if hdr.Typeflag != tar.TypeSymlink {
		if err := os.Chmod(path, hdrInfo.Mode()); err != nil {
			return err
		}
	}
	return nil
}

func timeToTimespec(time time.Time) (ts syscall.Timespec) {
	if time.IsZero() {
		// Return UTIME_OMIT special value
		ts.Sec = 0
		ts.Nsec = (1 << 30) - 2
		return
	}
	return syscall.NsecToTimespec(time.UnixNano())
}

func getFileUIDGID(stat interface{}) (Identity, error) {
	s, ok := stat.(*syscall.Stat_t)

	if !ok {
		return Identity{}, fmt.Errorf("cannot convert stat value to syscall.Stat_t")
	}
	return Identity{UID: int(s.Uid), GID: int(s.Gid)}, nil
}

// Empty returns true if there are no id mappings
func (i *IdentityMapping) Empty() bool {
	return len(i.uids) == 0 && len(i.gids) == 0
}

// ToContainer returns the container UID and GID for the host uid and gid
func (i *IdentityMapping) ToContainer(pair Identity) (int, int, error) {
	uid, err := toContainer(pair.UID, i.uids)
	if err != nil {
		return -1, -1, err
	}
	gid, err := toContainer(pair.GID, i.gids)
	return uid, gid, err
}

// toContainer takes an id mapping, and uses it to translate a
// host ID to the remapped ID. If no map is provided, then the translation
// assumes a 1-to-1 mapping and returns the passed in id
func toContainer(hostID int, idMap []IDMap) (int, error) {
	if idMap == nil {
		return hostID, nil
	}
	for _, m := range idMap {
		if (hostID >= m.HostID) && (hostID <= (m.HostID + m.Size - 1)) {
			contID := m.ContainerID + (hostID - m.HostID)
			return contID, nil
		}
	}
	return -1, fmt.Errorf("Host ID %d cannot be mapped to a container ID", hostID)
}

func getWhiteoutConverter(format WhiteoutFormat, inUserNS bool) tarWhiteoutConverter {
	if format == OverlayWhiteoutFormat {
		return overlayWhiteoutConverter{inUserNS: inUserNS}
	}
	return nil
}

func (overlayWhiteoutConverter) ConvertWrite(hdr *tar.Header, path string, fi os.FileInfo) (wo *tar.Header, err error) {
	// convert whiteouts to AUFS format
	if fi.Mode()&os.ModeCharDevice != 0 && hdr.Devmajor == 0 && hdr.Devminor == 0 {
		// we just rename the file and make it normal
		dir, filename := filepath.Split(hdr.Name)
		hdr.Name = filepath.Join(dir, WhiteoutPrefix+filename)
		hdr.Mode = 0600
		hdr.Typeflag = tar.TypeReg
		hdr.Size = 0
	}

	if fi.Mode()&os.ModeDir != 0 {
		// convert opaque dirs to AUFS format by writing an empty file with the prefix
		opaque, err := system.Lgetxattr(path, "trusted.overlay.opaque")
		if err != nil {
			return nil, err
		}
		if len(opaque) == 1 && opaque[0] == 'y' {
			if hdr.Xattrs != nil {
				delete(hdr.Xattrs, "trusted.overlay.opaque")
			}

			// create a header for the whiteout file
			// it should inherit some properties from the parent, but be a regular file
			wo = &tar.Header{
				Typeflag:   tar.TypeReg,
				Mode:       hdr.Mode & int64(os.ModePerm),
				Name:       filepath.Join(hdr.Name, WhiteoutOpaqueDir),
				Size:       0,
				Uid:        hdr.Uid,
				Uname:      hdr.Uname,
				Gid:        hdr.Gid,
				Gname:      hdr.Gname,
				AccessTime: hdr.AccessTime,
				ChangeTime: hdr.ChangeTime,
			}
		}
	}

	return
}

func (c overlayWhiteoutConverter) ConvertRead(hdr *tar.Header, path string) (bool, error) {
	base := filepath.Base(path)
	dir := filepath.Dir(path)

	// if a directory is marked as opaque by the AUFS special file, we need to translate that to overlay
	if base == WhiteoutOpaqueDir {
		err := unix.Setxattr(dir, "trusted.overlay.opaque", []byte{'y'}, 0)
		if err != nil {
			if c.inUserNS {
				if err = replaceDirWithOverlayOpaque(dir); err != nil {
					return false, fmt.Errorf("replaceDirWithOverlayOpaque(%q) failed with error %v", dir, err)
				}
			} else {
				return false, fmt.Errorf("setxattr(%q, trusted.overlay.opaque=y) with error %v", dir, err)
			}
		}
		// don't write the file itself
		return false, err
	}

	// if a file was deleted and we are using overlay, we need to create a character device
	if strings.HasPrefix(base, WhiteoutPrefix) {
		originalBase := base[len(WhiteoutPrefix):]
		originalPath := filepath.Join(dir, originalBase)

		if err := unix.Mknod(originalPath, unix.S_IFCHR, 0); err != nil {
			if c.inUserNS {
				// Ubuntu and a few distros support overlayfs in userns.
				//
				// Although we can't call mknod directly in userns (at least on bionic kernel 4.15),
				// we can still create 0,0 char device using mknodChar0Overlay().
				//
				// NOTE: we don't need this hack for the containerd snapshotter+unpack model.
				if err := mknodChar0Overlay(originalPath); err != nil {
					return false, fmt.Errorf("failed to mknodChar0UserNS(%q) with error %v", originalPath, err)
				}
			} else {
				return false, fmt.Errorf("failed to mknod(%q, S_IFCHR, 0) with error %v", originalPath, err)
			}
		}
		if err := os.Chown(originalPath, hdr.Uid, hdr.Gid); err != nil {
			return false, err
		}

		// don't write the file itself
		return false, nil
	}

	return true, nil
}

// replaceDirWithOverlayOpaque replaces path with a new directory with trusted.overlay.opaque
// xattr. The contents of the directory are preserved.
func replaceDirWithOverlayOpaque(path string) error {
	if path == "/" {
		return fmt.Errorf("replaceDirWithOverlayOpaque: path must not be \"/\"")
	}
	dir := filepath.Dir(path)
	tmp, err := ioutil.TempDir(dir, "rdwoo")
	if err != nil {
		return fmt.Errorf("failed to create a tmp directory under %s with error : %v", dir, err)
	}
	defer os.RemoveAll(tmp)
	// newPath is a new empty directory crafted with trusted.overlay.opaque xattr.
	// we copy the content of path into newPath, remove path, and rename newPath to path.
	newPath, err := createDirWithOverlayOpaque(tmp)
	if err != nil {
		return fmt.Errorf("createDirWithOverlayOpaque(%q) failed with %v", tmp, err)
	}
	if err := fs.CopyDir(newPath, path); err != nil {
		return fmt.Errorf("CopyDir(%q, %q) failed with error %v", newPath, path, err)
	}
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return os.Rename(newPath, path)
}

// createDirWithOverlayOpaque creates a directory with trusted.overlay.opaque xattr,
// without calling setxattr, so as to allow creating opaque dir in userns on Ubuntu.
func createDirWithOverlayOpaque(tmp string) (string, error) {
	lower := filepath.Join(tmp, "l")
	upper := filepath.Join(tmp, "u")
	work := filepath.Join(tmp, "w")
	merged := filepath.Join(tmp, "m")
	for _, s := range []string{lower, upper, work, merged} {
		if err := os.MkdirAll(s, 0700); err != nil {
			return "", fmt.Errorf("failed to mkdir %s with %v", s, err)
		}
	}
	dummyBase := "d"
	lowerDummy := filepath.Join(lower, dummyBase)
	if err := os.MkdirAll(lowerDummy, 0700); err != nil {
		return "", fmt.Errorf("failed to create a dummy lower directory %s with error %v", lowerDummy, err)
	}
	mOpts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lower, upper, work)
	// docker/pkg/mount.Mount() requires procfs to be mounted. So we use syscall.Mount() directly instead.
	if err := syscall.Mount("overlay", merged, "overlay", uintptr(0), mOpts); err != nil {
		return "", fmt.Errorf("failed to mount overlay (%s) on %s with error %v", mOpts, merged, err)
	}
	mergedDummy := filepath.Join(merged, dummyBase)
	if err := os.Remove(mergedDummy); err != nil {
		syscall.Unmount(merged, 0)
		return "", fmt.Errorf("failed to rmdir %s with error %v", mergedDummy, err)
	}
	// upperDummy becomes a 0,0-char device file here
	if err := os.Mkdir(mergedDummy, 0700); err != nil {
		syscall.Unmount(merged, 0)
		return "", fmt.Errorf("failed to mkdir %s with error %v", mergedDummy, err)
	}
	// upperDummy becomes a directory with trusted.overlay.opaque xattr
	// (but can't be verified in userns)
	if err := syscall.Unmount(merged, 0); err != nil {
		return "", fmt.Errorf("failed to unmount %s with error %v", merged, err)
	}
	upperDummy := filepath.Join(upper, dummyBase)
	return upperDummy, nil
}

// mknodChar0Overlay creates 0,0 char device by mounting overlayfs and unlinking.
// This function can be used for creating 0,0 char device in userns on Ubuntu.
//
// Steps:
// * Mkdir lower,upper,merged,work
// * Create lower/dummy
// * Mount overlayfs
// * Unlink merged/dummy
// * Unmount overlayfs
// * Make sure a 0,0 char device is created as upper/dummy
// * Rename upper/dummy to cleansedOriginalPath
func mknodChar0Overlay(cleansedOriginalPath string) error {
	dir := filepath.Dir(cleansedOriginalPath)
	tmp, err := ioutil.TempDir(dir, "mc0o")
	if err != nil {
		return fmt.Errorf("failed to create a tmp directory under %s with error %v", dir, err)
	}
	defer os.RemoveAll(tmp)
	lower := filepath.Join(tmp, "l")
	upper := filepath.Join(tmp, "u")
	work := filepath.Join(tmp, "w")
	merged := filepath.Join(tmp, "m")
	for _, s := range []string{lower, upper, work, merged} {
		if err := os.MkdirAll(s, 0700); err != nil {
			return fmt.Errorf("failed to mkdir %s with error %v", s, err)
		}
	}
	dummyBase := "d"
	lowerDummy := filepath.Join(lower, dummyBase)
	if err := ioutil.WriteFile(lowerDummy, []byte{}, 0600); err != nil {
		return fmt.Errorf("failed to create a dummy lower file %s with error %v", lowerDummy, err)
	}
	mOpts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lower, upper, work)
	// docker/pkg/mount.Mount() requires procfs to be mounted. So we use syscall.Mount() directly instead.
	if err := syscall.Mount("overlay", merged, "overlay", uintptr(0), mOpts); err != nil {
		return fmt.Errorf("failed to mount overlay (%s) on %s with error %v", mOpts, merged, err)
	}
	mergedDummy := filepath.Join(merged, dummyBase)
	if err := os.Remove(mergedDummy); err != nil {
		syscall.Unmount(merged, 0)
		return fmt.Errorf("failed to unlink %s with error %v", mergedDummy, err)
	}
	if err := syscall.Unmount(merged, 0); err != nil {
		return fmt.Errorf("failed to unmount %s with error %v", merged, err)
	}
	upperDummy := filepath.Join(upper, dummyBase)
	if err := isChar0(upperDummy); err != nil {
		return err
	}
	if err := os.Rename(upperDummy, cleansedOriginalPath); err != nil {
		return fmt.Errorf("failed to rename %s to %s with error %v", upperDummy, cleansedOriginalPath, err)
	}
	return nil
}

func isChar0(path string) error {
	osStat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat %s with error %v", path, err)
	}
	st, ok := osStat.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("got unsupported stat for %s", path)
	}
	if os.FileMode(st.Mode)&syscall.S_IFMT != syscall.S_IFCHR {
		return fmt.Errorf("%s is not a character device, got mode=%d", path, st.Mode)
	}
	if st.Rdev != 0 {
		return fmt.Errorf("%s is not a 0,0 character device, got Rdev=%d", path, st.Rdev)
	}
	return nil
}

// SplitPathDirEntry splits the given path between its directory name and its
// basename by first cleaning the path but preserves a trailing "." if the
// original path specified the current directory.
func SplitPathDirEntry(path string) (dir, base string) {
	cleanedPath := filepath.Clean(filepath.FromSlash(path))

	if specifiesCurrentDir(path) {
		cleanedPath += string(os.PathSeparator) + "."
	}

	return filepath.Dir(cleanedPath), filepath.Base(cleanedPath)
}

// specifiesCurrentDir returns whether the given path specifies
// a "current directory", i.e., the last path segment is `.`.
func specifiesCurrentDir(path string) bool {
	return filepath.Base(path) == "."
}

/*
 * Detect whether we are currently running in a user namespace.
 * Copied from github.com/lxc/lxd/shared/util.go
 */
func RunningInUserNS() bool {
	file, err := os.Open("/proc/self/uid_map")
	if err != nil {
		/*
		 * This kernel-provided file only exists if user namespaces are
		 * supported
		 */
		return false
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	l, _, err := buf.ReadLine()
	if err != nil {
		return false
	}

	line := string(l)
	var a, b, c int64
	fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
	/*
	 * We assume we are in the initial user namespace if we have a full
	 * range - 4294967295 uids starting at uid 0.
	 */
	if a == 0 && b == 0 && c == 4294967295 {
		return false
	}
	return true
}
