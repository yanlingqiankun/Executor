package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetMachineUnpauseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "unpause id",
		Short: "Unpause all processes within one or more machines",
		Args:  cobra.MinimumNArgs(1),
		Run:   machineUnpauseHandle,
	}
}

func machineUnpauseHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
	} else {
		for _, i := range args {
			UnpauseMachine(i)
		}
	}
}

func UnpauseMachine(id string) {

	id = CheckNameOrId(id)

	r, err := connection.Client.UnpauseMachine(context.Background(), &pb.MachineIdReq{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Cannot unpause the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id, "unpaused")
		}
	}
}
