package rotationalcipher

import (
	"testing"
)

var testCases = []struct {
	description   string
	inputPlain    string
	inputShiftKey int
	expected      string
}{
	{
		description:   "rotate a by 0, same output as input",
		inputPlain:    "a",
		inputShiftKey: 0,
		expected:      "a",
	},
	{
		description:   "rotate a by 1",
		inputPlain:    "a",
		inputShiftKey: 1,
		expected:      "b",
	},
	{
		description:   "rotate a by 26, same output as input",
		inputPlain:    "a",
		inputShiftKey: 26,
		expected:      "a",
	},
	{
		description:   "rotate n by 13 with wrap around alphabet",
		inputPlain:    "n",
		inputShiftKey: 13,
		expected:      "a",
	},
	{
		description:   "rotate capital letters",
		inputPlain:    "OMG",
		inputShiftKey: 5,
		expected:      "TRL",
	},
	{
		description:   "rotate spaces",
		inputPlain:    "O M G",
		inputShiftKey: 5,
		expected:      "T R L",
	},
	{
		description:   "rotate numbers",
		inputPlain:    "Testing 1 2 3 testing",
		inputShiftKey: 4,
		expected:      "Xiwxmrk 1 2 3 xiwxmrk",
	},
	{
		description:   "rotate punctuation",
		inputPlain:    "Let's eat, Grandma!",
		inputShiftKey: 21,
		expected:      "Gzo'n zvo, Bmviyhv!",
	},
	{
		description:   "rotate all letters",
		inputPlain:    "The quick brown fox jumps over the lazy dog.",
		inputShiftKey: 13,
		expected:      "Gur dhvpx oebja sbk whzcf bire gur ynml qbt.",
	},
}

func TestRotationalCipher(t *testing.T) {
	for _, testCase := range testCases {
		cipher := RotationalCipher(testCase.inputPlain, testCase.inputShiftKey)
		if cipher != testCase.expected {
			t.Fatalf("FAIL: %s\n\tRotationalCipher(%s, %d)\nexpected: %s, \ngot:      %s",
				testCase.description, testCase.inputPlain, testCase.inputShiftKey, testCase.expected, cipher)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkRotationalCipher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			RotationalCipher(testCase.inputPlain, testCase.inputShiftKey)
		}
	}
}
