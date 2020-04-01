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
	if state != libvirt.DOMAIN_SHUTDOWN || state != libvirt.DOMAIN_SHUTOFF || state != libvirt.DOMAIN_NOSTATE {
		return fmt.Errorf("vm is Running, don't start a running container")
	}

	return domain.Create()
}

func DeleteVM(UUID string) error {
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
	if state != libvirt.DOMAIN_SHUTDOWN || state != libvirt.DOMAIN_SHUTOFF || state != libvirt.DOMAIN_NOSTATE {
		return fmt.Errorf("vm is Running, don't destroy a running container")
	}


	//TODO remove in db
	if err := domain.Destroy(); err != nil {
		return err
	}
	return domain.Undefine()
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
