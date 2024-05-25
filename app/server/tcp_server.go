package server

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/command"
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
		cmdToExec, err := ds.parser.Parse(string(stream[:n]))
		if err != nil {
			fmt.Println("error:", err)
		}
		res, err := cmdToExec.Execute()
		if err != nil {
			return err
		}

		conn.Write([]byte(res.Output()))
	}
}
