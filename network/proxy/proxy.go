package proxy

import (
	"errors"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
)

const UDP_DEFAULT_GC_INTERVAL = 10
const UDP_DEFAULT_GC_TIMEOUT = 30

var ErrNetClosing = errors.New("use of closed network connection")

var logger = logging.GetLogger("proxy")

func init(){
	level := logging.GetLevel(conf.GetString("LogLevel"))
	logger.SetLevel(level)
}

func GetProxyManager() ProxyManager {
	return &Proxies{containerProxies: make([]Proxy, 0)}
}

func (p *Proxies) Add(src, dst, protocol string) error {
	p.containerProxies = append(p.containerProxies, getProxy(src, dst, protocol))
	return nil
}

func (p *Proxies) StartAll() (errs []error) {
	for _, containerProxy := range p.containerProxies {
		if containerProxy.IsConnected() {
			continue
		}
		if err := containerProxy.Start(); err != nil {
			errs = append(errs, err)
		}
	}
	return
}

func (p *Proxies) StopAll() (errs []error) {
	for _, containerProxy := range p.containerProxies {
		if !containerProxy.IsConnected() {
			continue
		}
		if err := containerProxy.Stop(); err != nil {
			errs = append(errs, err)
		}
	}
	return
}
