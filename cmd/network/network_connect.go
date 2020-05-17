package network

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/machine"
	"github.com/yanlingqiankun/Executor/pb"
)
var (
	mac           string
	machineName   string
	ipAddr        string
	interfaceName string
)
func GetNetworkConnectCmd() *cobra.Command {
	networkConnectCmd := &cobra.Command{
		Use:   "connect [flags] [network]",
		Short: "connect to the network",
		Long:  `connect machine to to specific network`,
		Args:  cobra.MinimumNArgs(1),
		Run:   networkConnectHandle,
	}
	networkConnectCmd.Flags().StringVar(&mac, "mac", "", "set mac for your new interface")
	networkConnectCmd.Flags().StringVar(&ipAddr, "ip", "", "set ip and mask for your interface")
	networkConnectCmd.Flags().StringVar(&machineName, "machine", "", "the name or id of machine which you want to connect to the network")
	networkConnectCmd.Flags().StringVar(&interfaceName, "dev", "", "set name for your new interface")
	return networkConnectCmd
}

func networkConnectHandle(cmd *cobra.Command, args []string) {
	if machineName == "" {
		_ = cmd.Help()
	} else {
		machineName := machine.CheckNameOrId(machineName)
		_, err := connection.Client.ConnectNetwork(context.Background(), &pb.NetworkConnectReq{
			Interface:            &pb.NetworkInterface{
				Name:                 interfaceName,
				Bridge:               args[0],
				Mac:                  mac,
				Gateway:              "",
			},
			Name:                 args[0],
			MachineId:            machineName,
		})
		if err != nil {
			fmt.Println("error to connect ", err.Error())
			return
		} else {
			fmt.Println("connect successfully , reboot machine to take effect")
		}
	}
}

