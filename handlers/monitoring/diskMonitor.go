package monitoring
import(
	"github.com/shirou/gopsutil/disk"
)

func GetDiskUsage() (*disk.UsageStat, error) {
	return disk.Usage("/")
}
