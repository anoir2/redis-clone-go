package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	respEndline = "\r\n"
)

type Parser interface {
	Parse(input string) ([]Command, error)
}

type RESPParser struct {
}

func NewRESPParser() *RESPParser {
	return &RESPParser{}
}

func (sp *RESPParser) Parse(rawInput string) ([]Command, error) {
	var rawCmds, err = sp.extractCommands(rawInput)
	if err != nil {
		return nil, err
	} else if len(rawCmds) == 0 {
		return nil, errors.New("No commands to parse found")
	}
	var cmdToReturn = make([]Command, 0, len(rawCmds))
	for _, rawCmd := range rawCmds {
		switch rawCmd {
		case "PING":
			cmdToReturn = append(cmdToReturn, NewPingCommand())
		case "COMMAND":
			cmdToReturn = append(cmdToReturn, NewCommandsCommand())
		default:
			return nil, errors.New("invalid command: " + rawCmd)
		}
	}

	return cmdToReturn, nil
}

func (sp *RESPParser) extractCommands(rawInput string) ([]string, error) {
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
