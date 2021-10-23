package wordsearch

import "fmt"

func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	positions := make(map[string][2][2]int, len(words))
	for _, word := range words {
		span, err := find(word, puzzle)
		if err != nil {
			return nil, err
		}
		positions[word] = span
	}
	return positions, nil
}

type dir struct {
	cx, cy int
}

var (
	north     = dir{0, -1}
	northeast = dir{1, -1}
	east      = dir{1, 0}
	southeast = dir{1, 1}
	south     = dir{0, 1}
	southwest = dir{-1, 1}
	west      = dir{-1, 0}
	northwest = dir{-1, -1}
)

var dirs = []dir{north, northeast, east, southeast, south, southwest, west, northwest}

var zeroSpan = [2][2]int{{0, 0}, {0, 0}}

func find(word string, puzzle []string) ([2][2]int, error) {
	for x := 0; x < len(puzzle[0]); x++ {
		for y := 0; y < len(puzzle); y++ {
			for _, dir := range dirs {
				if span, ok := try(word, puzzle, x, y, dir); ok {
					return span, nil
				}
			}
		}
	}
	return zeroSpan, fmt.Errorf("Didn't find %q", word)
}

func try(word string, puzzle []string, startX, startY int, dir dir) ([2][2]int, bool) {
	x, y := startX, startY
	for _, c := range word {
		if lc, ok := lookup(x, y, puzzle); !ok || lc != c {
			return zeroSpan, false
		}
		x += dir.cx
		y += dir.cy
	}
	// Need to remove the last increment to get the position of the last character.
	return [2][2]int{{startX, startY}, {x - dir.cx, y - dir.cy}}, true
}

func lookup(x, y int, puzzle []string) (rune, bool) {
	if x >= 0 && x < len(puzzle[0]) && y >= 0 && y < len(puzzle) {
		// Yes, yes, this won't work with multi-byte codepoints.
		return rune(puzzle[y][x]), true
	}
	return 0, false
}
