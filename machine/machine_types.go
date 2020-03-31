package machine

import (
	"github.com/yanlingqiankun/Executor/machine/types"
	"github.com/yanlingqiankun/Executor/network/proxy"
)

type Machine interface {
	Start() error                   // 运行machine
	Kill(string, bool) error        // 使用指定的信号来停止一个容器
	Pause() error                   // 暂停
	Unpause() error                 // 恢复
	Delete() error                  // 删除这个machine所占用的资源
	Stop(int) error                 // 停止
	Restart(int) error              // 重启
}

type Factory interface {
	SetName(string) error          // 容器的名字
	SetHostname(string)            // 设置主机名
	SetTTY(bool)                   // 设置是否使用终端
	SetNetworks([]*types.Network)        // 添加网络
	SetRoutes([]*types.Route)            // 添加路由
	SetImage(string)               // 添加镜像地址
	SetVolumes([]*types.Volume)          // 添加卷
	SetEntrypoint([]string)        // EntryPoint (entrypoint, cmd)
	SetCmd([]string)               // Cmd
	SetWorkingDir(string)          // WorkingDir 默认 "/"
	SetUser(string)                // User
	SetEnv([]string)               // Env
	SetTTYSize(width, height uint16)
	GetBase() *Base				   // get machine entry
	SetHosts([]string)            // 格式：{"hostname:192.168.0.2"}
	SetExposedPorts(info []proxy.ProxyInfo)
}

type Base struct {
	IsDocker 	   bool             `json:"is_docker"`
	ImageID 	   string			`json:"image_id"`
	ID             string          `json:"id"`
	Name           string          `json:"name"` // 名字
	RuntimeSetting *types.RuntimeConfig  `json:"runtime_config"` // 运行配置
}


