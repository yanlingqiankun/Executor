package daemon

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yanlingqiankun/Executor/image"
	"github.com/yanlingqiankun/Executor/pb"
)

const TIME_LAYOUT = "2006-01-02 15:04:05.999999999 -0700 MST"

func (s server) ImportImage(ctx context.Context, req *pb.ImportImageReq) (*pb.ImportImageResp, error) {
	return importImage(req)
}

func importImage (req *pb.ImportImageReq) (*pb.ImportImageResp, error) {
	var id string
	var err error
	if req.Type == "vm-iso" || req.Type == "docker-pull" || req.Type == "docker-repo" || req.Type == "vm-disk"{
		if req.Type == "vm-iso" {
			id, err = image.QEMUImageSave(req.Name, "iso", req.Path)
		} else if req.Type == "vm-disk" {
			id, err = image.QEMUImageSave(req.Name,"disk", req.Path)
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

func (s server) ListImage(context.Context, *empty.Empty) (*pb.ListImageResp, error) {
	return listImage()
}

func listImage() (*pb.ListImageResp, error) {
	var result = &pb.ListImageResp{
		Images: nil,
		Err:    nil,
	}
	images := image.ListImage()
	for _, key := range images {
		result.Images = append(result.Images, &pb.Image{
			Id:                   key.ID,
			CreateTime:           key.CreateTime.Format(TIME_LAYOUT),
			Name:                 key.Name,
			Type:                 key.Type,
			Machines:             key.Counter,
		})
	}
	return result, nil
}

func (s server) DeleteImage(ctx context.Context, req *pb.DeleteImageReq) (*pb.DeleteImageResp, error) {
	err := deleteImage(req.Id)
	if err != nil {
		logger.WithError(err).Errorf("failed to delete image %s ", req.Id)
		return &pb.DeleteImageResp{
			Err:                  &pb.Error{
				Code:                 1,
				Message:              err.Error(),
			},
		}, err
	} else {
		logger.Debugf("image %s has been deleted", req.Id)
		return &pb.DeleteImageResp{
			Err:                  &pb.Error{
				Code:                 0,
				Message:              "",
			},
		}, nil
	}
}

func deleteImage(id string) error {
	img, err := image.OpenImage(id)
	if err != nil {
		return err
	}
	return img.Remove()
}