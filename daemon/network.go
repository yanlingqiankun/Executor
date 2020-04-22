package daemon

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yanlingqiankun/Executor/network"
	"github.com/yanlingqiankun/Executor/pb"
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
