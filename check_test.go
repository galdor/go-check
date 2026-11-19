package check

import (
	"errors"
	"fmt"
	"testing"
)

var ErrTest = errors.New("test error")

type TestErr struct {
}

func (err TestErr) Error() string {
	return "test error"
}

func TestTrue(t *testing.T) {
	check := New(t)

	check.True(true)
	check.True(1 == 1)
}

func TestFalse(t *testing.T) {
	check := New(t)

	check.False(false)
	check.False(1 == 2)
}

func TestNil(t *testing.T) {
	check := New(t)

	check.Nil(nil)
	check.Nil((*int)(nil))
	check.Nil(([]string)(nil))
}

func TestNotNil(t *testing.T) {
	check := New(t)

	check.NotNil(42)
	check.NotNil(new(42))
	check.NotNil([]string{})
}

func TestEqual(t *testing.T) {
	check := New(t)

	check.Equal(nil, nil)
	check.Equal(42, 42)
	check.Equal(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 2, "a": 1},
	)
}

func TestOrdering(t *testing.T) {
	check := New(t)

	check.Lesser(1, 2)
	check.Lesser("abc", "bcd")

	check.LesserOrEqual(1, 2)
	check.LesserOrEqual(2, 2)
	check.LesserOrEqual("foo", "foo2")
	check.LesserOrEqual("foo", "foo")

	check.Greater(3.0, 1.5)
	check.Greater("abc", "abb")

	check.GreaterOrEqual(3.0, 1.5)
	check.GreaterOrEqual(3.0, 3.0)
	check.GreaterOrEqual("abc", "abb")
	check.GreaterOrEqual("abc", "abc")
}

func TestNotEqual(t *testing.T) {
	check := New(t)

	check.NotEqual(1, 2)
	check.NotEqual(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"a": 2, "b": 3},
	)
}

func TestError(t *testing.T) {
	check := New(t)

	check.Error(errors.New("test"))
}

func TestErrorIs(t *testing.T) {
	check := New(t)

	check.ErrorIs(ErrTest, ErrTest)
	check.ErrorIs(fmt.Errorf("%w", ErrTest), ErrTest)
}

func TestErrorAs(t *testing.T) {
	check := New(t)

	var testErr TestErr
	check.ErrorAs(TestErr{}, &testErr)
	check.ErrorAs(fmt.Errorf("%w", TestErr{}), &testErr)
}

func TestNotError(t *testing.T) {
	check := New(t)

	check.NotError(nil)
}
