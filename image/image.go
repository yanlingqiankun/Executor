package image

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"github.com/yanlingqiankun/Executor/stringid"
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
	level := logging.GetLevel(conf.GetString("LogLevel"))
	logger.SetLevel(level)
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

func (image *ImageEntry) GetPath() string {
	path := filepath.Join(conf.GetString("RootPath"), "images", image.ID)
	result := image.ID
	if image.Type == "disk" {
		id := stringid.GenerateRandomID()
		err := copy(filepath.Join(path, image.ID), filepath.Join(path, id))
		if err != nil {
			logger.WithError(err).Error("failed to create copy of image ", image.ID)
		}
		result = id
	}
	return filepath.Join(path, result)
}

func (image *ImageEntry) GetName() string {
	return image.Name
}

func (image *ImageEntry) Register() {
	db.register(image.ID)
}

func (image *ImageEntry) UnRegister() {
	db.unRegister(image.ID)
}

func (image *ImageEntry) Export(target string) error {
	tmpDir, err := ioutil.TempDir(conf.GetString("Temp"), "")
	if err != nil {
		logger.Error("failed to create temporary folder")
		return err
	}
	defer func() {
		if os.RemoveAll(tmpDir) != nil {
			logger.WithField("path", tmpDir).Error("failed to remove temporary folder")
		}
	}()

	if image.Type == "iso" || image.Type == "disk"{
		err := copy(image.GetPath(), filepath.Join(tmpDir, "image"))
		if err != nil {
			logger.WithError(err).Errorf("failed to copy image file to temp directory")
			return err
		}
	} else {
		//cli.
	}
	return nil
}


//被machine使用之前要先进行注册
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

type importStruct struct {
	Type string
	Name string
}

func ImportImage (path string) (string, error) {
	// if file not exist
	if !exists(path) {
		logger.Errorf("can't find the file : %s", path)
		return "", fmt.Errorf("can't find the file : %s", path)
	}

	tmpDir, err := ioutil.TempDir(conf.GetString("Temp"), "")
	if err != nil {
		logger.Error("failed to create temporary folder")
		return "", err
	}
	defer func() {
		if os.RemoveAll(tmpDir) != nil {
			logger.WithField("path", tmpDir).Error("failed to remove temporary folder")
		}
	}()
	if err := unTar(path, tmpDir, "tar"); err != nil {
		logger.WithField("error", err).Error("untar failed")
		return "", err
	} else {
		logger.Debug("untar success")
	}
	imageConfig, err := getImportStruct(filepath.Join(tmpDir, "config"))
	if err != nil {
		logger.WithError(err).Errorf("failed to get image configure file")
		return "", nil
	}
	if imageConfig.Type == "disk" || imageConfig.Type == "iso"{
		return QEMUImageSave(imageConfig.Name, imageConfig.Type, filepath.Join(tmpDir, "image"))
	} else if imageConfig.Type == "docker" {
		return ImportDocekrImage(context.Background(), imageConfig.Name, filepath.Join(tmpDir, "image"))
	} else {
		return "", errors.New("invalid image type to import")
	}
}

func getImportStruct(filename string) (importStruct, error){
	result := importStruct{}
	if data, err := ioutil.ReadFile(filename); err != nil {
		logger.WithField("file", filename).Error("failed to open the file for parsing")
		return importStruct{}, err
	} else {
		if err = json.Unmarshal(data, &result); err != nil {
			logger.WithError(err).Error("an error occurs while parsing the image configure file")
			return importStruct{}, err
		} else {
			return result, nil
		}
	}
}