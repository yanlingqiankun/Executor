package proxy

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

//定义初始化连接
func (t *TCP) set(config *Config) {
	t.config = config
	t.isConnect = false
}

//定义连接函数
func (t *TCP) accept() {
	logger.Debugf("TCP.Connect %v, backend: %v", t.config.FromAddr, t.config.ToAddr)
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			if !strings.Contains(err.Error(), ErrNetClosing.Error()) {
				logger.WithError(err).Error("failed to proxy to ", t.config.ToAddr)
			}
			return
		}
		go t.handle(conn)
	}
}

//定义通信函数
func (t *TCP) handle(srcConnection net.Conn) {
	if srcConnection == nil {
		return
	}
	defer func() {
		srcConnection.Close()
	}()

	addr := t.config.ToAddr
	dstConnection, err := net.Dial("tcp", addr)
	if err != nil {
		logger.Errorf("TCP connect to %v failed: %v\n", addr, err)
		return
	}
	defer func() {
		dstConnection.Close()
	}()

	var wg sync.WaitGroup
	wg.Add(2)
	go forward(&wg, srcConnection, dstConnection)
	go forward(&wg, dstConnection, srcConnection)
	wg.Wait()
}

func forward(wg *sync.WaitGroup, srcConnection net.Conn, dstConnection net.Conn) {
	if srcConnection == nil || dstConnection == nil {
		return
	}

	_, err := io.Copy(dstConnection, srcConnection)
	if err != nil {
		if err == io.EOF {
			srcConnection.(*net.TCPConn).CloseRead()
			dstConnection.(*net.TCPConn).CloseWrite()
		} else if strings.Contains(err.Error(), ErrNetClosing.Error()) {

		} else {
			logger.WithError(err).Error("TCP can't send data")
		}
	}
	wg.Done()
}

func (t *TCP) Start() error {
	var err error
	t.listener, err = net.Listen("tcp", t.config.FromAddr)
	if err != nil {
		logger.WithError(err).Error("TCP : failed to listen ", t.config.FromAddr)
		return fmt.Errorf("failed to listen %s with error : %v", t.config.FromAddr, err)
	}
	t.isConnect = true
	go t.accept()
	return nil
}

func (t *TCP) Stop() error {
	listener, ok := t.listener.(*net.TCPListener)
	if ok {
		logger.Debugf("TCP.Close %s, backends: %s", t.config.FromAddr, t.config.ToAddr)
		return listener.Close()
	}
	return fmt.Errorf("The listen does not exist")
}

func (t *TCP) IsConnected () bool {
	return t.isConnect
}
