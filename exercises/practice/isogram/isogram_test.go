package isogram

import "testing"

func TestIsIsogram(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsIsogram(tc.input); actual != tc.expected {
				t.Fatalf("IsIsogram(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogram(c.input)
		}
	}
}
