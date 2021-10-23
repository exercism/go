package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Maintenance note:  This section should generally match
// exercises/practice/matrix/.meta/example.go.

type Matrix [][]int

func New(s string) (*Matrix, error) {
	var err error
	lines := strings.Split(s, "\n")
	m := make(Matrix, len(lines))
	for i, l := range lines {
		ws := strings.Fields(l)
		if i > 0 && len(ws) != len(m[0]) {
			return nil, errors.New("rows have unequal length")
		}
		row := make([]int, len(ws))
		for i, w := range ws {
			if row[i], err = strconv.Atoi(w); err != nil {
				return nil, errors.New("invalid int in element data")
			}
		}
		m[i] = row
	}
	return &m, nil
}

func (m *Matrix) Set(row, col, val int) (ok bool) {
	mm := *m
	if row < 0 || row >= len(mm) || col < 0 {
		return false
	}
	if cols := len(mm[0]); col >= cols {
		return false
	}
	mm[row][col] = val
	return true
}

func (m *Matrix) Rows() [][]int {
	r := make([][]int, len(*m))
	for i, mr := range *m {
		r[i] = append([]int{}, mr...)
	}
	return r
}

func (m *Matrix) Cols() [][]int {
	mm := *m
	if len(mm) == 0 {
		return nil
	}
	c := make([][]int, len(mm[0]))
	for i := range c {
		col := make([]int, len(mm))
		for j := range col {
			col[j] = mm[j][i]
		}
		c[i] = col
	}
	return c
}

// End of the section affected by the maintenance-note

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
