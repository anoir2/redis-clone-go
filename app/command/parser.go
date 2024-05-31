package command

import (
	"errors"
	"fmt"
	"strconv"
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

func (sp *RedisCommandParser) extractCommands(rawInput string) ([]string, error) {
	var cmds = make([]string, 0, 10)
	if len(rawInput) == 0 || rawInput[0] != '*' {
		return nil, errors.New("no commands to extract")
	}

	var inputs = strings.Split(rawInput, respEndline)
	var argsNumAsStr = strings.TrimSpace(inputs[0])[1:]
	fmt.Println("args number", argsNumAsStr)

	var argsNum, err = strconv.Atoi(argsNumAsStr)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= argsNum*2; i += 2 {
		var argsNumAsStr = strings.TrimSpace(inputs[i])[1:]
		if len(inputs[i]) == 0 || inputs[i][0] != '$' {
			return nil, errors.New("no line size")
		}

		var cmdLen, err = strconv.Atoi(argsNumAsStr)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, inputs[i+1][:cmdLen])
	}

	return cmds, nil
}
