package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/codecrafters-io/redis-starter-go/app/serializer"
	"github.com/codecrafters-io/redis-starter-go/app/server"
	"os"
)

func main() {
	fmt.Println("Starting server")
	// Implement the command handler that have all the possible command
	// Implement Request handler that accept the commandhandler
	var parser = command.NewRedisCommandParser()
	var serializer = serializer.NewRESPSerializer(parser)
	var server server.RedisServer = server.NewDefaultTCPServer("0.0.0.0", 6379, serializer)
	err := server.Start()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Server terminated.")
}
