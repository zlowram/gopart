package main

import (
	"fmt"

	"github.com/zlowram/gopart"
)

func main() {
	fmt.Println("Is someone debugging me..?")
	if gopart.BeingDebugged() {
		fmt.Println("Oh yes, sir...")
	} else {
		fmt.Println("Nope :-)")
	}
}
