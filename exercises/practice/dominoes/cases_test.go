package dominoes

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var testCases = []struct {
	description string
	dominoes    []Domino
	valid       bool // true => can chain, false => cannot chain
}{
	{
		description: "empty input = empty output",
		dominoes:    []Domino{},
		valid:       true,
	},
	{
		description: "singleton input = singleton output",
		dominoes:    []Domino{{1, 1}},
		valid:       true,
	},
	{
		description: "singleton that can't be chained",
		dominoes:    []Domino{{1, 2}},
		valid:       false,
	},
	{
		description: "three elements",
		dominoes:    []Domino{{1, 2}, {3, 1}, {2, 3}},
		valid:       true,
	},
	{
		description: "can reverse dominoes",
		dominoes:    []Domino{{1, 2}, {1, 3}, {2, 3}},
		valid:       true,
	},
	{
		description: "can't be chained",
		dominoes:    []Domino{{1, 2}, {4, 1}, {2, 3}},
		valid:       false,
	},
	{
		description: "disconnected - simple",
		// This meets the requirement of being possibly-Euclidean.
		// All vertices have even degree.
		// Nevertheless, there is no chain here, as there's no way to get from 1 to 2.
		// This test (and the two following) prevent solutions from using the even-degree test as the sole criterion,
		// as that is not a sufficient condition.
		dominoes: []Domino{{1, 1}, {2, 2}},
		valid:    false,
	},
	{
		description: "disconnected - double loop",
		dominoes:    []Domino{{1, 2}, {2, 1}, {3, 4}, {4, 3}},
		valid:       false,
	},
	{
		description: "disconnected - single isolated",
		dominoes:    []Domino{{1, 2}, {2, 3}, {3, 1}, {4, 4}},
		valid:       false,
	},
	{
		description: "need backtrack",
		// Some solutions may make a chain out of (1, 2), (2, 3), (3, 1)
		// then decide that since there are no more dominoes containing a 1,
		// there is no chain possible.
		// There is indeed a chain here, so this test checks for this line of reasoning.
		// You need to place the (2, 4) after the (1, 2) rather than the (2, 3).
		dominoes: []Domino{{1, 2}, {2, 3}, {3, 1}, {2, 4}, {2, 4}},
		valid:    true,
	},
	{
		description: "separate loops",
		dominoes:    []Domino{{1, 2}, {2, 3}, {3, 1}, {1, 1}, {2, 2}, {3, 3}},
		valid:       true,
	},
	{
		description: "nine elements",
		dominoes:    []Domino{{1, 2}, {5, 3}, {3, 1}, {1, 2}, {2, 4}, {1, 6}, {2, 3}, {3, 4}, {5, 6}},
		valid:       true,
	},
}
