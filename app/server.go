package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type RedisServer interface {
	Start() error
}

type DefaultTCPServer struct {
	port int
	host string
}

func NewDefaultTCPServer(host string, port int) *DefaultTCPServer {
	return &DefaultTCPServer{port: port, host: host}
}

func (ds *DefaultTCPServer) Start() error {
	var address = ds.host + ":" + strconv.Itoa(ds.port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	_, err = l.Accept()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Logs from your program will appear here!")
	var server RedisServer = NewDefaultTCPServer("0.0.0.0", 6379)
	err := server.Start()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
