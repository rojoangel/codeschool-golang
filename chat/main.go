package main

import (
	"chat/util"
	"os"
)

func main() {

	args := os.Args
	if len(args) > 1 {

		util.RunGuest(hostIP)
	} else {

		util.RunHost(listenIP + ":" + listenPort)
	}
}
