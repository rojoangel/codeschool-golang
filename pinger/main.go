package main

import (
	"fmt"
	"pinger/util"
)

func main() {
	for {
		if util.PingChatServer() {
			break
		}
	}
	fmt.Println("Ping is done")
}
