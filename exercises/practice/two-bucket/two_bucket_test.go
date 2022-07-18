package twobucket

import "testing"

func TestSolve(t *testing.T) {
	for _, tc := range append(testCases, errorTestCases...) {
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
		"Invalid first bucket size",
		0, 5, 5, "one", "one", 1, 0, "invalid first bucket size",
	},
	{
		"Invalid second bucket size",
		3, 0, 3, "one", "one", 1, 0, "invalid second bucket size",
	},
	{
		"Invalid goal amount",
		1, 1, 0, "one", "one", 0, 1, "invalid goal amount",
	},
	{
		"Invalid start bucket name",
		3, 5, 1, "three", "one", 4, 5, "invalid start bucket name",
	},
}
