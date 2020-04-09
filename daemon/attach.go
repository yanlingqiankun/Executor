package daemon

import (
	"context"
	"fmt"
	"github.com/yanlingqiankun/Executor/machine"
	"github.com/yanlingqiankun/Executor/pb"
)

func (s server) AttachMachine(req pb.Executor_AttachMachineServer) error {
	return attacheMachine(req)
}

func attacheMachine(req pb.Executor_AttachMachineServer) error {
	r, err := req.Recv()
	if err != nil {
		logger.WithError(err).Debug("abort attach")
		return err
	}
	logger.WithField("id", r.Id).Debug("attach to machine")

	if m, err := machine.GetMachine(r.Id); err != nil {
		logger.WithError(err).WithField("id", r.Id).Error("can't open machine")
		return err
	} else {
		if !m.(*machine.Base).RuntimeConfig.Tty {
			return fmt.Errorf("the machine %s don't have a tty ", r.Id)
		}
		// TODO detach key
		stdin, stdout, _, err := m.GetStdio("ctrl-p")
		if err != nil {
			logger.WithError(err).Error("failed to get stdio of machine")
			return err
		}
		exitChan := make(chan bool, 2)
		go attachStdinHandle(r.Content, req, stdin, exitChan)
		go attachStdoutHandle(req, stdout, exitChan)
		<-exitChan
	}
	return nil
}

func attachStdinHandle(content []byte, req pb.Executor_AttachMachineServer, stdin chan []byte, exit chan bool) {
	remoteStdin := make(chan []byte, 1)
	if content != nil {
		remoteStdin <- content
	}

	go func() {
		for {
			if r, err := req.Recv(); err != nil {
				logger.WithError(err).Debug("stop")
				close(remoteStdin)
				return
			} else {
				remoteStdin <- r.Content
			}
		}
	}()
	for chunk := range remoteStdin {
		stdin <- chunk
	}
	close(stdin)
	exit <- true
}

func attachStdoutHandle(req pb.Executor_AttachMachineServer, stdout chan []byte, exit chan bool) {
	for buffer := range stdout {
		//fmt.Print(string(buffer))
		if err := req.Send(&pb.AttachStreamOut{Content: buffer}); err == nil {
			continue
		} else {
			// 可能这里需要丢弃chan内数据避免island阻塞
			logger.WithError(err).Debug("error while sending")
			return
		}
	}
	exit <- true
}


func (s server) ResizeMachineTTY(ctx context.Context, req *pb.ResizeTTYReq) (*pb.Error, error) {
	err := resizeMachineTTY(req.Id, req.Width, req.Height)
	if err != nil {
		return newErr(1, err), nil
	} else {
		return newErr(0, err), nil
	}
}

func resizeMachineTTY (id string, w, h uint32) error {
	m, err := machine.GetMachine(id)
	if err != nil {
		logger.WithError(err).Error("failed to get get machine")
		return err
	} else {
		return m.ResizeTTY(h, w)
	}
}

func (s server) CanAttachJudge(ctx context.Context, req *pb.MachineIdReq) (*pb.CanAttachJudgeResp, error) {
	tty, state, imageType:= canAttacheJudge(req.Id)
	return &pb.CanAttachJudgeResp{
		Tty:                  tty,
		State:                state,
		ImageType:            imageType,
	}, nil
}

func canAttacheJudge(id string) (bool, string, string) {
	m, err := machine.GetMachine(id)
	if err != nil {
		return false, "", ""
	} else {
		state := m.GetState()
		return m.(*machine.Base).RuntimeConfig.Tty, state, m.(*machine.Base).ImageType
	}
}