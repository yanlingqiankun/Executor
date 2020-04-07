package machine

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

type BaseVM struct {
	BaseInfo       *Base
	VMConfig      *libvirtxml.Domain
}
