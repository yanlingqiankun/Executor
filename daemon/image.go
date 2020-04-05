package daemon

import (
	"context"
	"fmt"
	"github.com/yanlingqiankun/Executor/image"
	"github.com/yanlingqiankun/Executor/pb"
)

func (s server) ImportImage(ctx context.Context, req *pb.ImportImageReq) (*pb.ImportImageResp, error) {
	return importImage(req)
}

func importImage (req *pb.ImportImageReq) (*pb.ImportImageResp, error) {
	var id string
	var err error
	if req.Type == "vm-iso" || req.Type == "docker-pull" || req.Type == "docker-repo" {
		if req.Type == "vm-iso" {
			id, err = image.QEMUImageSave(req.Name, "iso", req.Path)
		} else if req.Type == "docker-pull" {
			id, err = image.PullDockerImage(context.Background(), req.Name)
		} else if req.Type == "docker-repo" {
			id, err = image.GetImageFromDocker(req.Name)
		}
		if err != nil {
			return &pb.ImportImageResp{
				Id:                   "",
				Err:                  &pb.Error{
					Code:                 1,
					Message:              err.Error(),
				},
			},err
		} else {
			return &pb.ImportImageResp{
				Id:                   id,
				Err:                  &pb.Error{
					Code:                 0,
					Message:              "",
				},
			},err
		}
	} else {
		return &pb.ImportImageResp{
			Id:                   "",
			Err:                  &pb.Error{
				Code:                 1,
				Message:              "wrong type of image",
			},
		}, fmt.Errorf("wrong type of image : %s", req.Type)
	}
}