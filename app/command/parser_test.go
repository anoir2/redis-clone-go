//go:build unit
// +build unit

package command_test

import (
	"errors"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRESPParserPingCommandShouldReturnCommandParsed(t *testing.T) {
	var strCommand = "*1\r\n$4\r\nPING\n"
	var parser = command.NewRESPParser()

	var actualCommand, err = parser.Parse(strCommand)

	var expected = []command.Command{command.NewPingCommand()}
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestRESPParserInvalidCommandShouldReturnInvalidCmdErr(t *testing.T) {
	var strCommand = "*1\r\n$7\r\nINVALID\r\n"
	var parser = command.NewRESPParser()

	var actualCommand, actualErr = parser.Parse(strCommand)

	assert.Equal(t, errors.New("invalid command: INVALID"), actualErr)
	assert.Nil(t, actualCommand)
}

func TestRESPParserInputWithInvalidCharAsStartShouldReturnNoCmdErr(t *testing.T) {
	var strCommand = "@1\r\n$7\r\nINVALID\r\n"
	var parser = command.NewRESPParser()

	var actualCommand, actualErr = parser.Parse(strCommand)

	assert.Equal(t, errors.New("no commands to extract"), actualErr)
	assert.Nil(t, actualCommand)
}
