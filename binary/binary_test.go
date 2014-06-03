package binary

import (
	"testing"
)

var testCases = []struct {
	binary    string
	expected  int
	expectErr bool
}{
	{"1", 1, false},
	{"10", 2, false},
	{"11", 3, false},
	{"100", 4, false},
	{"1001", 9, false},
	{"10001101000", 1128, false},
	{"12", 0, true},
	{"hello", 0, true},
}

func TestParseBinary(t *testing.T) {
	for _, tt := range testCases {
		actual, err := ParseBinary(tt.binary)

		if actual != tt.expected {
			t.Fatalf("ParseBinary(%v): expected %v, actual %v", tt.binary, tt.expected, actual)
		}

		// if we expect an error and there isn't one
		if tt.expectErr && err == nil {
			t.Errorf("ParseBinary(%v): expected an error, but error is nil", tt.binary)
		}
		// if we don't expect an error and there is one
		if !tt.expectErr && err != nil {
			t.Errorf("ParseBinary(%v): expected no error, but error is: %s", tt.binary, err)
		}
	}
}

func BenchmarkBinary(b *testing.B) {
	b.StopTimer()
	for _, tt := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ParseBinary(tt.binary)
		}

		b.StopTimer()
	}
}
