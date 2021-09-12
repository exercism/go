package chessboard

import (
	"testing"
)

// newChessboard return a *Chessboard for tests
//
//   1 2 3 4 5 6 7 8
// A # _ # _ _ _ _ # A
// B _ _ _ _ # _ _ _ B
// C _ _ # _ _ _ _ _ C
// D _ _ _ _ _ _ _ _ D
// E _ _ _ _ _ # _ # E
// F _ _ _ _ _ _ _ _ F
// G _ _ _ # _ _ _ _ G
// H # # # # # # _ # H
//   1 2 3 4 5 6 7 8
func newChessboard() *Chessboard {
	return &Chessboard{
		'A': Rank{true, false, true, false, false, false, false, true},
		'B': Rank{false, false, false, false, true, false, false, false},
		'C': Rank{false, false, true, false, false, false, false, false},
		'D': Rank{false, false, false, false, false, false, false, false},
		'E': Rank{false, false, false, false, false, true, false, true},
		'F': Rank{false, false, false, false, false, false, false, false},
		'G': Rank{false, false, false, true, false, false, false, false},
		'H': Rank{true, true, true, true, true, true, false, true},
	}
}

func TestCountInRank(t *testing.T) {
	cb := newChessboard()
	for _, test := range []struct {
		in  byte
		out int
	}{
		{'A', 3},
		{'B', 1},
		{'C', 1},
		{'D', 0},
		{'E', 2},
		{'F', 0},
		{'G', 1},
		{'H', 7},
		{'Z', 0},
	} {
		if out := cb.CountInRank(test.in); out != test.out {
			t.Errorf(
				"CountInRank('%v') returned %v while %v was expected\n",
				string(test.in),
				out,
				test.out,
			)
		}
	}
}

func TestCountInFile(t *testing.T) {
	cb := newChessboard()
	for _, test := range []struct {
		in  int
		out int
	}{
		{1, 2},
		{2, 1},
		{3, 3},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 0},
		{8, 3},
		{100, 0},
	} {
		if out := cb.CountInFile(test.in); out != test.out {
			t.Errorf(
				"CountInFile('%v') returned %v while %v was expected\n",
				test.in,
				out,
				test.out,
			)
		}
	}
}

func TestCountAll(t *testing.T) {
	cb := newChessboard()
	wanted := 64
	if out := cb.CountAll(); out != wanted {
		t.Errorf("CountAll() returned %v while %v was expected", out, wanted)
	}
}

func TestCountOccupied(t *testing.T) {
	cb := newChessboard()
	wanted := 15
	if out := cb.CountOccupied(); out != wanted {
		t.Errorf("CountOccupied() returned %v while %v was expected", out, wanted)
	}
}
