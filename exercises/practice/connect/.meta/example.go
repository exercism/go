package connect

import (
	"errors"
	"fmt"
	"strings"
)

const (
	white = 1 << iota
	black
	connectedWhite
	connectedBlack
)

type colorFlags struct {
	color     int8
	connected int8
}

var flagsBlack = colorFlags{
	color:     black,
	connected: connectedBlack,
}

var flagsWhite = colorFlags{
	color:     white,
	connected: connectedWhite,
}

type coord struct {
	x int
	y int
}

type board struct {
	height int
	width  int
	fields [][]int8
}

func newBoard(lines []string) (board, error) {
	if len(lines) < 1 {
		return board{}, errors.New("No lines given")
	}
	height := len(lines)
	if len(lines[0]) < 1 {
		return board{}, errors.New("First line is empty string")
	}
	width := len(lines[0])
	// This trick for 2D arrays comes from Effective Go
	fields := make([][]int8, height)
	fieldsBacker := make([]int8, height*width)
	for i := range fields {
		fields[i], fieldsBacker = fieldsBacker[:width], fieldsBacker[width:]
	}
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case 'X':
				fields[y][x] = black
			case 'O':
				fields[y][x] = white
			}
			// No need for default, zero value already means no stone
		}
	}
	return board{
		height: height,
		width:  width,
		fields: fields,
	}, nil
}

// Whether there is a stone of the given color at the given location.
//
// Returns both whether there is a stone of the correct color and
// whether the connected flag was set for it.
func (b board) at(c coord, cf colorFlags) (isCorrectColor, isConnected bool) {
	f := b.fields[c.y][c.x]
	return f&cf.color == cf.color,
		f&cf.connected == cf.connected
}

func (b board) markConnected(c coord, cf colorFlags) {
	b.fields[c.y][c.x] |= cf.connected
}

func (b board) validCoord(c coord) bool {
	return c.x >= 0 && c.x < b.width && c.y >= 0 && c.y < b.height
}

func (b board) neighbors(c coord) []coord {
	coords := make([]coord, 0, 6)
	dirs := []coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, 1}, {1, -1}}
	for _, dir := range dirs {
		nc := coord{x: c.x + dir.x, y: c.y + dir.y}
		if b.validCoord(nc) {
			coords = append(coords, nc)
		}
	}
	return coords
}

func (b board) startCoords(cf colorFlags) []coord {
	if cf.color == white {
		coords := make([]coord, b.width)
		for i := 0; i < b.width; i++ {
			coords[i] = coord{x: i}
		}
		return coords
	}
	coords := make([]coord, b.height)
	for i := 0; i < b.height; i++ {
		coords[i] = coord{y: i}
	}
	return coords
}

func (b board) isTargetCoord(c coord, cf colorFlags) bool {
	if cf.color == white {
		return c.y == b.height-1
	}
	return c.x == b.width-1
}

func (b board) evaluate(c coord, cf colorFlags) bool {
	stone, connected := b.at(c, cf)
	if stone && !connected {
		b.markConnected(c, cf)
		if b.isTargetCoord(c, cf) {
			return true
		}
		for _, nc := range b.neighbors(c) {
			if b.evaluate(nc, cf) {
				return true
			}
		}
	}
	return false
}

// Helper for debugging.
func (b board) dump() {
	for y := 0; y < b.height; y++ {
		spaces := strings.Repeat(" ", y)
		chars := make([]string, b.width)
		for x := 0; x < b.width; x++ {
			switch {
			case b.fields[y][x]&white == white:
				if b.fields[y][x]&connectedWhite == connectedWhite {
					chars[x] = "O"
				} else {
					chars[x] = "o"
				}
			case b.fields[y][x]&black == black:
				if b.fields[y][x]&connectedBlack == connectedBlack {
					chars[x] = "X"
				} else {
					chars[x] = "x"
				}
			default:
				chars[x] = "."
			}
		}
		fmt.Printf("%s%s\n", spaces, strings.Join(chars, " "))
	}
}

// ResultOf evaluates the board and return the winner, "black" or
// "white". If there's no winnner ResultOf returns "".
func ResultOf(lines []string) (string, error) {
	board, err := newBoard(lines)
	if err != nil {
		return "", err
	}
	for _, c := range board.startCoords(flagsBlack) {
		if board.evaluate(c, flagsBlack) {
			return "X", nil
		}
	}
	for _, c := range board.startCoords(flagsWhite) {
		if board.evaluate(c, flagsWhite) {
			return "O", nil
		}
	}
	return "", nil
}
