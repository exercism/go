package series

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	digits   int
	input    string
	expected []string
}{
	{
		digits:   1,
		input:    "01234",
		expected: []string{"0", "1", "2", "3", "4"},
	},
	{
		digits:   1,
		input:    "92834",
		expected: []string{"9", "2", "8", "3", "4"},
	},
	{
		digits:   2,
		input:    "01234",
		expected: []string{"01", "12", "23", "34"},
	},
	{
		digits:   2,
		input:    "98273463",
		expected: []string{"98", "82", "27", "73", "34", "46", "63"},
	},
	{
		digits:   2,
		input:    "37103",
		expected: []string{"37", "71", "10", "03"},
	},
	{
		digits:   3,
		input:    "01234",
		expected: []string{"012", "123", "234"},
	},
	{
		digits:   3,
		input:    "31001",
		expected: []string{"310", "100", "001"},
	},
	{
		digits:   3,
		input:    "982347",
		expected: []string{"982", "823", "234", "347"},
	},
	{
		digits:   4,
		input:    "01234",
		expected: []string{"0123", "1234"},
	},
	{
		digits:   4,
		input:    "91274",
		expected: []string{"9127", "1274"},
	},
	{
		digits:   5,
		input:    "01234",
		expected: []string{"01234"},
	},
	{
		digits:   5,
		input:    "81228",
		expected: []string{"81228"},
	},
	{
		digits:   6,
		input:    "01234",
		expected: nil,
	},
	{
		digits:   len(cx) + 1,
		input:    cx,
		expected: nil,
	},
}

var cx = "01032987583"

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d digits in %s", tc.digits, tc.input), func(t *testing.T) {
			got := All(tc.digits, tc.input)
			if len(got) == 0 && len(tc.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("All(%d, %q)\n got: %v, want: %v", tc.digits, tc.input, got, tc.expected)
			}
		})
	}
}

func TestUnsafeFirst(t *testing.T) {
	for _, tc := range testCases {
		if len(tc.expected) == 0 {
			continue
		}
		t.Run(fmt.Sprintf("first with %d digits in %s", tc.digits, tc.input), func(t *testing.T) {
			if got := UnsafeFirst(tc.digits, tc.input); got != tc.expected[0] {
				t.Fatalf("UnsafeFirst(%d, %q) = %q, want: %q", tc.digits, tc.input, got, tc.expected[0])
			}
		})
	}
}
