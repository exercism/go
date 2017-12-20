package reverse

import (
	"testing"
	"testing/quick"
)

func TestReverse(t *testing.T) {
	for _, testCase := range testCases {
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
	for _, test := range testCases {
		for i := 0; i < b.N; i++ {
			String(test.input)
		}
	}
}
