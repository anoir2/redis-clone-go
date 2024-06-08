//go:build unit
// +build unit

package serializer_test

import (
	"errors"
	"github.com/codecrafters-io/redis-starter-go/app/serializer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayWithTwoCommandShouldReturnCommand(t *testing.T) {
	var strCommand = "*2\r\n$4\r\nPING\r\n$4\r\nPING\r\n"
	var parser = serializer.NewRESPSerializer()

	var actualCommand, err = parser.Deserialize(strCommand)

	var expected = []any{"PING", "PING"}
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestArrayWithACommandAndStringShouldReturnCommand(t *testing.T) {
	var strCommand = "*2\r\n$4\r\nPING\r\n$6\r\nSTRINg\r\n"
	var parser = serializer.NewRESPSerializer()

	var actualCommand, err = parser.Deserialize(strCommand)

	var expected = []any{"PING", "STRINg"}
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestArrayWithNestedArrayAndStringShouldReturnCommand(t *testing.T) {
	var strCommand = "*2\r\n*2\r\n$4\r\nPING\r\n$6\r\nSTRINg\r\n$5\r\nOTHER\r\n"
	var parser = serializer.NewRESPSerializer()

	var actualCommand, err = parser.Deserialize(strCommand)

	var expected = []any{[]any{"PING", "STRINg"}, "OTHER"}
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestWithInvalidCharInInputShouldReturnError(t *testing.T) {
	var strCommand = "@2\r\n$4\r\nPING\r\n$4\r\nPING\r\n"
	var parser = serializer.NewRESPSerializer()

	var actual, err = parser.Deserialize(strCommand)

	var expected = errors.New("invalid char")
	assert.Nil(t, actual)
	assert.Equal(t, expected, err)
}
