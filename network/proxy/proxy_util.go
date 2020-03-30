package proxy

import (
	"net"
	"strconv"
)

func getProxy(src, dst, p string) Proxy {
	config := Config{
		Protocol: p,
		FromAddr: src,
		ToAddr:   dst,
	}
	if p == "udp"{
		u := new(UDP)
		u.set(&config)
		return u
	} else if p == "tcp" {
		t := new(TCP)
		t.set(&config)
		return t
	}
	return nil
}

func IsSrcIP(str string) (result bool) {
	result = false
	if str == "" || net.ParseIP(str) != nil {
		result = true
	}
	return
}

func IsPort(str string) (result bool) {
	result = false
	i, err := strconv.Atoi(str)
	if err == nil && i > -1 && i < 65536 {
		result = true
	}
	return
}