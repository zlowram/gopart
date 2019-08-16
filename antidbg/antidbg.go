package antidbg

import (
	"fmt"
	"unsafe"

	"github.com/zlowram/gopart"
	windows "github.com/zlowram/gowin"
)

func BeingDebugged() bool {
	peb := windows.PebAddress()
	return peb.BeingDebugged == 1
}

func NtQueryInformationProcess() (debugged bool, err error) {
	handle, err := gopart.WindowsAPICall(
		"kernel32.dll",
		"GetCurrentProcess",
	)
	if err != nil {
		return false, err
	}

	var processDebugPort int
	ntstatus, err := gopart.WindowsAPICall(
		"ntdll.dll",
		"NtQueryInformationProcess",
		uintptr(handle),
		uintptr(windows.ProcessDebugPort),
		uintptr(unsafe.Pointer(&processDebugPort)),
		uintptr(unsafe.Sizeof(processDebugPort)),
		uintptr(0),
	)
	if ntstatus != 0 {
		return false, fmt.Errorf("NTSTATUS: %x", ntstatus)
	}
	return processDebugPort == -1, nil
}
