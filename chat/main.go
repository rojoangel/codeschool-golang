package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("I am the guest")
	} else {
		fmt.Println("I am the host")
	}
}
