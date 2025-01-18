package reverse

import (
	"testing"
	"testing/quick"
)

func TestReverse(t *testing.T) {
	for _, tc := range append(testCases, multiByteCases...) {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Reverse(tc.input); actual != tc.expected {
				t.Fatalf("Reverse(%q) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestReverseOfReverse(t *testing.T) {
	assertion := func(s string) bool {
		return s == Reverse(Reverse(s))
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkReverse(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Reverse(tc.input)
		}
	}
}

// mutiByteCases adds UTF-8 multi-byte case,
// since the canonical-data.json (generator data source for cases_test.go)
// doesn't have any such cases.
var multiByteCases = []reverseTestCase{
	{
		description: "a multi-byte test case",
		input:       "Hello, 世界",
		expected:    "界世 ,olleH",
	},
}
