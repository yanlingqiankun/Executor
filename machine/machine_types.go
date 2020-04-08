package machine

import (
	"github.com/yanlingqiankun/Executor/network/proxy"
	"time"
)

type Machine interface {
	Start() error                   // 运行machine
	Kill(string) error              // 使用指定的信号来停止一个容器
	Pause() error                   // 暂停
	Unpause() error                 // 恢复
	Delete() error                  // 删除这个machine所占用的资源
	Stop(int32) error                 // 停止
	Restart(int) error              // 重启
	GetImageID() string             //
	Rename(string) error            // rename
}

type Factory interface {
	SetName(string) error          // 容器的名字
	SetHostname(string)            // 设置主机名
	SetTTY(bool)                   // 设置是否使用终端
	SetNetworks([]*Network)        // 添加网络
	SetRoutes([]*Route)            // 添加路由
	SetImage(string, string, string)               // 添加镜像地址
	SetVolumes([]*Volume)          // 添加卷
	SetEntrypoint([]string)        // EntryPoint (entrypoint, cmd)
	SetCmd([]string)               // Cmd
	SetWorkingDir(string)          // WorkingDir 默认 "/"
	SetUser(string)                // User
	SetEnv([]string)               // Env
	SetTTYSize(width, height uint16)
	GetBase() (*Base, error)	  // get machine entry
	SetHosts([]string)            // 格式：{"hostname:192.168.0.2"}
	SetExposedPorts(info []proxy.ProxyInfo)
	Create() error
}

type Base struct {
	IsDocker       bool             `json:"is_docker"`
	ImageID 	   string			`json:"image_id"`
	ID             string          `json:"id"`
	ImageName      string           `json:"imageName"`
	ImagePath      string 			`json:"image_path"`
	ImageType 	   string           `json:"image_type"`
	Name           string          `json:"name"` // 名字
	CreateTime     time.Time           `json:"create_time"`
	RuntimeConfig *RuntimeConfig  `json:"runtime_config"` // 运行配置
}

type MachineInfo struct {
	ImageName    string
	ID           string
	Name         string
	CreateTime   string
	ImageType    string
	Status       string
	ImageId      string
}
