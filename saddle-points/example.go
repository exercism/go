package matrix

type Pair struct{ r, c int }

func (m *Matrix) Saddle() (p []Pair) {
	cols := m.Cols()
	colMin := make([][]int, len(cols))
	for c, col := range cols {
		colMin[c] = mins(col)
	}
	for r, row := range *m {
		for _, c := range maxs(row) {
			for _, rmin := range colMin[c] {
				if rmin == r {
					p = append(p, Pair{r, c})
					break
				}
			}
		}
	}
	return
}

// min returns indexes of s that are the minimum value
func mins(s []int) []int {
	if len(s) == 0 {
		return nil
	}
	last := len(s) - 1
	minv := s[last]
	xmin := []int{last}
	for x, v := range s[:last] {
		switch {
		case v < minv:
			minv = v
			xmin = xmin[:1]
			xmin[0] = x
		case v == minv:
			xmin = append(xmin, x)
		}
	}
	return xmin
}

// same code as above, but with >
func maxs(s []int) []int {
	if len(s) == 0 {
		return nil
	}
	last := len(s) - 1
	maxv := s[last]
	xmax := []int{last}
	for x, v := range s[:last] {
		switch {
		case v > maxv:
			maxv = v
			xmax = xmax[:1]
			xmax[0] = x
		case v == maxv:
			xmax = append(xmax, x)
		}
	}
	return xmax
}
