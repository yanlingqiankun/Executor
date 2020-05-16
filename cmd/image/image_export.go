package image

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

var target string

func GetImageExportCmd() *cobra.Command {
	imageExportCmd := &cobra.Command{
		Use:   "export id -o [target file]",
		Short: "export the image",
		Long:  `export specific image in your repo`,
		Args:  cobra.MinimumNArgs(1),
		Run:   imageExportHandle,
	}
	imageExportCmd.Flags().StringVarP(&target, "output", "o", "", "output file")
	return imageExportCmd
}

func imageExportHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 || target == ""{
		_ = cmd.Help()
	} else {
		for _, i := range args {
			exportImage(i)
		}
	}
}

func exportImage(id string) {
	id = CheckNameOrId(id)
	r, err := connection.Client.ExportImage(context.Background(), &pb.ExportImageReq{Id: id, Target:target})
	if err != nil {
		fmt.Printf("Cannot remove the image for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id, "exported")
		}
	}
}


