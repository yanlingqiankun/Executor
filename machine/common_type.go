package machine

import "github.com/yanlingqiankun/Executor/network/proxy"

// volume permission
const (
	ReadPermission  uint8 = 1
	WritePermission uint8 = 2
)

type ContainerVolume struct {
	Destination string `json:"destination"`
	RW          uint8  `json:"rw"`
	Source      string `json:"source"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}


type Network struct {
	// Type 设置 networks type, 通常设置为 veth 或者 loopback
	Type string `json:"type"`

	// Name 设置网络接口名字比如 eth0
	Name string `json:"name"`

	// Bridge 设置要使用的网桥
	Bridge string `json:"bridge"`

	// MacAddress 设置接口的 mac 地址
	MacAddress string `json:"mac_address"`

	// Address 设置包含 mask 的 IPV4 地址 比如 192.169.10.2/24
	Address []string `json:"address"`

	// Gateway 设置接口的默认 ipv4 网关地址
	Gateway string `json:"gateway"`

	// IPv6Address 设置包含 mask 的 IPV6 地址
	IPv6Address []string `json:"ipv6_address"`

	// IPv6Gateway 设置接口的默认 ipv6 网关地址
	IPv6Gateway string `json:"ipv6_gateway"`

	// 设置一対 veth 之间的 mtu
	// 注意：不会作用于：loopback
	Mtu int `json:"mtu"`

	// 设置一対 veth 之间的 tx_queuelen
	// 注意：不会作用于：loopback
	TxQueueLen int `json:"txqueuelen"`

	// 一対 veth 中在宿主机上的 veth 的名字
	HostInterfaceName string `json:"host_interface_name"`
}

type Route struct {
	// 设置 destination 和 mask 必须是 CIDR 格式
	// 支持 IPV4 和 IPV6 比如：192.168.1.2/24
	Destination string `json:"destination"`

	// 设置 source 和 mask 必须是 CIDR 格式
	// 支持 IPV4 和 IPV6 比如：192.168.1.3/24
	Source string `json:"source"`

	// 设置网关
	// 支持 IPV4 和 IPV6 比如：192.168.1.1
	Gateway string `json:"gateway"`

	// 设备名 比如 eth0
	InterfaceName string `json:"interface_name"`
}

type networkConfig struct {
	Pid      string     `json:"pid"`
	Networks []*Network `json:"networks"`
	Routes   []*Route   `json:"routes"`
}

type RuntimeConfig struct {
	Tty           bool                       `json:"tty"`
	Networks     []*Network                  `json:"networks"` // 网络设备配置
	Routes       []*Route                    `json:"routes"`   // 路由配置
	Volumes      map[string]*Volume          `json:"volumes"`
	ExposedPorts []proxy.ProxyInfo           `json:"exposed_ports"`
	ProxyManager proxy.ProxyManager
}

type Volume struct {
	Destination string `json:"destination"`
	RW          uint8  `json:"rw"`
	Source      string `json:"source"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

