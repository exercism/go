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
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Hey(tt.input)
		}
	}
}
