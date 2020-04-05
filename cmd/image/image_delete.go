package image

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetImageDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete id",
		Short: "delete the image",
		Long:  `delete specific image in your repo`,
		Args:  cobra.MinimumNArgs(1),
		Run:   imageDeleteHandle,
	}
}

func imageDeleteHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
	} else {
		for _, i := range args {
			deleteImage(i)
		}
	}
}

func deleteImage(id string) {
	id = CheckNameOrId(id)
	r, err := connection.Client.DeleteImage(context.Background(), &pb.DeleteImageReq{Id: id})
	if err != nil {
		fmt.Printf("Cannot remove the image for reason %v\n", err)
	} else {
		if !utils.PrintError(r.Err) {
			fmt.Println(id, "removed")
		}
	}
}

func CheckNameOrId(imageId string) string {
	imageListResp, err := connection.Client.ListImage(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	isNameFlag := false
	isIdFlag := false

	for _, v := range imageListResp.Images {
		if imageId == v.Name {
			isNameFlag = true
			imageId = v.Id
			break
		}

		if len(imageId) == 12 {
			if imageId == v.Id[:12] {
				isIdFlag = true
				imageId = v.Id
				break
			}
		} else {
			if imageId == v.Id {
				isIdFlag = true
				imageId = v.Id
				break
			}
		}
	}

	if !(isNameFlag || isIdFlag) {
		panic("No Image named: " + imageId)
	}
	return imageId
}

