package machine

import (
	"github.com/docker/docker/api/types/container"
	"github.com/yanlingqiankun/Executor/machine/types"
	"github.com/yanlingqiankun/Executor/network/proxy"
)

func CreateContainer(imageID string) Factory {
	container := &BaseContainer{
		ImageID:         imageID,
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
			Image:           "",
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
func (container *BaseContainer) SetName(name string) error {
	if name == "" {
		return nil
	}
	container.ContainerConfig.Domainname = name
	return nil
}

func (container *BaseContainer) SetImage(imageID string) {
	container.ImageID = imageID
}

func (container *BaseContainer) SetHostname(name string) {
	if name == "" {
		return
	}
	container.ContainerConfig.Hostname = name
}

func (container *BaseContainer) SetVolumes(volumes []*types.Volume) {
	if volumes == nil || len(volumes) == 0 {
		return
	}
	for _, v := range volumes {
		dest := v.Destination
		if dest == "" {
			continue
		}
		container.RuntimeConfig.Volumes[dest] = v
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
	container.RuntimeConfig.ProxyManager = proxy.GetProxyManager()
	if len(container.RuntimeConfig.Networks) == 0 {
		return
	}
	container.RuntimeConfig.ExposedPorts = info
}

func (container *BaseContainer) SetHosts(hosts []string) {
	container.HostConfig.ExtraHosts = hosts
}
