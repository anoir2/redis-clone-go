package request_handler

import (
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/codecrafters-io/redis-starter-go/app/serializer"
)

type RequestHandler interface {
	Handle(input string) (output string, err error)
}

type RequestHandlerImpl struct {
	serializer     serializer.Serializer
	commandHandler command.CommandHandler
}
