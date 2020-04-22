package volume

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetVolumeAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [name] [path]",
		Short: "add a volume",
		Long:  `add a volume in your Islands`,
		Args:  cobra.MinimumNArgs(1),
		Run:   volumeAddHandle,
	}
}

func volumeAddHandle(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		_ = cmd.Help()
	} else {
		addVolume(args[0], args[1])
	}
}

func addVolume(name string, path string) {
	r, err := connection.Client.AddVolume(context.Background(), &pb.AddVolumeReq{Name: name, Path: path})
	if err != nil {
		fmt.Printf("Cannot add volume for reason: %v\n", err)
	} else {
		if !utils.PrintError(r.Error) {
			fmt.Println(name)
		}
	}
}
