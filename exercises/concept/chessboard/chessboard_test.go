package chessboard

import (
	"fmt"
	"testing"
)

// newChessboard return a *Chessboard for tests
//
//   A B C D E F G H
// 8 # _ _ _ # _ _ # 8
// 7 _ _ _ _ _ _ _ _ 7
// 6 _ _ _ _ # _ _ # 6
// 5 _ # _ _ _ _ _ # 5
// 4 _ _ _ _ _ _ # # 4
// 3 # _ # _ _ _ _ # 3
// 2 _ _ _ _ _ _ _ # 2
// 1 # _ _ _ _ _ _ # 1
//   A B C D E F G H

func newChessboard() Chessboard {
	return Chessboard{
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{false, false, false, false, true, false, false, false},
		"C": File{false, false, true, false, false, false, false, false},
		"D": File{false, false, false, false, false, false, false, false},
		"E": File{false, false, false, false, false, true, false, true},
		"F": File{false, false, false, false, false, false, false, false},
		"G": File{false, false, false, true, false, false, false, false},
		"H": File{true, true, true, true, true, true, false, true},
	}
}

func TestCountInFile(t *testing.T) {
	cb := newChessboard()
	testCases := []struct {
		in       string
		expected int
	}{
		{in: "A", expected: 3},
		{in: "B", expected: 1},
		{in: "C", expected: 1},
		{in: "D", expected: 0},
		{in: "E", expected: 2},
		{in: "F", expected: 0},
		{in: "G", expected: 1},
		{in: "H", expected: 7},
		{in: "Z", expected: 0},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("Count of occupied squares in file %s", test.in), func(t *testing.T) {
			if got := CountInFile(cb, test.in); got != test.expected {
				t.Errorf("CountInFile(chessboard, %q) = %d, want: %d", test.in, got, test.expected)
			}
		})
	}
}

func TestCountInRank(t *testing.T) {
	cb := newChessboard()
	testCases := []struct {
		in       int
		expected int
	}{
		{in: 1, expected: 2},
		{in: 2, expected: 1},
		{in: 3, expected: 3},
		{in: 4, expected: 2},
		{in: 5, expected: 2},
		{in: 6, expected: 2},
		{in: 7, expected: 0},
		{in: 8, expected: 3},
		// cases not between 1 and 8, inclusive
		{in: 100, expected: 0},
		{in: 0, expected: 0},
		{in: -1, expected: 0},
		{in: -100, expected: 0},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("Count of occupied squares in rank %d", test.in), func(t *testing.T) {
			if got := CountInRank(cb, test.in); got != test.expected {
				t.Errorf("CountInRank(chessboard, %d) = %d, want: %d", test.in, got, test.expected)
			}
		})
	}
}

func TestCountAll(t *testing.T) {
	cb := newChessboard()
	expected := 64
	if got := CountAll(cb); got != expected {
		t.Errorf("CountAll(chessboard) = %d, want: %d", got, expected)
	}
}

func TestCountOccupied(t *testing.T) {
	cb := newChessboard()
	expected := 15
	if got := CountOccupied(cb); got != expected {
		t.Errorf("CountOccupied(chessboard) = %d, want: %d", got, expected)
	}
}
