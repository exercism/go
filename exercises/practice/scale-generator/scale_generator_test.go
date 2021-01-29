package scale

import (
	"fmt"
	"testing"
)

func TestScale(t *testing.T) {
	for _, test := range scaleTestCases {
		actual := Scale(test.tonic, test.interval)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("FAIL: %s - Scale(%q, %q)\nExpected: %q\nActual: %q", test.description, test.tonic, test.interval, test.expected, actual)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func BenchmarkScale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scaleTestCases {
			Scale(test.tonic, test.interval)
		}
	}
}
