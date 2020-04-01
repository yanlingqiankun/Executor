package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/yanlingqiankun/Executor/logging"
	"reflect"
)

type configType struct {
	LogLevel      string
	APIPath       string
	RootPath      string
	StoreNode	  string
	DockerRegistryPort 	string
	VMType        string
}

// 默认配置
var sysConfig = configType{
	LogLevel:      "info",
	APIPath:       "unix:///var/run/Executor.sock",
	RootPath:      "/var/lib/Executor",
	StoreNode:     "127.0.0.1",
	DockerRegistryPort: "5000",
	VMType:        "kvm",
}

var logger = logging.GetLogger("conf")

func init() {
	configFilePath := "/media/shado/Downloads/code/goCode/Executor/Executor.conf"
	if _, err := toml.DecodeFile(configFilePath, &sysConfig); err != nil {
		logger.WithError(err).WithField("path", configFilePath).Fatal("failed to load configurations")
	} else {
		level := logging.GetLevel(GetString("LogLevel"))
		logger.SetLevel(level)
		logger.WithField("path", configFilePath).Debug("Configuration file successfully loaded")
	}
}

// 根据item获取获取配置值
//
// 当获取的项目不存在时返回nil
func Get(item string) interface{} {
	return nil
}

// 根据item获取获取配置值
//
// 同Get()，但返回string类型的值
func GetString(item string) string {
	r := reflect.ValueOf(sysConfig)
	return r.FieldByName(item).String()
}

// 根据item获取获取配置值
//
// 同Get()，但返回int64类型的值
func GetInt(item string) int64 {
	r := reflect.ValueOf(sysConfig)
	return r.FieldByName(item).Int()
}

