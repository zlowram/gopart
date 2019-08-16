package main

import (
	"fmt"

	"github.com/zlowram/gopart/antidbg"
)

func main() {
	fmt.Printf("BeingDebugged flag from PEB: ")
	if antidbg.BeingDebugged() {
		fmt.Printf("Debugger detected!")
	} else {
		fmt.Printf("Debugger not detected")
	}

	fmt.Printf("NtQueryInformationProcess's ProcessDebugPort: ")
	detected, _ := antidbg.NtQueryInformationProcess()
	if detected {
		fmt.Printf("Debugger detected!")
	} else {
		fmt.Printf("Debugger not detected")
	}
}
