package rotationalcipher

import (
	"testing"
)

func TestRotationalCipher(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := RotationalCipher(tc.input, tc.shift)
			if got != tc.expected {
				t.Fatalf("RotationalCipher(%q, %d)\n got: %q\nwant: %q", tc.input, tc.shift, got, tc.expected)
			}
		})
	}
}

func BenchmarkRotationalCipher(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			RotationalCipher(tc.input, tc.shift)
		}
	}
}
