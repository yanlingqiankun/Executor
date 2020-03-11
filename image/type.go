package image

import "time"

const (
	TYPE_DOCKER_RAW = 1
	TYPE_DOCKER_LAYER = 2
	TYPE_KVM_IOS = 3
)

type image struct {
	Name string				`json:"name"`
	CreateTime time.Time	`json:"create-time"`
	sum string 				`json:"sha256"`
	Type int 				`json:"type"`
}
