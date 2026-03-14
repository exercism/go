package hangman

type Game struct{}

func NewGame(word string) *Game {
	panic("Please implement the NewGame() function")
}

func (g *Game) Guess(r rune) error {
	panic("Please implement the Guess() function")
}

func (g *Game) MaskedWord() string {
	panic("Please implement the MaskedWord() function")
}

func (g *Game) RemainingGuesses() int {
	panic("Please implement the RemainingGuesses() function")
}

func (g *Game) State() string {
	panic("Please implement the State() function")
}
