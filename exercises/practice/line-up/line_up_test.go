package lineup

import "testing"

func TestFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Format(tc.name, tc.number); actual != tc.expected {
				t.Fatalf("Format(%q, %d) = %q, want: %q", tc.name, tc.number, actual, tc.expected)
			}
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Format(tc.name, tc.number)
		}
	}
}

