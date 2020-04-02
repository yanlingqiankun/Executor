package machine

import (
	"encoding/xml"
	"github.com/docker/docker/pkg/stringid"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/image"
	"github.com/yanlingqiankun/Executor/network/proxy"
	"path/filepath"
	"strconv"
)


func CreateVM(imageID string) Factory {
	vmType := conf.GetString("VMType")
	VM := &BaseVM{
		ImageID:       imageID,
		VMConfig:      &libvirtxml.Domain{
			XMLName:              xml.Name{},
			Type:                 vmType,
			Memory:               &libvirtxml.DomainMemory{
				Value:    0,
				Unit:     "1024000",
				DumpCore: "",
			},
			CurrentMemory:        nil,
			BlockIOTune:          nil,
			MemoryTune:           nil,
			MemoryBacking:        nil,
			VCPU:                 &libvirtxml.DomainVCPU{
				Placement: "",
				CPUSet:    "",
				Current:   "",
				Value:     1,
			},

			OS:                  &libvirtxml.DomainOS{
				Type:        &libvirtxml.DomainOSType{
					Arch:    "x86_64",
					Machine: "pc",
					Type:    "",
				},
			},
			Devices:  &libvirtxml.DomainDeviceList{
				Graphics:     append(make([]libvirtxml.DomainGraphic, 0), libvirtxml.DomainGraphic{
					XMLName:     xml.Name{},
					SDL:         nil,
					VNC:         &libvirtxml.DomainGraphicVNC{
						Port:          -1,
						AutoPort:      "yes",
						WebSocket:     0,
						Keymap:        "en-us",
						Listen:        "0.0.0.0",
					},
				}),
			},
		},
		RuntimeConfig: &RuntimeConfig{},
	}
	VM.SetImage(imageID)
	return VM
}
// Factory interface
func (VM *BaseVM) SetName(name string) error {
	if name == "" {
		return nil
	}
	VM.VMConfig.Name = name
	return nil
}

func (VM *BaseVM) SetImage(imageID string) {
	VM.VMConfig.Devices = &libvirtxml.DomainDeviceList{}
	VM.VMConfig.Devices.Disks = append(VM.VMConfig.Devices.Disks, libvirtxml.DomainDisk{
		XMLName:      xml.Name{},
		Device:       "disk",
		RawIO:        "",
		SGIO:         "",
		Snapshot:     "",
		Model:        "",
		Driver:       &libvirtxml.DomainDiskDriver{
			Name:         "qemu",
			Type:         image.GetImageType(imageID),
			Cache:        "none",
			ErrorPolicy:  "",
			RErrorPolicy: "",
			IO:           "",
			IOEventFD:    "",
			EventIDX:     "",
			CopyOnRead:   "",
			Discard:      "",
			IOThread:     nil,
			DetectZeros:  "",
			Queues:       nil,
			IOMMU:        "",
			ATS:          "",
		},
		Auth:         nil,
		Source:       &libvirtxml.DomainDiskSource{
			File:          &libvirtxml.DomainDiskSourceFile{
				File:     filepath.Join(conf.GetString("RootPath"), "images", imageID, imageID),
				SecLabel: nil,
			},
			Block:         nil,
			Dir:           nil,
			Network:       nil,
			Volume:        nil,
			NVME:          nil,
			StartupPolicy: "",
			Index:         0,
			Encryption:    nil,
			Reservations:  nil,
		},
		BackingStore: nil,
		Geometry:     nil,
		BlockIO:      nil,
		Mirror:       nil,
		Target:       &libvirtxml.DomainDiskTarget{
			Dev:       "sda",
			Bus:       "usb",
			Tray:      "",
			Removable: "",
		},
		IOTune:       nil,
		ReadOnly:     nil,
		Shareable:    nil,
		Transient:    nil,
		Serial:       "",
		WWN:          "",
		Vendor:       "",
		Product:      "",
		Encryption:   nil,
		Boot:         nil,
		Alias:        nil,
		Address:      nil,
	})
}

func (VM *BaseVM) SetHostname(name string) {
	if name == "" {
		return
	}
	VM.VMConfig.Name = name
}

func (VM *BaseVM) SetVolumes(volumes []*Volume) {
	if volumes == nil || len(volumes) == 0 {
		return
	}
	//for _, v := range volumes {
	//	dest := v.Destination
	//	if dest == "" {
	//		continue
	//	}
	//
	//	VM.RuntimeConfig.Volumes[dest] = v
	//}

	// TODO drive filesystem to mount
}

func (VM *BaseVM) SetEntrypoint(entrypoint []string) {
	return
}

func (VM *BaseVM) SetCmd(cmd []string) {
	return
}

func (VM *BaseVM) SetWorkingDir(dir string) {
	return
}

func (VM *BaseVM) SetUser(user string) {
	if user == "" {
		return
	}
	return
}

func (VM *BaseVM) SetEnv(env []string) {
	if env == nil || len(env) == 0 {
		return
	}
	m := envSliceToMap(env)
	for k, v := range m {
		VM.VMConfig.OS.InitEnv = append(VM.VMConfig.OS.InitEnv, libvirtxml.DomainOSInitEnv{
			Name:  k,
			Value: v,
		})
	}
}

func (VM *BaseVM) SetTTY(tty bool) {
	return
}

func (VM *BaseVM) SetExposedPorts(info []proxy.ProxyInfo) {
	VM.RuntimeConfig.ProxyManager = proxy.GetProxyManager()
	if len(VM.RuntimeConfig.Networks) == 0 {
		return
	}
	VM.RuntimeConfig.ExposedPorts = info
}

func (VM *BaseVM) SetHosts(hosts []string) {
	//VM.HostConfig.ExtraHosts = hosts
	return
}

func (VM *BaseVM) SetTTYSize(width, height uint16) {
	return
}

func (VM *BaseVM) GetBase() (*Base, error) {
	return &Base{
		IsDocker:       true,
		ImageID:        VM.ImageID,
		ID:             VM.VMConfig.UUID,
		Name:           VM.VMConfig.Name,
		RuntimeSetting: VM.RuntimeConfig,
	}, nil
}


func (VM *BaseVM) SetNetworks(networks []*Network) {
	for idx, nw := range networks {
		if nw.Name == "" {
			nw.Name = "eth" + strconv.Itoa(idx)
		}

		if nw.HostInterfaceName == "" {
			nw.HostInterfaceName = "veth" + stringid.GenerateRandomID()[:6] + strconv.Itoa(idx)
		}

	}
	VM.RuntimeConfig.Networks = networks
}

func (VM *BaseVM) SetRoutes([]*Route) {
	panic("implement me")
}




