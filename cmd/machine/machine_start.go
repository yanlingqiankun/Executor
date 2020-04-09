package machine

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/io"
	"github.com/yanlingqiankun/Executor/cmd/streams"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
	"os"
	"sync"
)

var attach bool
var detachKey string

func GetMachineStartCmd() *cobra.Command {
	machineStartCmd := &cobra.Command{
		Use:   "start id",
		Short: "start the machine",
		Long:  `start specific machine in your computer`,
		Args:  cobra.MinimumNArgs(1),
		Run:   containerStartHandle,
	}

	machineStartCmd.Flags().BoolVarP(&attach, "attach", "a",false, "attach to the machine")
	machineStartCmd.Flags().StringVar(&detachKey, "detach-key", "ctrl-p", "the method to detach the machine")
	return machineStartCmd
}

func containerStartHandle(cmd *cobra.Command, args []string) {
	id := args[0]
	var before bool
	id = CheckNameOrId(id)


	resp, err := connection.Client.CanAttachJudge(context.Background(), &pb.MachineIdReq{
		Id: id,
	})
	if err != nil {
		fmt.Println(err)
	}
	if resp.ImageType == "docker" {
		before = true
	} else {
		before = false
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	group := &sync.WaitGroup{}
	// attach
	if attach && before {
		stream, err := connection.Client.AttachMachine(context.Background())

		if err != nil {
			fmt.Printf("Cannot attach to the container for reason %v\n", err)
		} else {
			if err := stream.Send(&pb.AttachStreamIn{
				Id:      id,
				Content: []byte{},
			}); err != nil {
				fmt.Println(err)
			}

			opts := io.AttachOptions{
				Tty:          true,
				AttachStdin:  true,
				AttachStdout: true,
				Id:           id,
				EscapeKeys:   detachKey,
				Container:    true,
			}

			islandsIo := io.IslandsIO{
				In:     streams.NewIn(os.NewFile(uintptr(0), "/dev/tty")),
				Out:    streams.NewOut(os.Stdout),
				Opts:   opts,
				GrpcIO: *io.NewGrpcIO(id, stream),
			}

			group.Add(1)
			go func(waitGroup *sync.WaitGroup) {
				islandsIo.Stream(ctx, waitGroup)
			}(group)
		}
	}


	r, err := connection.Client.StartMachine(context.Background(), &pb.StartMachineReq{
		Id: id,
	})

	if err != nil {
		fmt.Printf("Cannot start the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r) && !attach {
			fmt.Println(id)
		}
	}

	if attach && !before {
		stream, err := connection.Client.AttachMachine(context.Background())

		if err != nil {
			fmt.Printf("Cannot attach to the container for reason %v\n", err)
		} else {
			if err := stream.Send(&pb.AttachStreamIn{
				Id:      id,
				Content: []byte{},
			}); err != nil {
				fmt.Println(err)
			}

			opts := io.AttachOptions{
				Tty:          true,
				AttachStdin:  true,
				AttachStdout: true,
				Id:           id,
				EscapeKeys:   detachKey,
				Container:    true,
			}

			islandsIo := io.IslandsIO{
				In:     streams.NewIn(os.NewFile(uintptr(0), "/dev/tty")),
				Out:    streams.NewOut(os.Stdout),
				Opts:   opts,
				GrpcIO: *io.NewGrpcIO(id, stream),
			}

			group.Add(1)
			go func(waitGroup *sync.WaitGroup) {
				islandsIo.Stream(ctx, waitGroup)
			}(group)
		}
	}

	group.Wait()
}
