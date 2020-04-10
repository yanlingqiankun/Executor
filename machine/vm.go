package machine

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"time"
)

func StartVM(UUID string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		return err
	}
	defer domain.Free()
	domainInfo, err := domain.GetInfo()
	if err != nil {
		return err
	}
	state := domainInfo.State
	if state != libvirt.DOMAIN_SHUTDOWN && state != libvirt.DOMAIN_SHUTOFF && state != libvirt.DOMAIN_NOSTATE {
		return fmt.Errorf("vm is Running, don't start a running vm")
	}

	err = domain.Create()
	if err != nil {
		logger.WithError(err).Errorf("failed to start vm %s", UUID)
		return err
	}
	return nil
}

func DeleteVM(UUID string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		logger.WithError(err).Error("failed to find VM")
		return err
	}
	defer domain.Free()
	domainInfo, err := domain.GetInfo()
	if err != nil {
		return err
	}
	state := domainInfo.State
	if state != libvirt.DOMAIN_SHUTDOWN && state != libvirt.DOMAIN_SHUTOFF {
		return fmt.Errorf("vm is Running, can't destroy a running vm")
	}

	if ok, _ := domain.IsActive(); ok {
		if err := domain.Destroy(); err != nil {
			logger.WithError(err).Error("failed to destroy VM")
			return err
		}
	}
	if ok, _ := domain.IsPersistent(); ok {
		if err := domain.Undefine(); err != nil {
			logger.WithError(err).Error("failed to undefine VM")
			return err
		}
	}
	return nil
}

func PauseVM(UUID string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		return err
	}
	defer domain.Free()
	domainInfo, err := domain.GetInfo()
	if err != nil {
		return err
	}
	state := domainInfo.State
	if state != libvirt.DOMAIN_RUNNING {
		return fmt.Errorf("vm is not Running")
	}

	return domain.Suspend()
}

func UnpauseVM(UUID string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		return err
	}
	defer domain.Free()
	domainInfo, err := domain.GetInfo()
	if err != nil {
		return err
	}
	state := domainInfo.State
	if state != libvirt.DOMAIN_PAUSED {
		return fmt.Errorf("vm is not PAUSED")
	}

	return domain.Resume()
}

func KillVM(UUID string, signal string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		return err
	}
	defer domain.Free()
	domainInfo, err := domain.GetInfo()
	if err != nil {
		return err
	}
	state := domainInfo.State
	if state == libvirt.DOMAIN_SHUTOFF || state == libvirt.DOMAIN_SHUTDOWN {
		return fmt.Errorf("vm is not running")
	}
	if err := domain.Destroy(); err != nil {
		logger.WithError(err).Error("failed destroy vm")
		return err
	}
	return nil
}

func prestartHookVM(containerID string) error {
	// todo network and volume
	return nil
}

func poststopHook(containerID string) {
	return
}

func StopVM(timeout int32, UUID string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		return err
	}
	domainInfo, err := domain.GetInfo()
	if err != nil {
		return err
	}
	state := domainInfo.State
	if state == libvirt.DOMAIN_SHUTOFF || state == libvirt.DOMAIN_SHUTDOWN {
		return fmt.Errorf("vm is not running")
	}
	go func(){
		ticker := time.NewTicker(time.Duration(timeout)*time.Second)
		select {
		case <-ticker.C:
			if domain == nil {
				return
			}
			if ok, err := domain.IsActive(); err != nil {
				return
			} else if ok {
				domain.Destroy()
			}
		}
		ticker.Stop()
	}()
	err = domain.DestroyFlags(libvirt.DOMAIN_DESTROY_GRACEFUL)
	if err != nil {
		return err
	}
	return nil
}

func RestartVM(timeout int32, UUID string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		return err
	}
	domainInfo, err := domain.GetInfo()
	if err != nil {
		return err
	}
	state := domainInfo.State
	if state == libvirt.DOMAIN_SHUTOFF || state == libvirt.DOMAIN_SHUTDOWN {
		return fmt.Errorf("vm is not running")
	}
	defer domain.Free()

	return domain.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT)
}

func getVMInfo(id string) ([]byte, error){
	dom, err := libconn.LookupDomainByUUIDString(id)
	if err != nil {
		return nil, err
	}
	defer dom.Free()
	str, err := dom.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	var temp = libvirtxml.Domain{}
	err = xml.Unmarshal([]byte(str), &temp)
	if err != nil {
		return []byte{}, err
	}
	data, err := json.Marshal(temp)
	return data, err
}

func getVMState(id string) (string, error){
	states := []string{
		"nostate",
		"running",
		"blocked",
		"paused",
		"shutdown",
		"shutoff",
		"crashed",
		"pm-suspended",
	}
	dom, err := libconn.LookupDomainByUUIDString(id)
	if err != nil {
		return "", err
	}
	defer dom.Free()
	state, _, err := dom.GetState()
	if err != nil {
		return "", err
	} else if int(state) > len(states) - 1{
		return "", fmt.Errorf("unknown vm state")
	} else {
		return states[int(state)], nil
	}
}

func renameVM(id string, newName string) error {
	dom, err := libconn.LookupDomainByUUIDString(id)
	if err != nil {
		return err
	}
	defer dom.Free()
	err = dom.Rename(newName, 0)
	if err != nil {
		return err
	}
	return nil
}

type stdio map[chan []byte] *[2]chan []byte
var stdios map[string] *stdio
func getVMStdio(id string) (chan []byte, chan[]byte, chan[] byte, error) {
	//if _, ok := stdios[id]; ok {
	//
	//} else {
	domain, err := libconn.LookupDomainByUUIDString(id)
	if err != nil {
		return nil, nil, nil, err
	}
	defer domain.Free()
	stream, err := libconn.NewStream(0)
	if err != nil {
		return nil, nil, nil, err
	}
	err = domain.OpenConsole("", stream, libvirt.DOMAIN_CONSOLE_FORCE)
	if err != nil {
		logger.WithError(err).Error("failed to open console")
		return nil,nil, nil, err
	}
	stdin := make(chan []byte, 16)
	stdout := make(chan []byte, 16)
	stderr := make(chan []byte, 16)

	go stdinVMHandle(stream, stdin)
	go stdoutVMHandle(stream, stdout, stderr)

	return stdin, stdout, stderr, nil
	//}
}

func stdinVMHandle(stream *libvirt.Stream, in chan []byte) {
	defer func() {
		stream.Finish()
	}()
	for data := range in {
		_, err := stream.Send(data)
		if err != nil {
			logger.WithError(err).Error("close attach with error")
			return
		}
	}
}

func stdoutVMHandle(stream *libvirt.Stream, out chan []byte, stderr chan []byte) {
	for {
		data := make([]byte, 4096)
		n, err := stream.Recv(data)
		if err != nil {
			logger.WithError(err).Error("close attach with error")
			stream.Finish()
			stream.Free()
			close(out)
			close(stderr)
			return
		} else {
			out <- data[:n]
		}
	}
}
