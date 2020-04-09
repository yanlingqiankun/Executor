package io

import (
	"context"
	"fmt"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/streams"
	"github.com/yanlingqiankun/Executor/cmd/types"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type AttachOptions struct {
	Tty          bool
	AttachStdin  bool
	AttachStdout bool
	Id           string
	Container    bool
	EscapeKeys   string
}

type IslandsIO struct {
	In   streams.In
	Out  streams.Out
	Opts AttachOptions
	GrpcIO  GrpcIO
}

func (islands *IslandsIO) Stream(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()
	var once sync.Once

	// 设置输入
	if islands.Opts.AttachStdin && islands.Opts.Tty {
		oldState := islands.setInput()
		defer func() {
			terminal.Restore(0, oldState)
		}()
	}

	// 开启输入监听
	var inputDone <-chan struct{}
	if islands.Opts.AttachStdin {
		inputDone = islands.BeginInput()
	}

	// 开启输出监听
	var outputDone <-chan struct{}
	if islands.Opts.AttachStdout {
		outputDone = islands.BeginOutput()
	}

	// 检测tty变化
	if islands.Opts.Tty {
		go islands.MonitorTTY(ctx, once)
	}

	for {
		select {
		case <-outputDone:
			return
		case <-inputDone:
			return
		case <-ctx.Done():
			return
		}
	}
}

func (islands *IslandsIO) BeginInput() <-chan struct{} {
	inputDone := make(chan struct{}, 1)
	detach := make(chan error, 1)

	go func() {
		_, err := io.Copy(islands.GrpcIO, islands.In.Input)
		if err != nil {
			if _, ok := err.(EscapeError); ok {
				detach<-err
			} else {
				fmt.Println("input:" + err.Error())
			}
		}
		close(inputDone)
	}()
	return inputDone
}

func (islands *IslandsIO) BeginOutput() <-chan struct{} {
	outputDone := make(chan struct{}, 1)

	go func() {
		_, err := io.Copy(islands.Out.Output, islands.GrpcIO)
		if err != nil {
			fmt.Println(err)
		}
		close(outputDone)
	}()

	return outputDone
}

func (islands *IslandsIO) setInput() *terminal.State {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	escapeBytes := make([]byte, 3)
	if islands.Opts.EscapeKeys != "" {
		escapeBytes, err = ToBytes(islands.Opts.EscapeKeys)
		if err != nil {
			fmt.Println(err)
		}
	}

	islands.In.Input = NewEscapeProxy(islands.In.Input, escapeBytes)

	return oldState
}

func (islands *IslandsIO) MonitorTTY(ctx context.Context, once sync.Once) {

	var setTtySize func(string, [2]int, int)()
	if islands.Opts.Container {
		setTtySize = setContainerTtySize
	} else {
		setTtySize = setExecTtySize
	}

	//第一次要执行一次tty size
	once.Do(func() {
		height, width, err := terminal.GetSize(0)
		if err != nil {
			fmt.Println(err)
			return
		}
		sizes := [2]int{width, height}
		setTtySize(islands.Opts.Id, sizes, 0)
	})

	var winChSignalChan = make(chan os.Signal)
	signal.Notify(winChSignalChan, syscall.SIGWINCH)

	for {
		select {
		case <-ctx.Done():
			return
		case <-winChSignalChan:
			height, width, err := terminal.GetSize(0)
			if err != nil {
				fmt.Println(err)
				return
			}
			sizes := [2]int{width, height}
			setTtySize(islands.Opts.Id, sizes, 1)
		}
	}
}

func setContainerTtySize(id string, sizes [2]int, s int) {
	if r, err := connection.Client.ResizeMachineTTY(context.Background(), &pb.ResizeTTYReq{
		Id:     id,
		Height: uint32(sizes[0]),
		Width:  uint32(sizes[1]),
	}); err != nil {
		fmt.Println(err)
		return
	} else {
		if r.Code != 0 && s == 0 {
			time.Sleep(1*time.Second)
			setContainerTtySize(id, sizes, s)
		} else if utils.PrintError(r) {
		}
	}
}

func setExecTtySize(id string, sizes [2]int, s int) {
	//if r, err := connection.Client.ResizeExecTTY(context.Background(), &pb.ResizeTTYReq{
	//	Id:     id,
	//	Height: uint32(sizes[0]),
	//	Width:  uint32(sizes[1]),
	//}); err != nil {
	//	fmt.Println(err)
	//	return
	//} else {
	//	if !utils.PrintError(r) {
	//	}
	//}
}

type GrpcIO struct {
	stream      types.AttachStream
	id          string
	remainNum   int
	remainBytes []byte
}

func (g GrpcIO) Write(content []byte) (n int, err error) {
	err = g.stream.Send(&pb.AttachStreamIn{
		Id:      g.id,
		Content: content,
	})
	n = len(content)
	return
}

func (g GrpcIO) Read(content []byte) (n int, err error) {

	check := func(grpcIO GrpcIO, resp *pb.AttachStreamOut, copyNum int) {
		if copyNum != len(resp.Content) {
			g.remainNum = len(resp.Content) - copyNum
			g.remainBytes = resp.Content[g.remainNum:]
		}
	}

	if g.remainNum == 0 && len(g.remainBytes) == 0 {
		r, err := g.stream.Recv()
		if err != nil {
			if r != nil {
				if len(r.Content) != 0 {
					copyNum := copy(content, r.Content)
					check(g, r, copyNum)
					return copyNum, err
				}
			}
			return 0, err
		} else {
			copyNum := copy(content, r.Content)
			check(g, r, copyNum)
			return len(r.Content), err
		}
	} else {
		copyNum := copy(content, g.remainBytes)
		if copyNum != g.remainNum {
			g.remainNum = g.remainNum - copyNum
			g.remainBytes = g.remainBytes[g.remainNum:]
		}
		return copyNum, err
	}
}

func NewGrpcIO(id string, stream types.AttachStream) *GrpcIO {
	return &GrpcIO{
		stream: stream,
		id:     id,
	}
}
