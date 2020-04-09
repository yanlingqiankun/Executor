module github.com/yanlingqiankun/Executor/cmd

go 1.13

require (
	github.com/fatih/color v1.9.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/gosuri/uitable v0.0.4
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/spf13/cobra v0.0.7
	github.com/yanlingqiankun/Executor v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
	google.golang.org/grpc v1.28.0
)

replace github.com/yanlingqiankun/Executor => ../
