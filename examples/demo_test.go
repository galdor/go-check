package main

import (
	"errors"
	"testing"

	"go.n16f.net/check"
)

var ErrTest = errors.New("test error")

type TestError struct {
}

func (err TestError) Error() string {
	return "test error"
}

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

	check.Lesser(10.0, 10.0)
	check.LesserOrEqual(10.0, 3.14)
	check.Greater(-50, -50)
	check.GreaterOrEqual(-50, -10)

	check.Error(nil)
	check.ErrorIs(errors.New("random error"), ErrTest)
	var testErr TestError
	check.ErrorAs(errors.New("random error"), &testErr)
	check.NotError(errors.New("test error"))
}
