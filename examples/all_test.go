package main

import (
	"errors"
	"testing"

	"go.n16f.net/check"
)

func TestAll(t *testing.T) {
	check := check.New(t)

	check.True(1 == 2)

	check.Nil(42)
	check.Nil(map[string]int{})
	check.NotNil(nil)
	check.NotNil([]string(nil))

	check.Equal("foo", "bar")
	check.NotEqual(
		map[string]int{"a": 1},
		map[string]int{"a": 1},
	)

	check.Error(nil)
	check.NotError(errors.New("test error"))
}
