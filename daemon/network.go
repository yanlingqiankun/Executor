package daemon

import (
	"context"
	"github.com/yanlingqiankun/Executor/network"
	"github.com/yanlingqiankun/Executor/pb"
)

func (s server) CreateNetwork(ctx context.Context, req *pb.NetworkCreateReq) (*pb.NetworkCreateResp, error) {
	if err := createNetwork(req.Name, req.Subnet, req.Gateway); err != nil {
		return &pb.NetworkCreateResp{
			Id:                   "",
			Error:                &pb.Error{
				Message:              "failed to create network with error :" + err.Error(),
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