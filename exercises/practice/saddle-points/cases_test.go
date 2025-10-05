package matrix

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 2e820e1 Auto-format portions of some JSON files (#1967)

var testCases = []struct {
	description    string
	input          [][]int
	expectedOutput []Pair
}{
	{
		description: "Can identify single saddle point",
		input: [][]int{
			{9, 8, 7}, {5, 3, 2}, {6, 6, 7},
		},
		expectedOutput: []Pair{
			{
				2,
				1,
			},
		},
	},

	{
		description: "Can identify that empty matrix has no saddle points",
		input: [][]int{
			{},
		},
		expectedOutput: []Pair{},
	},

	{
		description: "Can identify lack of saddle points when there are none",
		input: [][]int{
			{1, 2, 3}, {3, 1, 2}, {2, 3, 1},
		},
		expectedOutput: []Pair{},
	},

	{
		description: "Can identify multiple saddle points in a column",
		input: [][]int{
			{4, 5, 4}, {3, 5, 5}, {1, 5, 4},
		},
		expectedOutput: []Pair{
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
		description: "Can identify multiple saddle points in a row",
		input: [][]int{
			{6, 7, 8}, {5, 5, 5}, {7, 5, 6},
		},
		expectedOutput: []Pair{
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
		description: "Can identify saddle point in bottom right corner",
		input: [][]int{
			{8, 7, 9}, {6, 7, 6}, {3, 2, 5},
		},
		expectedOutput: []Pair{
			{
				3,
				3,
			},
		},
	},

	{
		description: "Can identify saddle points in a non square matrix",
		input: [][]int{
			{3, 1, 3}, {3, 2, 4},
		},
		expectedOutput: []Pair{
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
		description: "Can identify that saddle points in a single column matrix are those with the minimum value",
		input: [][]int{
			{2}, {1}, {4}, {1},
		},
		expectedOutput: []Pair{
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
		description: "Can identify that saddle points in a single row matrix are those with the maximum value",
		input: [][]int{
			{2, 5, 3, 5},
		},
		expectedOutput: []Pair{
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
