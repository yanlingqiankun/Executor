package image

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"io/ioutil"
	"os"
	"path/filepath"
)

type imageDB map[string]*ImageEntry

var logger = logging.GetLogger("image")
var db imageDB
var imagefile string
var cli *client.Client

func init () {
	var err error
	db = make(map[string]*ImageEntry)
	rootPath := conf.GetString("RootPath")
	imagefile = filepath.Join(rootPath, "images", "imagedb.json")
	if err := load(); err != nil {
		logger.WithError(err).Errorf("failed to load image infomation")
	}
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logger.WithError(err).Error("failed to get docker client")
	}
}

func load () error {
	data, err := ioutil.ReadFile(imagefile)
	if err != nil {
		if os.IsExist(err) {
			return err
		} else {
			logger.Warning("can't find the file of image")
			return nil
		}
	} else {
		if err = json.Unmarshal(data, &db); err != nil {
			return err
		}
	}
	return nil
}

func(db *imageDB) save() error {
	file, err := os.OpenFile(imagefile, os.O_TRUNC | os.O_WRONLY | os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		return err
	}
	imageJson, err := json.Marshal(db)
	if err != nil {
		return err
	}
	_, err = file.Write(imageJson)
	return err
}

func (image *ImageEntry) GetType() (isDocker bool, imageType string) {
	isDocker = image.IsDockerImage
	imageType = image.Type
	return
}

func (image *ImageEntry) Remove() error {
	if _, ok := db[image.Name]; !ok {
		return fmt.Errorf("can not find the image")
	}
	delete(db, image.Name)
	logger.Info("The image %s has been removed", image.Name)
	return nil
}

func (image *ImageEntry) Rename(name string) error {
	if _, ok := db[image.Name]; !ok {
		return fmt.Errorf("can not find the image")
	}

	delete(db, image.Name)
	image.Name = name
	db[name] = image
	return nil
}
