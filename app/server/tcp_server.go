package server

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"io"
	"net"
	"strconv"
)

type DefaultTCPServer struct {
	port       int
	host       string
	serializer command.Serializer
}

func NewDefaultTCPServer(host string, port int, serializer command.Serializer) *DefaultTCPServer {
	return &DefaultTCPServer{port: port, host: host, serializer: serializer}
}

func (ds *DefaultTCPServer) Start() error {
	var address = ds.host + ":" + strconv.Itoa(ds.port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go ds.listen(conn)
	}
}

func (ds *DefaultTCPServer) listen(conn net.Conn) {
	defer conn.Close()

	var stream = make([]byte, 1024)
	for {
		n, err := conn.Read(stream)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			return
		}

		var input = string(stream[:n])
		fmt.Println(input)

		var output string
		cmdRes, err := ds.executeCommand(input)
		if err != nil {
			fmt.Println(err)
			cmdRes = []command.Result{}
			output = "Internal error\n"
		}

		for _, cmd := range cmdRes {
			output += cmd.Output()
		}
		fmt.Println("output:", output)

		_, err = conn.Write([]byte(output))
		if err != nil {
			fmt.Println("conn error:", err)
		}
	}
}

func (ds *DefaultTCPServer) executeCommand(input string) ([]command.Result, error) {
	deserializeOutput, err := ds.serializer.Deserialize(input)
	if err != nil {
		return []command.Result{}, err
	}

	cmdToExec, _ := deserializeOutput.([]any)

	var resToReturn = make([]command.Result, 0, len(cmdToExec))
	for _, cmd := range cmdToExec {
		cmd, ok := cmd.(command.Command)
		if ok {
			result, err := cmd.Execute()
			if err != nil {
				return []command.Result{}, err
			}

			resToReturn = append(resToReturn, result)
		}
	}

	return resToReturn, nil
}
