//go:build unit
// +build unit

package command_test

import (
	"errors"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRESPParserMultipleCommandsShouldReturnCommandsParsed(t *testing.T) {
	var strCommand = "PiNG"
	var parser = command.NewRedisCommandParser()

	var actualCommand, err = parser.Parse(strCommand)

	var expected = command.NewPingCommand()
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestRESPParserCommandsCommandShouldReturnCommandParsed(t *testing.T) {
	var strCommand = "COMMAND"
	var parser = command.NewRedisCommandParser()

	var actualCommand, err = parser.Parse(strCommand)

	var expected = command.NewCommandsCommand()
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestRESPParserInvalidCommandShouldReturnInvalidCmdErr(t *testing.T) {
	var strCommand = "INVALID"
	var parser = command.NewRedisCommandParser()

	var actualCommand, actualErr = parser.Parse(strCommand)

	assert.Equal(t, errors.New("invalid command: INVALID"), actualErr)
	assert.Nil(t, actualCommand)
}
