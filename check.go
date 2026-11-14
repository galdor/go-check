package check

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.n16f.net/pp"
)

type Check struct {
	T            testing.TB
	AbortOnError bool
}

func New(t testing.TB) *Check {
	return &Check{T: t}
}

func (c *Check) True(value bool, userMsg ...any) bool {
	c.T.Helper()

	if !value {
		c.Fail("value is not true", formatMessageAndArgs(userMsg...), nil)
		return false
	}

	return true
}

func (c *Check) False(value bool, userMsg ...any) bool {
	c.T.Helper()

	if value {
		c.Fail("value is not false", formatMessageAndArgs(userMsg...), nil)
		return false
	}

	return true
}

func (c *Check) Error(value error, userMsg ...any) bool {
	c.T.Helper()

	if value == nil {
		c.Fail("value is not an error", formatMessageAndArgs(userMsg...), nil)
		return false
	}

	return true
}

func (c *Check) NotError(value error, userMsg ...any) bool {
	c.T.Helper()

	if value != nil {
		c.Fail("value is an error", formatMessageAndArgs(userMsg...),
			[]string{
				formatValueLine("value:", prettyPrint(value)),
				formatValueLine("error:", value.Error()),
			})
		return false
	}

	return true
}

func (c *Check) Nil(value any, userMsg ...any) bool {
	c.T.Helper()

	if !isNil(value) {
		c.Fail("value is not nil", formatMessageAndArgs(userMsg...),
			[]string{
				formatValueLine("value:", prettyPrint(value)),
			})
		return false
	}

	return true
}

func (c *Check) NotNil(value any, userMsg ...any) bool {
	c.T.Helper()

	if isNil(value) {
		c.Fail("value is nil", formatMessageAndArgs(userMsg...), nil)
		return false
	}

	return true
}

func (c *Check) Equal(value1, value2 any, userMsg ...any) bool {
	c.T.Helper()

	if !cmp.Equal(value1, value2) {
		c.Fail("values are not equal", formatMessageAndArgs(userMsg...),
			[]string{
				formatValueLine("value:   ", prettyPrint(value2)),
				formatValueLine("expected:", prettyPrint(value1)),
			})

		return false
	}

	return true
}

func (c *Check) NotEqual(value1, value2 any, userMsg ...any) bool {
	c.T.Helper()

	if cmp.Equal(value1, value2) {
		c.Fail("values are equal", formatMessageAndArgs(userMsg...),
			[]string{
				formatValueLine("values:", prettyPrint(value1)),
			})

		return false
	}

	return true
}

func (c *Check) Fail(failureMsg, userMsg string, extraMsgs []string) {
	c.T.Helper()

	output := c.T.Output()

	c.T.Error(failureMsg)

	if userMsg != "" {
		c.printFailureLine(userMsg)
	}

	for _, msg := range extraMsgs {
		c.printFailureLine(msg)
	}

	// Keep multiple failures in the same test readable by adding an extra
	// newline character.
	fmt.Fprintln(output)

	if c.AbortOnError {
		c.T.FailNow()
	} else {
		c.T.Fail()
	}
}

func formatValueLine(prefix, value string) string {
	if strings.Contains(value, "\n") {
		return prefix + "\n" + indent(value, 1)
	} else {
		return prefix + " " + value
	}
}

func (c *Check) printFailureLine(line string) {
	// The "testing" package indents each testing.T.Errorf line with 4 spaces,
	// so we continue with the same indentation.

	fmt.Fprintln(c.T.Output(), indent(line, 1))

}

func formatMessageAndArgs(msg ...any) string {
	if len(msg) == 0 {
		return ""
	}

	return fmt.Sprintf(msg[0].(string), msg[1:]...)
}

func prettyPrint(value any) string {
	var buf bytes.Buffer

	if err := pp.PrintTo(&buf, value); err != nil {
		panic(fmt.Sprintf("cannot pretty print value %#v: %v", value, err))
	}

	s := buf.String()
	return strings.TrimRight(s, "\n")
}

func indent(s string, level int) string {
	padding := strings.Repeat(" ", level*4)
	return padding + strings.Replace(s, "\n", "\n"+padding, -1)
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	rValue := reflect.ValueOf(value)

	switch rValue.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return rValue.IsNil()
	}

	return false
}
