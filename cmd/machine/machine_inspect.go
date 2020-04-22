package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/pb"
)

func GetMachineInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect id",
		Short: "inspect the container",
		Long:  `inspect specific container in your executor`,
		Args:  cobra.MinimumNArgs(1),
		Run:   machineInspectHandle,
	}
}

func machineInspectHandle(cmd *cobra.Command, args []string) {
	id := CheckNameOrId(args[0])
	resp, err := connection.Client.InspectMachine(context.Background(), &pb.MachineIdReq{
		Id: id,
	})

	if err != nil {
		fmt.Printf("Cannot inspect the machine for reason %v\n", err)
		return
	} else {
		fmt.Println("name : ", resp.Name)
		fmt.Println("type : ", resp.Type)
		fmt.Println("runtime_config : ")
		fmt.Println(resp.RuntimeConfig)
		fmt.Println("spec : ")
		fmt.Println(resp.Spec)
	}
}
