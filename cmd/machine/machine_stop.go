package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

var time_out = 0

func GetMachineStopCmd() *cobra.Command {

	containerStopCmd := &cobra.Command{
		Use:   "stop id [options]",
		Short: "stop the machine",
		Long:  `stop specific machine in your computer`,
		Args:  cobra.MinimumNArgs(1),
		Run:   machineStopHandle,
	}
	containerStopCmd.Flags().IntVar(&time_out, "time", 10, "Seconds to wait for stop before killing it, default is 10")
	return containerStopCmd
}

func machineStopHandle(cmd *cobra.Command, args []string) {
	id := CheckNameOrId(args[0])
	r, err := connection.Client.StopMachine(context.Background(), &pb.StopMachineReq{
		Id:      id,
		Timeout: int32(time_out),
	})

	if err != nil {
		fmt.Printf("Cannot start the container for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id)
		}
	}
}

