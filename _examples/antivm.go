package main

import (
	"fmt"

	"github.com/zlowram/gopart/antivm"
)

func main() {
	fmt.Printf("Number of processors < 2: ")
	detected, err := antivm.NumberOfProcessors()
	if err != nil {
		fmt.Println("Oops, an error ocurred!", err)
		return
	}
	if detected {
		fmt.Println("VM detected!")
	} else {
		fmt.Println("VM not detected")
	}

	fmt.Printf("Physical memory < 5GB: ")
	detected, err = antivm.PhysicalMemory()
	if err != nil {
		fmt.Println("Oops, an error ocurred!", err)
		return
	}
	if detected {
		fmt.Println("VM detected!")
	} else {
		fmt.Println("VM not detected")
	}
}
