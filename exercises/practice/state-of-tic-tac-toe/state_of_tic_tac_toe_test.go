package stateoftictactoe

import (
	"testing"
)

func TestStateOfTicTacToe(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.description, func(t *testing.T) {
			result, err := StateOfTicTacToe(c.board)
			switch {
			case c.wantErr && err == nil:
				t.Fatalf("\n Board: %#v \n Expected error but got nil", c.board)

			case !c.wantErr && err != nil:
				t.Fatalf("\n Board: %#v \n Expected no errors but got error: %v", c.board, err)

			case c.expected != result:
				t.Fatalf("\n Board: %#v \n Expected: %#v \n Got: %#v", c.board, c.expected, result)
			}
		})
	}
}
