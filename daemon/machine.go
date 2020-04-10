package daemon

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
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
	logger.Debug("create machine req ", req)
	//if err := stringid.ValidateID(req.ImageId); err != nil {
	//	return "", err
	//}
	img, err := image.OpenImage(req.ImageId)
	if err != nil {
		return "", err
	}

	img.Register()
	var factory machine.Factory
	if isDocker, _ := img.GetType(); isDocker{
		factory = machine.CreateContainer(req.ImageId)
	} else {
		factory = machine.CreateVM(req.ImageId)
	}

	if req.Name == ""{
		req.Name = stringid.GenerateRandomID()[:12]
	}
	if err := factory.SetName(req.Name); err != nil {
		return "", err
	}
	//factory.SetImage(image.MountPoint())
	factory.SetImage(req.ImageId, img.GetPath(), img.GetName())
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
		logger.WithError(err).Error("failed to create the machine")
		return "", err
	}
	logger.Debugf("the machine name = %s has create successful", req.Name)
	return machine.AddMachine(factory)
}


func (s server) DeleteMachine(ctx context.Context, req *pb.DeleteMachineReq) (*pb.Error, error) {
	err := deleteMachine(req.Id)
	if err != nil {
		logger.WithError(err).Error("failed to delete machine ", req.Id)
		return newErr(1, err), err
	}
	return newErr(0, err), err
}

func deleteMachine(id string) error {
	m, err := machine.GetMachine(id)
	if err != nil {
		return err
	}
	imageId := m.GetImageID()
	if img, err := image.OpenImage(imageId); err != nil{
		logger.WithError(err).Error("failed to open image : ", imageId)
		return fmt.Errorf("failed to open image : %s", imageId)
	} else {
		img.UnRegister()
	}
	err = m.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s server) ListMachine(context.Context, *empty.Empty) (*pb.ListMachineResp, error) {
	return listMachine(), nil
}

func listMachine() *pb.ListMachineResp {
	machines := make([]*pb.MachineInfo, 0)
	for _, m := range machine.ListMachine() {
		machines = append(machines, &pb.MachineInfo{
			Id:                   m.ID,
			Name:                 m.Name,
			ImageName:            m.ImageName,
			ImageType:            m.ImageType,
			CreateTime:           m.CreateTime,
			Status:               m.Status,
			ImageId:              m.ImageId,
		})
	}
	return &pb.ListMachineResp{
		MachineInfos:         machines,
		Err:                  nil,
	}
}


func (s server) StartMachine(ctx context.Context, req *pb.StartMachineReq) (*pb.Error, error) {
	err := startMachine(req.Id)
	if err != nil {
		return newErr(1, err), err
	} else {
		return newErr(0, err), err
	}
}

func startMachine(id string) error {
	m, err := machine.GetMachine(id)
	if err != nil {
		return err
	}
	return m.Start()
}

func (s server) KillMachine(ctx context.Context,req *pb.KillMachineReq) (*pb.Error, error) {
	err := killMachine(req.Id, req.Signal)
	if err != nil {
		return newErr(1, err), err
	} else {
		return newErr(0, err), nil
	}
}

func killMachine(id string, signal string) error {
	m, err := machine.GetMachine(id)
	if err != nil {
		logger.WithError(err).Error("failed to get get machine")
		return err
	}
	err = m.Kill(signal)
	if err != nil {
		logger.WithError(err).Error("failed to kill machine")
		return err
	}
	logger.Debugf("%s has been kill", id)
	return nil
}


func (s server) StopMachine(ctx context.Context, req *pb.StopMachineReq) (*pb.Error, error) {
	err := stopMachine(req.Id, req.Timeout)
	if err != nil {
		return newErr(1, err), err
	} else {
		return newErr(0, err), err
	}
}

func stopMachine(id string, timeout int32) error {
	m, err := machine.GetMachine(id)
	if err != nil {
		logger.WithError(err).Error("failed to get get machine")
		return err
	}
	err = m.Stop(timeout)
	if err != nil {
		logger.WithError(err).Error("failed to stop machine")
		return err
	}
	return nil
}

func (s server) RenameMachine(ctx context.Context, req *pb.RenameMachineReq) (*pb.Error, error) {
	err := renameMachine(req.Id, req.Name)
	if err != nil {
		return newErr(1, err), err
	} else {
		return newErr(0, err), err
	}
}

func renameMachine(id, newName string) error {
	m, err := machine.GetMachine(id)
	if err != nil {
		logger.WithError(err).Error("failed to get get machine")
		return err
	}
	err = m.Rename(newName)
	if err != nil {
		logger.WithError(err).Error("failed rename machine")
		return err
	}
	return nil
}

func (s server) RestartMachine(ctx context.Context, req *pb.RestartMachineReq) (*pb.Error, error) {
	if err := restartMachine(req.Id, req.Timeout); err != nil {
		return newErr(1, err), err
	} else {
		return newErr(0, err), err
	}
}

func restartMachine(id string, timeout int32) error {
	m, err := machine.GetMachine(id)
	if err != nil {
		logger.WithError(err).Error("failed to get get machine")
		return err
	}
	err = m.Restart(timeout)
	if err != nil {
		logger.WithError(err).Error("failed to restart machine")
		return err
	}
	logger.Debugf("machine %s has restarted", id)
	return nil
}

func (s server) InspectMachine(ctx context.Context, req *pb.MachineIdReq) (*pb.InspectMachineResp, error) {
	return inspectmachine(req.Id)
}

func inspectmachine(id string) (*pb.InspectMachineResp, error) {
	m, err := machine.GetMachine(id)
	if err != nil {
		logger.WithError(err).Error("failed to get get machine")
		return nil, err
	}
	name, runtimeSetting, spec, machineType, err := m.Inspect()
	if err != nil {
		return &pb.InspectMachineResp{}, err
	} else {
		return &pb.InspectMachineResp{
			Name:                 name,
			Id:                   id,
			Type:                 machineType,
			Spec:                 spec,
			RuntimeConfig:        runtimeSetting,
		}, nil
	}
}
