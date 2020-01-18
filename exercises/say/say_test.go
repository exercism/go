package say

import (
	"testing"
)

func TestSay(t *testing.T) {
	for _, tc := range testCases {
		actual, ok := Say(tc.input)
		switch {
		case tc.expectError:
			if ok {
				t.Fatalf("FAIL: %s\nExpected error but received: %v", tc.description, actual)
			}
		case !ok:
			t.Fatalf("FAIL: %s\nDid not expect an error", tc.description)
		case actual != tc.expected:
			t.Fatalf("FAIL: %s\nExpected: %v\nActual: %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkSay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Say(tc.input)
		}
	}
}
