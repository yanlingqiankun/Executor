package machine

import (
	"encoding/json"
	"fmt"
	"github.com/vishvananda/netlink"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/machine/types"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

func (container *BaseContainer) SetNetworks(networks []*types.Network) {
	for idx, nw := range networks {
		if nw.Name == "" {
			nw.Name = "eth" + strconv.Itoa(idx)
		}

		if nw.HostInterfaceName == "" {
			nw.HostInterfaceName = "veth" + container.ID[:6] + strconv.Itoa(idx)
		}

	}
	container.RuntimeConfig.Networks = networks

	// 添加默认路由
	if len(networks) > 0 {
		container.SetRoutes([]*Route{
			{
				Gateway:       networks[0].Gateway,
				InterfaceName: networks[0].Name,
			},
		})
	}
}

func (container *BaseContainer) SetRoutes(routes []*Route) {
	container.RuntimeConfig.Routes = routes
}

func (nw *Network) connectBridge() error {
	bridgeName := nw.Bridge
	// 通过接口名获取到 Linux Bridge 接口的对象和接口属性
	br, err := netlink.LinkByName(bridgeName)
	if err != nil {
		return err
	}
	// 创建 veth 接口的配置
	linkAttrs := netlink.NewLinkAttrs()
	linkAttrs.Name = nw.HostInterfaceName
	linkAttrs.MasterIndex = br.Attrs().Index

	// 配置容器内部的接口
	vethPair := netlink.Veth{
		LinkAttrs: linkAttrs,
		PeerName:  "isl-" + nw.HostInterfaceName,
	}

	// 调用 net link LinkAdd 方法创建出这个 veth 接口对
	// 因为上面指定了 link Master Index 是网络对应的 Linux Bridge
	// 所以外面那一端就己经挂载到了网络对应的 Linux Bridge 上了
	if err = netlink.LinkAdd(&vethPair); err != nil {
		return err
	}

	// 调用 net link LinkSetUp 方法，设置 veth 启动
	// 相当于 ip link set xxx up 命令
	if err = netlink.LinkSetUp(&vethPair); err != nil {
		return err
	}
	return nil
}

func (nw *Network) setIn(pid string) error {
	device, err := netlink.LinkByName("isl-" + nw.HostInterfaceName)
	f, err := os.OpenFile(fmt.Sprintf("/proc/%s/ns/net", pid), os.O_RDONLY, 0)
	defer f.Close()
	if err != nil {
		return err
	}

	nsFD := f.Fd()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// 修改veth peer 另外一端移到容器的namespace中
	if err = netlink.LinkSetNsFd(device, int(nsFD)); err != nil {
		return err
	}

	return nil
}

func (container *BaseContainer) configNetwork() error {
	if container.RuntimeConfig.Networks == nil || len(container.RuntimeConfig.Networks) == 0 {
		return nil
	}
	// 将每个 container 连上网桥
	// 并把另外一端放到 container 里面
	for _, nw := range container.RuntimeConfig.Networks {
		if err := nw.connectBridge(); err != nil {
			logger.WithError(err).Error(nw.HostInterfaceName + " failed to connect bridge")
			return err
		}

		if err := nw.setIn(strconv.Itoa(container.process.Pid)); err != nil {
			logger.WithError(err).Error("failed to set veth in container")
			return err
		}

	}

	config := networkConfig{
		Pid:      strconv.Itoa(container.process.Pid),
		Networks: container.RuntimeSetting.Networks,
		Routes:   container.RuntimeSetting.Routes,
	}
	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		logger.WithError(err).Error("failed to generate container network config data")
		return err
	}

	configPath := filepath.Join(filepath.Dir(container.Rootfs), "network.json")

	if err = ioutil.WriteFile(configPath, data, 0600); err != nil {
		logger.WithError(err).Error("failed to write network.json for islandNet")
		return err
	}

	islandNet := exec.Command(conf.GetString("IslandNet"), configPath, conf.GetString("LogLevel"))
	islandNet.Stdout = os.Stdout
	islandNet.Stderr = os.Stderr

	if err = islandNet.Run(); err != nil {
		logger.WithError(err).Error("failed to run ", islandNet.String())
		return err
	}

	return nil

}
