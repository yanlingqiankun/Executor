package image

import (
	"encoding/json"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"io/ioutil"
	"os"
	"path/filepath"
)

type imageDB map[string]image

var logger = logging.GetLogger("image")
var db imageDB
var imagefile string

func init () {
	rootPath := conf.GetString("RootPath")
	imagefile = filepath.Join(rootPath, "images", "imagedb.json")
	if err := load(); err != nil {
		logger.WithError(err).Errorf("failed to load image infomation")
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
