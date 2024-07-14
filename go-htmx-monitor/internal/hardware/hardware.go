package hardware

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

const megabyteDiv uint64 = 1024 * 1024
const gigabyteDiv uint64 = megabyteDiv * 1024

func GetSystemSection() (string, error) {
	runTimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	hostStat, err := host.Info()
	if err != nil {
		return "", err
	}

	html := "<div class='system-data'><table class='w-full text-sm bg-zinc-800 rounded-md'><tbody>"
	html = html + "<tr><td class='py-1 px-2'>Operating System:</td> <td class='py-1 px-2'><i class='fa fa-brands fa-linux'></i> " + runTimeOS + "</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Platform:</td><td class='py-1 px-2'> <i class='fa fa-brands fa-fedora'></i> " + hostStat.Platform + "</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Hostname:</td><td class='py-1 px-2'>" + hostStat.Hostname + "</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Kernel Arch:</td><td class='py-1 px-2'>" + hostStat.KernelArch + "</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Kernel Version:</td><td class='py-1 px-2'>" + hostStat.KernelVersion + "</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Boot Time:</td><td class='py-1 px-2'>" + time.Unix(int64(hostStat.BootTime), 0).String() + "</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Number of processes running:</td><td class='py-1 px-2'>" + strconv.FormatUint(hostStat.Procs, 10) + "</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Total memory:</td><td class='py-1 px-2'>" + strconv.FormatUint(vmStat.Total/megabyteDiv, 10) + " MB</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Free memory:</td><td class='py-1 px-2'>" + strconv.FormatUint(vmStat.Free/megabyteDiv, 10) + " MB</td></tr>"
	html = html + "<tr><td class='py-1 px-2'>Percentage used memory:</td><td class='py-1 px-2'>" + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%</td></tr></tbody></table>"
	html = html + "</div>"

	return html, nil
}

func GetDiskSection() (string, error) {

	partitionsStat, err := disk.Partitions(true)
	if err != nil {
		return "", err
	}

	html := "<div class='disk-data'><table class='w-full text-sm bg-zinc-800 rounded-md'><tbody>"

	for _, partition := range partitionsStat {
		diskStat, err := disk.Usage(partition.Device)
		if err != nil {
			return "", err
		}
		html = html + "<tr><td class='py-1 px-2'>Disk Path:</td><td class='py-1 px-2'>" + diskStat.Path + "</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Total disk space:</td><td class='py-1 px-2'>" + strconv.FormatUint(diskStat.Total/gigabyteDiv, 10) + " GB</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Used disk space:</td><td class='py-1 px-2'>" + strconv.FormatUint(diskStat.Used/gigabyteDiv, 10) + " GB</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Free disk space:</td><td class='py-1 px-2'>" + strconv.FormatUint(diskStat.Free/gigabyteDiv, 10) + " GB</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Percentage disk space usage:</td><td class='py-1 px-2'>" + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%</td></tr>"
	}

	return html, nil
}

func GetCpuSection() (string, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting CPU info", err)

	}
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		return "", err
	}

	html := "<div class='cpu-data'><table class='w-full text-sm bg-zinc-800 rounded-md'><tbody>"
	if len(cpuStat) != 0 {
		html = html + "<tr><td class='py-1 px-2'>Model Name:</td><td class='py-1 px-2'>" + cpuStat[0].ModelName + "</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Family:</td><td class='py-1 px-2'>" + cpuStat[0].Family + "</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Physical ID:</td><td class='py-1 px-2'>" + cpuStat[0].PhysicalID + "</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Vendor ID:</td><td class='py-1 px-2'>" + cpuStat[0].VendorID + "</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>CPU Cores:</td><td class='py-1 px-2'>" + strconv.Itoa(int(cpuStat[0].Cores)) + "</td></tr>"
		html = html + "<tr><td class='py-1 px-2'>Speed:</td><td class='py-1 px-2'>" + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz</td></tr>"
	}

	html = html + "<tr><td class='py-1 px-2'>Cores: </td><td class='py-1 px-2 grid grid-cols-2 gap-3'>"
	for idx, cpupercent := range percentage {
		html = html + "<div>CPU [" + strconv.Itoa(idx) + "]: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%</div>"
	}
	html = html + "</td></tr></tbody></table></div>"

	return html, nil
}
