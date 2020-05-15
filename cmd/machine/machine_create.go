package machine

import (
	"context"
	"fmt"
	"github.com/docker/go-units"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/image"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
	"os"
	"strings"
)


type CgroupsSetting struct {
	memoryString      string
	memoryReservation string
	memorySwap        string
	oomKillDisable    bool
	kernelMemory      string

	cpuShares          int64
	cpuPeriod          int64
	cpuRealtimePeriod  int64
	cpuRealtimeRuntime int64
	cpuQuota           int64
	cpusetCpus         string
	cpusetMems         string
	swappiness         int64
	cgroupParent       string
}


var (
	hostname        string
	env             []string
	mountInfo       []string
	stopSignal      string
	tty             bool
	extraHostsValue []string
	dns             []string
	gateway         string
	mac             string
	network         string
	name            string
	ipAddr          string
	exposedPorts    []string
	cgroupsSetting CgroupsSetting
)

func GetMachineCreateCmd() *cobra.Command {
	machineCreateCmd := &cobra.Command{
		Use:   "create id [options]",
		Short: "create a machine",
		Long:  `create a machine in your server`,
		Args:  cobra.MinimumNArgs(1),
		Run:   machineCreateHandle,
	}

	machineCreateCmd.Flags().StringSliceVarP(&env, "env", "e", []string{}, "Set environment variables")
	machineCreateCmd.Flags().StringVar(&hostname, "hostname", "", "Machine host name")
	machineCreateCmd.Flags().StringArrayVar(&mountInfo, "mount", []string{}, "Bind mount a volume")
	machineCreateCmd.Flags().StringVar(&stopSignal, "stop-signal", "SIGTERM", " Signal to stop a machine")
	machineCreateCmd.Flags().BoolVarP(&tty, "tty", "t", false, "Allocate a pseudo-TTY")
	machineCreateCmd.Flags().StringSliceVar(&extraHostsValue, "add-host", []string{}, "Add a custom host-to-IP mapping (host:ip)")
	machineCreateCmd.Flags().StringSliceVar(&dns, "dns", []string{}, "Set custom DNS servers")
	machineCreateCmd.Flags().StringVar(&mac, "mac", "", "set mac for your machine")
	machineCreateCmd.Flags().StringVar(&network, "network", "", "set network for your machine")
	machineCreateCmd.Flags().StringVar(&ipAddr, "ip", "", "set ip and mask for your machine")
	machineCreateCmd.Flags().StringVar(&name, "name", "", "Assign a name to the machine")
	machineCreateCmd.Flags().StringArrayVar(&exposedPorts, "exposed-ports", nil, "Publish a machine's port(s) to the host")

	flags := machineCreateCmd.Flags()
	flags.StringVar(&cgroupsSetting.cpusetCpus, "cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	flags.StringVar(&cgroupsSetting.cpusetMems, "cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	flags.Int64Var(&cgroupsSetting.cpuPeriod, "cpu-period", 0, "Limit CPU CFS (Completely Fair Scheduler) period")
	flags.Int64Var(&cgroupsSetting.cpuQuota, "cpu-quota", 0, "Limit CPU CFS (Completely Fair Scheduler) quota")
	flags.Int64Var(&cgroupsSetting.cpuRealtimePeriod, "cpu-rt-period", 0, "Limit CPU real-time period in microseconds")
	flags.Int64Var(&cgroupsSetting.cpuRealtimeRuntime, "cpu-rt-runtime", 0, "Limit CPU real-time runtime in microseconds")
	flags.Int64VarP(&cgroupsSetting.cpuShares, "cpu-shares", "c", 0, "CPU shares (relative weight)")
	flags.StringVar(&cgroupsSetting.kernelMemory, "kernel-memory", "", "Kernel memory limit")
	flags.StringVarP(&cgroupsSetting.memoryString, "memory", "m", "", "Memory limit")
	flags.StringVar(&cgroupsSetting.memoryReservation, "memory-reservation", "", "Memory soft limit")
	flags.StringVar(&cgroupsSetting.memorySwap, "memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	flags.Int64Var(&cgroupsSetting.swappiness, "memory-swappiness", -1, "Tune container memory swappiness (0 to 100)")
	flags.BoolVar(&cgroupsSetting.oomKillDisable, "oom-kill-disable", false, "Disable OOM Killer")
	flags.StringVar(&cgroupsSetting.cgroupParent, "cgroup-parent", "", "Optional parent cgroup for the machine")

	return machineCreateCmd
}

func machineCreateHandle(cmd *cobra.Command, args []string) {

	imageId := args[0]

	//判断是id还是name
	imageId = image.CheckNameOrId(imageId)

	//处理volumes
	volumes := MountHandle(mountInfo)

	//ExtraHost
	extraHostsParam := ExtraHostDeal(extraHostsValue)

	// interface
	var interfaceSlice []*pb.NetworkInterface

	if network != "" {
		interfaceSlice = append(interfaceSlice, &pb.NetworkInterface{
			Name:                 "",
			Bridge:               network,
			Mac:                  "",
			Gateway:              "",
		})
	}

	if ipAddr != "" && network != "" {
		interfaceSlice[0].Address = append(interfaceSlice[0].Address, &pb.NetworkAddress{
			Ip:                   ipAddr,
			Mask:                 -1,
		})
	}

	// 设置环境变量
	defaultEnv := []string{}
	term := os.Getenv("TERM")
	defaultEnv = append(defaultEnv, "TERM="+term)
	env = append(env, defaultEnv...)

	//检查args长度
	command := make([]string, 0)
	if len(args) >= 2 {
		command = args[1:]
	}

	//ExposedPorts
	err, exposedPortsStruct := ExposedPortsHandle(exposedPorts)
	if err != nil {
		fmt.Printf("Cannot create the container for reason %v\n", err)
		return
	}

	resources, err := parse_resource()

	r, err := connection.Client.CreateMachine(context.Background(), &pb.CreateMachineReq{
		ImageId: imageId,
		Name:    name,
		Network: &pb.Network{
			Hostname:   hostname,
			ExtraHosts: extraHostsParam,
			Interfaces: interfaceSlice,
		},
		Env: env,
		Volumes:    volumes,
		//StopSignal: stopSignal,
		Tty:        tty,
		Cmd:        command,
		Resources:resources,
		ExposedPorts:exposedPortsStruct,
	})
	if err != nil {
		fmt.Printf("Cannot create the machine for reason %v\n", err)
	} else {
		if !utils.PrintError(r.Err) {
			id := strings.ReplaceAll(r.Id, "-", "")[:12]
			fmt.Println(id)
		}
	}
}
//
func CheckNameOrId(machineId string) string {
	imageListResp, err := connection.Client.ListMachine(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	isNameFlag := false
	isIdFlag := false

	for _, v := range imageListResp.MachineInfos {
		if machineId == v.Name {
			isNameFlag = true
			machineId = v.Id
			break
		}

		if len(machineId) == 12 {
			if machineId == v.Id[:12] {
				isIdFlag = true
				machineId = v.Id
				break
			}
		} else {
			if machineId == v.Id {
				isIdFlag = true
				machineId = v.Id
				break
			}
		}
	}

	if !(isNameFlag || isIdFlag) {
		panic("No machine named: " + machineId)
	}

	return machineId
}

func MountHandle(mountStrList []string) (volumes []*pb.MachineVolume) {

	for _, mountStr := range mountStrList {
		mountSlice := strings.Split(mountStr, ",")

		verifyMap := map[string]string{
			"name": "",
			"source": "",
			"target": "",
			"ro":     "",
		}

		readOnly := false

		for _, val := range mountSlice {
			if strings.Contains(val, "=") {
				kv := strings.Split(val, "=")
				k := strings.TrimSpace(kv[0])
				v := strings.TrimSpace(kv[1])
				if _, ok := verifyMap[k]; ok {
					verifyMap[k] = v
				}
			} else {
				if val == "ro" {
					readOnly = true
				}
			}
		}

		volumes = append(volumes, &pb.MachineVolume{
			Source:      verifyMap["source"],
			Destination: verifyMap["target"],
			Readonly:    readOnly,
		})

	}
	return
}

func ExtraHostDeal(extraHosts []string) []*pb.HostEntry {
	var extra []*pb.HostEntry

	for _, val := range extraHosts {
		hostsList := strings.Split(val, ":")
		if len(hostsList) < 2 {
			panic("wrong format of the extra hosts")
		}

		extra = append(extra, &pb.HostEntry{
			Ip:   hostsList[1],
			Host: hostsList[0],
		})
	}
	return extra
}

func PrintmachineNetworkError() {
	fmt.Println("executor: please check bridge, gateway and network address, you may lost one of them.")
	fmt.Println("See `executor machine create --help`.")
}

func ExposedPortsHandle(portMap []string) (err error, exposedPorts map[string]*pb.PortBindings){
	exposedPorts = make(map[string]*pb.PortBindings)
	for _, exposedPort := range portMap {
		tmpSlice := strings.Split(exposedPort, ":")
		if len(tmpSlice) == 1 {
			// machinePort
			if tmpExposedPorts, ok := exposedPorts[tmpSlice[0]]; ok {
				exposedPorts[tmpSlice[1]].PortBindings = append(tmpExposedPorts.PortBindings, &pb.PortBinding{
					HostIp:               "",
					HostPort:             "",
				})
			} else {
				tmp := make([]*pb.PortBinding, 1)
				exposedPorts[tmpSlice[0]] = &pb.PortBindings{PortBindings:tmp}
				exposedPorts[tmpSlice[0]].PortBindings = []*pb.PortBinding{
					{
						HostIp:               "",
						HostPort:             "",
					},
				}
			}
		} else if len(tmpSlice) == 2 {
			// hostPort:machinePort
			if tmpExposedPorts, ok := exposedPorts[tmpSlice[1]]; ok {
				exposedPorts[tmpSlice[1]].PortBindings = append(tmpExposedPorts.PortBindings, &pb.PortBinding{
					HostIp:               "",
					HostPort:             tmpSlice[0],
				})
			} else {
				tmp := make([]*pb.PortBinding, 1)
				exposedPorts[tmpSlice[1]] = &pb.PortBindings{PortBindings:tmp}
				exposedPorts[tmpSlice[1]].PortBindings = []*pb.PortBinding{
					{
						HostIp:               "",
						HostPort:             tmpSlice[0],
					},
				}
			}
		} else if len(tmpSlice) == 3 {
			// hostIP:hostPort:machinePort
			if tmpExposedPorts, ok := exposedPorts[tmpSlice[2]]; ok {
				exposedPorts[tmpSlice[1]].PortBindings = append(tmpExposedPorts.PortBindings, &pb.PortBinding{
					HostIp:               tmpSlice[0],
					HostPort:             tmpSlice[1],
				})
			} else {
				tmp := make([]*pb.PortBinding, 1)
				exposedPorts[tmpSlice[2]] = &pb.PortBindings{PortBindings:tmp}
				exposedPorts[tmpSlice[2]].PortBindings = []*pb.PortBinding{
					{
						HostIp:               tmpSlice[0],
						HostPort:             tmpSlice[1],
					},
				}
			}
		} else {
			return fmt.Errorf("error exposed-ports"), nil
		}
	}
	return
}

func parse_resource() (*pb.Resources, error) {
	var memory int64
	var err error
	if cgroupsSetting.memoryString != "" {
		memory, err = units.RAMInBytes(cgroupsSetting.memoryString)
		if err != nil {
			return nil, err
		}
	}

	var memoryReservation int64
	if cgroupsSetting.memoryReservation != "" {
		memoryReservation, err = units.RAMInBytes(cgroupsSetting.memoryReservation)
		if err != nil {
			return nil, err
		}
	}

	var memorySwap int64
	if cgroupsSetting.memorySwap != "" {
		if cgroupsSetting.memorySwap == "-1" {
			memorySwap = -1
		} else {
			memorySwap, err = units.RAMInBytes(cgroupsSetting.memorySwap)
			if err != nil {
				return nil, err
			}
		}
	}

	var kernelMemory int64
	if cgroupsSetting.kernelMemory != "" {
		kernelMemory, err = units.RAMInBytes(cgroupsSetting.kernelMemory)
		if err != nil {
			return nil, err
		}
	}

	swappiness := cgroupsSetting.swappiness
	if swappiness != -1 && (swappiness < 0 || swappiness > 100) {
		return nil, fmt.Errorf("invalid value: %d. Valid memory swappiness range is 0-100", swappiness)
	}

	resources := &pb.Resources{
		CgroupParent:       cgroupsSetting.cgroupParent,
		Memory:             memory,
		MemoryReservation:  memoryReservation,
		MemorySwap:         memorySwap,
		MemorySwappiness:   swappiness,
		KernelMemory:       kernelMemory,
		OomKillDisable:     cgroupsSetting.oomKillDisable,
		CPUShares:          cgroupsSetting.cpuShares,
		CPUPeriod:          cgroupsSetting.cpuPeriod,
		CpusetCpus:         cgroupsSetting.cpusetCpus,
		CpusetMems:         cgroupsSetting.cpusetMems,
		CPUQuota:           cgroupsSetting.cpuQuota,
		CPURealtimePeriod:  cgroupsSetting.cpuRealtimePeriod,
		CPURealtimeRuntime: cgroupsSetting.cpuRealtimeRuntime,
	}

	return resources, nil
}

