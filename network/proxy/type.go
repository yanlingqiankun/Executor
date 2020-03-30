package proxy

import (
	"net"
	"sync"
)

//定义配置参数
type Config struct {
	Protocol 	string
	FromAddr 	string
	ToAddr 		string
}


//定义TCP参数
type TCP struct {
	config 		*Config
	listener 	net.Listener
	isConnect	bool
}

//定义UDP参数
type UDP struct {
	config    *Config
	listener  *net.UDPConn
	sessions  *sync.Map
	closeChan chan struct{}
	isConnect bool
}

//定义连接参数
type udpSession struct {
	Conn	*net.UDPConn
	Active	int64
}

type ProxyInfo struct {
	Src      string
	DstPort  string
	Protocol string
}

//代理服务器接口
type Proxy interface {
	Start() error
	Stop() error
	IsConnected() bool
}

type ProxyManager interface {
	Add(src, dst, protocol string) error    // add a new proxy
	StartAll() []error     // start all proxy
	StopAll() []error      // stop all proxy
}

type Proxies struct {
	containerProxies []Proxy
}


