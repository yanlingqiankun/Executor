package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
)

var commitName string

func GetMachineCommitCmd() *cobra.Command {
	machineCommitCmd := &cobra.Command{
		Use:   "commit id --name ",
		Short: "commit the machine",
		Long:  `commit specific machine to the image repo`,
		Args:  cobra.MinimumNArgs(1),
		Run:   machineCommitHandle,
	}

	machineCommitCmd.Flags().StringVar(&commitName, "name", "", "name of the new image to store in image repo")
	return machineCommitCmd
}

func machineCommitHandle(cmd *cobra.Command, args []string) {
	if commitName == "" {
		fmt.Println("invalid name")
		return
	}
	id := CheckNameOrId(args[0])

	r, err := connection.Client.CommitMachine(context.Background(), &pb.CommitMachineReq{
		Id:     id,
		Name: commitName,
	})

	if err != nil {
		fmt.Printf("Cannot commit the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r.Err) {
			fmt.Println(r.Id)
		}
	}
}