package monitoring

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)


type SysMetrics struct{
	CPUUsage float64
	RamUsage *mem.VirtualMemoryStat
	DiskUsage *disk.UsageStat
	NetworkUsage []net.IOCountersStat
}

func GetSysMetrics()(*SysMetrics, error){
	cpuUsage, err := GetCPUUsage()
	if err != nil {
		return nil, err
	}

	ramUsage, err := GetRAMUsage()
	if err != nil {
		return nil, err
	}
	diskUsage, err := GetDiskUsage()
	if err != nil {
		return nil, err
	}

	networkUsage, err := GetNetworkUsage()
	if err != nil {
		return nil, err
	}

	return &SysMetrics{
		CPUUsage:     cpuUsage,
		RamUsage:     ramUsage,
		DiskUsage:    diskUsage,
		NetworkUsage: networkUsage,
	}, nil

}

func BytesToGB(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024 / 1024
}


func PrintSystemMetrics(metrics *SysMetrics) {
	fmt.Println("YOUR SYSTEM INFORMATION")
	fmt.Println("__________________________\n")

	fmt.Printf("CPU USAGE: %.2f%%\n", metrics.CPUUsage)
	fmt.Printf("MEMORY USAGE: %.2f%%\n", metrics.RamUsage.UsedPercent)
	fmt.Printf("DISK USAGE: %.2f%%\n", metrics.DiskUsage.UsedPercent)

	fmt.Println("__________________________\n")
	fmt.Printf("TOTAL RAM: %.2f GB\n", BytesToGB(metrics.RamUsage.Total))
	fmt.Printf("AVAILABLE RAM: %.2f GB\n", BytesToGB(metrics.RamUsage.Available))
	fmt.Printf("USED RAM: %.2f GB\n", BytesToGB(metrics.RamUsage.Used))
	fmt.Printf("FREE RAM: %.2f GB\n", BytesToGB(metrics.RamUsage.Free))
	fmt.Printf("ACTIVE RAM: %.2f GB\n", BytesToGB(metrics.RamUsage.Active))
	fmt.Printf("INACTIVE RAM: %.2f GB\n", BytesToGB(metrics.RamUsage.Inactive))
	fmt.Printf("WIRED RAM: %.2f GB\n", BytesToGB(metrics.RamUsage.Wired))
	fmt.Println("__________________________\n")
	fmt.Printf("TOTAL DISK: %.2f GB\n", BytesToGB(metrics.DiskUsage.Total))
	fmt.Printf("FREE DISK: %.2f GB\n", BytesToGB(metrics.DiskUsage.Free))
	fmt.Printf("USED DISK: %.2f GB\n", BytesToGB(metrics.DiskUsage.Used))
	fmt.Println("__________________________\n")
	fmt.Printf("TOTAL NETWORK: %.2f GB\n", BytesToGB(metrics.NetworkUsage[0].BytesSent))
}