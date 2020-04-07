package machine

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/yanlingqiankun/Executor/network/proxy"
)

func CreateContainer(imageID string) Factory {
	container := &BaseContainer{
		Base: &Base{
			IsDocker:      true,
			ImageID:       imageID,
			ID:            "",
			ImagePath:     "",
			ImageType:     "",
			Name:          "",
			RuntimeConfig: &RuntimeConfig{},
		},
		ContainerConfig: &container.Config{
			Hostname:        "",
			Domainname:      "",
			User:            "",
			AttachStdin:     false,
			AttachStdout:    false,
			AttachStderr:    false,
			ExposedPorts:    nil,
			Tty:             false,
			OpenStdin:       false,
			StdinOnce:       false,
			Env:             nil,
			Cmd:             nil,
			Healthcheck:     nil,
			ArgsEscaped:     false,
			Image:           imageID,
			Volumes:         nil,
			WorkingDir:      "",
			Entrypoint:      nil,
			NetworkDisabled: true,
			MacAddress:      "",
			OnBuild:         nil,
			Labels:          nil,
			StopSignal:      "",
			StopTimeout:     nil,
			Shell:           nil,
		},
		HostConfig: nil,
	}
	return container
}


// Factory interface
func (container *BaseContainer) Create() error {
	body, err := cli.ContainerCreate(context.Background(), container.ContainerConfig, container.HostConfig, nil, container.Base.Name)
	if err != nil {
		return err
	}
	container.Base.ID = body.ID
	logger.Debugf("The container %s created successfully", container.Base.ID)
	return nil
}

func (container *BaseContainer) SetName(name string) error {
	if name == "" {
		return nil
	}
	container.Base.Name = name
	return nil
}

func (container *BaseContainer) SetImage(imageID string, path string) {
	container.Base.ImageType = "docker"
	container.Base.ImagePath = path
	container.ContainerConfig.Image = imageID
	container.Base.ImageID = imageID
}

func (container *BaseContainer) SetHostname(name string) {
	if name == "" {
		return
	}
	container.ContainerConfig.Hostname = name
}

func (container *BaseContainer) SetVolumes(volumes []*Volume) {
	if volumes == nil || len(volumes) == 0 {
		return
	}
	for _, v := range volumes {
		dest := v.Destination
		if dest == "" {
			continue
		}
		container.Base.RuntimeConfig.Volumes[dest] = v
	}
}

func (container *BaseContainer) SetEntrypoint(entrypoint []string) {
	if entrypoint == nil || len(entrypoint) == 0 {
		return
	}

	container.ContainerConfig.Entrypoint = entrypoint
}

func (container *BaseContainer) SetCmd(cmd []string) {
	if cmd == nil || len(cmd) == 0 {
		return
	}

	container.ContainerConfig.Cmd = cmd
}

func (container *BaseContainer) SetWorkingDir(dir string) {
	if dir == "" {
		return
	}

	container.ContainerConfig.WorkingDir = dir
}

func (container *BaseContainer) SetUser(user string) {
	if user == "" {
		return
	}
	container.ContainerConfig.User = user
}

func (container *BaseContainer) SetEnv(env []string) {
	if env == nil || len(env) == 0 {
		return
	}
	container.ContainerConfig.Env = env
}

func (container *BaseContainer) SetTTY(tty bool) {
	container.ContainerConfig.Tty = tty
}

func (container *BaseContainer) SetExposedPorts(info []proxy.ProxyInfo) {
	container.Base.RuntimeConfig.ProxyManager = proxy.GetProxyManager()
	if len(container.Base.RuntimeConfig.Networks) == 0 {
		return
	}
	container.Base.RuntimeConfig.ExposedPorts = info
}

func (container *BaseContainer) SetHosts(hosts []string) {
	container.HostConfig.ExtraHosts = hosts
}

func (container *BaseContainer) SetTTYSize(width, height uint16) {

}

func (container *BaseContainer) GetBase() (*Base, error) {
	return container.Base, nil
}
