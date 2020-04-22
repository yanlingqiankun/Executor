package daemon

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yanlingqiankun/Executor/logging"
	"github.com/yanlingqiankun/Executor/pb"
	"github.com/yanlingqiankun/Executor/volume"
)

func (s *server) ListVolume(ctx context.Context, _ *empty.Empty) (*pb.ListVolumeResp, error) {
	logger.Debug("list local volumes")
	volumes := make([]*pb.Volume, 0)
	for _, val := range volume.List() {
		volumes = append(volumes, &pb.Volume{
			Name: val.Name,
			Path: val.Path,
			CreateTime: val.CreateTime.Format(TIME_LAYOUT),
		})
	}
	return &pb.ListVolumeResp{Volumes: volumes}, nil
}

func (s *server) AddVolume(ctx context.Context, req *pb.AddVolumeReq) (*pb.AddVolumeResp, error) {
	logger.WithFields(logging.Fields{
		"name": req.Name,
		"path": req.Path,
	}).Debug("add a volume")

	if v, err := volume.Add(req.Path, req.Name); err != nil {
		return &pb.AddVolumeResp{Error: newErr(3, err)}, err
	} else {
		return &pb.AddVolumeResp{Volume: v.Name}, nil
	}
}

func (s *server) CreateVolume(ctx context.Context, req *pb.CreateVolumeReq) (*pb.CreateVolumeResp, error) {
	logger.WithFields(logging.Fields{
		"name": req.Name,
	}).Debug("create a volume")

	if v, err := volume.Create(req.Name); err != nil {
		return &pb.CreateVolumeResp{Error: newErr(4, err)}, err
	} else {
		return &pb.CreateVolumeResp{Volume: &pb.Volume{
			Name: v.Name,
			Path: v.Path,
		}}, nil
	}
}

func (s *server) DeleteVolume(ctx context.Context, req *pb.DeleteVolumeReq) (*pb.Error, error) {
	if err := volume.Delete(req.Name); err != nil {
		return newErr(5, err), err
	} else {
		return newErr(0, err), nil
	}
}
