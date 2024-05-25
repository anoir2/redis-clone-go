//go:build unit
// +build unit

package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePingCommand(t *testing.T) {
	var strCommand = "PING"
	var parser = NewParser()

	var actualCommand, err = parser.Parse(strCommand)

	var expected = NewPingCommand()
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}

func TestParseInvalidCommand(t *testing.T) {
	var strCommand = "INVALID-CMD"
	var parser = NewParser()

	var actualCommand, err = parser.Parse(strCommand)

	var expected = NewUnknownCommand()
	assert.Nil(t, err)
	assert.Equal(t, expected, actualCommand)
}
