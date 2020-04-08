package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetMachineStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start id",
		Short: "start the machine",
		Long:  `start specific machine in your computer`,
		Args:  cobra.MinimumNArgs(1),
		Run:   containerStartHandle,
	}
}

func containerStartHandle(cmd *cobra.Command, args []string) {
	id := args[0]
	id = CheckNameOrId(id)
	r, err := connection.Client.StartMachine(context.Background(), &pb.StartMachineReq{
		Id: id,
	})

	if err != nil {
		fmt.Printf("Cannot start the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r) {
			fmt.Println(id)
		}
	}
}
