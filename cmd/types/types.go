package types

import "github.com/yanlingqiankun/Executor/pb"

type AttachStream interface {
	Send(*pb.AttachStreamIn) error
	Recv() (*pb.AttachStreamOut, error)
}
