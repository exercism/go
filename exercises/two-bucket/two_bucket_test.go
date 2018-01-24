package twobucket

import "testing"

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		g, m, other := Solve(tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket)
		if g != tc.goalBucket || m != tc.moves || other != tc.otherBucket {
			t.Fatalf("FAIL: %s\nSolve(%d, %d, %d, %q)\nExpected: %q, %d, %d\nActual: %q, %d, %d",
				tc.description,
				tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket,
				tc.goalBucket, tc.moves, tc.otherBucket,
				g, m, other)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Solve(tc.bucketOne, tc.bucketTwo, tc.goal, tc.startBucket)
		}
	}
}
