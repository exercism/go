//go:build bonus
// +build bonus

package twobucket

import "testing"

var noSolutionCases = []bucketTestCase{
	{
		description:   "No solution case 1",
		bucketOne:     2,
		bucketTwo:     4,
		goal:          1,
		startBucket:   "two",
		goalBucket:    "",
		moves:         0,
		otherBucket:   0,
		expectedError: "no solution",
	},
	{
		description:   "No solution case 2",
		bucketOne:     3,
		bucketTwo:     6,
		goal:          1,
		startBucket:   "one",
		goalBucket:    "",
		moves:         0,
		otherBucket:   0,
		expectedError: "no solution",
	},
}

func TestSolve_bonus(t *testing.T) {
	for _, tc := range noSolutionCases {
		runTestCase(t, tc)
	}
}
