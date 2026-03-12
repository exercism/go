package camicia

import "slices"

type Outcome struct {
	finishes bool
	cards    int
	tricks   int
}

var Penalty = map[string]int{"A": 4, "K": 3, "Q": 2, "J": 1}

type Game struct {
	hands     [2][]string
	middle    []string
	seen      [][]string
	curPlayer int
	cards     int
	tricks    int
	loop      bool
}

// Return if the game should stop.
// The game stops if one player has all the cards of if we are in a loop.
func (g *Game) gameOver() bool {
	if len(g.middle) == 0 && len(g.hands[0])*len(g.hands[1]) == 0 {
		return true
	}
	state := append(g.hands[0], "SEP")
	state = append(state, g.hands[1]...)

	g.loop = false
	for _, prior := range g.seen {
		if slices.Equal(state, prior) {
			g.loop = true
		}
	}
	g.seen = append(g.seen, state)
	return g.loop
}

func (g *Game) endTurn() {
	g.curPlayer = 1 - g.curPlayer
}

// Add the middle pile to the player's hand.
func (g *Game) collectMiddle() {
	g.hands[g.curPlayer] = append(g.hands[g.curPlayer], g.middle...)
	g.tricks++
	g.middle = g.middle[:0]
}

// Move one card from the current player's hand to the middle.
func (g *Game) playOneCard() string {
	card := g.hands[g.curPlayer][0]
	g.hands[g.curPlayer] = g.hands[g.curPlayer][1:]
	g.middle = append(g.middle, card)
	g.cards++
	return card
}

// Play out a penalty.
func (g *Game) resolvePenalty(penalty int) {
	for penalty > 0 && len(g.hands[g.curPlayer]) > 0 {
		card := g.playOneCard()
		penalty--
		if card != "-" {
			g.endTurn()
			penalty = Penalty[card]
		}
	}

	// Collect the pile.
	g.endTurn()
	g.collectMiddle()
}

// Return the outcome of the game.
func (g *Game) play() Outcome {
	for !g.gameOver() {
		if len(g.hands[g.curPlayer]) == 0 {
			g.endTurn()
			g.collectMiddle()
			break
		}

		card := g.playOneCard()
		g.endTurn()
		if card != "-" {
			g.resolvePenalty(Penalty[card])
		}
	}
	return Outcome{
		finishes: !g.loop,
		cards:    g.cards,
		tricks:   g.tricks,
	}
}

func SimulateGame(playerA, playerB []string) Outcome {
	hands := [2][]string{}
	for i, cards := range [][]string{playerA, playerB} {
		hand := []string{}
		for _, card := range cards {
			if _, ok := Penalty[card]; ok {
				hand = append(hand, card)
			} else {
				hand = append(hand, "-")
			}
		}
		hands[i] = hand
	}
	g := &Game{
		hands:  hands,
		middle: []string{},
		seen:   [][]string{},
	}
	got := g.play()
	return got
}
