package wordy

import "testing"

func TestAnswer(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			actual, ok := Answer(tc.question)
			switch {
			case tc.expectError:
				if ok {
					t.Fatalf("Answer(%q) expected error, got: %d", tc.question, actual)
				}
			case !ok:
				t.Fatalf("Answer(%q) returned ok=%t, want: %d", tc.question, ok, tc.expected)
			case actual != tc.expected:
				t.Fatalf("Answer(%q) = %d, want: %d", tc.question, actual, tc.expected)
			}
		})
	}
}

// Benchmark combined time to answer all questions.
func BenchmarkAnswer(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Answer(test.question)
		}
	}
}
