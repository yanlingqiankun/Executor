package image

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
	"path/filepath"
)

var imageType string
var name string

func GetImageImportCmd() *cobra.Command {
	imageImportCmd := &cobra.Command{
		Use:   "import [OPTIONS] path",
		Short: "import your image",
		Long:  `import the image into your repo by specify type`,
		Args:  cobra.MinimumNArgs(0),
		Run:   importHandle,
	}
	imageImportCmd.Flags().StringVar(&imageType, "type", "", "image type : vm-iso vm-disk docker-pull docker-repo docker-import docker-save raw")
	imageImportCmd.Flags().StringVar(&name, "name", "", "image name in repo")
	return imageImportCmd
}

func importHandle(cmd *cobra.Command, args []string) {
	// 检查路径是否存在
	if imageType == "" {
		_ = cmd.Help()
		return
	}

	if name == "" {
		if imageType != "docker-save" && imageType != "raw" {
			_ = cmd.Help()
			return
		}
	}

	if len(args) == 0 {
		if imageType != "docker-pull" && imageType != "docker-repo" {
			_ = cmd.Help()
			return
		}
	} else {
		if !utils.PathExist(args[0]) {
			fmt.Println("path doesn't exit")
			return
		}
	}

	var path string
	if len(args) == 1 {
		path = args[0]
	}

	var r *pb.ImportImageResp
	var err error

	dir, file := filepath.Split(path)
	absDir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println(err.Error())
	}
	path = filepath.Join(absDir, file)

	r, err = connection.Client.ImportImage(context.Background(), &pb.ImportImageReq{
		Path: path,
		Name: name,
		Type: imageType,
	})
	if err != nil {
		fmt.Printf("Cannot import for reason %v\n", err)
	} else {
		if !utils.PrintError(r.Err) {
			fmt.Println(r.Id)
		}
	}
}
