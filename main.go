package main

import (
	"fmt"
	test1 "github.com/cancerballs/test1"
)

func main() {
	// Start the HTTP and TCP server
	go test1.SetupServers(":9999", ":8888")

	// each Listen represents a subscriber of the channel and receives the
	// same command
	for {
		test1.Listen(func(cmd test1.Cmd) { fmt.Println(1, cmd) })
		test1.Listen(func(cmd test1.Cmd) { fmt.Println(2, cmd) })
		test1.Listen(func(cmd test1.Cmd) { fmt.Println(3, cmd) })
	}
}
