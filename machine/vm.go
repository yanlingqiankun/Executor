package machine

import (
	"fmt"
	"github.com/libvirt/libvirt-go"
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
		return fmt.Errorf("vm is Running, don't start a running container")
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
		return fmt.Errorf("vm is Running, can't destroy a running container")
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
	return domain.Destroy()
}

func prestartHookVM(containerID string) error {
	// todo network and volume
	return nil
}

func poststopHook(containerID string) {
	return
}

func StopVM(timeout int, UUID string) error {
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
	return domain.Shutdown()
}

func RestartVM(timeout int, UUID string) error {
	domain, err := libconn.LookupDomainByUUIDString(UUID)
	if err != nil {
		return err
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
	return []byte(str), err
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
