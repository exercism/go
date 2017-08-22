package bob

import "testing"

func TestHey(t *testing.T) {
	for _, tt := range testCases {
		actual := Hey(tt.input)
		if actual != tt.expected {
			msg := `
	ALICE (%s): %q
	BOB: %s

	Expected Bob to respond: %s`
			t.Fatalf(msg, tt.description, tt.input, actual, tt.expected)
		}
	}
}

func BenchmarkHey(b *testing.B) {
	for _, tt := range testCases {
		for i := 0; i < b.N; i++ {
			Hey(tt.input)
		}
	}
}

// This test versioning is specific to Exercism,
// you don't need to worry about this now.
const targetTestVersion = 3

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}
