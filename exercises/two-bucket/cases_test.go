package twobucket

// Source: exercism/problem-specifications
// Commit: edbc86b two-bucket: apply lowerCamelCase rule to variables
// Problem Specifications Version: 1.4.0

type bucketTestCase struct {
	description   string
	bucketOne     int
	bucketTwo     int
	goal          int
	startBucket   string
	goalBucket    string
	moves         int
	otherBucket   int
	errorExpected bool // always false for generated test cases.
}

var testCases = []bucketTestCase{
	{
		"Measure using bucket one of size 3 and bucket two of size 5 - start with bucket one",
		3, 5, 1, "one", "one", 4, 5, false,
	},
	{
		"Measure using bucket one of size 3 and bucket two of size 5 - start with bucket two",
		3, 5, 1, "two", "two", 8, 3, false,
	},
	{
		"Measure using bucket one of size 7 and bucket two of size 11 - start with bucket one",
		7, 11, 2, "one", "one", 14, 11, false,
	},
	{
		"Measure using bucket one of size 7 and bucket two of size 11 - start with bucket two",
		7, 11, 2, "two", "two", 18, 7, false,
	},
	{
		"Measure one step using bucket one of size 1 and bucket two of size 3 - start with bucket two",
		1, 3, 3, "two", "two", 1, 0, false,
	},
	{
		"Measure using bucket one of size 2 and bucket two of size 3 - start with bucket one and end with bucket two",
		2, 3, 3, "one", "two", 2, 2, false,
	},
}
