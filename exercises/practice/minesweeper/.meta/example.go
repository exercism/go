package minesweeper

func Annotate(b []string) []string {
	result := make([]string, len(b))
	dirX := []int{}
	dirY := []int{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				dirX = append(dirX, i)
				dirY = append(dirY, j)
			}
		}
	}
	inGrid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < len(b) && j < len(b[0])
	}
	countForSquare := func(i, j int) int {
		result := 0
		for k := range dirX {
			i1, j1 := i+dirX[k], j+dirY[k]
			if inGrid(i1, j1) && b[i1][j1] == '*' {
				result++
			}
		}
		return result
	}
	for i, row := range b {
		resultRow := make([]rune, len(b[i]))
		for j, c := range row {
			if row[j] == '*' {
				resultRow[j] = c
			} else {
				count := countForSquare(i, j)
				if count == 0 {
					resultRow[j] = ' '
				} else {
					resultRow[j] = rune('0' + count)
				}
			}
		}
		result[i] = string(resultRow)
	}
	return result
}
