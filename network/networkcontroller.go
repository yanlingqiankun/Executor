package network

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"github.com/vishvananda/netlink"
	"github.com/yanlingqiankun/Executor/conf"
	"github.com/yanlingqiankun/Executor/logging"
	"github.com/yanlingqiankun/Executor/util"
	"net"
	"os"
	"path"
	"path/filepath"
	"text/tabwriter"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05.999999999 -0700 MST"

var (
	defaultNetworkPath string
	bridges            = map[string]NetworkDriver{}
	networks           = map[string]*Network{}
	ipAllocator        = &IPAM{}
)

var logger = logging.GetLogger("network")

var libconn *libvirt.Connect

func init() {
	var err error
	networkPath := filepath.Join(conf.GetString("RootPath"), "Excutor-net")
	//defaultNetworkPath = filepath.Join(networkPath, "Network")
	defaultNetworkPath = filepath.Join(networkPath, "network")
	ipAllocator.SubnetAllocatorPath = filepath.Join(networkPath, "IPAM")
	libconn, err = libvirt.NewConnect("qemu:///system")
	if err != nil {
		logger.Fatalf("failed to connect to qemu")
	}
	if err = Init(); err != nil {
		logger.Fatal("failed to init network : ", err.Error())
	}
	if err = ipAllocator.Init(); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Debug("network init")
}

func (nw *Network) dump(dumpPath string) error {
	if _, err := os.Stat(dumpPath); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dumpPath, 0600)
		} else {
			return err
		}
	}

	nwPath := path.Join(dumpPath, nw.Name)
	nwFile, err := os.OpenFile(nwPath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("dump error：", err)
	}
	defer nwFile.Close()

	nwJson, err := json.Marshal(nw)
	if err != nil {
		return fmt.Errorf("dump error：", err)
	}

	_, err = nwFile.Write(nwJson)
	if err != nil {
		return fmt.Errorf("dump error：", err)
	}
	return nil
}

func (nw *Network) remove(dumpPath string) error {
	if _, err := os.Stat(path.Join(dumpPath, nw.Name)); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	} else {
		return os.Remove(path.Join(dumpPath, nw.Name))
	}
}

func (nw *Network) load(dumpPath string) error {
	nwConfigFile, err := os.Open(dumpPath)
	defer nwConfigFile.Close()
	if err != nil {
		return err
	}
	nwJson := make([]byte, 2000)
	n, err := nwConfigFile.Read(nwJson)
	if err != nil {
		return err
	}

	err = json.Unmarshal(nwJson[:n], nw)
	if err != nil {
		return fmt.Errorf("Error load network info", err)
	}
	return nil
}

func Init() error {
	if _, err := os.Stat(defaultNetworkPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(defaultNetworkPath, 0600)
		} else {
			return err
		}
	}

	filepath.Walk(defaultNetworkPath, func(nwPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		_, nwName := path.Split(nwPath)
		nw := &Network{
			Name: nwName,
		}

		if err := nw.load(nwPath); err != nil {
			return fmt.Errorf("error load network: %s", err)
		}
		networks[nw.Name] = nw
		if bridge, _ := libconn.LookupNetworkByName(nw.Driver); bridge != nil {
			return nil
		}
		return nw.createBridge()
	})
	return nil
}

func CreateNetwork(subnet, name, gateway string) error {
	if _, ok := networks[name]; ok {
		return fmt.Errorf("The network exists")
	}
	_, cidr, err := net.ParseCIDR(subnet)
	if err != nil {
		return fmt.Errorf("Can't get the subnet")
	}
	var gatewayIP net.IP
	if gateway == "" {
		tmp, err := ipAllocator.allocate(cidr)
		if err != nil {
			logger.Error("failed to alloc ip for network")
			return err
		}
		gatewayIP = tmp
	} else {
		IP := net.ParseIP(gateway)
		if IP == nil {
			return fmt.Errorf("The gateway is invalid")
		}
		if !cidr.Contains(IP) {
			return fmt.Errorf("The gateway not in the subnet")
		}
		gatewayIP = IP
	}
	nw := &Network{
		Name:       name,
		Subnet:     cidr,
		Driver:     name,
		CreateTime: time.Now(),
		GateWay:    gatewayIP,
	}
	networks[name] = nw
	if err := nw.dump(defaultNetworkPath); err != nil {
		return fmt.Errorf("failed to save network information :", err.Error())
	}
	if link, _ := netlink.LinkByName(nw.Driver); link != nil {
		return fmt.Errorf("Warning", nw.Name, " exists")
	}
	return nw.createBridge()
}

func listNetwork() {
	w := tabwriter.NewWriter(os.Stdout, 12, 4, 3, ' ', 0)
	fmt.Fprint(w, "NAME\tSubnet\tBridge\tGateway\n")
	for _, nw := range networks {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			nw.Name,
			nw.Subnet.String(),
			nw.Driver,
			nw.GateWay.String(),
		)
	}
	if err := w.Flush(); err != nil {
		fmt.Println("Flush error %v", err)
		return
	}
}

func DeleteNetwork(networkName string, force bool) error {
	nw, ok := networks[networkName]
	if !ok {
		return fmt.Errorf("No Such Network: %s", networkName)
	}
	if !force {
		if ipAllocator.getCounter(nw.Subnet.String()) > 1 {
			return fmt.Errorf("can remove a network in use, please use -f to forcely remove it")
		}
	}

	if err := deleteBridge(networks[networkName]); err != nil {
		return fmt.Errorf("Error Remove Network DriverError: %s", err)
	}
	if err := ipAllocator.releaseSubnet(nw.Subnet.String()); err != nil {
		return err
	}

	return nw.remove(defaultNetworkPath)
}

func inspectNetwork(networkName string) error {
	nw, ok := networks[networkName]
	if !ok {
		return fmt.Errorf("No Such Network: %s", networkName)
	}
	networkinfo := NetworkInfo{
		Name:       nw.Name,
		Driver:     nw.Driver,
		CreateTime: nw.CreateTime.Format(TIME_LAYOUT),
		Config: NetworkConf{
			Gateway: nw.GateWay.String(),
			Subnet:  nw.Subnet.String(),
		},
	}
	output, err := json.MarshalIndent(networkinfo, " ", "\t")
	if err != nil {
		return fmt.Errorf("failed to inspect network with error: %v", err)
	}
	fmt.Println(string(output))
	return nil
}

//返回一个network的信息，如果没有则返回nil
func GetNetworkInfo(netname string) *NetworkInfo {
	nw, ok := networks[netname]
	if !ok {
		return nil
	}
	return &NetworkInfo{
		Name:       nw.Name,
		Driver:     nw.Driver,
		CreateTime: nw.CreateTime.Format(TIME_LAYOUT),
		Config: NetworkConf{
			Gateway: nw.GateWay.String(),
			Subnet:  nw.Subnet.String(),
		},
	}
}

func AllocateIP(netname string) (net.IP, error) {
	nw, ok := networks[netname]
	if !ok {
		return nil, fmt.Errorf("The network is not exists")
	}
	return ipAllocator.allocate(nw.Subnet)
}

func RegisterIP(netname string, VMName string, ipaddr net.IP) error {
	//check if the ip valid
	nw, ok := networks[netname]
	if !ok {
		return fmt.Errorf("The network is not exists")
	}
	if !nw.Subnet.Contains(ipaddr) {
		return fmt.Errorf("The IP is invalid")
	}

	libnet, err := libconn.LookupNetworkByName(netname)
	if err != nil {
		return err
	}
	libhost := libvirtxml.NetworkDHCPHost{
		XMLName: xml.Name{},
		ID:      util.GetBytesSha256([]byte(VMName)),
		MAC:     "",
		Name:    VMName,
		IP:      ipaddr.String(),
	}
	hostStr, err := libhost.Marshal()
	if err != nil {
		return err
	}
	err = libnet.Update(libvirt.NETWORK_UPDATE_COMMAND_ADD_FIRST,libvirt.NETWORK_SECTION_IP_DHCP_HOST, -1,hostStr,libvirt.NETWORK_UPDATE_AFFECT_LIVE | libvirt.NETWORK_UPDATE_AFFECT_CONFIG)
	if err != nil {
		return err
	}
	return ipAllocator.register(nw.Subnet, ipaddr)
}

func ReleaseIP(netname string, VMName string, ipaddr net.IP) error {
	//check if the ip valid
	nw, ok := networks[netname]
	if !ok {
		return fmt.Errorf("The network is not exists")
	}
	if !nw.Subnet.Contains(ipaddr) {
		return fmt.Errorf("The IP is invalid")
	}
	libhost := libvirtxml.NetworkDHCPHost{
		XMLName: xml.Name{},
		ID:      util.GetBytesSha256([]byte(VMName)),
		MAC:     "",
		Name:    VMName,
		IP:      ipaddr.String(),
	}
	hostStr, err := libhost.Marshal()
	if err != nil {
		return err
	}
	libnet, err := libconn.LookupNetworkByName(netname)
	if err != nil {
		return err
	}
	err = libnet.Update(libvirt.NETWORK_UPDATE_COMMAND_DELETE,libvirt.NETWORK_SECTION_IP_DHCP_HOST, -1,hostStr,libvirt.NETWORK_UPDATE_AFFECT_LIVE | libvirt.NETWORK_UPDATE_AFFECT_CONFIG)
	if err != nil {
		return err
	}
	return ipAllocator.release(nw.Subnet, &ipaddr)
}

func SetRoute() (err error) {
	panic("set route")
	return
}
