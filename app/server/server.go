package server

import (
	"net"
)

type RedisServer interface {
	Start() error
	listen(conn net.Conn) error
}
