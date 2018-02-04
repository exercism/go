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

func TestSolve_bonus(t *testing.T) {
	for _, tc := range noSolutionCases {
		runTestCase(t, tc)
	}
}
