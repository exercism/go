package affinecipher

import (
	"testing"
)

func runTests(
	t *testing.T,
	name string,
	op func(string, int, int) (string, error),
	testCases []testCase,
) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := op(tc.inputPhrase, tc.inputA, tc.inputB)
			if err != nil {
				if !tc.expectError {
					t.Fatalf("%s(%s, %d, %d) returned unexpected error %v", name, tc.inputPhrase, tc.inputA, tc.inputB, err)
				}
				return
			}
			if tc.expectError {
				t.Fatalf("%s(%s, %d, %d) expected error, got %v", name, tc.inputPhrase, tc.inputA, tc.inputB, got)
			}
			if tc.expected != got {
				t.Fatalf("%s(%s, %d, %d) = %s, expected: %s", name, tc.inputPhrase, tc.inputA, tc.inputB, got, tc.expected)
			}
		})
	}
}
func TestEncode(t *testing.T) {
	runTests(t, "Encode", Encode, encodeTests)
}

func TestDecode(t *testing.T) {
	runTests(t, "Decode", Decode, decodeTests)
}

func runBenchmark(
	b *testing.B,
	op func(string, int, int) (string, error),
	testCases []testCase,
) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
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
