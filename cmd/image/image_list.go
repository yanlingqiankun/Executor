package image

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"strings"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05.999999999 -0700 MST"

func GetImageListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all images",
		Long:  `list all images in your repo`,
		Args:  cobra.MaximumNArgs(0),
		Run:   imageList,
	}
}

func imageList(cmd *cobra.Command, args []string) {
	r, err := connection.Client.ListImage(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Printf("cannot list the images for reason: %v\n", err)
	} else {
		table := uitable.New()
		table.MaxColWidth = 50

		table.AddRow("ID", "TIME", "NAME", "MACHINES", "TYPE")

		for _, v := range r.Images {
			if t, err := time.Parse(TIME_LAYOUT, strings.Split(v.CreateTime, " m=")[0]); err != nil {
				fmt.Println(err)
				return
			} else {
				table.AddRow(
					utils.CheckLength(v.Id),
					t.Format("2006-01-02 15:04:05"),
					utils.CheckLength(v.Name),
					v.Machines,
					v.Type,
				)
			}
		}
		fmt.Println(table)
	}
}
