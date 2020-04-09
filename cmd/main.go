package main

import (
	"fmt"
	"github.com/yanlingqiankun/Executor/cmd/attach"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/image"
	"github.com/yanlingqiankun/Executor/cmd/machine"
	"github.com/yanlingqiankun/Executor/cmd/network"
	"os"
)

var ApiPath = ""
var DefaultApiPath = "unix:///var/run/Executor.sock"

func init() {
	// 获取api path
	if api := os.Getenv("EXECUTOR_API"); api == "" {
		ApiPath = DefaultApiPath
	} else {
		ApiPath = api
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	connection.InitClient(ApiPath)

	rootCmd := GetRootCmd()
	rootCmd.AddCommand(
		network.GetNetworkCmd(),
		image.GetImageCmd(),
		machine.GetMachineCmd(),
		attach.GetAttachCmd(),
	)
	_ = rootCmd.Execute()
}
