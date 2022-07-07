package stateoftictactoe

type state string

const (
	state_win     state = "win"
	state_ongoing state = "ongoing"
	state_draw    state = "draw"
)

func StateOfTicTacToe(board []string) (state, error) {
	panic("Please implement the StateOfTicTacToe function")
}
