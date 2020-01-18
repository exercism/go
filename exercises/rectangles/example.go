package rectangles

func checkHorizontalSide(diagram []string, start, end, rowIndex int) bool {
	for i := start; i < end; i++ {
		s := diagram[rowIndex][i]
		if s != '+' && s != '-' {
			return false
		}
	}
	return true
}

func checkVerticalSide(diagram []string, start, end, colIndex int) bool {
	for i := start; i < end; i++ {
		s := diagram[i][colIndex]
		if s != '+' && s != '|' {
			return false
		}
	}
	return true
}

// Count counts rectangles in given diagram as lines
func Count(diagram []string) int {
	var count, rows, cols = 0, len(diagram), 0
	if rows > 0 {
		cols = len(diagram[0])
	}

	for y := 0; y < rows-1; y++ {
		for x := 0; x < cols-1; x++ {
			if diagram[y][x] == '+' {
				for j := y + 1; j < rows; j++ {
					for i := x + 1; i < cols; i++ {
						if diagram[j][i] == '+' && diagram[y][i] == '+' && diagram[j][x] == '+' {
							if checkHorizontalSide(diagram, x+1, i, y) &&
								checkHorizontalSide(diagram, x+1, i, j) &&
								checkVerticalSide(diagram, y+1, j, x) &&
								checkVerticalSide(diagram, y+1, j, i) {
								count++
							}
						}
					}
				}
			}
		}
	}
	return count
}
