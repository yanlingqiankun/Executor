package daemon

import (
	"context"
	"fmt"
	"github.com/yanlingqiankun/Executor/image"
	"github.com/yanlingqiankun/Executor/machine"
	"github.com/yanlingqiankun/Executor/network/proxy"
	"github.com/yanlingqiankun/Executor/pb"
	"github.com/yanlingqiankun/Executor/stringid"
	"strings"
)

func (s server) CreateMachine(ctx context.Context, req *pb.CreateMachineReq) (*pb.CreateMachineResp, error) {
	id, err := createMachine(req)
	if err != nil {
		return &pb.CreateMachineResp{
			Id:                   "",
			Err:                  newErr(1, err),
		}, err
	} else {
		return &pb.CreateMachineResp{
			Id:                   id,
			Err:                  newErr(0, err),
		}, nil
	}
}

func createMachine (req *pb.CreateMachineReq) (string, error) {
	logger.Debug("create container req ", req)
	if err := stringid.ValidateID(req.ImageId); err != nil {
		return "", err
	}
	img, err := image.OpenImage(req.ImageId)
	if err != nil {
		return "", err
	}

	img.Register()
	factory := machine.CreateVM(req.ImageId)
	if err := factory.SetName(req.Name); err != nil {
		return "", err
	}
	//factory.SetImage(image.MountPoint())

	factory.SetTTY(req.Tty)
	factory.SetCmd(req.Cmd)
	factory.SetEnv(req.Env)
	factory.SetWorkingDir(req.WorkingDir)

	//resources, err := getContainerResources(req.Resources)
	//if err != nil {
	//	return &pb.CreateContainerResp{Err: newErr(1, err)}, err
	//}
	//factory.SetCgroups(island.CgroupsConfig{
	//	CgroupsPath:    req.Resources.CgroupParent,
	//	LinuxResources: resources,
	//})

	//if req.Network != nil {
	//	factory.SetHostname(req.Network.Hostname)
	//	factory.SetHosts(convertHostsFromPB(req.Network.ExtraHosts))
	//	factory.SetDNS(req.Network.Dns)
	//
	//	containerInterface := make([]*island.Network, len(req.Network.Interfaces))
	//	for index, i := range req.Network.Interfaces {
	//		address := make([]string, len(i.Address))
	//		for index, addr := range i.Address {
	//			address[index] = fmt.Sprintf("%s/%d", addr.Ip, addr.Mask)
	//		}
	//		containerInterface[index] = &island.Network{
	//			Name:       i.Name,
	//			Bridge:     i.Bridge,
	//			MacAddress: i.Mac,
	//			Address:    address,
	//			Gateway:    i.Gateway,
	//		}
	//	}
	//	factory.SetNetworks(containerInterface)
	//}

	//if req.Volumes != nil {
	//	volumes := make([]*island.ContainerVolume, len(req.Volumes))
	//	for index, v := range req.Volumes {
	//		flag := island.ReadPermission
	//		if !v.Readonly {
	//			flag |= island.WritePermission
	//		}
	//		volumes[index] = &island.ContainerVolume{
	//			Destination: v.Destination,
	//			RW:          flag,
	//			Source:      v.Source,
	//			Name:        v.Name,
	//		}
	//	}
	//	factory.SetVolumes(volumes)
	//}

	if req.ExposedPorts != nil {
		proxies := make([]proxy.ProxyInfo, 0)
		for dst, portBinds := range req.ExposedPorts {
			// judge dst port[/tcp | /udp]
			tmp := strings.Split(dst, "/")
			var dstPort, protocol string
			allProtocol := false
			if len(tmp) == 1 {
				if !proxy.IsPort(tmp[0]) {
					return "", fmt.Errorf("error exposed-ports")
				}
				allProtocol = true
				dstPort = tmp[0]
			} else if len(tmp) == 2 {
				if !proxy.IsPort(tmp[0]) || !(tmp[1] == "tcp" || tmp[1] == "udp") {
					return "", fmt.Errorf("error exposed-ports")
				}
				dstPort = tmp[0]
				protocol = tmp[1]
			} else {
				return"", fmt.Errorf("error exposed-ports")
			}

			for _, portBind := range portBinds.PortBindings {
				if !proxy.IsPort(portBind.HostPort) || !proxy.IsSrcIP(portBind.HostIp) {
					return "", fmt.Errorf("error exposed-ports")
				}
				if !allProtocol {
					proxies = append(proxies, proxy.ProxyInfo{
						Src:      portBind.HostIp + ":" + portBind.HostPort,
						DstPort:  dstPort,
						Protocol: protocol,
					})
				} else {
					proxies = append(proxies, proxy.ProxyInfo{
						Src:      portBind.HostIp + ":" + portBind.HostPort,
						DstPort:  dstPort,
						Protocol: "udp",
					})
					proxies = append(proxies, proxy.ProxyInfo{
						Src:      portBind.HostIp + ":" + portBind.HostPort,
						DstPort:  dstPort,
						Protocol: "tcp",
					})
				}
			}
		}
		factory.SetExposedPorts(proxies)
	}

	//if container, err := factory.Save(); err != nil {
	//	return &pb.CreateContainerResp{Err: newErr(1, err)}, err
	//} else {
	//	return &pb.CreateContainerResp{Id: container.GetID()}, nil
	//}
	err = factory.Create()
	if err != nil {
		return "", err
	}
	return machine.AddMachine(factory)
}