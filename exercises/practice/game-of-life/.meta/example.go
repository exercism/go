package gameoflife

func get(matrix [][]int, x, y int) int {
	if x < 0 || y < 0 || y >= len(matrix) || x >= len(matrix[0]) {
		return 0
	}
	return matrix[y][x]
}

func Tick(matrix [][]int) [][]int {
	offsets := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /* 0,0 */, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	out := make([][]int, len(matrix))
	for y := range matrix {
		out[y] = make([]int, len(matrix[0]))
		for x := range matrix[0] {
			sum := 0
			for _, offset := range offsets {
				sum += get(matrix, x+offset[0], y+offset[1])
			}
			if sum == 3 || (get(matrix, x, y) == 1 && sum == 2) {
				out[y][x] = 1
			}
		}
	}
	return out
}
