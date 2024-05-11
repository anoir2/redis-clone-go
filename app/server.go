package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type RedisServer interface {
	Start() error
	listen(conn net.Conn) error
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

	conn, err := l.Accept()
	if err != nil {
		return err
	}

	defer conn.Close()

	err = ds.listen(conn)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DefaultTCPServer) listen(conn net.Conn) error {
	var stream = make([]byte, 1024)
	for {
		n, err := conn.Read(stream)
		if err != nil {
			return err
		}
		fmt.Println(string(stream[:n]))
	}
}

func main() {
	fmt.Println("Starting server")
	var server RedisServer = NewDefaultTCPServer("0.0.0.0", 6379)
	err := server.Start()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Server terminated.")
}
