package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/server"
	"os"
)

func main() {
	fmt.Println("Starting server")
	var server server.RedisServer = server.NewDefaultTCPServer("0.0.0.0", 6379)
	err := server.Start()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Server terminated.")
}
