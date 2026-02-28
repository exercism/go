package armstrong

import (
	"testing"
)

func TestArmstrong(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsNumber(tc.input); actual != tc.expected {
				t.Fatalf("IsNumber(%d) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkIsNumber(b *testing.B) {
	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			IsNumber(tc.input)
		}
	}
}
