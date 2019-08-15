package main

import (
	"fmt"

	"github.com/zlowram/gopart"
)

func main() {
	fmt.Println("Am I being run in a VM..?")
	vm, err := gopart.PhysicalMemory()
	if err != nil {
		fmt.Println("Oops, an error ocurred!", err)
		return
	}
	if vm {
		fmt.Println("Oh yes, sir...")
	} else {
		fmt.Println("Nope :-)")
	}
}
