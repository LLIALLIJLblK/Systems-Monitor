package monitoring

import(
	"github.com/shirou/gopsutil/net"
)


func GetNetworkUsage() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}
