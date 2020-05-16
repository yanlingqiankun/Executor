package image

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/yanlingqiankun/Executor/conf"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ImportDocekrImage(ctx context.Context, name, file string) (string, error) {
	fileReader, err := os.Open(file)
	if err != nil {
		logger.WithError(err).Errorf("failed to open image")
		return "", err
	}
	imageImportSource := types.ImageImportSource{
		Source:     fileReader,
		SourceName: "-",
	}
	out, err := cli.ImageImport(ctx, imageImportSource, name, types.ImageImportOptions{})
	if err != nil {
		logger.WithError(err).Errorf("failed import docker image")
		return "", err
	}
	out.Close()
	return GetImageFromDocker(name)
}

func SaveDockerImage(ctx context.Context, path string) (string, error) {
	// if file not exist
	if !exists(path) {
		logger.Errorf("can't find the file : %s", path)
		return "", fmt.Errorf("can't find the file : %s", path)
	}

	logger.WithField("image", path).Info("start importing docker image")
	tmpDir, err := ioutil.TempDir(conf.GetString("Temp"), "")
	if err != nil {
		logger.Error("failed to create temporary folder")
		return "", nil
	}
	defer func() {
		if os.RemoveAll(tmpDir) != nil {
			logger.WithField("path", tmpDir).Error("failed to remove temporary folder")
		}
	}()

	// 把docker镜像解压到临时目录然后进行解析
	logger.Debug("untar docker image to " + tmpDir)
	if err := unTar(path, tmpDir, "tar"); err != nil {
		logger.WithField("error", err).Error("untar failed")
	} else {
		logger.Debug("untar success")
	}

	manifest, err := parseDockerManifest(filepath.Join(tmpDir, "manifest.json"))
	if err != nil {
		return "", err
	}
	var name string
	if manifest.RepoTags != nil && len(manifest.RepoTags) > -1 {
		repo := manifest.RepoTags[0]
		nameAndTag := strings.Split(repo, ":")
		if len(nameAndTag) > 1 {
			if nameAndTag[1] != "latest" {
				logger.Errorf("not support docker image with tag : %s", nameAndTag[1])
				return "",fmt.Errorf("not support docker image with tag : %s", nameAndTag[1])
			}
		}
		name = nameAndTag[0]
	}
	fileReader, err := os.Open(path)
	if err != nil {
		return "", err
	}
	resp, err := cli.ImageLoad(ctx, fileReader, true)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	logger.Debugf("%s has load to docker repo", name)
	return GetImageFromDocker(name)
}

type dockerManifest struct {
	Config   string
	RepoTags []string
	Layers   []string
}

func parseDockerManifest(filename string) (dockerManifest, error) {
	logger.WithField("file", filename).Debug("start parsing manifest")
	var manifest []dockerManifest
	if data, err := ioutil.ReadFile(filename); err != nil {
		logger.WithField("file", filename).Error("failed to open the file for parsing")
		return dockerManifest{}, err
	} else {
		if err = json.Unmarshal(data, &manifest); err != nil {
			logger.WithError(err).Error("an error occurs while parsing the docker manifest")
			return dockerManifest{}, err
		} else {
			if len(manifest) == 1 {
				return manifest[0], nil
			}
			return dockerManifest{}, fmt.Errorf("invalid docker manifest format")
		}
	}
}

func PullDockerImage(ctx context.Context, name string) (string, error) {
	id := CheckNameOrID(name)
	if id != "" {
		logger.Debug("The docker image has been in repository : ", name)
		return id, nil
	}
	buf := bytes.Buffer{}
	storeNode := conf.GetString("StoreNode")
	registryPort := conf.GetString("DockerRegistryPort")
	refStr := storeNode + ":" + registryPort + "/" + name
	resp, err := cli.ImagePull(ctx, refStr, types.ImagePullOptions{All: false})
	if err != nil {
		return  "", err
	}
	defer resp.Close()
	io.Copy(&buf, resp)
	logger.Infoln("docker pull information : \n", buf.String())
	if err := cli.ImageTag(ctx, refStr, name); err != nil {
		return "", err
	}
	imageId, err := getDockerImageID(name)
	if err != nil {
		return "", err
	}
	db[imageId] = &ImageEntry{
		Name:          name,
		ID:            imageId,
		CreateTime:    time.Now(),
		Type:          "docker",
		IsDockerImage: true,
		Counter:       0,
	}
	return imageId,db.save()
}

func GetImageFromDocker(name string) (string, error) {
	id, _ := getDockerImageID(name)
	if id == "" {
		return "", fmt.Errorf("can't find %s in docker repo", name)
	}
	if _, ok := db[id]; ok {
		logger.Error("The image has be in the image repo")
		return "", fmt.Errorf("The image has be in the image repo")
	} else {
		db[id] = &ImageEntry{
			Name:          name,
			ID:            id,
			CreateTime:    time.Now(),
			Type:          "docker",
			IsDockerImage: true,
			Counter:       0,
		}
		return id, db.save()
	}
}

func getDockerImageID (name string) (string, error) {
	inspect, _, err := cli.ImageInspectWithRaw(context.Background(), name)
	if err != nil {
		return "", err
	}
	logger.Debugf("get %s id = %s in docker repo", name, inspect.ID)
	return strings.ReplaceAll(inspect.ID, "sha256:", ""), nil
}
