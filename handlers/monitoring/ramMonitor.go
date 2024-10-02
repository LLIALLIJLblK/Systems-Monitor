package monitoring

import(
	"github.com/shirou/gopsutil/mem"
)


func GetRAMUsage() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}
