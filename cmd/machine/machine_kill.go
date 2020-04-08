package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

var signal = ""

func GetMachineKillCmd() *cobra.Command {
	machineKillCmd := &cobra.Command{
		Use:   "kill id",
		Short: "kill the machine",
		Long:  `kill specific machine in your computer`,
		Args:  cobra.MinimumNArgs(1),
		Run:   machineKillHandle,
	}

	machineKillCmd.Flags().StringVarP(&signal, "signal", "s", "KILL", "Signal to send to the machine , default is \"KILL\"")
	return machineKillCmd
}

func machineKillHandle(cmd *cobra.Command, args []string) {

	id := CheckNameOrId(args[0])

	r, err := connection.Client.KillMachine(context.Background(), &pb.KillMachineReq{
		Id:     id,
		Signal: signal,
	})

	if err != nil {
		fmt.Printf("Cannot kill the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id + " killed")
		}
	}
}

