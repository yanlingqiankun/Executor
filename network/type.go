package network

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"net"
	"time"
)

type Network struct {
	Name       string
	Subnet     *net.IPNet
	Driver     string
	CreateTime time.Time
	GateWay    net.IP
}

type NetworkDriver interface {
	Name() string
	Create(IpRange *net.IPNet) error
}

type BridgeNetworkDriver struct {
	BridgeName string
	IP         net.IP
	Mask       net.IPMask
}

type IPAM struct {
	SubnetAllocatorPath string
	Subnets             map[string]string
}

type NetworkConf struct {
	Gateway string
	Subnet  string
}

type NetworkInfo struct {
	Name       string
	Driver     string
	CreateTime string
	Config     NetworkConf
}

type inspectInfo struct {
	Name                string              `json:"name,omitempty"`
	Bridge              *libvirtxml.NetworkBridge      `json:"bridge"`
	MAC                 *libvirtxml.NetworkMAC         `json:"mac"`
	IPs                 []libvirtxml.NetworkIP         `json:"ip"`
}
