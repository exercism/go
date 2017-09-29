package minesweeper

import (
	"bytes"
	"testing"
)

var validBoards = []string{
	// zero size board
	`
++
++`,

	// empty board
	`
+---+
|   |
|   |
|   |
+---+`,

	// board full of mines
	`
+---+
|***|
|***|
|***|
+---+`,

	// surrounded
	`
+---+
|***|
|*8*|
|***|
+---+`,

	// horizontal line
	`
+-----+
|1*2*1|
+-----+`,

	// vertical line
	`
+-+
|1|
|*|
|2|
|*|
|1|
+-+`,

	// cross
	`
+-----+
| 2*2 |
|25*52|
|*****|
|25*52|
| 2*2 |
+-----+`,

	// Large 6x6 board with 8 bombs
	`
+------+
|1*22*1|
|12*322|
| 123*2|
|112*4*|
|1*22*2|
|111111|
+------+`,

	// Large 5x5 board with 7 bombs
	`
+-----+
|1*2*1|
|11322|
| 12*2|
|12*4*|
|1*3*2|
+-----+`,

	// Small 1x5 board with 2 bombs
	`
+-+
|*|
|2|
|*|
|1|
| |
+-+`,

	// 1x1 square with 1 bomb
	`
+-+
|*|
+-+`,

	// 2x2 square with 4 bombs
	`
+--+
|**|
|**|
+--+`,

	// 5x5 square with 2 bombs
	`
+-----+
|  111|
|  1*1|
|  111|
|111  |
|1*1  |
+-----+`,
}

var badBoards = []string{
	// Unaligned
	`
+-+
| |
|*  |
|  |
+-+`,

	// incomplete border
	`
+-----+
*   * |
+-- --+`,

	// Unknown characters
	`
+-----+
|X  * |
+-----+`,
}

// constructor.  take board as a string, return Board with digits cleared.
func clear(s string) Board {
	b := []byte(s)
	for i, c := range b {
		if c >= '0' && c <= '9' {
			b[i] = ' '
		}
	}
	return bytes.Split(b, []byte{'\n'})[1:]
}

func TestValid(t *testing.T) {
	for _, ref := range validBoards {
		b := clear(ref)
		t.Log(b)
		if err := b.Count(); err != nil {
			var _ error = err
			t.Fatalf("Count() returned %q, want nil.", err)
		}
		if res := b.String(); res != ref {
			t.Fatalf("Count() result: %s,\n  want:%s.", res, ref)
		}
	}
}

func TestBad(t *testing.T) {
	for _, ref := range badBoards {
		b := clear(ref)
		t.Log(b)
		if b.Count() == nil {
			t.Fatal("Count() returned nil, want error.")
		}
	}
}

func BenchmarkCount(b *testing.B) {
	m := clear(`
+------+
|1*22*1|
|12*322|
| 123*2|
|112*4*|
|1*22*2|
|111111|
+------+`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Count()
	}
}
