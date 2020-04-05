package machine

import (
	"fmt"
	"github.com/docker/docker/client"
	"github.com/libvirt/libvirt-go"
	"github.com/vishvananda/netns"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"path/filepath"
	"runtime"
	"strconv"
)

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
		logger.Fatalf("failed to connect to qemu")
	}
	db.init()
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
	db.Store(baseEntry.ID, baseEntry)
	return baseEntry.ID, nil
}

func (m *Base) Start() error {
	if m.IsDocker {
		return StartContainer(m.ID)
	} else {
		return StartVM(m.ID)
	}
}

func (m *Base) Kill(string) error {
	if m.IsDocker {
		return KillContainer(m.ID, "SIGKILL")
	} else {
		return StartVM(m.ID)
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
		return DeleteContainer(m.ID)
	} else {
		return DeleteVM(m.ID)
	}
}

func (m *Base) Stop(timeout int) error {
	if m.IsDocker {
		return StopContainer(timeout, m.ID)
	} else {
		return StartVM(m.ID)
	}
}

func (m *Base) Restart(timeout int) error {
	if m.IsDocker {
		return RestartContainer(timeout, m.ID)
	} else {
		return StartVM(m.ID)
	}
}

func (container *Base) configNetwork() error {
	if !container.IsDocker {
		return nil
	}
	if container.RuntimeSetting.Networks == nil || len(container.RuntimeSetting.Networks) == 0 {
		return nil
	}
	// 将每个 container 连上网桥
	// 并把另外一端放到 container 里面
	pid, err:= 	netns.GetFromDocker(container.ID)
	if err != nil {
		return fmt.Errorf("failed to get pid to set network of container : %v", err)
	}
	for _, nw := range container.RuntimeSetting.Networks {
		if err := nw.connectBridge(); err != nil {
			logger.WithError(err).Error(nw.HostInterfaceName + " failed to connect bridge")
			return err
		}

		if err := nw.setIn(strconv.Itoa(int(pid))); err != nil {
			logger.WithError(err).Error("failed to set veth in container")
			return err
		}

	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Save the current network namespace
	origns, _ := netns.Get()
	defer origns.Close()

	err = netns.Set(pid)
	if err != nil {
		logger.WithError(err).Errorf("failed to change net namespace")
		return err
	}
	defer netns.Set(origns)

	// TODO set route for container

	return nil
}

func CreateMachine(imageID string, machineType string) (Factory, error) {
	if machineType == "container" {
		return CreateContainer(imageID), nil
	} else if machineType == "vm" {
		return CreateVM(imageID), nil
	} else {
		logger.Errorf("Error machine type")
		return nil, fmt.Errorf("Error machine type")
	}
}