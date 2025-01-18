package stateoftictactoe

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	panic("Please implement the StateOfTicTacToe function")
}
