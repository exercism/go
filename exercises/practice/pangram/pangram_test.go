package pangram

import (
	"testing"
)

func TestPangram(t *testing.T) {
	for _, test := range testCases {
		if actual := IsPangram(test.input); actual != test.expected {
			t.Fatalf("FAIL: %s\nInput %q expected [%t], actual [%t]", test.description, test.input, test.expected, actual)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func BenchmarkPangram(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsPangram(test.input)
		}
	}
}
