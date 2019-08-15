package gopart

import (
	"errors"
	"fmt"
	"unsafe"

	windows "github.com/zlowram/gowin"
)

func BeingDebugged() bool {
	peb := windows.PebAddress()
	return peb.BeingDebugged == 1
}

func NtQueryInformationProcess() (debugged bool, err error) {
	handle, err := WindowsAPICall(
		"kernel32.dll",
		"GetCurrentProcess",
	)
	if err != nil {
		return false, err
	}

	var processDebugPort int
	ntstatus, err := WindowsAPICall(
		"ntdll.dll",
		"NtQueryInformationProcess",
		uintptr(handle),
		uintptr(windows.ProcessDebugPort),
		uintptr(unsafe.Pointer(&processDebugPort)),
		uintptr(unsafe.Sizeof(processDebugPort)),
		uintptr(0),
	)
	if ntstatus != 0 {
		return false, errors.New(fmt.Sprint(ntstatus))
	}
	return processDebugPort == -1, nil
}
