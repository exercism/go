package reverse

import (
	"testing"
	"testing/quick"
)

func TestReverse(t *testing.T) {
	for _, testCase := range append(testCases, multiByteCases...) {
		if res := String(testCase.input); res != testCase.expected {
			t.Fatalf("FAIL: %s(%s)\nExpected: %q\nActual: %q",
				testCase.description, testCase.input, testCase.expected, res)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestReverseOfReverse(t *testing.T) {
	assertion := func(s string) bool {
		return s == String(String(s))
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			String(test.input)
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
