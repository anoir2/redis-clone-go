package command

import (
	"errors"
	"strings"
)

type CommandParser interface {
	Parse(input string) (Command, error)
}

type RedisCommandParser struct {
}

func NewRedisCommandParser() *RedisCommandParser {
	return &RedisCommandParser{}
}

func (sp *RedisCommandParser) Parse(cmd string) (Command, error) {
	switch strings.ToUpper(cmd) {
	case "PING":
		return NewPingCommand(), nil
	case "COMMAND":
		return NewCommandsCommand(), nil
	}

	return nil, errors.New("invalid command: " + cmd)
}
