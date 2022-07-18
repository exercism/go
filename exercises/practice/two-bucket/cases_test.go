package twobucket

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
		moves:         '\x04',
		otherBucket:   '\x05',
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 3 and bucket two of size 5 - start with bucket two",
		bucketOne:     3,
		bucketTwo:     5,
		goal:          1,
		startBucket:   "two",
		goalBucket:    "two",
		moves:         '\b',
		otherBucket:   '\x03',
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 7 and bucket two of size 11 - start with bucket one",
		bucketOne:     7,
		bucketTwo:     11,
		goal:          2,
		startBucket:   "one",
		goalBucket:    "one",
		moves:         '\x0e',
		otherBucket:   '\v',
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 7 and bucket two of size 11 - start with bucket two",
		bucketOne:     7,
		bucketTwo:     11,
		goal:          2,
		startBucket:   "two",
		goalBucket:    "two",
		moves:         '\x12',
		otherBucket:   '\a',
		expectedError: "",
	},
	{
		description:   "Measure one step using bucket one of size 1 and bucket two of size 3 - start with bucket two",
		bucketOne:     1,
		bucketTwo:     3,
		goal:          3,
		startBucket:   "two",
		goalBucket:    "two",
		moves:         '\x01',
		otherBucket:   '\x00',
		expectedError: "",
	},
	{
		description:   "Measure using bucket one of size 2 and bucket two of size 3 - start with bucket one and end with bucket two",
		bucketOne:     2,
		bucketTwo:     3,
		goal:          3,
		startBucket:   "one",
		goalBucket:    "two",
		moves:         '\x02',
		otherBucket:   '\x02',
		expectedError: "",
	},
	{
		description:   "Not possible to reach the goal",
		bucketOne:     6,
		bucketTwo:     15,
		goal:          5,
		startBucket:   "one",
		goalBucket:    "",
		moves:         '\x00',
		otherBucket:   '\x00',
		expectedError: "impossible",
	},
	{
		description:   "With the same buckets but a different goal, then it is possible",
		bucketOne:     6,
		bucketTwo:     15,
		goal:          9,
		startBucket:   "one",
		goalBucket:    "two",
		moves:         '\n',
		otherBucket:   '\x00',
		expectedError: "",
	},
	{
		description:   "Goal larger than both buckets is impossible",
		bucketOne:     5,
		bucketTwo:     7,
		goal:          8,
		startBucket:   "one",
		goalBucket:    "",
		moves:         '\x00',
		otherBucket:   '\x00',
		expectedError: "impossible",
	},
}
