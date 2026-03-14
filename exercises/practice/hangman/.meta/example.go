package hangman

import (
	"errors"
	"strings"
)

type Game struct {
	state     string
	word      string
	guesses   map[rune]bool
	remaining int
}

func NewGame(word string) *Game {
	return &Game{
		state:     "Ongoing",
		word:      word,
		guesses:   make(map[rune]bool),
		remaining: 9,
	}
}

func (g *Game) Guess(r rune) error {
	if g.state == "Win" {
		return errors.New("cannot guess after the game is won")
	} else if g.state == "Lose" {
		return errors.New("cannot guess after the game is lost")
	}
	if g.guesses[r] || !strings.ContainsRune(g.word, r) {
		if g.remaining > 0 {
			g.remaining--
		} else {
			g.state = "Lose"
		}
	}
	g.guesses[r] = true
	if g.MaskedWord() == g.word {
		g.state = "Win"
	}
	return nil
}

func (g *Game) MaskedWord() string {
	var sb strings.Builder
	for _, r := range g.word {
		if g.guesses[r] {
			sb.WriteRune(r)
		} else {
			sb.WriteRune('_')
		}
	}
	return sb.String()
}

func (g *Game) RemainingGuesses() int {
	return g.remaining
}

func (g *Game) State() string {
	return g.state
}
