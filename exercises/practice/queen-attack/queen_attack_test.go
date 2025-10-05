package queenattack

import (
	"fmt"
	"testing"
)

// Arguments to CanQueenAttack are in algebraic notation.
// See http://en.wikipedia.org/wiki/Algebraic_notation_(chess)

type testCase struct {
	description string
	pos1, pos2  string
	expected    bool
}

var validTestCases = []testCase{
	{description: "no attack", pos1: "b3", pos2: "d7", expected: false},
	{description: "no attack", pos1: "a1", pos2: "f8", expected: false},
	{description: "same file", pos1: "b4", pos2: "b7", expected: true},
	{description: "same rank", pos1: "e4", pos2: "b4", expected: true},
	{description: "common diagonals", pos1: "a1", pos2: "f6", expected: true},
	{description: "common diagonals", pos1: "a6", pos2: "b7", expected: true},
	{description: "common diagonals", pos1: "d1", pos2: "f3", expected: true},
	{description: "common diagonals", pos1: "f1", pos2: "a6", expected: true},
	{description: "common diagonals", pos1: "a1", pos2: "h8", expected: true},
	{description: "common diagonals", pos1: "a8", pos2: "h1", expected: true},
}

func TestCanQueenAttackValid(t *testing.T) {
	for _, tc := range validTestCases {
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	allTestCases := append(validTestCases, invalidTestCases...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range allTestCases {
			CanQueenAttack(test.pos1, test.pos2)
		}
	}
}
