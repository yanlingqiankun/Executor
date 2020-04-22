package network

import (
	"encoding/xml"
	"fmt"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"net"
	"strings"
)

func (d *BridgeNetworkDriver) Name() string {
	return d.BridgeName
}

func (d *BridgeNetworkDriver) Create(IpRange *net.IPNet) error {
	err := d.initBridge(IpRange, d.IsIsolated)
	if err != nil {
		return fmt.Errorf("error init bridge: %v", err)
	}

	return err
}

func (nw *Network) createBridge() error {
	bridges[nw.Driver] = &BridgeNetworkDriver{
		BridgeName: nw.Driver,
		IP:         nw.GateWay,
		Mask:       nw.Subnet.Mask,
		IsIsolated: nw.IsIsolated,
	}
	return bridges[nw.Name].Create(nw.Subnet)
}

func (d *BridgeNetworkDriver) initBridge(IpRange *net.IPNet, isIsolated bool) error {
	// try to get bridge by name, if it already exists then just exit
	bridgeName := d.BridgeName
	if err := createBridgeInterface(bridgeName, d.IP.String(), net.IP(d.Mask).String(), isIsolated); err != nil {
		return fmt.Errorf("Error add bridge： %s, Error: %v", bridgeName, err)
	}

	//// Set bridge IP
	//gatewayIP := &net.IPNet{
	//	IP:   d.IP,
	//	Mask: d.Mask,
	//}
	//
	//if err := setInterfaceIP(bridgeName, gatewayIP.String()); err != nil && !strings.Contains(err.Error(), "exists") {
	//	return fmt.Errorf("Error assigning address: %s on bridge: %s with an error of: %v", gatewayIP, bridgeName, err)
	//}
	//if err := setInterfaceUP(bridgeName); err != nil {
	//	return fmt.Errorf("Error set bridge up: %s, Error: %v", bridgeName, err)
	//}
	//
	//if err := setRoute(d.BridgeName, IpRange); err != nil && !strings.Contains(err.Error(), "exists") {
	//	return fmt.Errorf("Error add route: %v", err)
	//}

	// Setup iptables
	//if err := setupIPTables(bridgeName, IpRange); err != nil {
	//	return fmt.Errorf("Error setting iptables for %s: %v", bridgeName, err)
	//}

	return nil
}

// deleteBridge deletes the bridge
func deleteBridge(n *Network) error {
	bridgeName := n.Driver
	l, err := libconn.LookupNetworkByName(bridgeName)
	if err != nil {
		return fmt.Errorf("Getting link with name %s failed: %v", bridgeName, err)
	}
	if ok, _ := l.IsActive(); ok {
		if err := l.Destroy(); err != nil {
			return err
		}
	}
	if ok, _ := l.IsPersistent(); ok {
		if err := l.Undefine(); err != nil {
			return err
		}
	}
	return nil
}

func createBridgeInterface(bridgeName, ip, mask string, isIsolated bool) error {
	_, err := net.InterfaceByName(bridgeName)
	if err == nil || !strings.Contains(err.Error(), "no such network interface") {
		return err
	}

	netXML := &libvirtxml.Network{
		XMLName:             xml.Name{},
		IPv6:                "",
		TrustGuestRxFilters: "",
		Name:                bridgeName,
		UUID:                "",
		Metadata:            nil,

		Bridge:              &libvirtxml.NetworkBridge{
			Name:            bridgeName,
			STP:             "on",
			Delay:           "",
			MACTableManager: "",
			Zone:            "",
		},
		//Routes: append(make([]libvirtxml.NetworkRoute,0), libvirtxml.NetworkRoute{
		//	Family:  "ipv4",
		//	Address: "0.0.0.0",
		//	Netmask: "",
		//	Prefix:  0,
		//	Gateway: ip,
		//	Metric:  "",
		//}),
	}


	if !isIsolated {
		netXML.Forward = &libvirtxml.NetworkForward{
			Mode:       "nat",
		}
	} else {

	}
	netXML.IPs = make([]libvirtxml.NetworkIP, 1)
	netXML.IPs[0] = libvirtxml.NetworkIP{
		Address:  ip,
		Family:   "",
		Netmask:  mask,
		Prefix:   0,
		LocalPtr: "",
		DHCP:     &libvirtxml.NetworkDHCP{
			Ranges: nil,
			Hosts:  nil,
			Bootp:  nil,
		},
		TFTP:     nil,
	}

	netStr, err := netXML.Marshal()
	if err != nil {
		return err
	}
	libnet, err := libconn.NetworkDefineXML(netStr)
	if err != nil {
		return err
	}
	if err := libnet.Create(); err != nil {
		return err
	}
	return libnet.SetAutostart(true)
}
//
//func setInterfaceUP(interfaceName string) error {
//
//	iface, err := netlink.LinkByName(interfaceName)
//	if err != nil {
//		return fmt.Errorf("Error retrieving a link named [ %s ]: %v", iface.Attrs().Name, err)
//	}
//
//	if err := netlink.LinkSetUp(iface); err != nil {
//		return fmt.Errorf("Error enabling interface for %s: %v", interfaceName, err)
//	}
//	return nil
//}
//
//// Set the IP addr of a netlink interface
//func setInterfaceIP(name string, rawIP string) error {
//	retries := 2
//	var iface netlink.Link
//	var err error
//	for i := 0; i < retries; i++ {
//		iface, err = netlink.LinkByName(name)
//		if err == nil {
//			break
//		}
//		fmt.Printf("error retrieving new bridge netlink link [ %s ]... retrying\n", name)
//		time.Sleep(2 * time.Second)
//	}
//	if err != nil {
//		return fmt.Errorf(
//			"Abandoning retrieving the new bridge link from netlink, Run [ ip link ] to troubleshoot the error: %v", err)
//	}
//	ipNet, err := netlink.ParseIPNet(rawIP)
//	if err != nil {
//		return err
//	}
//	addr := &netlink.Addr{IPNet: ipNet, Label: ""}
//	return netlink.AddrAdd(iface, addr)
//}
//
//func setRoute(bridgeName string, subnet *net.IPNet) error {
//	iface, err := netlink.LinkByName(bridgeName)
//	if err != nil {
//		return err
//	}
//	route := netlink.Route{
//		LinkIndex: iface.Attrs().Index,
//		Dst: &net.IPNet{
//			IP:   subnet.IP,
//			Mask: subnet.Mask,
//		},
//		Src: nil,
//	}
//	if err := netlink.RouteAdd(&route); err != nil {
//		return err
//	}
//	return nil
//}
//
//func setupIPTables(bridgeName string, IPRange *net.IPNet) error {
//	iptable, err := iptables.New()
//	if err != nil {
//		return fmt.Errorf("iptables has error : %v", err)
//	}
//	if err = iptable.Append("nat", "POSTROUTING",
//		"-s", IPRange.String(), "!", "-o", bridgeName, "-j", "MASQUERADE"); err != nil {
//		return fmt.Errorf("failed to set iptables with error :", err.Error())
//	}
//	if err = iptable.Append("filter", "FORWARD",
//		"-i", bridgeName, "-j", "ACCEPT"); err != nil {
//		return fmt.Errorf("failed to set iptables with error :", err.Error())
//	}
//	if err = iptable.Append("filter", "FORWARD",
//		"-o", bridgeName, "-j", "ACCEPT"); err != nil {
//		return fmt.Errorf("failed to set iptables with error :", err.Error())
//	}
//	return nil
//}
