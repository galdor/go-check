package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestify(t *testing.T) {
	assert := assert.New(t)

	assert.True(1 == 2)
	assert.Equal("foo", "bar", "strings are not equal")
	assert.Equal(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 2},
	)
	assert.NoError(errors.New("test error"))
}
