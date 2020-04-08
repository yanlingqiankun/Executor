package machine

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/image"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"strings"
	"time"
)

var quiet bool

func GetMachineListCmd() *cobra.Command {
	containerListCmd := &cobra.Command{
		Use:   "list",
		Short: "list the machines",
		Long:  `list machines in your computer`,
		Args:  cobra.MinimumNArgs(0),
		Run:   containerListHandle,
	}

	containerListCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Only display numeric IDs")
	return containerListCmd
}

func containerListHandle(cmd *cobra.Command, args []string) {
	r, err := connection.Client.ListMachine(context.Background(), &empty.Empty{})

	if err != nil {
		fmt.Printf("Cannot start the container for reason %v\n", err)
	} else {
		if !utils.PrintError(r.Err) {

			if quiet {
				for _, v := range r.MachineInfos {
					fmt.Printf("%s\t", utils.CheckLength(v.Id))
				}
				return
			}

			table := uitable.New()
			table.MaxColWidth = 50

			table.AddRow("ID", "NAME", "IMAGE", "IMAGEID", "IMAGETYPE", "CREATED", "STATUS")

			for _, v := range r.MachineInfos {
				if t, err := time.Parse(image.TIME_LAYOUT, strings.Split(v.CreateTime, " m=")[0]); err != nil {
					fmt.Println(err)
					return
				} else {
					table.AddRow(
						utils.CheckLength(v.Id),
						v.Name,
						v.ImageName,
						utils.CheckLength(v.ImageId),
						v.ImageType,
						t.Format("2006-01-02 15:04:05"),
						v.Status,
					)
				}
			}
			fmt.Println(table)
		}
	}

}
