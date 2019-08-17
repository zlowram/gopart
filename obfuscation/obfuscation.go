package obfuscation

import (
	"syscall"
	"unsafe"

	windows "github.com/zlowram/gowin"
)

type WindowsApi struct {
	Hash    func(string) string
	Decrypt func(string) string
}

func NewWindowsApi(hash, decrypt func(string) string) *WindowsApi {
	return &WindowsApi{
		Hash:    hash,
		Decrypt: decrypt,
	}
}

func (w *WindowsApi) Call(module string, function string, params ...uintptr) (ret int32, err error) {
	// Get the function address
	functionAddr, err := w.windowsAPIFunction(module, function)
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

func (w *WindowsApi) windowsAPIFunction(moduleName, functionName string) (addr uintptr, err error) {
	module, err := w.loadModule(moduleName)
	if err != nil {
		return uintptr(0), err
	}
	exports, err := module.Exports()
	if err != nil {
		return uintptr(0), err
	}
	hashedExports := make(map[string]*windows.Export, len(exports))
	for name, export := range exports {
		hashedExports[w.Hash(name)] = export
	}
	return uintptr(hashedExports[functionName].Addr), nil
}

func (w *WindowsApi) loadModule(name string) (module *windows.Module, err error) {
	peb := windows.PebAddress()

	moduleName := w.Decrypt(name)

	// Check if already loaded
	module, err = peb.Module(moduleName)
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
	moduleToLoad := windows.NewUnicodeString(moduleName)
	r0, _, err := syscall.Syscall(uintptr(loadLibrary.Addr), 1, uintptr(unsafe.Pointer(moduleToLoad.Buffer)), 0, 0)
	addr := uint64(r0)
	if addr == 0 {
		return nil, err
	}

	return windows.NewModule(moduleName, addr), nil
}
