package queenattack

import "fmt"

func CanQueenAttack(w, b string) (attack bool, err error) {
	if err = valSq(w); err != nil {
		return false, err
	}
	if err = valSq(b); err != nil {
		return false, err
	}
	if w == b {
		return false, fmt.Errorf("queens on same square")
	}
	if w[0] == b[0] || // same file
		w[1] == b[1] { // same rank
		return true, nil
	}
	df := w[0] - b[0]
	dr := w[1] - b[1]
	return df == dr || // same white-kingside diagonal
			df+dr == 0, // same white-queenside diagonal
		nil
}

func valSq(s string) error {
	if len(s) != 2 ||
		s[0] < 'a' || s[0] > 'h' ||
		s[1] < '1' || s[1] > '8' {
		return fmt.Errorf("invalid square: %q", s)
	}
	return nil
}
