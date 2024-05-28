package server

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"io"
	"net"
	"strconv"
)

type DefaultTCPServer struct {
	port   int
	host   string
	parser command.Parser
}

func NewDefaultTCPServer(host string, port int, parser command.Parser) *DefaultTCPServer {
	return &DefaultTCPServer{port: port, host: host, parser: parser}
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
		if err != nil && err != io.EOF {
			return err
		}
	}

	return nil
}

func (ds *DefaultTCPServer) listen(conn net.Conn) error {
	defer conn.Close()

	var stream = make([]byte, 1024)
	for {
		n, err := conn.Read(stream)
		if err != nil {
			return err
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
	cmdsToExec, err := ds.parser.Parse(input)
	if err != nil {
		return []command.Result{}, err
	}

	var resToReturn = make([]command.Result, 0, len(cmdsToExec))
	for _, cmd := range cmdsToExec {
		result, err := cmd.Execute()
		if err != nil {
			return []command.Result{}, err
		}

		resToReturn = append(resToReturn, result)
	}

	return resToReturn, nil
}
