package main

import (
	"fmt"
	test1 "github.com/cancerballs/test1"
)

func main() {
	// each Listen represents a subscriber of the channel and receives the
	// same command
	test1.Listen(func(cmd test1.Cmd) { fmt.Println(1, cmd) })
	test1.Listen(func(cmd test1.Cmd) { fmt.Println(2, cmd) })
	test1.Listen(func(cmd test1.Cmd) { fmt.Println(3, cmd) })
	// Start the HTTP and TCP server
	test1.SetupServers(":9999", ":8888")
}
