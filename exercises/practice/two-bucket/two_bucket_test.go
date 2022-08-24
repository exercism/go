package twobucket

import "testing"

func TestSolve(t *testing.T) {
	for _, tc := range append(testCases, errorTestCases...) {
		runTestCase(t, tc)
	}
}

func runTestCase(t *testing.T, tc bucketTestCase) {
	t.Run(tc.description, func(t *testing.T) {
		g, m, other, err := Solve(tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket)
		switch {
		case tc.expectedError != "":
			if err == nil {
				t.Fatalf("Solve(%d,%d,%d,%q) expected error, got:%q,%d,%d", tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket, g, m, other)
			}
		case err != nil:
			t.Fatalf("Solve(%d,%d,%d,%q) returned error: %v, want:%q,%d,%d", tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket, err, tc.goalBucket, tc.moves, tc.otherBucket)
		case g != tc.goalBucket || m != tc.moves || other != tc.otherBucket:
			t.Fatalf("Solve(%d,%d,%d,%q) = %q,%d,%d, want:%q,%d,%d", tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket, g, m, other, tc.goalBucket, tc.moves, tc.otherBucket)
		}
	})
}

func BenchmarkSolve(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range append(testCases, errorTestCases...) {
			Solve(tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket)
		}
	}
}

var errorTestCases = []bucketTestCase{
	{
		description:   "Invalid first bucket size",
		bucketOne:     0,
		bucketTwo:     5,
		goal:          5,
		startBucket:   "one",
		goalBucket:    "one",
		moves:         1,
		otherBucket:   0,
		expectedError: "invalid first bucket size",
	},
	{
		description:   "Invalid second bucket size",
		bucketOne:     3,
		bucketTwo:     0,
		goal:          3,
		startBucket:   "one",
		goalBucket:    "one",
		moves:         1,
		otherBucket:   0,
		expectedError: "invalid second bucket size",
	},
	{
		description:   "Invalid goal amount",
		bucketOne:     1,
		bucketTwo:     1,
		goal:          0,
		startBucket:   "one",
		goalBucket:    "one",
		moves:         0,
		otherBucket:   1,
		expectedError: "invalid goal amount",
	},
	{
		description:   "Invalid start bucket name",
		bucketOne:     3,
		bucketTwo:     5,
		goal:          1,
		startBucket:   "three",
		goalBucket:    "one",
		moves:         4,
		otherBucket:   5,
		expectedError: "invalid start bucket name",
	},
}
