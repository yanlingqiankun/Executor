package daemon

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"syscall"
)

const apiPattern = "(tcp|unix)://(.+)"

var rpcServer = grpc.NewServer()

type server struct{}

func StartServer() {
	api := conf.GetString("APIPath")
	reg := regexp.MustCompile(apiPattern)
	r := reg.FindStringSubmatch(api)
	if r == nil || len(r) != 3 {
		logger.WithField("APIPath", api).Fatal("invalid APIPath")
	}

	//检查unix socket 是否占用, 占用则删除
	if r[1] == "unix" {
		if s, err := os.Stat(r[2]); err == nil {
			// 如果对应路径是个文件夹则报错
			if s.IsDir() {
				logger.WithField("APIPath", api).Fatal("invalid APIPath, target is a directory")
			} else {
				if err := os.Remove(r[2]); err != nil {
					logger.WithError(err).Fatal("failed to bind the domain socket")
				}
			}
		}
	}

	s, err := net.Listen(r[1], r[2])
	if err != nil {
		logger.WithField("APIPath", api).WithError(err).Fatal("failed to bind the domain socket")
	}

	if r[1] == "unix" {
		if group, err := user.LookupGroup("executor"); err == nil {
			gid, _ := strconv.Atoi(group.Gid)
			syscall.Chown(r[2], syscall.Getuid(), gid)
			syscall.Chmod(r[2], 0770)
		} else {
			logger.WithError(err).Error("failed to get current groups")
		}
	}

	pb.RegisterExecutorServer(rpcServer, &server{})
	reflection.Register(rpcServer)
	logger.WithField("APIPath", api).Info("rpc server started!")
	if err := rpcServer.Serve(s); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

