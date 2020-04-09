package daemon

import (
	"context"
	"github.com/yanlingqiankun/Executor/machine"
	"github.com/yanlingqiankun/Executor/pb"
)

func (s server) AttachMachine(req pb.Executor_AttachMachineServer) error {
	panic("implement me")
}
//
//func attacheMachine(req pb.Executor_AttachMachineServer) error {
//	r, err := req.Recv()
//	if err != nil {
//		logger.WithError(err).Debug("abort attach")
//		return err
//	}
//	logger.WithField("id", r.Id).Debug("attach to machine")
//
//	if m, err := machine.GetMachine(r.Id); err != nil {
//		logger.WithError(err).WithField("id", r.Id).Error("can't open machine")
//		return err
//	} else {
//		if !m.(*machine.Base).RuntimeConfig.Tty {
//			return fmt.Errorf("the machine %s don't have a tty ", r.Id)
//		}
//		stdin, stdout, _ := m.GetStdio()
//		exitChan := make(chan bool, 2)
//		go attachStdinHandle(r.Content, req, stdin, exitChan)
//		go attachStdoutHandle(req, stdout, exitChan)
//		<-exitChan
//	}
//	return nil
//}
//


func (s server) ResizeMachineTTY(ctx context.Context, req *pb.ResizeTTYReq) (*pb.Error, error) {
	err := resizeMachineTTY(req.Id, req.Width, req.Height)
	if err != nil {
		return newErr(1, err), err
	} else {
		return newErr(0, err), err
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
