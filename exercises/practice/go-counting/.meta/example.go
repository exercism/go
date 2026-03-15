package gocounting

import "errors"

type AllTerritories struct {
	Black [][2]int
	White [][2]int
	None  [][2]int
}

type Game struct {
	board         []string
	width, height int
}

func NewGame(board []string) *Game {
	height := len(board)
	width := 0
	if height > 0 {
		width = len(board[0])
	}
	return &Game{board, width, height}
}

func (g *Game) onBoard(x, y int) bool {
	return x >= 0 && y >= 0 && x < g.width && y < g.height
}

func (g *Game) neighbors(x, y int) [][2]int {
	offsets := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ns := make([][2]int, 0, 4)
	for _, offset := range offsets {
		nx, ny := x+offset[0], y+offset[1]
		if g.onBoard(nx, ny) {
			ns = append(ns, [2]int{nx, ny})
		}
	}
	return ns
}

func (g *Game) Territory(x, y int) (string, [][2]int, error) {
	if !g.onBoard(x, y) {
		return "", nil, errors.New("invalid coordinate")
	}
	if g.board[y][x] != ' ' {
		return "NONE", nil, nil
	}

	var territory [][2]int
	seen := map[[2]int]bool{[2]int{x, y}: true}
	colors := map[byte]bool{' ': true}
	q := [][2]int{[2]int{x, y}}

	for len(q) != 0 {
		x, y = q[0][0], q[0][1]
		q = q[1:]
		if g.board[y][x] == ' ' {
			territory = append(territory, [2]int{x, y})
			for _, n := range g.neighbors(x, y) {
				if seen[n] {
					continue
				}
				seen[n] = true
				q = append(q, n)
			}
		} else {
			colors[g.board[y][x]] = true
		}
	}

	owner := "NONE"
	if len(colors) > 2 {
	} else if colors['B'] {
		owner = "BLACK"
	} else if colors['W'] {
		owner = "WHITE"
	}
	return owner, territory, nil
}

func (g *Game) Territories() AllTerritories {
	t := AllTerritories{}
	seen := make(map[[2]int]bool)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if seen[[2]int{x, y}] {
				continue
			}
			owner, terr, _ := g.Territory(x, y)
			if owner == "BLACK" {
				t.Black = append(t.Black, terr...)
			} else if owner == "WHITE" {
				t.White = append(t.White, terr...)
			} else {
				t.None = append(t.None, terr...)
			}
			for _, p := range terr {
				seen[p] = true
			}
		}
	}
	return t
}
