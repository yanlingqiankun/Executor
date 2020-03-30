package proxy

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

//定义初始化连接
func (t *UDP) set(config *Config) {
	t.config = config
	t.sessions = new(sync.Map)
	t.closeChan = make(chan struct{}, 1)
	t.isConnect = false
}

//定义连接函数
func (t *UDP) accept()  {
	logger.Debugf("UDP.Start %v, backend: %v", t.config.FromAddr, t.config.ToAddr)

	go t.gc()
	data := make([]byte, 4096)
	for {
		n, remoteAddr, err := t.listener.ReadFromUDP(data)
		if err != nil {
			if !strings.Contains(err.Error(), ErrNetClosing.Error()) {
				logger.WithError(err).Error("UDP.Listener.Read error")
			}
			return
		}
		var dstConnection *net.UDPConn

		// the session in the sessions
		c, ok := t.sessions.Load(remoteAddr.String())
		if ok {
			dstConnection = c.(*udpSession).Conn
			_, err = dstConnection.Write(data[:n])
			if err != nil {
				logger.Errorf("failed to sent udp data to [%v] : %v", dstConnection.RemoteAddr().String(), err)
			} else {
				atomic.StoreInt64(&c.(*udpSession).Active, time.Now().Unix())
				continue
			}
		}

		// don't have the session
		udpAddr, _ := net.ResolveUDPAddr("udp", t.config.ToAddr)
		dstConnection, err = net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			logger.Errorf("failed to connect udp server [%v] : %v\n", remoteAddr.String(), err)
			break
		}

		_, err = dstConnection.Write(data[:n])
		if err != nil {
			logger.Errorf("failed to sent udp date to %v : %v\n", dstConnection.RemoteAddr().String(), err)
		} else {
			t.sessions.Store(remoteAddr.String(), &udpSession{Conn: dstConnection, Active: time.Now().Unix()})
		}

		go t.response(dstConnection, t.listener, remoteAddr)
	}
}

func (t *UDP)response(dstConnection, remoteConnection *net.UDPConn, addr *net.UDPAddr) {
	data := make([]byte, 4096)
	for {
		n, _, err := dstConnection.ReadFromUDP(data)
		if err != nil {
			if !strings.Contains(err.Error(), ErrNetClosing.Error()) {
				logger.Errorf("failed to receive udp data from %v : %v\n", dstConnection.RemoteAddr().String(), err)
			}
			break
		}

		_, err = remoteConnection.WriteToUDP(data[:n], addr)
		if s, ok := t.sessions.Load(addr.String()); ok {
			atomic.StoreInt64(&s.(*udpSession).Active, time.Now().Unix())
		}
		if err != nil {
			logger.Errorf("failed to sent udp data to %v: %v\n", addr.String(), err)
			return
		}
	}
}

//定义超时终止
func (t *UDP) gc() {
	ticker := time.NewTicker(time.Second * UDP_DEFAULT_GC_INTERVAL)
	for {
		select {
		case <- t.closeChan:
			ticker.Stop()
			return
		case <- ticker.C:
			t.sessions.Range(func(key, val interface{}) bool {
				conn, ok := val.(*udpSession)
				if !ok {
					t.sessions.Delete(key)
					return false
				}
				active := atomic.LoadInt64(&conn.Active)
				if active + UDP_DEFAULT_GC_TIMEOUT < time.Now().Unix() {
					conn.Conn.Close()
					t.sessions.Delete(key)
				}
				return true
			})
		}
	}
}

func (t *UDP) closeAllSessions() {
	t.sessions.Range(func (key, val interface{}) bool {
		conn, ok := val.(*udpSession)
		if !ok {
			return true
		}
		conn.Conn.Close()
		return true
	})
	t.sessions = nil
}

func (t *UDP) Start() error {
	addr, err := net.ResolveUDPAddr("udp", t.config.FromAddr)
	if err != nil {
		logger.WithError(err).Error("UDP.Start error")
		return fmt.Errorf("UDP.Start error: %v\n", err)
	}
	t.listener, err = net.ListenUDP("udp", addr)
	if err != nil {
		logger.WithError(err).Error("UDP.Start error")
		return fmt.Errorf("UDP.Start error: %v\n", err)
	}
	t.isConnect = true
	go t.accept()
	return nil
}

func (t *UDP) Stop() error {
	if t.isConnect{
		err := t.listener.Close()
		if err != nil{
			return err
		}
		close(t.closeChan)
		t.closeAllSessions()
		logger.Debugf("UDP.Stop %s, backends: %s", t.config.FromAddr, t.config.ToAddr)
		return nil
	}
	return fmt.Errorf("The listen does not exist")
}

func (t *UDP) IsConnected () bool {
	return t.isConnect
}

