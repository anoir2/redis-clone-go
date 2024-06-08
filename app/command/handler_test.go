//go:build unit
// +build unit

package command_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	assert.Nil(t, nil)
}
