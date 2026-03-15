package gocounting

type AllTerritories struct {
	Black [][2]int
	White [][2]int
	None  [][2]int
}

type Game struct{}

func NewGame(board []string) *Game {
	panic("Please implement the NewGame() function")
}

func (g *Game) Territory(x, y int) (string, [][2]int, error) {
	panic("Please implement the Territory() function")
}

func (g *Game) Territories() AllTerritories {
	panic("Please implement the Territories() function")
}
