package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetMachineRenameCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rename id name",
		Short: "rename the machine",
		Long:  `rename specific machine in your computer`,
		Args:  cobra.MinimumNArgs(2),
		Run:   machineRenameHandle,
	}
}

func machineRenameHandle(cmd *cobra.Command, args []string) {
	id := args[0]
	name := args[1]

	id = CheckNameOrId(id)

	r, err := connection.Client.RenameMachine(context.Background(), &pb.RenameMachineReq{
		Id:   id,
		Name: name,
	})

	if err != nil {
		fmt.Printf("Cannot rename the container for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id)
		}
	}
}
