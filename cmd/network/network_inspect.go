package network

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetNetworkInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect name",
		Short: "inspect the network",
		Long:  `inspect specific network information`,
		Args:  cobra.MinimumNArgs(1),
		Run:   networkInspectHandle,
	}
}

func networkInspectHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
	} else {
		resp, err := connection.Client.InspectNetwork(context.Background(), &pb.NetworkInspectReq{
			Name:                 args[0],
		})
		if err != nil {
			fmt.Println("error ", err.Error())
			return
		}
		if !utils.PrintError(resp.Error) {
			fmt.Println(resp.NetInfo)
		}
	}
}
