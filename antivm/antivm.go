package antivm

import (
	"math"
	"unsafe"

	"github.com/zlowram/gopart"
	windows "github.com/zlowram/gowin"
)

func NumberOfProcessors() (ret bool, err error) {
	var systemInfo windows.SystemInfo
	_, err = gopart.WindowsAPICall(
		"kernel32.dll",
		"GetSystemInfo",
		uintptr(unsafe.Pointer(&systemInfo)),
	)
	return systemInfo.NumberOfProcessors < 2, nil
}

func PhysicalMemory() (ret bool, err error) {
	var memory windows.MemoryStatusEx
	memory.Length = uint32(unsafe.Sizeof(memory))
	_, err = gopart.WindowsAPICall(
		"kernel32.dll",
		"GlobalMemoryStatusEx",
		uintptr(unsafe.Pointer(&memory)),
	)
	return memory.TotalPhys/uint64(math.Pow(2, 30)) < 5, nil
}
