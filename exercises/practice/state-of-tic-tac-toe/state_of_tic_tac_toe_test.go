package stateoftictactoe

import (
	"testing"
)

func TestStateOfTicTacToe(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.description, func(t *testing.T) {
			result, err := StateOfTicTacToe(c.board)
			if c.wantErr {
				if err == nil {
					t.Fatalf(`
					Board: %#v
					Expected error but got nil`, c.board)
				}
				return
			}
			if c.expected != result {
				t.Fatalf(`
					Board: %#v
					Expected: %#v
					Got: %#v`, c.board, c.expected, result)
			}
		})
	}
}
