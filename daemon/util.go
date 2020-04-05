package daemon

import "github.com/yanlingqiankun/Executor/pb"

func newErr(code uint32, err error) *pb.Error {
	if err == nil {
		return &pb.Error{
			Code:                 0,
			Message:              "",
		}
	}
	return &pb.Error{
		Code:    code,
		Message: err.Error(),
	}
}