package image

import (
	"time"
)

type ImageEntry struct {
	Name          string            `json:"name"`
	ID            string            `json:"id"`
	CreateTime    time.Time         `json:"create-time"`
	Type          string            `json:"type"`		// docker_save, docker_raw, kvm_ios, kvm_qcow2
	IsDockerImage bool              `json:"is_docker_image"`
	Counter       int32 				`json:"counter"`
}

type Image interface {
	Remove() error
	Rename() error
	GetType() error
}
