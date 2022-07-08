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
		[][]int{[]int{9, 8, 7}, []int{5, 3, 2}, []int{6, 6, 7}},
		[]Pair{
			{
				2,
				1,
			},
		},
	},

	{
		"Can identify that empty matrix has no saddle points",
		[][]int{[]int{}},
		[]Pair{},
	},

	{
		"Can identify lack of saddle points when there are none",
		[][]int{[]int{1, 2, 3}, []int{3, 1, 2}, []int{2, 3, 1}},
		[]Pair{},
	},

	{
		"Can identify multiple saddle points in a column",
		[][]int{[]int{4, 5, 4}, []int{3, 5, 5}, []int{1, 5, 4}},
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
		[][]int{[]int{6, 7, 8}, []int{5, 5, 5}, []int{7, 5, 6}},
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
		[][]int{[]int{8, 7, 9}, []int{6, 7, 6}, []int{3, 2, 5}},
		[]Pair{
			{
				3,
				3,
			},
		},
	},

	{
		"Can identify saddle points in a non square matrix",
		[][]int{[]int{3, 1, 3}, []int{3, 2, 4}},
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
		[][]int{[]int{2}, []int{1}, []int{4}, []int{1}},
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
		[][]int{[]int{2, 5, 3, 5}},
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
