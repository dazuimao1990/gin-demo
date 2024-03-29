package utils

import (
	"net"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetEnvDefault(key, def string) (res string) {
	res = os.Getenv(key)
	if res == "" {
		return def
	} else {
		return res
	}
}

func GetEniIP() (ip string) {
	addrs, err := net.InterfaceAddrs()
	var Ip []string
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				Ip = append(Ip, ipnet.IP.String())
			}
		}
	}
	return strings.Join(Ip[:], ",")
}

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetMemTotal() uint64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.Total / 1024 / 1024 / 1024
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}
