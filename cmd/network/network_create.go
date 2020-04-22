package network

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/pb"
)

var subnet string
var gateway string
var isolated bool

func GetNetworkCreateCmd() *cobra.Command {
	networkCreateCmd := &cobra.Command{
		Use:   "create name --subnet [--gateway -i]",
		Short: "create a network",
		Long:  `create a network for your machine`,
		Args:  cobra.MinimumNArgs(1),
		Run:   networkCreateHandle,
	}

	networkCreateCmd.Flags().StringVar(&subnet, "subnet", "", "subnet of network")
	networkCreateCmd.Flags().StringVar(&gateway, "gateway", "", "gateway of network")
	networkCreateCmd.Flags().BoolVarP(&isolated, "isolated", "i", false, "create a isolated network")

	_ = networkCreateCmd.MarkFlagRequired("subnet")

	return networkCreateCmd
}

func networkCreateHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
		return
	}

	resp, err := connection.Client.CreateNetwork(context.Background(),&pb.NetworkCreateReq{
		Name:                 args[0],
		Subnet:               subnet,
		Gateway:              gateway,
		Isolated:            isolated,
	})
	if err != nil {
		fmt.Println("failed to create network : ", err.Error())
		return
	}

	if resp != nil {
		fmt.Printf("The network %s was created success\n", args[0])
	}
}
