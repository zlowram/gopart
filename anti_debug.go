package gopart

import (
	"errors"
	"fmt"
	"unsafe"

	windows "github.com/zlowram/gowin"
)

// Same as antivm. I would move these functions to a new package called
// antidbg (github.com/zlowram/gopart/antidbg).

// Missing docs.
func BeingDebugged() bool {
	peb := windows.PebAddress()
	return peb.BeingDebugged == 1
}

// Missing docs.
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
		// use fmt.Errorf() here
		return false, errors.New(fmt.Sprint(ntstatus))
	}
	return processDebugPort == -1, nil
}
