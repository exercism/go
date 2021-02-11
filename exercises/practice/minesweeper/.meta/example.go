package minesweeper

import (
	"bytes"
	"errors"
)

func (b Board) Count() error {
	if len(b) < 2 {
		return errors.New("need top and bottom border")
	}
	last := len(b) - 1
	if len(b[last]) != len(b[0]) {
		return errors.New("top and bottom border must be same size")
	}
	if err := border(b[0]); err != nil {
		return err
	}
	if err := border(b[last]); err != nil {
		return err
	}
	w := len(b[0])
	lc := w - 1
	for r := 1; r < last; r++ {
		row := b[r]
		if len(row) != w {
			return errors.New("all rows must be same size")
		}
		if row[0] != '|' || row[lc] != '|' {
			return errors.New("left and right borders must be '|'")
		}
		for c := 1; c < lc; c++ {
			switch row[c] {
			default:
				return errors.New("only ' ' and '*' allowed")
			case '*':
			case ' ':
				n := 0
				for rm, rz := r-1, r+1; rm <= rz; rm++ {
					rowm := b[rm]
					for cm, cz := c-1, c+1; cm <= cz; cm++ {
						if rowm[cm] == '*' {
							n++
						}
					}
				}
				if n > 0 {
					row[c] = '0' + byte(n)
				}
			}
		}
	}
	return nil
}

func border(b []byte) error {
	switch {
	case len(b) < 2 || b[0] != '+' || b[len(b)-1] != '+':
		return errors.New("need '+' in corners")
	case bytes.Count(b, []byte{'-'}) != len(b)-2:
		return errors.New("top and bottom borders must be '-'")
	}
	return nil
}
