package stateoftictactoe

import (
	"testing"
)

func TestStateOfTicTacToe(t *testing.T) {
	for _, c := range testCases {
		result, err := StateOfTicTacToe(c.board)
		if c.expectedErr != "" {
			if err == nil || c.expectedErr != err.Error() {
				t.Fatalf(`FAIL: %s
	Expected error: %s
	Got: %v`, c.description, c.expectedErr, err)
			}
		} else if c.expected != result {
			t.Fatalf(`FAIL: %s
    Board: %#v
	Expected: %s
    Got: %#v`, c.description, c.board, c.expected, result)
		}
		t.Logf("PASS: %s", c.description)
	}
}
