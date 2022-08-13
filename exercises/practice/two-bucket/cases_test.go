package twobucket

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 4482b50 two-bucket: test inability to reach the goal (#1580)

type bucketTestCase struct {
	description   string
	bucketOne     int
	bucketTwo     int
	goal          int
	startBucket   string
	goalBucket    string
	moves         int
	otherBucket   int
	expectedError string
}

var testCases = []bucketTestCase{
	{
		description:   "Measure using bucket one of size 3 and bucket two of size 5 - start with bucket one",
		bucketOne:     3,
		bucketTwo:     5,
		goal:          1,
		startBucket:   "one",
		goalBucket:    "one",
		moves:         4,
		otherBucket:   5,
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 3 and bucket two of size 5 - start with bucket two",
		bucketOne:     3,
		bucketTwo:     5,
		goal:          1,
		startBucket:   "two",
		goalBucket:    "two",
		moves:         8,
		otherBucket:   3,
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 7 and bucket two of size 11 - start with bucket one",
		bucketOne:     7,
		bucketTwo:     11,
		goal:          2,
		startBucket:   "one",
		goalBucket:    "one",
		moves:         14,
		otherBucket:   11,
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 7 and bucket two of size 11 - start with bucket two",
		bucketOne:     7,
		bucketTwo:     11,
		goal:          2,
		startBucket:   "two",
		goalBucket:    "two",
		moves:         18,
		otherBucket:   7,
		expectedError: "",
	},
	{
		description:   "Measure one step using bucket one of size 1 and bucket two of size 3 - start with bucket two",
		bucketOne:     1,
		bucketTwo:     3,
		goal:          3,
		startBucket:   "two",
		goalBucket:    "two",
		moves:         1,
		otherBucket:   0,
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 2 and bucket two of size 3 - start with bucket one and end with bucket two",
		bucketOne:     2,
		bucketTwo:     3,
		goal:          3,
		startBucket:   "one",
		goalBucket:    "two",
		moves:         2,
		otherBucket:   2,
		expectedError: "",
	},
	{
		description:   "Not possible to reach the goal",
		bucketOne:     6,
		bucketTwo:     15,
		goal:          5,
		startBucket:   "one",
		goalBucket:    "",
		moves:         0,
		otherBucket:   0,
		expectedError: "impossible",
	},
	{
		description:   "With the same buckets but a different goal, then it is possible",
		bucketOne:     6,
		bucketTwo:     15,
		goal:          9,
		startBucket:   "one",
		goalBucket:    "two",
		moves:         10,
		otherBucket:   0,
		expectedError: "",
	},
	{
		description:   "Goal larger than both buckets is impossible",
		bucketOne:     5,
		bucketTwo:     7,
		goal:          8,
		startBucket:   "one",
		goalBucket:    "",
		moves:         0,
		otherBucket:   0,
		expectedError: "impossible",
	},
}
