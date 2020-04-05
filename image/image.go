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
var imageFile string
var imagePath string
var cli *client.Client

func init () {
	var err error
	db = make(map[string]*ImageEntry)
	rootPath := conf.GetString("RootPath")
	imagePath = filepath.Join(rootPath, "images")
	if _, err := os.Stat(imagePath); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(imagePath, 0600)
			if err != nil {
				return
			}
		}
	}
	imageFile = filepath.Join(imagePath, "imagedb.json")
	if err := load(); err != nil {
		logger.WithError(err).Errorf("failed to load image information")
	}
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logger.WithError(err).Error("failed to get docker client")
	}
}

func load () error {
	data, err := ioutil.ReadFile(imageFile)
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
	file, err := os.OpenFile(imageFile, os.O_TRUNC | os.O_WRONLY | os.O_CREATE, 0600)
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

func OpenImage(id string) (Image, error) {
	if image, ok := db[id]; ok {
		return image, nil
	} else {
		return nil, fmt.Errorf("can't find the image")
	}
}

func ListImage() []*ImageEntry {
	imageEntrys := make([]*ImageEntry, len(db))
	var i = 0
	for _, k := range db {
		imageEntrys[i] = k
		i ++
	}
	return imageEntrys
}

func (image *ImageEntry) GetType() (isDocker bool, imageType string) {
	isDocker = image.IsDockerImage
	imageType = image.Type
	return
}

func (image *ImageEntry) Remove() error {
	if i, ok := db[image.ID]; !ok {
		return fmt.Errorf("can not find the image")
	} else {
		if i.Counter != 0 {
			return fmt.Errorf("the image was by use, please delete machine firstly")
		}
		if !i.IsDockerImage {
			_ = os.Remove(filepath.Join(imagePath, image.ID))
		}
		delete(db, image.ID)
		logger.Info("The image %s has been removed :", image.Name)
		return db.save()
	}
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

//被容器使用之前要先进行注册
func (db imageDB) register(id string) error {
	if image, ok := db[id]; ok {
		image.Counter ++
		db.save()
		logger.Debug(image.ID, " registered success")
		return nil
	}
	return fmt.Errorf("The image not exist")
}

//当容器被销毁要取消对其注册
func (db imageDB) unRegister(id string) error {
	if image, ok := db[id]; ok {
		if image.Counter > 0 {
			image.Counter --
			db.save()
		}
		logger.Debug(image.ID, " unRegister success")
		return nil
	}
	return fmt.Errorf("The image not exist")
}


func GetImageType(imageID string) string {
	if image, ok := db[imageID]; ok {
		return image.Type
	}
	return ""
}

// method to get imageId of a name or id return id
func CheckNameOrID (args string) string {
	if image, ok := db[args]; ok {
		return image.ID
	} else {
		for key, value := range db {
			if value.Name == args {
				return key
			}
		}
	}
	return ""
}