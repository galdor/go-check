package main

import (
	"errors"
	"testing"

	"go.n16f.net/check"
)

func TestCheck(t *testing.T) {
	check := check.New(t)

	check.True(1 == 2)
	check.Equal("foo", "bar", "strings are not equal")
	check.Equal(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 2},
	)
	check.NotError(errors.New("test error"))
}
