package allyourbase

import (
	"reflect"
	"testing"
)

func TestConvertToBase(t *testing.T) {
	for _, c := range testCases {
		output, err := ConvertToBase(c.inputBase, c.inputDigits, c.outputBase)
		if c.err != "" {
			if err == nil || c.err != err.Error() {
				t.Fatalf(`FAIL: %s
	Expected error: %s
	Got: %v`, c.description, c.err, err)
			}
		} else if !reflect.DeepEqual(c.expected, output) {
			t.Fatalf(`FAIL: %s
    Input base: %d
    Input digits: %#v
    Output base: %d
    Expected output digits: %#v
    Got: %#v`, c.description, c.inputBase, c.inputDigits, c.outputBase, c.expected, output)
		}
		t.Logf("PASS: %s", c.description)
	}
}
