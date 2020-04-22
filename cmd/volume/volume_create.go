package volume

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetVolumeCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create [name]",
		Short: "create a volume",
		Long:  `create a volume in your Executor`,
		Args:  cobra.MinimumNArgs(1),
		Run:   volumeCreateServe,
	}

}

func volumeCreateServe(cmd *cobra.Command, args []string) {
	name := ""
	if len(args) == 1 {
		name = args[0]
	}

	r, err := connection.Client.CreateVolume(context.Background(), &pb.CreateVolumeReq{Name: name})
	if err != nil {
		fmt.Printf("Cannot remove the volume for reason %v\n", err)
	} else {
		if !utils.PrintError(r.Error) {
			fmt.Println(name)
		}
	}
}
