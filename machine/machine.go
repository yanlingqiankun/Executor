package machine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"github.com/vishvananda/netns"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/image"
	"github.com/yanlingqiankun/Executor/logging"
	"github.com/yanlingqiankun/Executor/network"
	"github.com/yanlingqiankun/Executor/network/proxy"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05.999999999 -0700 MST"

var logger = logging.GetLogger("machine")
var machineRootDir string
var db = new(machineDB)
var cli *client.Client
var libconn *libvirt.Connect

func init() {
	level := logging.GetLevel(conf.GetString("LogLevel"))
	logger.SetLevel(level)
	machineRootDir = filepath.Join(conf.GetString("RootPath"), "machines")
	client, err := client.NewClientWithOpts(client.FromEnv)
	cli = client
	if err != nil {
		logger.Fatal("error to init docker ", err.Error())
	}
	libconn, err = libvirt.NewConnect("qemu:///system")
	if err != nil {
		logger.Fatalf("failed to connect to xen")
	}
	db.init()
}

func ListMachine() []MachineInfo {
	list := make([]MachineInfo, 0)
	var err error
	db.Range(func(k, v interface{}) bool {
		item := v.(*dbItem)
		m := item.machine
		state := ""
		if m.IsDocker {
			state, err = getContainerState(m.ID)
			if err != nil {
				logger.WithError(err).Errorf("failed to get %s state", m.ID)
				state = "unknown"
			}
		} else {
			state, err = getVMState(m.ID)
			if err != nil {
				logger.WithError(err).Errorf("failed to get %s state", m.ID)
				state = "unknown"
			}
		}
		info := MachineInfo{
			ID:      m.ID,
			Name:    m.Name,
			ImageName: m.ImageName,
			ImageType : m.ImageType,
			CreateTime: m.CreateTime.Format(TIME_LAYOUT),
			Status:  state,
			ImageId: m.ImageID,
		}
		list = append(list, info)
		return true
	})
	return list
}

func GetMachine(id string) (Machine, error) {
	machine, ok := db.getItem(id)
	if !ok {
		return nil, fmt.Errorf("can't find the machine")
	}
	return machine.machine, nil
}

func AddMachine(f Factory) (string, error) {
	baseEntry, err := f.GetBase()
	if err != nil {
		return "", err
	}
	db.add(baseEntry, filepath.Join(machineRootDir, baseEntry.ID, "machine.json"))
	db.save(true, baseEntry.ID)
	return baseEntry.ID, nil
}

func (m *Base) Start() error {
	if m.IsDocker {
		err := StartContainer(m.ID)
		if err != nil {
			return err
		}
		nsHandle, err := netns.GetFromDocker(m.ID)
		if err != nil {
			return fmt.Errorf("failed to get %s's network namespace for %v", m.ID, err)
		}
		// set container network
		if m.RuntimeConfig.Networks != nil && len(m.RuntimeConfig.Networks) > 0 {
			for i, interf := range m.RuntimeConfig.Networks {
				err := interf.connectBridge(nsHandle, i)
				if err != nil {
					logger.WithError(err).Errorf("%s failed to connect to %s", interf.Name, interf.Bridge)
					return err
				}
			}
		}
	} else {
		err := StartVM(m.ID)
		if err != nil {
			return err
		}
	}
	m.prestartHook()
	return nil
}

func (m *Base) prestartHook() {
	logger.Debugf("%s prestart hook ", m.ID)
	if len(m.RuntimeConfig.Networks) > 0 && len(m.RuntimeConfig.ExposedPorts) > 0 {
		dstIP := strings.Split(m.RuntimeConfig.Networks[0].Address[0], "/")[0]
		m.RuntimeConfig.ProxyManager = proxy.GetProxyManager()
		for _, exposedPort := range m.RuntimeConfig.ExposedPorts {
			m.RuntimeConfig.ProxyManager.Add(exposedPort.Src, dstIP+":"+exposedPort.DstPort, exposedPort.Protocol)
		}
		errs := m.RuntimeConfig.ProxyManager.StartAll()
		if len(errs) > 0 {
			for _, err := range errs {
				logger.WithError(err).Error("failed to start proxy")
			}
		}
	}
	go m.monitor()
}

func (m *Base) monitor() {
	defer m.poststopHook()
	if m.IsDocker {
		okC, errC := cli.ContainerWait(context.Background(), m.ID, container.WaitConditionNextExit)
		select {
		case <-okC:
			logger.Debugf("%s container exit", m.ID)
		case err := <-errC:
			logger.WithError(err).Errorf("%s container exit with error", m.ID)
		}
	} else {
		for {
			time.Sleep(1 * time.Second)
			state, err := getVMState(m.ID)
			if err != nil {
				logger.WithError(err).Errorf("the vm id is %s maybe exit", m.ID)
				break
			} else {
				if state == "shutdown" || state == "shutoff" || state == "crashed" {
					logger.Debugf("the vm id is %s exit with state %s", m.ID, state)
					break
				}
			}
		}
	}
}

func (m *Base) poststopHook() {
	logger.Debugf("%s poststop hook ", m.ID)
	if m.RuntimeConfig.ProxyManager != nil {
		errs := m.RuntimeConfig.ProxyManager.StopAll()
		if len(errs) > 0 {
			for _, err := range errs {
				logger.WithError(err).Error("failed to stop proxy")
			}
		}
	}
}

func (m *Base) Kill(signal string) error {
	if m.IsDocker {
		return KillContainer(m.ID, signal)
	} else {
		return KillVM(m.ID, signal)
	}
}

func (m *Base) Pause() error {
	if m.IsDocker {
		return PauseContainer(m.ID)
	} else {
		return PauseVM(m.ID)
	}
}

func (m *Base) Unpause() error {
	if m.IsDocker {
		return UnpauseContainer(m.ID)
	} else {
		return UnpauseVM(m.ID)
	}
}

func (m *Base) Delete() error {
	if m.IsDocker {
		if err := DeleteContainer(m.ID); err != nil {
			return err
		}
	} else {
		if err := DeleteVM(m.ID); err != nil {
			logger.WithError(err).Error("failed to delete VM ", m.ID)
			return err
		}
		if m.ImageType == "disk"{
			err := os.Remove(m.ImagePath)
			if err != nil {
				return err
			}
		}
	}

	// release ip
	if m.RuntimeConfig.Networks != nil && len(m.RuntimeConfig.Networks) > 0 {
		for _, interf := range m.RuntimeConfig.Networks {
			for _, addr := range interf.Address {
				tmp := strings.Split(addr, "/")
				if len(tmp) > 0 {
					addr = tmp[0]
				}
				err := network.ReleaseIP(interf.Bridge, m.Name, net.ParseIP(addr))
				if err != nil {
					logger.WithError(err).Errorf("failed to release ", addr)
				}
			}
		}
	}

	db.Delete(m.ID)
	db.save(false, "")
	logger.Debugf("%s has removed sussessfully", m.ID)
	return nil
}

func (m *Base) Stop(timeout int32) error {
	if timeout < 1 {
		return fmt.Errorf("timeout must bigger than 0")
	}
	if m.IsDocker {
		return StopContainer(timeout, m.ID)
	} else {
		return StopVM(timeout, m.ID)
	}
}

func (m *Base) Restart(timeout int32) error {
	if m.IsDocker {
		return RestartContainer(timeout, m.ID)
	} else {
		return RestartVM(timeout, m.ID)
	}
}

func (m *Base) GetImageID() string {
	return m.ImageID
}

func (container *Base) configNetwork() error {
	//if !container.IsDocker {
	//	return nil
	//}
	//if container.RuntimeSetting.Networks == nil || len(container.RuntimeSetting.Networks) == 0 {
	//	return nil
	//}
	//// 将每个 container 连上网桥
	//// 并把另外一端放到 container 里面
	//pid, err:= 	netns.GetFromDocker(container.ID)
	//if err != nil {
	//	return fmt.Errorf("failed to get pid to set network of container : %v", err)
	//}
	//for _, nw := range container.RuntimeSetting.Networks {
	//	if err := nw.connectBridge(); err != nil {
	//		logger.WithError(err).Error(nw.HostInterfaceName + " failed to connect bridge")
	//		return err
	//	}
	//
	//	if err := nw.setIn(strconv.Itoa(int(pid))); err != nil {
	//		logger.WithError(err).Error("failed to set veth in container")
	//		return err
	//	}
	//
	//}
	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()
	//
	//// Save the current network namespace
	//origns, _ := netns.Get()
	//defer origns.Close()
	//
	//err = netns.Set(pid)
	//if err != nil {
	//	logger.WithError(err).Errorf("failed to change net namespace")
	//	return err
	//}
	//defer netns.Set(origns)

	// TODO set route for container

	return nil
}

func (m *Base) Rename(newName string) error {
	var err error
	if m.IsDocker {
		err = renameContainer(m.ID, newName)
	} else {
		err = renameVM(m.ID, newName)
	}
	if err != nil {
		logger.WithError(err).Error("failed to rename machine")
		return err
	}
	m.Name = newName
	db.save(true, m.ID)
	logger.Debugf("rename machine %s to %s", m.ID, m.Name)
	return nil
}

func (m *Base) ResizeTTY(h uint32, w uint32) error {
	if m.IsDocker {
		return resizeTTY(m.ID, h, w)
	} else{
		return nil
	}
}

func (m *Base) GetStdio(detachKey string) (chan []byte, chan []byte, chan []byte, error) {
	if m.IsDocker {
		return getContainerStdio(m.ID, detachKey,m.RuntimeConfig.Tty)
	} else {
		return getVMStdio(m.ID)
	}
}

func (m *Base) GetState() string {
	var state string
	var err error
	if m.IsDocker{
		state, err = getContainerState(m.ID)
	} else {
		state, err = getVMState(m.ID)
	}
	if err != nil {
		return "unknown"
	} else {
		return state
	}
}

func (m *Base) Inspect() (name string, runtimeSetting string, spec string, machineType string, err error) {
	var data []byte
	var container = types.ContainerJSON{}
	var vm = libvirtxml.Domain{}
	if m.IsDocker {
		data, err = getContainerInfo(m.ID)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &container)
		if err != nil {
			return
		}
		data, err = json.MarshalIndent(container, "", "\t")
	} else {
		data, err = getVMInfo(m.ID)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &vm)
		if err != nil {
			return
		}
		data, err = json.MarshalIndent(vm, "", "\t")
	}
	if err != nil {
		return
	}
	spec = string(data)
	data, err = json.MarshalIndent(m.RuntimeConfig, "", "\t")
	if err != nil {
		return
	}
	runtimeSetting = string(data)
	name = m.Name
	if m.IsDocker {
		machineType = "container"
	} else {
		machineType = "vm"
	}
	return
}

func (m *Base) Commit(name string) (string, error) {
	var id string
	var err error
	if m.ImageType == "iso" {
		logger.Errorf("can't commit machine base on iso image")
		return "", errors.New("can't commit machine base on iso image")
	} else if m.ImageType == "disk" {
		// vm-disk file
		id, err = image.QEMUImageSave(name, "disk", m.ImagePath)
		if err != nil {
			logger.WithError(err).Errorf("failed to commit disk machine")
			return "", err
		}
	} else if m.ImageType == "docker" {
		_, err := cli.ContainerCommit(context.Background(), m.ID, types.ContainerCommitOptions{
			Reference: name + ":latest",
		})
		if err != nil {
			logger.WithError(err).Errorf("failed to commit docker machine")
			return "", err
		}
		id, err = image.GetImageFromDocker(name)
		if err != nil {
			logger.WithError(err).Errorf("failed to commit docker machine")
			return "", err
		}
	}
	logger.Debugf("the image which id is %s has been committed", id)
	return id, nil
}

func (m *Base) ConnectNetWork(nw *Network) error {
	if m.IsDocker {
		return ContainerConnectNetwork(m.ID, nw)
	} else {
		return VMConnectNetwork(m.ID, nw)
	}
}

func (m *Base) GetName() string {
	return m.Name
}
