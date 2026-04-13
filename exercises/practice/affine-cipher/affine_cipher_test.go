package affinecipher

import (
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tc := range encodeTests {
		t.Run(tc.description, func(t *testing.T) {
			if got, err := Encode(tc.inputPhrase, tc.inputA, tc.inputB); err != nil {
				if !tc.expectError {
					t.Fatalf("Encode(%s, %d, %d) returned unexpected error %v", tc.inputPhrase, tc.inputA, tc.inputB, err)
				}
			} else if tc.expectError {
				t.Fatalf("Encode(%s, %d, %d) expected error, got %q", tc.inputPhrase, tc.inputA, tc.inputB, got)
			} else if tc.expected != got {
				t.Fatalf("Encode(%s, %d, %d) = %q, want: %q", tc.inputPhrase, tc.inputA, tc.inputB, got, tc.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range decodeTests {
		t.Run(tc.description, func(t *testing.T) {
			if got, err := Decode(tc.inputPhrase, tc.inputA, tc.inputB); err != nil {
				if !tc.expectError {
					t.Fatalf("Decode(%s, %d, %d) returned unexpected error %v", tc.inputPhrase, tc.inputA, tc.inputB, err)
				}
			} else if tc.expectError {
				t.Fatalf("Decode(%s, %d, %d) expected error, got %q", tc.inputPhrase, tc.inputA, tc.inputB, got)
			} else if tc.expected != got {
				t.Fatalf("Decode(%s, %d, %d) = %q, want: %q", tc.inputPhrase, tc.inputA, tc.inputB, got, tc.expected)
			}
		})
	}
}

func runBenchmark(
	b *testing.B,
	op func(string, int, int) (string, error),
	testCases []testCase,
) {
	for range b.N {
		for _, tc := range testCases {
			op(tc.inputPhrase, tc.inputA, tc.inputB)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	runBenchmark(b, Encode, encodeTests)
}

func BenchmarkDecode(b *testing.B) {
	runBenchmark(b, Decode, decodeTests)
}
