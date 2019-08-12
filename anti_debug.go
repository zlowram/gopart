package gopart

import (
	windows "github.com/zlowram/gowin"
)

func BeingDebugged() bool {
	peb := windows.PebAddress()
	return peb.BeingDebugged == 1
}
