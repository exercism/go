package connect

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var testCases = []struct {
	description string
	board       []string
	expected    string
}{
	{
		description: "an empty board has no winner",
		board: []string{
			". . . . .",
			" . . . . .",
			"  . . . . .",
			"   . . . . .",
			"    . . . . .",
		},
		expected: "",
	},
	{
		description: "X can win on a 1x1 board",
		board: []string{
			"X",
		},
		expected: "X",
	},
	{
		description: "O can win on a 1x1 board",
		board: []string{
			"O",
		},
		expected: "O",
	},
	{
		description: "only edges does not make a winner",
		board: []string{
			"O O O X",
			" X . . X",
			"  X . . X",
			"   X O O O",
		},
		expected: "",
	},
	{
		description: "illegal diagonal does not make a winner",
		board: []string{
			"X O . .",
			" O X X X",
			"  O X O .",
			"   . O X .",
			"    X X O O",
		},
		expected: "",
	},
	{
		description: "nobody wins crossing adjacent angles",
		board: []string{
			"X . . .",
			" . X O .",
			"  O . X O",
			"   . O . X",
			"    . . O .",
		},
		expected: "",
	},
	{
		description: "X wins crossing from left to right",
		board: []string{
			". O . .",
			" O X X X",
			"  O X O .",
			"   X X O X",
			"    . O X .",
		},
		expected: "X",
	},
	{
		description: "O wins crossing from top to bottom",
		board: []string{
			". O . .",
			" O X X X",
			"  O O O .",
			"   X X O X",
			"    . O X .",
		},
		expected: "O",
	},
	{
		description: "X wins using a convoluted path",
		board: []string{
			". X X . .",
			" X . X . X",
			"  . X . X .",
			"   . X X . .",
			"    O O O O O",
		},
		expected: "X",
	},
	{
		description: "X wins using a spiral path",
		board: []string{
			"O X X X X X X X X",
			" O X O O O O O O O",
			"  O X O X X X X X O",
			"   O X O X O O O X O",
			"    O X O X X X O X O",
			"     O X O O O X O X O",
			"      O X X X X X O X O",
			"       O O O O O O O X O",
			"        X X X X X X X X O",
		},
		expected: "X",
	},
}
