package image

import (
	"github.com/yanlingqiankun/Executor/conf"
	"io"
	"os"
	"path/filepath"
	"time"
)

func QEMUImageSave(imageName string, imageType string,  fileName string) (string, error) {
	var err error
	defer returnWithError("failed to save image", err)
	imageId, err := getSha256(fileName)
	if err != nil {
		return "", err
	}
	if _, ok := db[imageId]; !ok {
		// new image
		imageDir := filepath.Join(conf.GetString("RootPath"), "images", imageId)
		if err := os.MkdirAll(imageDir, 0700); err != nil {
			return "", err
		}
		if err = copy(fileName, filepath.Join(imageDir, imageId)); err != nil {
			return "", err
		}
		db[imageId] = &ImageEntry{
			Name:          imageName,
			ID:            imageId,
			CreateTime:    time.Now(),
			Type:          imageType,
			IsDockerImage: false,
			Counter:       0,
		}
		db.save()
		logger.Debugf("The image %s has been saved ", imageId)
	} else {
		// the image in the db
		logger.Debug("The image already in repo, ignored :", imageId)
	}
	return imageId, nil
}

func copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return nil
}

