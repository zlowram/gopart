package gopart

import (
	"math"
	"unsafe"

	windows "github.com/zlowram/gowin"
)

// I don't like this naming. When called it looks like
// gowin.NumberOfProcessors(), which seems to return the number of processors.
// I would move these functions to a new package called "antivm" within gopart
// (github.com/zlowram/gopart/antivm). So, when called it would look like
// antivm.NumberOfProcessors() or antivm.PhysicalMemory() which makes much more
// sense to me.

// Missing docs.
func NumberOfProcessors() (ret bool, err error) {
	var systemInfo windows.SystemInfo
	_, err = WindowsAPICall(
		"kernel32.dll",
		"GetSystemInfo",
		uintptr(unsafe.Pointer(&systemInfo)),
	)
	return systemInfo.NumberOfProcessors < 2, nil
}

// Missing docs.
func PhysicalMemory() (ret bool, err error) {
	var memory windows.MemoryStatusEx
	memory.Length = uint32(unsafe.Sizeof(memory))
	_, err = WindowsAPICall(
		"kernel32.dll",
		"GlobalMemoryStatusEx",
		uintptr(unsafe.Pointer(&memory)),
	)
	return memory.TotalPhys/uint64(math.Pow(2, 30)) < 5, nil
}
