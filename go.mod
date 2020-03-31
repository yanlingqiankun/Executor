module github.com/yanlingqiankun/Executor

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/containerd/containerd v1.3.3 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/libvirt/libvirt-go v6.0.0+incompatible
	github.com/libvirt/libvirt-go-xml v6.0.0+incompatible
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
	golang.org/x/sys v0.0.0-20200327173247-9dae0f8f5775 // indirect
	google.golang.org/grpc v1.27.1 // indirect
)

replace github.com/docker/docker v1.13.1 => github.com/docker/engine v0.0.0-20191113042239-ea84732a7725
