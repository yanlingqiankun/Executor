package attach

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/io"
	"github.com/yanlingqiankun/Executor/cmd/machine"
	"github.com/yanlingqiankun/Executor/cmd/streams"
	"github.com/yanlingqiankun/Executor/pb"
	"os"
	"sync"
)

var (
	inChan   = make(chan []byte)
	outChan  = make(chan []byte)
	exitChan = make(chan int)
	once     sync.Once
)

var (
	detachKeys string
)

func GetAttachCmd() *cobra.Command {
	attachCmd := &cobra.Command{
		Use:   "attach [options]",
		Short: "Attach to machine",
		Long:  `Attach local standard input, output, and error streams to a running machine`,
		Args:  cobra.MinimumNArgs(1),
		Run:   attachHandle,
	}

	attachCmd.Flags().StringVar(&detachKeys, "detach-keys", "ctrl-p,ctrl-q", "Override the key sequence for detaching a container.")
	return attachCmd
}

func attachHandle(cmd *cobra.Command, args []string) {

	id := machine.CheckNameOrId(args[0])
	up, err := CheckContainerUp(id)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if !up {
			fmt.Printf("Machine %s is not running\n", id)
			return
		}
	}

	group := &sync.WaitGroup{}
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

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		opts := io.AttachOptions{
			Tty:          true,
			AttachStdin:  true,
			AttachStdout: true,
			Id:           id,
			EscapeKeys:   detachKeys,
			Container:    true,
		}

		islandsIo := io.IslandsIO{
			In:   streams.NewIn(os.NewFile(uintptr(0), "/dev/tty")),
			Out:  streams.NewOut(os.Stdout),
			Opts: opts,
			GrpcIO:  *io.NewGrpcIO(id, stream),
		}

		group.Add(1)
		go func(waitGroup *sync.WaitGroup) {
			islandsIo.Stream(ctx, group)
		}(group)

		group.Wait()
	}
}

func GetNetworkOutput(group *sync.WaitGroup, stream pb.Executor_AttachMachineClient, out chan []byte) {
	defer func() {
		group.Done()
	}()
	for {
		if outCon, err := stream.Recv(); err != nil {
			close(outChan)
			return
		} else {
			out <- outCon.Content
		}
	}
}
//
func CheckContainerUp(machineID string) (bool, error) {
	if resp, err := connection.Client.CanAttachJudge(context.Background(), &pb.MachineIdReq{
		Id: machineID,
	}); err != nil {
		fmt.Println(err)
		return false, err
	} else {
		if resp.Tty && resp.State == "running" && resp.ImageType != "iso" {
			return true, nil
		}
		return false, nil
	}
}
