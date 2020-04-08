package machine

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
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

func StopContainer(timeout int, containerID string) error {
	duration := time.Duration(timeout)*time.Second
	if err := cli.ContainerStop(context.Background(), containerID, &duration); err != nil {
		logger.WithError(err).Error("failed to stop container")
		return err
	}
	return nil
}

func RestartContainer(timeout int, containerID string) error {
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