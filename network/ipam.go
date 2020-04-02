package network

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path"
	"strings"
)

func (ipam *IPAM) Init() error {
	// 存放网段中地址分配信息的数组
	ipam.Subnets = make(map[string]string)

	// 从文件中加载已经分配的网段信息
	err := ipam.load()
	if err != nil {
		fmt.Printf("Error init ipam info, %v\n", err)
	}
	return err
}

func (ipam *IPAM) releaseSubnet(subnet string) error {
	if _, ok := ipam.Subnets[subnet]; ok {
		delete(ipam.Subnets, subnet)
		return nil
	}
	return fmt.Errorf("can't find the subnet")
}

func (ipam *IPAM) getCounter(subnet string) (counter int) {
	counter = 0
	if str, ok := ipam.Subnets[subnet]; ok {
		for c := range str {
			if str[c] == '1' {
				counter ++
			}
		}
	}
	return
}

func (ipam *IPAM) load() error {
	if _, err := os.Stat(ipam.SubnetAllocatorPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}
	subnetConfigFile, err := os.Open(ipam.SubnetAllocatorPath)
	defer subnetConfigFile.Close()
	if err != nil {
		return err
	}
	subnetJson := make([]byte, 2000)
	n, err := subnetConfigFile.Read(subnetJson)
	if err != nil {
		return err
	}

	err = json.Unmarshal(subnetJson[:n], &(ipam.Subnets))
	if err != nil {
		return fmt.Errorf("Error load allocation info, %v", err)
	}
	return nil
}

func (ipam *IPAM) dump() error {
	ipamConfigFileDir, _ := path.Split(ipam.SubnetAllocatorPath)
	if _, err := os.Stat(ipamConfigFileDir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(ipamConfigFileDir, 0644)
		} else {
			return err
		}
	}
	subnetConfigFile, err := os.OpenFile(ipam.SubnetAllocatorPath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	defer subnetConfigFile.Close()
	if err != nil {
		return err
	}

	ipamConfigJson, err := json.Marshal(ipam.Subnets)
	if err != nil {
		return err
	}

	_, err = subnetConfigFile.Write(ipamConfigJson)
	if err != nil {
		return err
	}

	return nil
}

func (ipam *IPAM) allocate(subnet *net.IPNet) (ip net.IP, err error) {

	_, subnet, _ = net.ParseCIDR(subnet.String())

	one, size := subnet.Mask.Size()

	if _, exist := (ipam.Subnets)[subnet.String()]; !exist {
		(ipam.Subnets)[subnet.String()] = strings.Repeat("0", 1<<uint8(size-one))
	}

	for c := range (ipam.Subnets)[subnet.String()] {
		if (ipam.Subnets)[subnet.String()][c] == '0' {
			ipalloc := []byte((ipam.Subnets)[subnet.String()])
			ipalloc[c] = '1'
			(ipam.Subnets)[subnet.String()] = string(ipalloc)
			ip = subnet.IP
			for t := uint(4); t > 0; t -= 1 {
				[]byte(ip)[4-t] += uint8(c >> ((t - 1) * 8))
			}
			ip[3] += 1
			break
		}
	}

	ipam.dump()
	return
}

func (ipam *IPAM) release(subnet *net.IPNet, ipaddr *net.IP) error {

	_, subnet, _ = net.ParseCIDR(subnet.String())
	c := 0
	releaseIP := ipaddr.To4()
	releaseIP[3] -= 1
	defer func() {
		releaseIP[3] += 1
	}()
	for t := uint(4); t > 0; t -= 1 {
		c += int(releaseIP[t-1]-subnet.IP[t-1]) << ((4 - t) * 8)
	}
	ipalloc := []byte(ipam.Subnets[subnet.String()])
	ipalloc[c] = '0'
	ipam.Subnets[subnet.String()] = string(ipalloc)

	return ipam.dump()
}

func (ipam *IPAM) register(subnet *net.IPNet, ipaddr net.IP) error {
	_, subnet, _ = net.ParseCIDR(subnet.String())

	c := 0
	registerIP := ipaddr.To4()
	registerIP[3] -= 1
	defer func() {
		registerIP[3] += 1
	}()
	for t := uint(4); t > 0; t -= 1 {
		c += int(registerIP[t-1]-subnet.IP[t-1]) << ((4 - t) * 8)
	}
	ipalloc := []byte((ipam.Subnets)[subnet.String()])
	if ipalloc[c] == '1' {
		return fmt.Errorf("failed to register : The IP has been registered")
	}
	ipalloc[c] = '1'
	(ipam.Subnets)[subnet.String()] = string(ipalloc)
	return ipam.dump()
}
