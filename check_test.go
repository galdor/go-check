package check

import "testing"

func TestTrue(t *testing.T) {
	check := New(t, false)

	check.True(true)
	check.True(1 == 1)
}

func TestFalse(t *testing.T) {
	check := New(t, false)

	check.False(false)
	check.False(1 == 2)
}

func TestEqual(t *testing.T) {
	check := New(t, false)

	check.Equal(nil, nil)
	check.Equal(42, 42)
	check.Equal(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 2, "a": 1},
	)
}
