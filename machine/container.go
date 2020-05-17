package machine

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/yanlingqiankun/Executor/machine/io"
	"time"
)

func StartContainer(containerID string) error {
	if checkStatus(containerID, StatusRunning, StatusRemoving) {
		return fmt.Errorf("container is Running, don't start a running container")
	}

	err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		logger.WithError(err).Errorf("failed to start container %s", containerID)
		return fmt.Errorf("failed to start container with error %s", err.Error())
	}
	return nil
}

func DeleteContainer(containerID string) error {
	if checkStatus(containerID, StatusRunning) {
		return fmt.Errorf("can not delete running container")
	}
	if err := cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{}); err != nil{
		logger.WithField("id", containerID).WithError(err).Error("failed to delete a container")
		return err
	}
	//TODO remove in db
	return nil
}

func PauseContainer(containerID string) error {
	if !checkStatus(containerID, StatusRunning) {
		return fmt.Errorf("can not pause running container")
	}

	if err := cli.ContainerPause(context.Background(), containerID); err != nil {
		return err
	}
	return nil
}

func UnpauseContainer(containerID string) error {
	if !checkStatus(containerID, StatusPaused) {
		return fmt.Errorf("can not resume paused container")
	}
	if err := cli.ContainerUnpause(context.Background(), containerID); err != nil {
		return err
	}
	return nil
}

func KillContainer(containerID string, signal string) error {
	if !checkStatus(containerID, StatusRunning, StatusPaused) {
		return fmt.Errorf("can't kill the container not running or paused")
	}
	if err := cli.ContainerKill(context.Background(), containerID, signal); err != nil {
		return fmt.Errorf("failed to kill container %s with error : %s", containerID, err.Error())
	}
	return nil
}

func prestartHookContainer(containerID string) error {
	// todo network and volume
	return nil
}

func poststopHookContainer(containerID string) {
	return
}

func StopContainer(timeout int32, containerID string) error {
	if !checkStatus(containerID, StatusRunning, StatusPaused) {
		return fmt.Errorf("can't stop the container not running or paused")
	}
	duration := time.Duration(timeout)*time.Second
	if err := cli.ContainerStop(context.Background(), containerID, &duration); err != nil {
		logger.WithError(err).Error("failed to stop container")
		return err
	}
	return nil
}

func RestartContainer(timeout int32, containerID string) error {
	if !checkStatus(containerID, StatusRunning, StatusPaused) {
		return fmt.Errorf("can't stop the container not running or paused")
	}
	duration := time.Duration(timeout) * time.Second
	if err := cli.ContainerRestart(context.Background(), containerID, &duration); err != nil {
		logger.WithError(err).Error("failed to restart container")
		return err
	}
	return nil
}

func getContainerInfo(id string) ([]byte, error){
	_, data, err:= cli.ContainerInspectWithRaw(context.Background(), id, false)
	if err != nil {
		logger.WithError(err).Errorf("failed to get information of container %s ", id)
		return nil, err
	} else {
		return data, nil
	}
}

func getContainerState(id string) (string, error) {
	var inspcet types.ContainerJSON
	data, err := getContainerInfo(id)
	if err != nil {
		return "", err
	} else {
		err = json.Unmarshal(data, &inspcet)
		if err != nil {
			return "", err
		}
	}
	return inspcet.State.Status, nil
}

func renameContainer(id string, newName string) error {
	return cli.ContainerRename(context.Background(), id, newName)
}

func resizeTTY (id string, h, w uint32) error {
	err := cli.ContainerResize(context.Background(), id, types.ResizeOptions{
		Height: uint(h),
		Width:  uint(w),
	})
	if err != nil {
		logger.WithError(err).Error("failed to set container's tty size")
	}
	return err
}

func getContainerStdio (id string, detachKey string, openStdin bool) (chan []byte, chan []byte, chan []byte, error) {
	stream, err := cli.ContainerAttach(context.Background(), id, types.ContainerAttachOptions{
		Stream:     true,
		Stdin:      openStdin,
		Stdout:     true,
		Stderr:     true,
		DetachKeys: detachKey,
		Logs:       false,
	})
	if err != nil {
		logger.WithError(err).Error("failed to attach a container")
		return nil, nil, nil, err
	} else {
		stdin := make(chan []byte, 16)
		stdout := make(chan []byte, 16)
		stderr := make(chan []byte, 16)

		go stdinContainerHandle(&stream, stdin, detachKey)
		go stdoutContainerHandle(&stream, stdout, stderr)

		return stdin, stdout, stderr, nil
	}
}

func stdinContainerHandle(response *types.HijackedResponse, in chan []byte, detachKey string) {
	defer func(){
		detacheByte, err := io.ToBytes(detachKey)
		if err != nil {
			response.Close()
			return
		}
		response.Conn.Write(detacheByte)
		response.CloseWrite()
	}()
	for data := range in{
		if _, err := response.Conn.Write(data); err != nil {
			logger.WithError(err).Error("close attach with error")
			return
		}
	}
}

func stdoutContainerHandle(response *types.HijackedResponse, out chan []byte, stderr chan []byte) {
	for {
		data := make([]byte, 4096)
		n, err := response.Reader.Read(data)
		if err != nil {
			logger.WithError(err).Error("close attache with error")
			response.Close()
			close(out)
			close(stderr)
			return
		} else {
			out <- data[:n]
		}
	}
}

func ContainerConnectNetwork(id string, network *Network) error {
	item, exist:= db.getItem(id)
	if !exist {
		return fmt.Errorf("%s can't find in the machine repo")
	}
	c := item.machine
	c.RuntimeConfig.Networks = append(c.RuntimeConfig.Networks, network)
	db.save(true, id)
	return nil
}
