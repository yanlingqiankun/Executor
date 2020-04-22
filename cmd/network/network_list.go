package network

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"strings"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05.999999999 -0700 MST"

func GetNetworkListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all Network",
		Long:  `list all Networks in your Executor`,
		Args:  cobra.MinimumNArgs(0),
		Run:   networkList,
	}
}

func networkList(cmd *cobra.Command, args []string) {
	r, err := connection.Client.ListNetwork(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Printf("cannot list the Network for reason: %v\n", err)
	} else {
		table := uitable.New()
		table.MaxColWidth = 50

		table.AddRow("Name", "TIME", "SUBNET", "GATEWAY", "TYPE")

		for _, v := range r.Networks {
			if t, err := time.Parse(TIME_LAYOUT, strings.Split(v.CreateTime, " m=")[0]); err != nil {
				fmt.Println(err)
				return
			} else {
				table.AddRow(
					v.Name,
					t.Format("2006-01-02 15:04:05"),
					v.Subnet,
					v.Gateway,
					v.Type,
				)
			}
		}
		fmt.Println(table)
	}
}

