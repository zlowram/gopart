package main

import (
	"fmt"

	"github.com/zlowram/gopart"
)

func main() {
	fmt.Println("Is someone debugging me..?")
	debugged, err := gopart.NtQueryInformationProcess()
	if err != nil {
		fmt.Println("Oops, an error ocurred!", err)
		return
	}
	if debugged {
		fmt.Println("Oh yes, sir...")
	} else {
		fmt.Println("Nope :-)")
	}
}
