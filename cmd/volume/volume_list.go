package volume

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

func GetVolumeListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all volumes",
		Long:  `list all volumes in your Islands`,
		Args:  cobra.MinimumNArgs(0),
		Run:   volumeListHandle,
	}
}

func volumeListHandle(cmd *cobra.Command, args []string) {
	r, err := connection.Client.ListVolume(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Printf("cannot list volumes for reason: %v", err)
	}

	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("NAME", "PATH", "CREATETIME")

	for _, v := range r.Volumes {
		if t, err := time.Parse(TIME_LAYOUT, strings.Split(v.CreateTime, " m=")[0]); err != nil {
			fmt.Println(err)
			return
		} else {
			table.AddRow(v.Name, v.Path, t.Format("2006-01-02 15:04:05"))
		}
	}
	fmt.Println(table)
}
