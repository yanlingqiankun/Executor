package machine

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"time"
)

func Start(containerID string) error {
	if checkStatus(containerID, StatusRunning, StatusRemoving) {
		return fmt.Errorf("container is Running, don't start a running container")
	}

	err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		return fmt.Errorf("failed to start container with error %s", err.Error())
	}
	return nil
}

func Delete(containerID string) error {
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

func Pause(containerID string) error {
	if !checkStatus(containerID, StatusRunning) {
		return fmt.Errorf("can not pause running container")
	}

	if err := cli.ContainerPause(context.Background(), containerID); err != nil {
		return err
	}
	return nil
}

func Unpause(containerID string) error {
	if !checkStatus(containerID, StatusPaused) {
		return fmt.Errorf("can not resume paused container")
	}
	if err := cli.ContainerUnpause(context.Background(), containerID); err != nil {
		return err
	}
	return nil
}

func Kill(containerID string, signal string) error {
	if err := cli.ContainerKill(context.Background(), containerID, signal); err != nil {
		return fmt.Errorf("failed to kill container %s with error : %s", containerID, err.Error())
	}
	return nil
}

func prestartHook(containerID string) error {
	// todo network and volume
	return nil
}

func poststopHook(containerID string) {
	return
}

func Stop(timeout int, containerID string) error {
	duration := time.Duration(timeout)*time.Second
	if err := cli.ContainerStop(context.Background(), containerID, &duration); err != nil {
		logger.WithError(err).Error("failed to stop container")
		return err
	}
	return nil
}

func Restart(timeout int, containerID string) error {
	duration := time.Duration(timeout) * time.Second
	if err := cli.ContainerRestart(context.Background(), containerID, &duration); err != nil {
		logger.WithError(err).Error("failed to restart container")
		return err
	}
	return nil
}