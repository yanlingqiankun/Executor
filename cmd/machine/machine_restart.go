package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

var restartTimeOut = 0

func GetMachineRestartCmd() *cobra.Command {
	machineRestartCmd := &cobra.Command{
		Use:   "restart id",
		Short: "Restart the machine",
		Args:  cobra.MinimumNArgs(1),
		Run:   machineRestartHandle,
	}
	machineRestartCmd.Flags().IntVar(&restartTimeOut, "time", 10, "Seconds to wait for stop before killing the machine, default is 10")
	return machineRestartCmd
}

func machineRestartHandle(cmd *cobra.Command, args []string) {
	id := CheckNameOrId(args[0])
	r, err := connection.Client.RestartMachine(context.Background(), &pb.RestartMachineReq{
		Id:      id,
		Timeout: int32(restartTimeOut),
	})

	if err != nil {
		fmt.Printf("Cannot restart the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id)
		}
	}
}
