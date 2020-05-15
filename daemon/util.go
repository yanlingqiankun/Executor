package daemon

import (
	"crypto/rand"
	"fmt"
	"github.com/docker/docker/api/types/blkiodev"
	"github.com/docker/docker/api/types/container"
	"github.com/yanlingqiankun/Executor/pb"
	"golang.org/x/sys/unix"
)

func newErr(code uint32, err error) *pb.Error {
	if err == nil {
		return &pb.Error{
			Code:                 0,
			Message:              "",
		}
	}
	return &pb.Error{
		Code:    code,
		Message: err.Error(),
	}
}

func convertHostsFromPB(hosts []*pb.HostEntry) []string {
	machineHosts := make([]string, len(hosts))
	for index, entry := range hosts {
		machineHosts[index] = entry.Host + ":" + entry.Ip
	}
	return machineHosts
}

func getMac() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		logger.WithError(err).Error("failed to get a rand mac")
		return ""
	}
	// Set the local bit
	buf[0] |= 2
	return fmt.Sprintf("00:%02x:%02x:%02x:%02x:%02x",  buf[1], buf[2], buf[3], buf[4], buf[5])
}


func getMachineResources(config *pb.Resources) (container.Resources, error) {
	errorResult := container.Resources{}
	if config == nil {
		return errorResult, nil
	}
	readBpsDevice, err := getBlkioThrottleDevices(config.BlkioDeviceReadBps)
	if err != nil {
		return errorResult, err
	}
	writeBpsDevice, err := getBlkioThrottleDevices(config.BlkioDeviceWriteBps)
	if err != nil {
		return errorResult, err
	}
	readIOpsDevice, err := getBlkioThrottleDevices(config.BlkioDeviceReadIOps)
	if err != nil {
		return errorResult, err
	}
	writeIOpsDevice, err := getBlkioThrottleDevices(config.BlkioDeviceWriteIOps)
	if err != nil {
		return errorResult, err
	}
	weightDevices, err := getBlkioWeightDevices(config)
	if err != nil {
		return errorResult, err
	}

	result := container.Resources{
		CPUShares:            config.CPUShares,
		Memory:               config.Memory,
		NanoCPUs:             config.NanoCPUs,
		CgroupParent:         "",
		BlkioWeight:          uint16(config.BlkioWeight),
		BlkioWeightDevice:    weightDevices,
		BlkioDeviceReadBps:   readBpsDevice,
		BlkioDeviceWriteBps:  writeBpsDevice,
		BlkioDeviceReadIOps:  readIOpsDevice,
		BlkioDeviceWriteIOps: writeIOpsDevice,
		CPUPeriod:            config.CPUPeriod,
		CPUQuota:             config.CPUQuota,
		CPURealtimePeriod:    config.CPURealtimePeriod,
		CPURealtimeRuntime:   config.CPURealtimeRuntime,
		CpusetCpus:           config.CpusetCpus,
		CpusetMems:           config.CpusetMems,
		Devices:              nil,
		DeviceCgroupRules:    nil,
		DeviceRequests:       nil,
		KernelMemory:         config.KernelMemory,
		KernelMemoryTCP:      config.KernelMemoryTCP,
		MemoryReservation:    config.MemoryReservation,
		MemorySwap:           config.MemorySwap,
		MemorySwappiness:     nil,
		OomKillDisable:       nil,
		PidsLimit:            &config.PidsLimit,
		Ulimits:              nil,
	}

	//
	//pidLimits := getPidsLimit(config)
	//blkioWeight := uint16(config.BlkioWeight)
	//specResources := &specs.LinuxResources{
	//	Memory: memoryResources,
	//	CPU:    cpuResource,
	//	Pids:   pidLimits,
	//	BlockIO: &specs.LinuxBlockIO{
	//		Weight:                  &blkioWeight,
	//		WeightDevice:            weightDevices,
	//		ThrottleReadBpsDevice:   readBpsDevice,
	//		ThrottleWriteBpsDevice:  writeBpsDevice,
	//		ThrottleReadIOPSDevice:  readIOpsDevice,
	//		ThrottleWriteIOPSDevice: writeIOpsDevice,
	//	},
	//}

	return result, nil
}

func getBlkioThrottleDevices(devs []*pb.ThrottleDevice) ([]*blkiodev.ThrottleDevice, error) {
	var throttleDevices []*blkiodev.ThrottleDevice
	var stat unix.Stat_t

	for _, d := range devs {
		if err := unix.Stat(d.Path, &stat); err != nil {
			return nil, err
		}
		d := &blkiodev.ThrottleDevice{Rate: d.Rate, Path:d.Path}
		throttleDevices = append(throttleDevices, d)
	}

	return throttleDevices, nil
}
//
//func getMemoryResources(config *pb.Resources) *specs.LinuxMemory {
//	memory := specs.LinuxMemory{}
//
//	if config.Memory > 0 {
//		memory.Limit = &config.Memory
//	}
//
//	if config.MemoryReservation > 0 {
//		memory.Reservation = &config.MemoryReservation
//	}
//
//	if config.MemorySwap > 0 {
//		memory.Swap = &config.MemorySwap
//	}
//
//	if config.MemorySwappiness != 0 {
//		swappiness := uint64(config.MemorySwappiness)
//		memory.Swappiness = &swappiness
//	}
//
//	if config.OomKillDisable != false {
//		memory.DisableOOMKiller = &config.OomKillDisable
//	}
//
//	if config.KernelMemory != 0 {
//		memory.Kernel = &config.KernelMemory
//	}
//
//	if config.KernelMemoryTCP != 0 {
//		memory.KernelTCP = &config.KernelMemoryTCP
//	}
//
//	return &memory
//}
//
//func getPidsLimit(config *pb.Resources) *specs.LinuxPids {
//	if config.PidsLimit == 0 {
//		return nil
//	}
//	if config.PidsLimit <= 0 {
//		// 可以使用负数
//		// 但是底层 runtime tool 统一 -1
//		return &specs.LinuxPids{Limit: -1}
//	}
//	return &specs.LinuxPids{Limit: config.PidsLimit}
//}
//
//func getCPUResources(config *pb.Resources) (*specs.LinuxCPU, error) {
//	cpu := specs.LinuxCPU{}
//
//	if config.CPUShares < 0 {
//		return nil, fmt.Errorf("shares: invalid argument")
//	}
//	if config.CPUShares >= 0 {
//		shares := uint64(config.CPUShares)
//		cpu.Shares = &shares
//	}
//
//	if config.CpusetCpus != "" {
//		cpu.Cpus = config.CpusetCpus
//	}
//
//	if config.CpusetMems != "" {
//		cpu.Mems = config.CpusetMems
//	}
//
//	if config.NanoCPUs > 0 {
//		// https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt
//		period := uint64(100 * time.Millisecond / time.Microsecond)
//		quota := config.NanoCPUs * int64(period) / 1e9
//		cpu.Period = &period
//		cpu.Quota = &quota
//	}
//
//	if config.CPUPeriod != 0 {
//		period := uint64(config.CPUPeriod)
//		cpu.Period = &period
//	}
//
//	if config.CPUQuota != 0 {
//		q := config.CPUQuota
//		cpu.Quota = &q
//	}
//
//	if config.CPURealtimePeriod != 0 {
//		period := uint64(config.CPURealtimePeriod)
//		cpu.RealtimePeriod = &period
//	}
//
//	if config.CPURealtimeRuntime != 0 {
//		c := config.CPURealtimeRuntime
//		cpu.RealtimeRuntime = &c
//	}
//
//	return &cpu, nil
//}

func getBlkioWeightDevices(config *pb.Resources) ([]*blkiodev.WeightDevice, error) {
	var stat unix.Stat_t
	var blkioWeightDevices []*blkiodev.WeightDevice

	for _, weightDevice := range config.BlkioWeightDevice {
		if err := unix.Stat(weightDevice.Path, &stat); err != nil {
			return nil, err
		}
		weight := uint16(weightDevice.Weight)
		d := &blkiodev.WeightDevice{Weight: weight, Path:weightDevice.Path}
		blkioWeightDevices = append(blkioWeightDevices, d)
	}

	return blkioWeightDevices, nil
}
