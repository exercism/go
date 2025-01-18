package prime

import "testing"

func TestNth(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Nth(tc.input)
			switch {
			case tc.err != "":
				if err == nil {
					t.Fatalf("Nth(%d) expected error: %q, got: %d", tc.input, tc.err, actual)
				}
			case err != nil:
				t.Fatalf("Nth(%d) returned error: %v, want: %d", tc.input, err, tc.expected)
			case actual != tc.expected:
				t.Fatalf("Nth(%d) = %d, want: %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkNth(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Nth(10001)
	}
}
