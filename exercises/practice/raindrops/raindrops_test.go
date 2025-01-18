package raindrops

import "testing"

func TestConvert(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Convert(tc.input); actual != tc.expected {
				t.Fatalf("Convert(%d) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkConvert(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Convert(test.input)
		}
	}
}
