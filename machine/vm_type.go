package machine

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"github.com/yanlingqiankun/Executor/conf"
)


func createVM(imageID string) *BaseVM {
	return &BaseVM{
		ImageID:  imageID,
		VMConfig: &libvirtxml.Domain{Type: conf.GetString("VMType")},
	}
}

type BaseVM struct {
	ImageID       string
	VMConfig      *libvirtxml.Domain
	RuntimeConfig *RuntimeConfig
}


