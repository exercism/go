package say

// The steps are interesting, but all that matters is the final exam.

import (
	"math"
	"testing"
)

var tests = []struct {
	uint64
	string
}{
	{1, "one"},
	{14, "fourteen"},
	{20, "twenty"},
	{22, "twenty-two"},
	{100, "one hundred"},
	{120, "one hundred twenty"},
	{123, "one hundred twenty-three"},
	{1000, "one thousand"},
	{1234, "one thousand two hundred thirty-four"},
	{1000000, "one million"},
	{1000002, "one million two"},
	{1002345, "one million two thousand three hundred forty-five"},
	{1e9, "one billion"},
	{987654321123, "nine hundred eighty-seven billion " +
		"six hundred fifty-four million " +
		"three hundred twenty-one thousand " +
		"one hundred twenty-three"},
	{0, "zero"},
	{math.MaxUint64, "eighteen quintillion " +
		"four hundred forty-six quadrillion " +
		"seven hundred forty-four trillion " +
		"seventy-three billion " +
		"seven hundred nine million " +
		"five hundred fifty-one thousand " +
		"six hundred fifteen"},
}

func TestSay(t *testing.T) {
	for _, test := range tests {
		if s := Say(test.uint64); s != test.string {
			t.Errorf("Say(%d) = %q.  Want %q.", test.uint64, s, test.string)
		}
	}
}
