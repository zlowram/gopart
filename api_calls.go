package gopart

import (
	"fmt"
	"syscall"
	"unsafe"

	windows "github.com/zlowram/gowin"
)

// Missing docs.
func WindowsAPICall(module string, function string, params ...uintptr) (ret int32, err error) {
	// Get the function address
	functionAddr, err := windowsAPIFunction(module, function)
	if err != nil {
		return int32(-1), err
	}

	// Fill non-used parameters with 0's
	numberOfParams := len(params)
	additionalParams := 6 - numberOfParams
	for i := 0; i < additionalParams; i++ {
		params = append(params, 0)
	}

	// Call the function
	r0, _, err := syscall.Syscall6(functionAddr, uintptr(numberOfParams), params[0], params[1], params[2], params[3], params[4], params[5])
	ret = int32(r0)
	if ret == 0 {
		return ret, err
	}

	return ret, nil
}

func windowsAPIFunction(moduleName, functionName string) (addr uintptr, err error) {
	module, err := loadModule(moduleName)
	if err != nil {
		return uintptr(0), err
	}
	export, err := module.Export(functionName)
	if err != nil {
		return uintptr(0), err
	}
	return uintptr(export.Addr), nil
}

func loadModule(name string) (module *windows.Module, err error) {
	peb := windows.PebAddress()

	// Check if already loaded
	module, err = peb.Module(name)
	if err == nil {
		return module, nil
	}

	// Get the LoadLibraryW function address
	kernel32, err := peb.Module("kernel32.dll")
	if err != nil {
		return nil, err
	}
	loadLibrary, err := kernel32.Export("LoadLibraryW")
	if err != nil {
		return nil, err
	}

	// Call to LoadLibraryW to load the module
	moduleToLoad := windows.NewUnicodeString(name)
	r0, _, err := syscall.Syscall(uintptr(loadLibrary.Addr), 1, uintptr(unsafe.Pointer(moduleToLoad.Buffer)), 0, 0)
	addr := uint64(r0)
	if addr == 0 {
		// Never print in a package (unless you export a log.Logger)
		// Have you forgotten returning with error here?
		fmt.Println(err)
	}

	return windows.NewModule(name, addr), nil
}
