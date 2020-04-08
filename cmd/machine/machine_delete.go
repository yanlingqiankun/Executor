package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetMachineDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete id",
		Short: "delete the machine",
		Long:  `delete specific machine in your computer`,
		Args:  cobra.MinimumNArgs(1),
		Run:   machineDeleteHandle,
	}
}

func machineDeleteHandle(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		_ = cmd.Help()
	} else {
		for _, i := range args {
			removeMachine(i)
		}
	}
}

func removeMachine(id string) {

	id = CheckNameOrId(id)

	r, err := connection.Client.DeleteMachine(context.Background(), &pb.DeleteMachineReq{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Cannot remove the container for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id, "removed")
		}
	}
}
