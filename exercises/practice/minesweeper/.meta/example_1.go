package minesweeper

func Annotate(b []string) []string {
	dirX := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	dirY := []int{1, 1, 1, 0, -1, -1, -1, 0}
	inGrid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < len(b) && j < len(b[0])
	}
	result := make([][]rune, len(b))
	for i := range result {
		result[i] = make([]rune, len(b[i]))
		for j := range result[i] {
			result[i][j] = ' '
		}
	}
	incSquare := func(i, j int) {
		switch result[i][j] {
		case ' ':
			result[i][j] = '1'
		case '*':
			return
		default:
			result[i][j]++
		}
	}
	for i := range b {
		for j, c := range b[i] {
			if c == '*' {
				result[i][j] = c
				for k := range dirX {
					i1, j1 := i+dirX[k], j+dirY[k]
					if inGrid(i1, j1) {
						incSquare(i1, j1)
					}
				}
			}
		}
	}
	resultStrings := make([]string, len(result))
	for i := range resultStrings {
		resultStrings[i] = string(result[i])
	}
	return resultStrings
}

/*
func IsValid(board []string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	width := len(board[0])
	for i := range board {
		if len(board[i]) != width {
			return false
		}
		for _, c := range board[i] {
			switch c {
			case ' ', '*':
				continue
			default:
				return false
			}
		}
	}
	return true
}
*/
