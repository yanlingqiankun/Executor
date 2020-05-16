package machine

import (
	"encoding/xml"
	"github.com/docker/docker/api/types/container"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/image"
	"github.com/yanlingqiankun/Executor/network/proxy"
	"github.com/yanlingqiankun/Executor/stringid"
	"github.com/yanlingqiankun/Executor/util"
	"runtime"
	"strconv"
	"time"
)

func CreateVM(imageID string) Factory {
	uuid, err := stringid.GetStanderUUID()
	if err != nil {
		logger.WithError(err).Error("failed to get uuid")
	}
	vmType := conf.GetString("VMType")
	VM := &BaseVM{
		BaseInfo:      &Base{
			IsDocker:      false,
			ImageID:       imageID,
			ID:            "",
			ImagePath:     "",
			ImageType:     "",
			Name:          "",
			RuntimeConfig: &RuntimeConfig{},
		},
		VMConfig:      &libvirtxml.Domain{
			XMLName:              xml.Name{},
			Type:                 vmType,
			UUID:                 uuid,
			Memory:               &libvirtxml.DomainMemory{
				Value:    1024000,
				Unit:     "",
				DumpCore: "",
			},
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
					Type:    "hvm",
				},
				BootDevices:append(make([]libvirtxml.DomainBootDevice, 0), libvirtxml.DomainBootDevice{Dev:"hd"}),
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
	}
	return VM
}
// Factory interface
func (VM *BaseVM) Create() error {
	xmlStr, err := VM.VMConfig.Marshal()
	if err != nil {
		return err
	}
	_, err = libconn.DomainDefineXML(xmlStr)
	if err != nil {
		return err
	}
	id := util.UUIDTOID(VM.VMConfig.UUID)
	VM.BaseInfo.ID = id
	//if VM.BaseInfo.RuntimeConfig.Networks != nil {
	//	for _ ,net := range VM.BaseInfo.RuntimeConfig.Networks {
	//		network.Restart(net.Bridge)
	//	}
	//}
	logger.Debugf("The VM %s created successfully", id)
	return nil
}

func (VM *BaseVM) SetName(name string) error {
	if name == "" {
		return nil
	}
	VM.VMConfig.Name = name
	VM.BaseInfo.Name = name
	return nil
}

func (VM *BaseVM) SetImage(imageID string, path string, name string) {
	VM.BaseInfo.ImageName = name
	imageType := "iso"
	device := "cdrom"
	driver := &libvirtxml.DomainDiskDriver{}
	if image.GetImageType(imageID) == "disk" {
		imageType = "qcow2"
		driver = &libvirtxml.DomainDiskDriver{
			Name:         "qemu",
			Type:         imageType,
		}
		device = "disk"
		VM.BaseInfo.ImagePath = path
	} else {
		driver = nil
	}
	VM.BaseInfo.ImageType = image.GetImageType(imageID)
	if VM.VMConfig.Devices == nil {
		VM.VMConfig.Devices = &libvirtxml.DomainDeviceList{}
	}
	VM.VMConfig.Devices.Disks = append(VM.VMConfig.Devices.Disks, libvirtxml.DomainDisk{
		XMLName:      xml.Name{},
		Device:       device,
		Driver:       driver,
		Auth:         nil,
		Source:       &libvirtxml.DomainDiskSource{
			File:          &libvirtxml.DomainDiskSourceFile{
				File:     path,
				SecLabel: nil,
			},
		},
		Target:       &libvirtxml.DomainDiskTarget{
			Dev:       "hdb",
			Bus:       "ide",
			Tray:      "",
			Removable: "",
		},
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
	VM.BaseInfo.RuntimeConfig.Volumes = make(map[string]*Volume)
	for _, v := range volumes {
		dest := v.Destination
		if dest == "" {
			continue
		}
		var readOnly *libvirtxml.DomainFilesystemReadOnly
		readOnly = nil
		if v.RW == WritePermission {
			readOnly = &libvirtxml.DomainFilesystemReadOnly{}
		}
		VM.VMConfig.Devices.Filesystems = append(VM.VMConfig.Devices.Filesystems, libvirtxml.DomainFilesystem{
			XMLName:        xml.Name{},
			AccessMode:     "passthrough",
			Model:          "",
			Driver:         &libvirtxml.DomainFilesystemDriver{
				Type:     "path",
				Format:   "",
				Name:     "",
				WRPolicy: "immediate",
				IOMMU:    "",
				ATS:      "",
			},
			Source:         &libvirtxml.DomainFilesystemSource{
				Mount:    &libvirtxml.DomainFilesystemSourceMount{Dir:v.Source},
			},
			Target:         &libvirtxml.DomainFilesystemTarget{Dir:v.Destination},
			ReadOnly:       readOnly,
		})

		VM.BaseInfo.RuntimeConfig.Volumes[dest] = v
	}

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
	return
}

func (VM *BaseVM) SetTTY(tty bool) {
	if VM.VMConfig.Devices == nil {
		VM.VMConfig.Devices = &libvirtxml.DomainDeviceList{}
	}

	var port uint
	port = 0
	VM.VMConfig.Devices.Serials = append(make([]libvirtxml.DomainSerial, 0), libvirtxml.DomainSerial{
		XMLName:  xml.Name{},
		Source:   nil,
		Protocol: nil,
		Target:   &libvirtxml.DomainSerialTarget{
			Type:  "",
			Port:  &port,
			Model: nil,
		},
		Log:      nil,
		Alias:    nil,
		Address:  nil,
	})

	VM.VMConfig.Devices.Consoles = append(make([]libvirtxml.DomainConsole, 0), libvirtxml.DomainConsole{
		XMLName:  xml.Name{},
		TTY:      "pty",
		Source:   nil,
		Protocol: nil,
		Target:   &libvirtxml.DomainConsoleTarget{
			Type: "serial",
			Port: &port,
		},
		Log:      nil,
		Alias:    nil,
		Address:  nil,
	})
	VM.BaseInfo.RuntimeConfig.Tty = tty
	return
}

func (VM *BaseVM) SetExposedPorts(info []proxy.ProxyInfo) {
	VM.BaseInfo.RuntimeConfig.ProxyManager = proxy.GetProxyManager()
	if len(VM.BaseInfo.RuntimeConfig.Networks) == 0 {
		return
	}
	VM.BaseInfo.RuntimeConfig.ExposedPorts = info
}

func (VM *BaseVM) SetHosts(hosts []string) {
	//VM.HostConfig.ExtraHosts = hosts
	return
}

func (VM *BaseVM) SetTTYSize(width, height uint16) {
	return
}

func (VM *BaseVM) GetBase() (*Base, error) {
	VM.BaseInfo.CreateTime = time.Now()
	return VM.BaseInfo, nil
}

func (VM *BaseVM) SetNetworks(networks []*Network) {
	for idx, nw := range networks {
		if nw.Name == "" {
			nw.Name = "eth" + stringid.GenerateRandomID()[:5] + strconv.Itoa(idx)
		}

		VM.VMConfig.Devices.Interfaces = append(VM.VMConfig.Devices.Interfaces, libvirtxml.DomainInterface{
			XMLName:             xml.Name{},
			Managed:             "",
			TrustGuestRXFilters: "",
			MAC:                 &libvirtxml.DomainInterfaceMAC{Address:nw.MacAddress},
			Source:              &libvirtxml.DomainInterfaceSource{
				Network:   &libvirtxml.DomainInterfaceSourceNetwork{
					Network:   nw.Bridge,
				},
			},
			Route:               nil,
			Target:              &libvirtxml.DomainInterfaceTarget{
				Dev:     nw.Name,
			},
		})
	}
	VM.BaseInfo.RuntimeConfig.Networks = networks
}

func (VM *BaseVM) SetRoutes([]*Route) {
	return
}

func (VM *BaseVM) SetCgroups(res container.Resources) {
	logger.Debugf("prepare to set resource limit for VM")
	VM.VMConfig.CPUTune = &libvirtxml.DomainCPUTune{}
	//CPU
	if res.CpusetCpus != "" {
		VM.VMConfig.VCPU = &libvirtxml.DomainVCPU{
			Placement: "",
			CPUSet:    res.CpusetCpus,
			Current:   "",
			Value:     runtime.NumCPU(),
		}
	}
	if res.CPUPeriod >= 1000 {
		VM.VMConfig.CPUTune.Period = &libvirtxml.DomainCPUTunePeriod{Value:uint64(res.CPUPeriod)}
	}
	if res.CPUQuota >= 1000 {
		VM.VMConfig.CPUTune.Quota =  &libvirtxml.DomainCPUTuneQuota{Value:res.CPUQuota}
	}
	if res.CPURealtimePeriod >= 1000 {
		VM.VMConfig.CPUTune.GlobalPeriod = &libvirtxml.DomainCPUTunePeriod{Value:uint64(res.CPURealtimeRuntime)}
	}
	if res.CPURealtimeRuntime >= 1000 {
		VM.VMConfig.CPUTune.GlobalQuota = &libvirtxml.DomainCPUTuneQuota{Value:res.CPURealtimeRuntime}
	}

	//Memory
	if res.Memory != 0 {
		VM.VMConfig.Memory = &libvirtxml.DomainMemory{
			Value:    uint(res.Memory),
			Unit:     "",
			DumpCore: "",
		}
	}

	//Device
	VM.VMConfig.BlockIOTune = &libvirtxml.DomainBlockIOTune{}
	if res.BlkioWeight != 0 {
		VM.VMConfig.BlockIOTune.Weight = uint(res.BlkioWeight)
	}
	deviceList := getDeviceLimits(res)
	for _, v := range deviceList {
		VM.VMConfig.BlockIOTune.Device = append(VM.VMConfig.BlockIOTune.Device, libvirtxml.DomainBlockIOTuneDevice{
			Path:          v.path,
			Weight:        v.weight,
			ReadIopsSec:   v.readIOpsSec,
																																																																																																			WriteIopsSec:  v.writeIOpsSec,
			ReadBytesSec:  v.readBytesSec,
			WriteBytesSec: v.writeBytesSec,
		})
	}
}

type deviceLimit struct {
	path string
	weight uint
	readBytesSec uint
	writeBytesSec uint
	readIOpsSec uint
	writeIOpsSec uint
}

func getDeviceLimits(res container.Resources) map[string]deviceLimit {
	deviceList := make(map[string]deviceLimit)
	if res.BlkioWeightDevice != nil {
		for _, d := range res.BlkioWeightDevice {
			deviceList[d.Path] = deviceLimit{
				path:          d.Path,
				weight:        uint(d.Weight),
			}
		}
	}
	if res.BlkioDeviceReadBps != nil {
		for _, d := range res.BlkioDeviceReadBps {
			if device, ok := deviceList[d.Path]; ok {
				device.readBytesSec = uint(d.Rate)
			} else {
				deviceList[d.Path] = deviceLimit{
					path:          d.Path,
					readBytesSec:  uint(d.Rate),
				}
			}
		}
	}
	if res.BlkioDeviceWriteBps != nil {
		for _, d := range res.BlkioDeviceWriteBps {
			if device, ok := deviceList[d.Path]; ok {
				device.writeBytesSec = uint(d.Rate)
			} else {
				deviceList[d.Path] = deviceLimit{
					path:          d.Path,
					writeBytesSec:  uint(d.Rate),
				}
			}
		}
	}

	if res.BlkioDeviceReadIOps != nil {
		for _, d := range res.BlkioDeviceReadIOps {
			if device, ok := deviceList[d.Path]; ok {
				device.readIOpsSec = uint(d.Rate)
			} else {
				deviceList[d.Path] = deviceLimit{
					path:          d.Path,
					readIOpsSec:  uint(d.Rate),
				}
			}
		}
	}
	if res.BlkioDeviceWriteIOps != nil {
		for _, d := range res.BlkioDeviceWriteIOps {
			if device, ok := deviceList[d.Path]; ok {
				device.writeIOpsSec = uint(d.Rate)
			} else {
				deviceList[d.Path] = deviceLimit{
					path:          d.Path,
					writeIOpsSec:  uint(d.Rate),
				}
			}
		}
	}
	return deviceList
}
