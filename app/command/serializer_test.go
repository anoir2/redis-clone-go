//go:build unit
// +build unit

package command_test

import (
	"errors"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayWithTwoCommandShouldReturnCommand(t *testing.T) {
	var strCommand = "*2\r\n$4\r\nPING\r\n$4\r\nPING\r\n"
	var parser = command.NewRESPSerializer(command.NewRedisCommandParser())

	var actualCommand, err = parser.Deserialize(strCommand)

	var expected = []any{command.NewPingCommand(), command.NewPingCommand()}
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestArrayWithACommandAndStringShouldReturnCommand(t *testing.T) {
	var strCommand = "*2\r\n$4\r\nPING\r\n$6\r\nSTRINg\r\n"
	var parser = command.NewRESPSerializer(command.NewRedisCommandParser())

	var actualCommand, err = parser.Deserialize(strCommand)

	var expected = []any{command.NewPingCommand(), "STRINg"}
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestWithInvalidCharInInputShouldReturnError(t *testing.T) {
	var strCommand = "@2\r\n$4\r\nPING\r\n$4\r\nPING\r\n"
	var parser = command.NewRESPSerializer(command.NewRedisCommandParser())

	var actual, err = parser.Deserialize(strCommand)

	var expected = errors.New("invalid char")
	assert.Nil(t, actual)
	assert.Equal(t, expected, err)
}
