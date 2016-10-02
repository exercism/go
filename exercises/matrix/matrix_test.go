// For the Matrix exercise in Go you have to do a few things not mentioned
// in the README.
//
// 1. You must implement a constructor and methods Rows() and Cols() as
//    described in the README, but Rows() and Cols must return results that
//    are independent from the original matrix.  That is, you should be able
//    to do as you please with the results without affecting the matrix.
//
// 2. You must implement a method Set(row, col, val) for setting a matrix
//    element.
//
// 3. As usual in Go, you must detect and return error conditions.

package matrix

import (
	"reflect"
	"testing"
)

const targetTestVersion = 1

var tests = []struct {
	in   string
	ok   bool
	rows [][]int
	cols [][]int
}{
	{"1 2\n10 20",
		true,
		[][]int{
			{1, 2},
			{10, 20},
		},
		[][]int{
			{1, 10},
			{2, 20},
		},
	},
	{"9 7\n8 6",
		true,
		[][]int{
			{9, 7},
			{8, 6},
		},
		[][]int{
			{9, 8},
			{7, 6},
		},
	},
	{"9 8 7\n19 18 17",
		true,
		[][]int{
			{9, 8, 7},
			{19, 18, 17},
		},
		[][]int{
			{9, 19},
			{8, 18},
			{7, 17},
		},
	},
	{"1 4 9\n16 25 36",
		true,
		[][]int{
			{1, 4, 9},
			{16, 25, 36},
		},
		[][]int{
			{1, 16},
			{4, 25},
			{9, 36},
		},
	},
	{"1 2 3\n4 5 6\n7 8 9\n 8 7 6",
		true,
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{8, 7, 6},
		},
		[][]int{
			{1, 4, 7, 8},
			{2, 5, 8, 7},
			{3, 6, 9, 6},
		},
	},
	{"89 1903 3\n18 3 1\n9 4 800",
		true,
		[][]int{
			{89, 1903, 3},
			{18, 3, 1},
			{9, 4, 800},
		},
		[][]int{
			{89, 18, 9},
			{1903, 3, 4},
			{3, 1, 800},
		},
	},
	{"1 2 3", // valid, 1 row, 3 columns
		true,
		[][]int{
			{1, 2, 3},
		},
		[][]int{
			{1},
			{2},
			{3},
		},
	},
	{"1\n2\n3", // valid, 3 rows, 1 column
		true,
		[][]int{
			{1},
			{2},
			{3},
		},
		[][]int{
			{1, 2, 3},
		},
	},
	{"0", // valid, 1 row, 1 column
		true,
		[][]int{
			{0},
		},
		[][]int{
			{0},
		},
	},
	{"9223372036854775808", false, nil, nil}, // overflows int64
	{"1 2\n10 20 30", false, nil, nil},       // uneven rows
	{"\n3 4\n5 6", false, nil, nil},          // first row empty
	{"1 2\n\n5 6", false, nil, nil},          // middle row empty
	{"1 2\n3 4\n", false, nil, nil},          // last row empty
	{"2.7", false, nil, nil},                 // non-int
	{"cat", false, nil, nil},                 // non-numeric
	// undefined
	// {"\n\n", // valid?, 3 rows, 0 columns
	// {"",     // valid?, 0 rows, 0 columns
}

func TestNew(t *testing.T) {
	for _, test := range tests {
		m, err := New(test.in)
		switch {
		case err != nil:
			var _ error = err
			if test.ok {
				t.Fatalf("New(%q) returned error %q.  Error not expected",
					test.in, err)
			}
		case !test.ok:
			t.Fatalf("New(%q) = %v, %v.  Expected non-nil error.",
				test.in, m, err)
		case m == nil:
			t.Fatalf("New(%q) = %v, want non-nil *Matrix",
				test.in, m)
		}
	}
}

func TestRows(t *testing.T) {
	for _, test := range tests {
		if !test.ok {
			continue
		}
		m, err := New(test.in)
		if err != nil {
			t.Skip("Need working New for TestRows")
		}
		r := m.Rows()
		if len(r) == 0 && len(test.rows) == 0 {
			continue // agreement, and nothing more to test
		}
		if !reflect.DeepEqual(r, test.rows) {
			t.Fatalf("New(%q).Rows() = %v, want %v", test.in, r, test.rows)
		}
		if len(r[0]) == 0 {
			continue // not currently in test data, but anyway
		}
		r[0][0]++
		if !reflect.DeepEqual(m.Rows(), test.rows) {
			t.Fatalf("Matrix.Rows() returned slice based on Matrix " +
				"representation.  Want independent copy of element data.")
		}
	}
}

func TestCols(t *testing.T) {
	for _, test := range tests {
		if !test.ok {
			continue
		}
		m, err := New(test.in)
		if err != nil {
			t.Skip("Need working New for TestCols")
		}
		c := m.Cols()
		if len(c) == 0 && len(test.cols) == 0 {
			continue // agreement, and nothing more to test
		}
		if !reflect.DeepEqual(c, test.cols) {
			t.Fatalf("New(%q).Cols() = %v, want %v", test.in, c, test.cols)
		}
		if len(c[0]) == 0 {
			continue // not currently in test data, but anyway
		}
		c[0][0]++
		if !reflect.DeepEqual(m.Cols(), test.cols) {
			t.Fatalf("Matrix.Cols() returned slice based on Matrix " +
				"representation.  Want independent copy of element data.")
		}
	}
}

func TestSet(t *testing.T) {
	s := "1 2 3\n4 5 6\n7 8 9"
	m, err := New(s)
	if err != nil {
		t.Skip("Need working New for TestSet")
	}
	xr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if !reflect.DeepEqual(m.Rows(), xr) {
		t.Skip("Need working Rows for TestSet")
	}
	xc := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	if !reflect.DeepEqual(m.Cols(), xc) {
		t.Skip("Need working Cols for TestSet")
	}
	// test each corner, each side, and an interior element
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			m, _ = New(s)
			val := 10 + r*3 + c
			if ok := m.Set(r, c, val); !ok {
				t.Fatalf("Matrix(%q).Set(%d, %d, %d) returned !ok, want ok.",
					s, r, c, val)
			}
			xr = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
			xc = [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
			xr[r][c] = val
			xc[c][r] = val
			if res := m.Rows(); !reflect.DeepEqual(res, xr) {
				t.Fatalf("Matrix(%q).Set(%d, %d, %d), Rows() = %v, want %v",
					s, r, c, val, res, xr)
			}
			if res := m.Cols(); !reflect.DeepEqual(res, xc) {
				t.Fatalf("Matrix(%q).Set(%d, %d, %d), Cols() = %v, want %v",
					s, r, c, val, res, xc)
			}
		}
	}
	// test 1 and 2 off each corner and side
	m, _ = New(s)
	for _, r := range []int{-2, -1, 0, 3, 4} {
		for _, c := range []int{-2, -1, 0, 3, 4} {
			if r == 0 && c == 0 {
				continue
			}
			if ok := m.Set(r, c, 0); ok {
				t.Fatalf("Matrix(%q).Set(%d, %d, 0) = ok, want !ok", s, r, c)
			}
		}
	}
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}
