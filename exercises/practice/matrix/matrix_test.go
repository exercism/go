package matrix

import (
	"reflect"
	"testing"
)

type testCase struct {
	description string
	in          string
	ok          bool
	rows        [][]int
	cols        [][]int
}

var validTestCases = []testCase{
	{
		description: "2 rows, 2 columns",
		in:          "1 2\n10 20",
		ok:          true,
		rows: [][]int{
			{1, 2},
			{10, 20},
		},
		cols: [][]int{
			{1, 10},
			{2, 20},
		},
	},
	{
		description: "2 rows, 2 columns",
		in:          "9 7\n8 6",
		ok:          true,
		rows: [][]int{
			{9, 7},
			{8, 6},
		},
		cols: [][]int{
			{9, 8},
			{7, 6},
		},
	},
	{
		description: "2 rows, 3 columns",
		in:          "9 8 7\n19 18 17",
		ok:          true,
		rows: [][]int{
			{9, 8, 7},
			{19, 18, 17},
		},
		cols: [][]int{
			{9, 19},
			{8, 18},
			{7, 17},
		},
	},
	{
		description: "2 rows, 3 columns",
		in:          "1 4 9\n16 25 36",
		ok:          true,
		rows: [][]int{
			{1, 4, 9},
			{16, 25, 36},
		},
		cols: [][]int{
			{1, 16},
			{4, 25},
			{9, 36},
		},
	},
	{
		description: "4 rows, 3 columns",
		in:          "1 2 3\n4 5 6\n7 8 9\n 8 7 6",
		ok:          true,
		rows: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{8, 7, 6},
		},
		cols: [][]int{
			{1, 4, 7, 8},
			{2, 5, 8, 7},
			{3, 6, 9, 6},
		},
	},
	{
		description: "3 rows, 3 columns",
		in:          "89 1903 3\n18 3 1\n9 4 800",
		ok:          true,
		rows: [][]int{
			{89, 1903, 3},
			{18, 3, 1},
			{9, 4, 800},
		},
		cols: [][]int{
			{89, 18, 9},
			{1903, 3, 4},
			{3, 1, 800},
		},
	},
	{
		description: "1 row, 3 columns",
		in:          "1 2 3",
		ok:          true,
		rows: [][]int{
			{1, 2, 3},
		},
		cols: [][]int{
			{1},
			{2},
			{3},
		},
	},
	{
		description: "3 rows, 1 column",
		in:          "1\n2\n3",
		ok:          true,
		rows: [][]int{
			{1},
			{2},
			{3},
		},
		cols: [][]int{
			{1, 2, 3},
		},
	},
	{
		description: "1 row, 1 column",
		in:          "0",
		ok:          true,
		rows: [][]int{
			{0},
		},
		cols: [][]int{
			{0},
		},
	},
	// undefined
	// {"\n\n", // valid?, 3 rows, 0 columns
	// {"",     // valid?, 0 rows, 0 columns
}

var invalidTestCases = []testCase{
	{description: "int64 overflow", in: "9223372036854775808", ok: false, rows: nil, cols: nil},
	{description: "uneven rows", in: "1 2\n10 20 30", ok: false, rows: nil, cols: nil},
	{description: "first row empty", in: "\n3 4\n5 6", ok: false, rows: nil, cols: nil},
	{description: "middle row empty", in: "1 2\n\n5 6", ok: false, rows: nil, cols: nil},
	{description: "last row empty", in: "1 2\n3 4\n", ok: false, rows: nil, cols: nil},
	{description: "non integer", in: "2.7", ok: false, rows: nil, cols: nil},
	{description: "non numeric", in: "cat", ok: false, rows: nil, cols: nil},
}

func TestNew(t *testing.T) {
	for _, tc := range append(validTestCases, invalidTestCases...) {
		t.Run(tc.description, func(t *testing.T) {
			got, err := New(tc.in)
			switch {
			case !tc.ok:
				if err == nil {
					t.Fatalf("New(%q) expected error, got: %v", tc.in, got)
				}
			case err != nil:
				t.Fatalf("New(%q) returned error %q.  Error not expected", tc.in, err)
			case got == nil:
				t.Fatalf("New(%q) = %v, want non-nil *Matrix", tc.in, got)
			}
		})
	}
}

func TestRows(t *testing.T) {
	for _, tc := range validTestCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := New(tc.in)
			if err != nil {
				t.Fatalf("error in test setup: TestRows needs working New and valid matrix")
			}
			rows := got.Rows()
			if len(rows) == 0 && len(tc.rows) == 0 {
				return // agreement, and nothing more to test
			}
			if !reflect.DeepEqual(rows, tc.rows) {
				t.Fatalf("New(%q).Rows() = %v (type %T), want: %v (type %T)", tc.in, rows, rows, tc.rows, tc.rows)
			}
			if len(rows[0]) == 0 {
				return // not currently in test data, but anyway
			}
			rows[0][0]++
			if !reflect.DeepEqual(got.Rows(), tc.rows) {
				t.Fatalf("Matrix.Rows() returned slice based on Matrix representation. Want independent copy of element data.")
			}
		})
	}
}

func TestCols(t *testing.T) {
	for _, tc := range validTestCases {
		t.Run(tc.description, func(t *testing.T) {
			m, err := New(tc.in)
			if err != nil {
				t.Fatalf("error in test setup: TestCols needs working New and valid matrix")
			}
			cols := m.Cols()
			if len(cols) == 0 && len(tc.cols) == 0 {
				return // agreement, and nothing more to test
			}
			if !reflect.DeepEqual(cols, tc.cols) {
				t.Fatalf("New(%q).Cols() = %v (type %T), want: %v (type %T)", tc.in, cols, cols, tc.cols, tc.cols)
			}
			if len(cols[0]) == 0 {
				return // not currently in test data, but anyway
			}
			cols[0][0]++
			if !reflect.DeepEqual(m.Cols(), tc.cols) {
				t.Fatalf("Matrix.Cols() returned slice based on Matrix representation. Want independent copy of element data.")
			}
		})
	}
}

func TestSet(t *testing.T) {
	s := "1 2 3\n4 5 6\n7 8 9"
	m, err := New(s)
	if err != nil {
		t.Fatalf("error in test setup: TestSet needs working New and valid matrix")
	}
	xr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if !reflect.DeepEqual(m.Rows(), xr) {
		t.Fatalf("error in test setup: TestSet needs working Rows")
	}
	xc := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	if !reflect.DeepEqual(m.Cols(), xc) {
		t.Fatalf("error in test setup: TestSet needs working Cols and valid matrix")
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

func BenchmarkNew(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	var matrix Matrix
	for i := 0; i < b.N; i++ {
		var err error
		matrix, err = New("1 2 3 10 11\n4 5 6 11 12\n7 8 9 12 13\n 8 7 6 13 14")
		if err != nil {
			b.Fatalf("Failed to create the matrix: %v", err)
		}
	}
	if matrix == nil {
		b.Fatalf("No matrix parsed")
	}
}

func BenchmarkRows(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	matrix, err := New("1 2 3\n4 5 6\n7 8 9\n 8 7 6")
	if err != nil {
		b.Fatalf("Failed to create the matrix: %v", err)
	}
	b.ResetTimer()
	var rows [][]int
	for i := 0; i < b.N; i++ {
		rows = matrix.Rows()
	}
	if len(rows) != 4 {
		b.Fatalf("Incorrect number of rows returned: %v", rows)
	}
}

func BenchmarkCols(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	matrix, err := New("1 2 3 10 11\n4 5 6 11 12\n7 8 9 12 13\n 8 7 6 13 14")
	if err != nil {
		b.Fatalf("Failed to create the matrix: %v", err)
	}
	b.ResetTimer()
	var cols [][]int
	for i := 0; i < b.N; i++ {
		cols = matrix.Cols()
	}
	if len(cols) != 5 {
		b.Fatalf("Incorrect number of columns returned: %v", cols)
	}
}
