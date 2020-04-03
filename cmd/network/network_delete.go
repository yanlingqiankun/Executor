package network

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

var force bool

func GetNetworkDeleteCmd() *cobra.Command {
	deleteCmd :=  &cobra.Command{
		Use:   "delete name  [-f]",
		Short: "delete the network",
		Long:  `delete specific network and drive of the network`,
		Args:  cobra.MinimumNArgs(1),
		Run:   networkDeleteHandle,
	}

	deleteCmd.Flags().BoolVarP(&force, "all", "a", false, "delete network force")
	return deleteCmd
}

func networkDeleteHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
	} else {
		resp, err := connection.Client.DeleteNetwork(context.Background(), &pb.NetworkDeleteReq{
			Name:  args[0],
			Force: force,
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if !utils.PrintError(resp.Error) {
			fmt.Println("The network has been deleted")
		}
	}
}
