package main

import (
	"unsafe"

	"github.com/zlowram/gopart"
	windows "github.com/zlowram/gowin"
)

func main() {
	messageTitle := windows.NewUnicodeString("GOPART")
	message := windows.NewUnicodeString("This call to MessageBoxW was done by finding the function address via PEB")
	gopart.WindowsAPICall(
		"user32.dll",
		"MessageBoxW",
		uintptr(0),
		uintptr(unsafe.Pointer(message.Buffer)),
		uintptr(unsafe.Pointer(messageTitle.Buffer)),
		uintptr(0),
	)
}
