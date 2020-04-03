package daemon

import (
	"context"
	"github.com/yanlingqiankun/Executor/network"
	"github.com/yanlingqiankun/Executor/pb"
)

func (s server) CreateNetwork(ctx context.Context, req *pb.NetworkCreateReq) (*pb.NetworkCreateResp, error) {
	if err := createNetwork(req.Name, req.Subnet, req.Gateway); err != nil {
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

func createNetwork (name, subnet, gateway string) error {
	return network.CreateNetwork(name, subnet, gateway)
}

func (s server) DeleteNetwork(ctx context.Context, req *pb.NetworkDeleteReq) (*pb.NetworkDeleteResp, error) {
	if err := deleteNetwork(req.Name, req.Force); err != nil {
		logger.WithError(err).Error("failed to delete network")
		return &pb.NetworkDeleteResp{
			Error:                &pb.Error{Code : 1,Message:"failed to create network with error :" + err.Error()},
		}, err
	} else {
		return nil, nil
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
			Error:                &pb.Error{},
			NetInfo:              str,
		}, nil
	}
}

func inspectNetwork (name string) (string, error) {
	return network.InspectNetwork(name)
}
