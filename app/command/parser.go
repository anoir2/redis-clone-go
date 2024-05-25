package command

import (
	"strings"
)

type Parser interface {
	Parse(input string) (Command, error)
}

type StandardParser struct {
}

func NewParser() *StandardParser {
	return &StandardParser{}
}

func (sp *StandardParser) Parse(rawInput string) (Command, error) {
	input := strings.TrimSpace(rawInput)
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, "\r")
	switch input {
	case "PING":
		return NewPingCommand(), nil
	}

	return NewUnknownCommand(), nil
}
