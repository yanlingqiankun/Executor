package machine

import (
	"github.com/docker/docker/api/types/container"
)


const (
	StatusRunning    = "running"
	StatusCreated    = "created"
	StatusExited     = "exited"
	StatusRemoving   = "removing"
	StatusDead       = "dead"
	StatusRestarting = "restarting"
	StatusPaused     = "paused"
)

type BaseContainer struct {
	Base 			*Base
	ContainerConfig  *container.Config  `json:"container_config"`
	HostConfig 	     *container.HostConfig  `json:"host_config"`
}


