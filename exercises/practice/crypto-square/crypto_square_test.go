package cryptosquare

import "testing"

func TestEncode(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := Encode(tc.input); got != tc.expected {
				t.Fatalf("Encode(%q):\ngot: %q\nwant: %q", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Encode(tc.input)
		}
	}
}
