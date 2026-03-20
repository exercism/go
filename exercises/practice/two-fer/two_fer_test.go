package twofer

import "testing"

func TestShareWith(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := ShareWith(tc.input)
			if got != tc.expected {
				t.Fatalf("ShareWith(%q)\n got: %q\nwant: %q", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkShareWith(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			ShareWith(test.input)
		}
	}
}
