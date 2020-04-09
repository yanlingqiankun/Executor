package machine

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
	"github.com/yanlingqiankun/Executor/cmd/connection"
	"github.com/yanlingqiankun/Executor/cmd/image"
	utils "github.com/yanlingqiankun/Executor/cmd/util"
	"github.com/yanlingqiankun/Executor/pb"
	"os"
	"strings"
)

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
	bridge          string
	cidr            string
	name            string
	ipAddr          string
)

func GetMachineCreateCmd() *cobra.Command {
	machineCreateCmd := &cobra.Command{
		Use:   "create id [command]",
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
	machineCreateCmd.Flags().StringVar(&mac, "mac", "", "set mac for your container")
	machineCreateCmd.Flags().StringVar(&bridge, "network", "", "set network for your machine")
	machineCreateCmd.Flags().StringVar(&ipAddr, "ip", "", "set ip and mask for your machine")
	machineCreateCmd.Flags().StringVar(&name, "name", "", "Assign a name to the machine")

	return machineCreateCmd
}

func machineCreateHandle(cmd *cobra.Command, args []string) {

	imageId := args[0]

	//判断是id还是name
	imageId = image.CheckNameOrId(imageId)

	//处理volumes
	//volumes := MountHandle(mountInfo)

	//ExtraHost
	//extraHostsParam := ExtraHostDeal(extraHostsValue)

	//get network info
	//networkInfo := network.GetNetworkInfo(bridge)
	//if networkInfo != nil {
	//	cidr = networkInfo.Config.Subnet
	//	gateway = networkInfo.Config.Gateway
	//	bridge = networkInfo.Driver
	//	if ipAddr == "" {
	//		ip, err := network.AllocateIP(bridge)
	//		if err != nil {
	//			fmt.Println("ip allocate failed")
	//			return
	//		}
	//
	//		if err := network.RegisterIP(bridge, ip); err != nil {
	//			fmt.Println("ip register failed")
	//			return
	//		}
	//
	//		defer func() {
	//			if err := network.ReleaseIP(bridge, ip); err != nil {
	//				fmt.Println("ip release failed")
	//			}
	//		}()
	//
	//		ipAddr = ip.String()
	//	}
	//}

	//cidr
	//networkAddress := []*pb.NetworkAddress{}
	//mask := 0
	//
	//if cidr != "" {
	//	reg, err := regexp.Compile(`([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})/([0-9]{0,2})`)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	subStr := reg.FindAllStringSubmatch(cidr, 1)[0]
	//
	//	mask, err = strconv.Atoi(subStr[2])
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	networkAddress = append(networkAddress, &pb.NetworkAddress{
	//		Ip:   ipAddr,
	//		Mask: int32(mask),
	//	})
	//}
	//
	//// interface
	//var interfaceSlice []*pb.NetworkInterface
	//
	////network interfaces
	//if !(bridge == "" || gateway == "" || len(networkAddress) == 0) {
	//	networkInterfaces := pb.NetworkInterface{
	//		Name:    "",
	//		Bridge:  bridge,
	//		Mac:     mac,
	//		Gateway: gateway,
	//		Address: networkAddress,
	//	}
	//	interfaceSlice = append(interfaceSlice, &networkInterfaces)
	//} else if (bridge == "" || gateway == "" || len(networkAddress) == 0) && !(bridge == "" && gateway == "" && len(networkAddress) == 0) {
	//	PrintmachineNetworkError()
	//	return
	//} else {
	//}

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

	r, err := connection.Client.CreateMachine(context.Background(), &pb.CreateMachineReq{
		ImageId: imageId,
		Name:    name,
		//Network: &pb.Network{
		//	Hostname:   hostname,
		//	ExtraHosts: extraHostsParam,
		//	Dns:        dns,
		//	Interfaces: interfaceSlice,
		//},
		Env: env,
		//Volumes:    volumes,
		//StopSignal: stopSignal,
		Tty:        tty,
		Cmd:        command,
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
//
//func MountHandle(mountStrList []string) (volumes []*pb.machineVolume) {
//
//	for _, mountStr := range mountStrList {
//		mountSlice := strings.Split(mountStr, ",")
//
//		verifyMap := map[string]string{
//			"name":   "",
//			"source": "",
//			"target": "",
//			"ro":     "",
//		}
//
//		readOnly := false
//
//		for _, val := range mountSlice {
//			if strings.Contains(val, "=") {
//				kv := strings.Split(val, "=")
//				k := strings.TrimSpace(kv[0])
//				v := strings.TrimSpace(kv[1])
//				if _, ok := verifyMap[k]; ok {
//					verifyMap[k] = v
//				}
//			} else {
//				if val == "ro" {
//					readOnly = true
//				}
//			}
//		}
//
//		volumes = append(volumes, &pb.machineVolume{
//			Name:        verifyMap["name"],
//			Source:      verifyMap["source"],
//			Destination: verifyMap["target"],
//			Readonly:    readOnly,
//		})
//
//	}
//	return
//}

//func ExtraHostDeal(extraHosts []string) []*pb.HostEntry {
//	var extra []*pb.HostEntry
//
//	for _, val := range extraHosts {
//		hostsList := strings.Split(val, ":")
//		if len(hostsList) < 2 {
//			panic("wrong format of the extra hosts")
//		}
//
//		extra = append(extra, &pb.HostEntry{
//			Ip:   hostsList[1],
//			Host: hostsList[0],
//		})
//	}
//	return extra
//}

func PrintmachineNetworkError() {
	fmt.Println("islands: please check bridge, gateway and network address, you may lost one of them.")
	fmt.Println("See `islands machine create --help`.")
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