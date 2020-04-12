package daemon

import (
	"crypto/rand"
	"fmt"
	"github.com/yanlingqiankun/Executor/pb"
)

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

func convertHostsFromPB(hosts []*pb.HostEntry) []string {
	machineHosts := make([]string, len(hosts))
	for index, entry := range hosts {
		machineHosts[index] = entry.Host + ":" + entry.Ip
	}
	return machineHosts
}

func getMac() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		logger.WithError(err).Error("failed to get a rand mac")
		return ""
	}
	// Set the local bit
	buf[0] |= 2
	return fmt.Sprintf("00:%02x:%02x:%02x:%02x:%02x",  buf[1], buf[2], buf[3], buf[4], buf[5])
}
