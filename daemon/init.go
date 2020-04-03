package daemon

import (
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"golang.org/x/sys/unix"
	"io"
	"os"
	"os/signal"
)

var logger = logging.GetLogger("daemon")
var exitChan = make(chan os.Signal, 1)

func init() {
	logger.SetLevel(logging.GetLevel(conf.GetString("LogLevel")))

	if os.Geteuid() != 0 {
		logger.Error("executor is not running with root privileges")
	}
	if err := os.MkdirAll(conf.GetString("RootPath"), 0700); err != nil {
		logger.WithError(err).WithField("RootPath", conf.GetString("RootPath")).Error("failed to create RootPath")
	}
	// 创建pid文件
	//initPIDLock()

	// 检查ip forward是否开启
	initIpv4Forward()

	//	监听全局退出信号
	signal.Notify(exitChan, unix.SIGINT, unix.SIGTERM, unix.SIGQUIT)
	go ExecutorExitHandle(exitChan)
}
//
//func initPIDLock() {
//	file, err := os.Create(filepath.Join(conf.GetString("RootPath"), "executor.pid"))
//	if err != nil {
//		logger.WithError(err).Fatal("failed to open pid file")
//	}
//	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
//	if err != nil {
//		logger.Fatal("service already started, exiting...")
//	}
//	if _, err = file.WriteString(fmt.Sprintf("%d", os.Getpid())); err != nil {
//		logger.WithError(err).Fatal("cannot write to pid file")
//	}
//}

func initIpv4Forward() {
	file, _ := os.OpenFile("/proc/sys/net/ipv4/ip_forward", os.O_RDWR, 0)
	forward := make([]byte, 1)
	_, _ = file.Read(forward)
	_, _ = file.Seek(0, io.SeekStart)
	if forward[0] == '0' {
		logger.Debug("enable ipv4 forward")
		_, _ = file.Write([]byte("1\n"))
	}
	_ = file.Close()
}


func ExecutorExitHandle(exit chan os.Signal) {
	signal := <-exit

	logger.WithField("signal", signal).Info("Clean up the executor")
	CleanUp()
}

func CleanUp() {
	// 关闭grpcserver
	logger.Info("shutdown rpc server")
	rpcServer.Stop()

	logger.Info("receive signal, stop containers")
	//island.IslandCleanUp()

	logger.Info("receive signal, clean up the images")
	// Todo: 停止image
}

