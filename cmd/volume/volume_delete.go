package volume

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

var (
	purge bool
)

func GetVolumeDeleteCmd() *cobra.Command {
	volumeDeleteCmd := &cobra.Command{
		Use:   "delete name",
		Short: "delete the volume",
		Long:  `delete specific volume in your repo`,
		Args:  cobra.MinimumNArgs(1),
		Run:   volumeDeleteServe,
	}
	volumeDeleteCmd.Flags().BoolVarP(&purge, "purge", "p", false, "delete dir when try to delete volume")
	return volumeDeleteCmd
}

func volumeDeleteServe(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
	} else {
		for _, i := range args {
			deleteVolume(i)
		}
	}
}

func deleteVolume(name string) {
	r, err := connection.Client.DeleteVolume(context.Background(), &pb.DeleteVolumeReq{Name:name, Purge:purge})
	if err != nil {
		fmt.Printf("Cannot remove the image for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(name, "removed")
		}
	}
}
