package gigasecond

// Source: exercism/x-common
// Commit: f362340 Merge pull request #36 from soniakeys/gigasecond-test-data

// Add one gigasecond to the input.
var addCases = []struct {
	in   string
	want string
}{
	{
		"2011-04-25",
		"2043-01-01T01:46:40",
	},
	{
		"1977-06-13",
		"2009-02-19T01:46:40",
	},
	{
		"1959-07-19",
		"1991-03-27T01:46:40",
	},
}
