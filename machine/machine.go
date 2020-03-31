package machine

import (
	"github.com/docker/docker/client"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"path/filepath"
)

var logger = logging.GetLogger("machine")
var machineRootDir string
var db = new(machineDB)
var cli *client.Client

func init() {
	level := logging.GetLevel(conf.GetString("LogLevel"))
	logger.SetLevel(level)
	machineRootDir = filepath.Join(conf.GetString("RootPath"), "machines")
	client, err := client.NewClientWithOpts(client.FromEnv)
	cli = client
	if err != nil {
		logger.Fatal("error to init docker ", err.Error())
	}
	db.init()
}