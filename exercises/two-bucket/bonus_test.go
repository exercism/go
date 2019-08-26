
// +build bonus

package twobucket

import "testing"

var noSolutionCases = []bucketTestCase{
	{
		"No solution case 1",
		2, 4, 1, "two", "", 0, 0, true,
	},
	{
		"No solution case 2",
		3, 6, 1, "one", "", 0, 0, true,
	},
}

var bonusCases = []bucketTestCase{
	{"Adding Complexity",
		333333, 666667, 1, "one", "one", 1999992, 666667, false,
	},
}

func TestSolve_noSolution(t *testing.T) {
	for _, tc := range noSolutionCases {
		runTestCase(t, tc)
	}
}

func TestSolve_bonus(t *testing.T) {
	for _, tc := range bonusCases {
		runTestCase(t, tc)
	}
}
