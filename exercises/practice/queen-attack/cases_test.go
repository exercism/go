package queenattack

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 647c1ed queen-attack: Add test case to catch an incomplete diagonal check (#1885)

type testCase struct {
	description string
	pos1, pos2  string
	expected    bool
}

var testCases = []testCase{
	{
		description: "cannot attack",
		pos1:        "e3",
		pos2:        "g7",
		expected:    false,
	},
	{
		description: "can attack on same row",
		pos1:        "e3",
		pos2:        "g3",
		expected:    true,
	},
	{
		description: "can attack on same column",
		pos1:        "f5",
		pos2:        "f3",
		expected:    true,
	},
	{
		description: "can attack on first diagonal",
		pos1:        "c3",
		pos2:        "e1",
		expected:    true,
	},
	{
		description: "can attack on second diagonal",
		pos1:        "c3",
		pos2:        "b4",
		expected:    true,
	},
	{
		description: "can attack on third diagonal",
		pos1:        "c3",
		pos2:        "b2",
		expected:    true,
	},
	{
		description: "can attack on fourth diagonal",
		pos1:        "h2",
		pos2:        "g1",
		expected:    true,
	},
	{
		description: "cannot attack if falling diagonals are only the same when reflected across the longest falling diagonal",
		pos1:        "b5",
		pos2:        "f3",
		expected:    false,
	},
}
