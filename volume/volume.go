package volume

import (
	"encoding/json"
	"fmt"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"github.com/yanlingqiankun/Executor/stringid"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Volume struct {
	Name string			// 卷名称
	Path string			// 卷位于宿主机的路径
	CreateTime time.Time// 创建时间
	containers []string	// 使用此卷的容器
}

var containersMutex sync.Mutex

//var volumeMap = make(map[string] *Volume)
var volumeMap = new(sync.Map)
var logger = logging.GetLogger("volume")
var configPath string

func init () {
	level := logging.GetLevel(conf.GetString("LogLevel"))
	logger.SetLevel(level)
	volumePath := filepath.Join(conf.GetString("RootPath"), "volumes")
	if err := os.MkdirAll(volumePath, 0755); err != nil {
		logger.WithError(err).Error("failed to init volume dir")
	}
	configPath = filepath.Join(volumePath, "volumes.json")
	if err := loadVolumes(); err != nil {
		logger.WithError(err).Error("failed to load volume information")
	}
}

func loadVolumes() error {
	data, err := ioutil.ReadFile(configPath)
	normalMap := make(map[string]*Volume)
	if err != nil {
		if os.IsExist(err) {
			return err
		} else {
			logger.Warning("can't find the file of volume")
			return nil
		}
	} else {
		if err = json.Unmarshal(data, &normalMap); err != nil {
			return err
		}
		volumeMap = volumeMapToSyncMap(normalMap)
	}
	return nil
}

func dumpVolumes() error {
	file, err := os.OpenFile(configPath, os.O_TRUNC | os.O_WRONLY | os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		return err
	}
	normalMap := syncMapVolumeMap(volumeMap)
	volumeJson, err := json.Marshal(normalMap)
	if err != nil {
		return err
	}
	_, err = file.Write(volumeJson)
	return err
}

// 返回所有已在系统内注册的卷
func List() []*Volume {
	list := make([]*Volume,0)
	var i = 0
	volumeMap.Range(func(_, item interface{}) bool {
		list = append(list, item.(*Volume))
		i ++
		return  true
	})
	return list
}

// 建立一个卷
func Create(name string) (*Volume, error) {
	if name == "" {
		name = stringid.GenerateNonCryptoID()
	}
	path := filepath.Join(conf.GetString("RootPath"), "volumes",name)
	volume := new(Volume)
	volume.Name = name
	volume.Path = path
	volume.CreateTime = time.Now()
	volume.containers = make ([]string, 0)

	if _, ok := volumeMap.LoadOrStore(name, volume); ok {
		logger.Error("the volume exists")
		return nil, fmt.Errorf("the volume exists")
	}

	//new volume
	if err := os.MkdirAll(path, 0755); err == nil {
		if err := dumpVolumes(); err != nil {
			logger.WithError(err).Error("failed to save volume information")
			volumeMap.Delete(name)
			return nil, err
		}
		return volume, nil
	} else {
		volumeMap.Delete(name)
		return nil, err
	}
}

// 添加一个外部存储目录作为卷，在添加前目录必须存在
func Add(path string, name string) (*Volume, error) {
	volume := new(Volume)
	volume.Path = path
	if name == "" {
		name = stringid.GenerateNonCryptoID()
	}
	volume.Name = name
	volume.CreateTime = time.Now()

	if _, ok := volumeMap.LoadOrStore(name, volume); ok {
		logger.Error("the volume exists")
		return nil, fmt.Errorf("the volume exists")
	}
	if filepath.IsAbs(path) == false {
		volumeMap.Delete(name)
		return nil, fmt.Errorf("volume must use absolute path")
	}
	if info, err := os.Stat(path); err != nil {
		volumeMap.Delete(name)
		return nil, err
	} else {
		if info.IsDir() {
			if err := dumpVolumes(); err != nil {
				logger.WithError(err).Error("failed to save volume information")
				return nil, err
			}
			return volume, dumpVolumes()
		} else {
			volumeMap.Delete(name)
			return nil, fmt.Errorf("volume should be a directory")
		}
	}
}

// 删除卷
// 此操作会删除整个目录
// 成功时返回nil，否则返回错误
func Delete(name string) error {
	if volume, ok := volumeMap.Load(name); ok {
		if len(volume.(*Volume).containers) != 0 {
			logger.Error("failed to remove volume : The volume was in use")
			return fmt.Errorf("The volume was in use")
		}
		if err := os.RemoveAll(volume.(*Volume).Path); err != nil {
			return err
		} else {
			volumeMap.Delete(name)
			if err := dumpVolumes(); err != nil {
				logger.WithError(err).Error("failed to save volume information")
				return err
			}
			return nil
		}
	} else {
		return fmt.Errorf("volume doesn't exist")
	}
}

func Open(name string) (*Volume, error) {
	if volume, ok := volumeMap.Load(name); ok {
		return volume.(*Volume), nil
	}
	return nil, fmt.Errorf("the volume not exists")
}

func (v *Volume) MountPoint() string {
	return v.Path
}

func (v *Volume) Register(containerid string) error {
	if contain(v.containers, containerid) > -1 {
		return nil
	}
	if volume, ok := volumeMap.Load(v.Name); ok {
		containersMutex.Lock()
		volume.(*Volume).containers = append(v.containers, containerid)
		containersMutex.Unlock()
	}
	return dumpVolumes()
}

func (v *Volume) UnRegister(containerid string) error {
	index := contain(v.containers, containerid)
	if index > -1 {
		if volume, ok := volumeMap.Load(v.Name); ok {
			containersMutex.Lock()
			volume.(*Volume).containers = append(volume.(*Volume).containers[0:index], volume.(*Volume).containers[index+1:len(volume.(*Volume).containers)]...)
			containersMutex.Unlock()
		}
	} else {
		return fmt.Errorf("the container don't have the volume")
	}
	return dumpVolumes()
}

func contain(containers []string, container string) int {
	i := len(containers) - 1
	for ; i >= 0; i -- {
		if containers[i] == container {
			return i
		}
	}
	return i
}

func volumeMapToSyncMap(normalMap map[string] *Volume) (syncMap *sync.Map){
	syncMap = new(sync.Map)
	for key, value := range normalMap {
		syncMap.Store(key, value)
	}
	return
}

func syncMapVolumeMap(syncMap *sync.Map) (normalMap map[string] *Volume) {
	normalMap = make(map[string]*Volume)
	syncMap.Range(func(key, value interface{}) bool {
		normalMap[key.(string)] = value.(*Volume)
		return true
	})
	return
}

