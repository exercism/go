package matrix

// Source: exercism/problem-specifications
// Commit: 2e820e1 Auto-format portions of some JSON files (#1967)

var testCases = []struct {
	description    string
	input          [][]int
	expectedOutput []Pair
}{

	{
		"Can identify single saddle point",
		[][]int{
			{9, 8, 7}, {5, 3, 2}, {6, 6, 7},
		},
		[]Pair{
			{
				2,
				1,
			},
		},
	},

	{
		"Can identify that empty matrix has no saddle points",
		[][]int{
			{},
		},
		[]Pair{},
	},

	{
		"Can identify lack of saddle points when there are none",
		[][]int{
			{1, 2, 3}, {3, 1, 2}, {2, 3, 1},
		},
		[]Pair{},
	},

	{
		"Can identify multiple saddle points in a column",
		[][]int{
			{4, 5, 4}, {3, 5, 5}, {1, 5, 4},
		},
		[]Pair{
			{
				1,
				2,
			},
			{
				2,
				2,
			},
			{
				3,
				2,
			},
		},
	},

	{
		"Can identify multiple saddle points in a row",
		[][]int{
			{6, 7, 8}, {5, 5, 5}, {7, 5, 6},
		},
		[]Pair{
			{
				2,
				1,
			},
			{
				2,
				2,
			},
			{
				2,
				3,
			},
		},
	},

	{
		"Can identify saddle point in bottom right corner",
		[][]int{
			{8, 7, 9}, {6, 7, 6}, {3, 2, 5},
		},
		[]Pair{
			{
				3,
				3,
			},
		},
	},

	{
		"Can identify saddle points in a non square matrix",
		[][]int{
			{3, 1, 3}, {3, 2, 4},
		},
		[]Pair{
			{
				1,
				3,
			},
			{
				1,
				1,
			},
		},
	},

	{
		"Can identify that saddle points in a single column matrix are those with the minimum value",
		[][]int{
			{2}, {1}, {4}, {1},
		},
		[]Pair{
			{
				2,
				1,
			},
			{
				4,
				1,
			},
		},
	},

	{
		"Can identify that saddle points in a single row matrix are those with the maximum value",
		[][]int{
			{2, 5, 3, 5},
		},
		[]Pair{
			{
				1,
				2,
			},
			{
				1,
				4,
			},
		},
	},
}
