package machine

import (
	"fmt"
	"github.com/docker/docker/pkg/stringid"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
	"net"
	"os"
	"runtime"
	"strconv"
)

func (container *BaseContainer) SetNetworks(networks []*Network) {
	for idx, nw := range networks {
		if nw.Name == "" {
			nw.Name = "eth" + strconv.Itoa(idx)
		}

		if nw.HostInterfaceName == "" {
			nw.HostInterfaceName = "veth" + stringid.GenerateRandomID()[:6] + strconv.Itoa(idx)
		}

	}
	container.Base.RuntimeConfig.Networks = networks

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
	container.Base.RuntimeConfig.Routes = routes
}

func (nw *Network) connectBridge(nns netns.NsHandle, index int) error {
	bridgeName := nw.Bridge
	// 通过接口名获取到 Linux Bridge 接口的对象和接口属性
	br, err := netlink.LinkByName(bridgeName)
	if err != nil {
		return err
	}
	// 创建 veth 接口的配置
	linkAttrs := netlink.NewLinkAttrs()
	if nw.HostInterfaceName == "" {
		nw.HostInterfaceName = "veth" + stringid.GenerateRandomID()[:6] + strconv.Itoa(index)
	}
	linkAttrs.Name = nw.HostInterfaceName
	linkAttrs.MasterIndex = br.Attrs().Index

	// 配置容器内部的接口
	vethPair := netlink.Veth{
		LinkAttrs: linkAttrs,
		PeerName:  "exl-" + strconv.Itoa(index),
	}

	newName := "eth" + strconv.Itoa(index)

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

	device, err := netlink.LinkByName("exl-" + strconv.Itoa(index))
	if err = netlink.LinkSetNsFd(device, int(nns)); err != nil {
		return err
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Save the current network namespace
	origns, _ := netns.Get()
	defer origns.Close()

	err = netns.Set(nns)
	if err != nil {
		return err
	}
	defer netns.Set(origns)


	// new name
	if err = netlink.LinkSetName(device, newName); err != nil {
		logger.WithError(err).Errorf("failed to set interface to %s", newName)
		return err
	}

	//get newname device
	device, err = netlink.LinkByName(newName)
	if err != nil {
		logger.WithError(err).Error("failed to get interface ", newName)
		return err
	}

	if err = netlink.LinkSetUp(device); err != nil {
		logger.WithError(err).Errorf("failed to set %s up", newName)
		return err
	}

	for i, ip := range nw.Address {
		var label string
		if i == 0 {
			// main nic
			label = ""
		} else {
			// sub nic
			label = ":" + strconv.Itoa(i)
		}
		ipNet, err := netlink.ParseIPNet(ip)
		if err != nil {
			logger.WithError(err).Error("failed to set container ip")
		}
		addr := &netlink.Addr{
			IPNet: ipNet,
			Label: newName + label,
		}
		err = netlink.AddrAdd(device, addr)
		if err != nil {
			logger.WithError(err).Error("failed to set container ip")
		}
	}

	if index == 0 {
		// set default route
		_, ipnet, err := net.ParseCIDR("0.0.0.0/0")
		if err != nil {
			logger.WithError(err).Error("failed to set default route")
			return err
		}
		gw := net.ParseIP(nw.Gateway)
		if gw == nil {
			logger.WithError(err).Error("failed to set gateway")
			return fmt.Errorf("failed to set gateway")
		}
		route := netlink.Route{
			LinkIndex: device.Attrs().Index,
			Dst: &net.IPNet{
				IP:   ipnet.IP,
				Mask: ipnet.Mask,
			},
			Src: nil,
			Gw:gw,
		}
		if err := netlink.RouteAdd(&route); err != nil {
			logger.WithError(err).Error("failed to set default route")
			return err
		}
	}
	return nil
}

func (nw *Network) setIn(pid string) error {
	device, err := netlink.LinkByName("exl-" + nw.HostInterfaceName)
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
