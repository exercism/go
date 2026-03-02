package queenattack

import (
	"fmt"
	"testing"
)

// Arguments to CanQueenAttack are in algebraic notation.
// See http://en.wikipedia.org/wiki/Algebraic_notation_(chess)

func TestCanQueenAttackValid(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s, white queen: %s, black queen: %s", tc.description, tc.pos1, tc.pos2), func(t *testing.T) {
			got, err := CanQueenAttack(tc.pos1, tc.pos2)
			if err != nil {
				t.Fatalf("CanQueenAttack(%q, %q) returned unexpected error %v", tc.pos1, tc.pos2, err)
			}
			if got != tc.expected {
				t.Fatalf("CanQueenAttack(%q, %q) = %v, want: %v", tc.pos1, tc.pos2, got, tc.expected)
			}
		})
	}
}

var invalidTestCases = []testCase{
	{description: "same square", pos1: "b4", pos2: "b4"},
	{description: "position off board", pos1: "a8", pos2: "b9"},
	{description: "position off board", pos1: "a0", pos2: "b1"},
	{description: "position off board", pos1: "g3", pos2: "i5"},
	{description: "invalid position", pos1: "here", pos2: "there"},
	{description: "empty position", pos1: "", pos2: ""},
}

func TestCanQueenAttackInvalid(t *testing.T) {
	for _, tc := range invalidTestCases {
		t.Run(fmt.Sprintf("%s, white queen: %s, black queen: %s", tc.description, tc.pos1, tc.pos2), func(t *testing.T) {
			got, err := CanQueenAttack(tc.pos1, tc.pos2)
			if err == nil {
				t.Fatalf("CanQueenAttack(%q, %q) expected error, got %v", tc.pos1, tc.pos2, got)
			}
		})
	}
}

// Benchmark combined time for all test cases
func BenchmarkCanQueenAttack(b *testing.B) {
	allTestCases := append(testCases, invalidTestCases...)
	b.ResetTimer()
	for range b.N {
		for _, test := range allTestCases {
			CanQueenAttack(test.pos1, test.pos2)
		}
	}
}
