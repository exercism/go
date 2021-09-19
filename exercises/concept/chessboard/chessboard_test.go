package chessboard

import (
	"testing"
)

// newChessboard return a *Chessboard for tests
//
//   1 2 3 4 5 6 7 8
// 1 # _ # _ _ _ _ # 1
// 2 _ _ _ _ # _ _ _ 2
// 3 _ _ # _ _ _ _ _ 3
// 4 _ _ _ _ _ _ _ _ 4
// 5 _ _ _ _ _ # _ # 5
// 6 _ _ _ _ _ _ _ _ 6
// 7 _ _ _ # _ _ _ _ 7
// 8 # # # # # # _ # 8
//   1 2 3 4 5 6 7 8
func newChessboard() *Chessboard {
	return &Chessboard{
		1: Rank{true, false, true, false, false, false, false, true},
		2: Rank{false, false, false, false, true, false, false, false},
		3: Rank{false, false, true, false, false, false, false, false},
		4: Rank{false, false, false, false, false, false, false, false},
		5: Rank{false, false, false, false, false, true, false, true},
		6: Rank{false, false, false, false, false, false, false, false},
		7: Rank{false, false, false, true, false, false, false, false},
		8: Rank{true, true, true, true, true, true, false, true},
	}
}

func TestCountInRank(t *testing.T) {
	cb := newChessboard()
	for _, test := range []struct {
		in  int
		out int
	}{
		{1, 3},
		{2, 1},
		{3, 1},
		{4, 0},
		{5, 2},
		{6, 0},
		{7, 1},
		{8, 7},
		{9, 0},
	} {
		if out := cb.CountInRank(test.in); out != test.out {
			t.Errorf(
				"CountInRank(%v) returned %v while %v was expected\n",
				test.in,
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
				"CountInFile(%v) returned %v while %v was expected\n",
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
