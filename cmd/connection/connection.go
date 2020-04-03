package connection

import (
	"github.com/yanlingqiankun/Executor/pb"
	"google.golang.org/grpc"
)

var Client pb.ExecutorClient

func InitClient(apiPath string) {
	if conn, err := grpc.Dial(apiPath, grpc.WithInsecure()); err != nil {
		panic(err)
	} else {
		Client = pb.NewExecutorClient(conn)
	}
}
