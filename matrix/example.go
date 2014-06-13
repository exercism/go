package matrix

import (
	"errors"
	"strconv"
	"strings"
)

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
