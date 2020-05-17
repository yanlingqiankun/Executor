package daemon

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yanlingqiankun/Executor/machine"
	"github.com/yanlingqiankun/Executor/network"
	"github.com/yanlingqiankun/Executor/pb"
	"net"
	"strings"
)

func (s server) CreateNetwork(ctx context.Context, req *pb.NetworkCreateReq) (*pb.NetworkCreateResp, error) {
	if err := createNetwork(req.Name, req.Subnet, req.Gateway, req.Isolated); err != nil {
		logger.WithError(err).Error("failed to create network")
		return &pb.NetworkCreateResp{
			Id:                   "",
			Error:                &pb.Error{
				Code: 				1,
				Message:          "failed to create network with error :" + err.Error(),
			},
		}, err
	}
	return &pb.NetworkCreateResp{
		Id:                   req.Name,
		Error:                nil,
	}, nil
}

func createNetwork (name, subnet, gateway string, isolated bool) error {
	return network.CreateNetwork(name, subnet, gateway, isolated)
}

func (s server) DeleteNetwork(ctx context.Context, req *pb.NetworkDeleteReq) (*pb.NetworkDeleteResp, error) {
	if err := deleteNetwork(req.Name, req.Force); err != nil {
		logger.WithError(err).Error("failed to delete network")
		return &pb.NetworkDeleteResp{
			Error:                &pb.Error{Code : 1,Message:"failed to create network with error :" + err.Error()},
		}, err
	} else {
		return &pb.NetworkDeleteResp{Error:&pb.Error{Code:0}}, nil
	}
}

func deleteNetwork(name string, force bool) error {
	return network.DeleteNetwork(name, force)
}

func (s server) InspectNetwork(ctx context.Context, req *pb.NetworkInspectReq) (*pb.NetworkInspectResp, error) {
	str, err := inspectNetwork(req.Name)
	if err != nil {
		logger.WithError(err).Error("failed to inspect network")
		return &pb.NetworkInspectResp{
			Error:                &pb.Error{
				Code:                 1,
				Message:              err.Error(),
			},
			NetInfo:              "",
		}, err
	} else {
		return &pb.NetworkInspectResp{
			Error:                &pb.Error{Code:0},
			NetInfo:              str,
		}, nil
	}
}

func inspectNetwork (name string) (string, error) {
	return network.InspectNetwork(name)
}

func (s server) ListNetwork(context.Context, *empty.Empty) (*pb.NetworkListResp, error) {
	return &pb.NetworkListResp{
		Networks:             ListNetwork(),
		Error:                &pb.Error{Code:0},
	}, nil
}

func ListNetwork () []*pb.NetworkInfo {
	return network.ListNetwork()
}

func (s server) ConnectNetwork(ctx context.Context, req *pb.NetworkConnectReq) (*pb.Error, error) {
	err := connectNetwork(req.Name, req.MachineId, req.Interface)
	if err != nil {
		logger.WithError(err).Errorf("failed to connect network")
		return newErr(1, err), err
	} else {
		logger.Debugf("connect %s to %s successfully", req.MachineId, req.Name)
		return newErr(0, err), err
	}
}

func connectNetwork (netName string, machineID string, i *pb.NetworkInterface) error {
	m, err := machine.GetMachine(machineID)
	if err != nil {
		logger.WithError(err).Errorf("failed to get get machine %s", machineID)
		return err
	}

	prefix, err := network.GetPrefix(i.Bridge)
	if err != nil {
		return err
	}
	gateWay, err := network.GetGateWay(i.Bridge)
	if err != nil {
		return err
	}
	address := make([]string, len(i.Address))
	if i.Mac == "" {
		i.Mac = getMac()
	}
	if len(address) == 0 {
		addr, err := network.AllocateIP(i.Bridge, m.GetName(), i.Mac)
		if err != nil {
			logger.WithError(err).Error("failed to get ip")
			return err
		}
		address = append(address, fmt.Sprintf("%s/%d", addr.String(), prefix))
	} else {
		for index, addr := range i.Address {
			err := network.RegisterIP(i.Bridge, m.GetName(), net.ParseIP(addr.Ip), i.Mac)
			if err != nil {
				return err
			}
			address[index] = fmt.Sprintf("%s/%d", addr.Ip, prefix)
		}
	}

	machineInterface := &machine.Network{
		Name:       i.Name,
		Bridge:     i.Bridge,
		MacAddress: i.Mac,
		Address:    address,
		Gateway:    gateWay,
	}
	err = m.ConnectNetWork(machineInterface)
	if err != nil {
		logger.WithError(err).Errorf("failed to connect %s to %s", m.GetName(), i.Name)
		// unRegister ip when connect failed
		for _, addr := range machineInterface.Address {
			tmp := strings.Split(addr, "/")
			if len(tmp) > 0 {
				addr = tmp[0]
			}
			err := network.ReleaseIP(machineInterface.Bridge, m.GetName(), net.ParseIP(addr))
			if err != nil {
				logger.WithError(err).Errorf("failed to release ", addr)
			}
		}
		return err
	}
	return nil
}