package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetMachinePauseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pause id",
		Short: "Pause all processes within one or more machines",
		Args:  cobra.MinimumNArgs(1),
		Run:   containerPauseHandle,
	}
}

func containerPauseHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
	} else {
		for _, i := range args {
			pauseMachine(i)
		}
	}
}

func pauseMachine(id string) {

	id = CheckNameOrId(id)

	r, err := connection.Client.PauseMachine(context.Background(), &pb.MachineIdReq{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Cannot pause the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id, "paused")
		}
	}
}
