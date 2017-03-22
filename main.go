package main

import (
	"fmt"
	_ "github.com/cancerballs/test1"
	handlers "github.com/cancerballs/test1/handlers"
	utils "github.com/cancerballs/test1/utils"
	"net/http"
)

func main() {
	// start http server
	go http.ListenAndServe(":9999", nil)
	// start tcp server
	go handlers.TcpServer()

	for {
		fmt.Println("channel result: ", <-utils.Channel_cmd)
	}
}
