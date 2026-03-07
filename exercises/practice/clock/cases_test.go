package clock

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var timeTestCases = []struct {
	description string
	h, m        int
	expected    string
}{
	{
		description: "on the hour",
		h:           8,
		m:           0,
		expected:    "08:00",
	},
	{
		description: "past the hour",
		h:           11,
		m:           9,
		expected:    "11:09",
	},
	{
		description: "midnight is zero hours",
		h:           24,
		m:           0,
		expected:    "00:00",
	},
	{
		description: "hour rolls over",
		h:           25,
		m:           0,
		expected:    "01:00",
	},
	{
		description: "hour rolls over continuously",
		h:           100,
		m:           0,
		expected:    "04:00",
	},
	{
		description: "sixty minutes is next hour",
		h:           1,
		m:           60,
		expected:    "02:00",
	},
	{
		description: "minutes roll over",
		h:           0,
		m:           160,
		expected:    "02:40",
	},
	{
		description: "minutes roll over continuously",
		h:           0,
		m:           1723,
		expected:    "04:43",
	},
	{
		description: "hour and minutes roll over",
		h:           25,
		m:           160,
		expected:    "03:40",
	},
	{
		description: "hour and minutes roll over continuously",
		h:           201,
		m:           3001,
		expected:    "11:01",
	},
	{
		description: "hour and minutes roll over to exactly midnight",
		h:           72,
		m:           8640,
		expected:    "00:00",
	},
	{
		description: "negative hour",
		h:           -1,
		m:           15,
		expected:    "23:15",
	},
	{
		description: "negative hour rolls over",
		h:           -25,
		m:           0,
		expected:    "23:00",
	},
	{
		description: "negative hour rolls over continuously",
		h:           -91,
		m:           0,
		expected:    "05:00",
	},
	{
		description: "negative minutes",
		h:           1,
		m:           -40,
		expected:    "00:20",
	},
	{
		description: "negative minutes roll over",
		h:           1,
		m:           -160,
		expected:    "22:20",
	},
	{
		description: "negative minutes roll over continuously",
		h:           1,
		m:           -4820,
		expected:    "16:40",
	},
	{
		description: "negative sixty minutes is previous hour",
		h:           2,
		m:           -60,
		expected:    "01:00",
	},
	{
		description: "negative hour and minutes both roll over",
		h:           -25,
		m:           -160,
		expected:    "20:20",
	},
	{
		description: "negative hour and minutes both roll over continuously",
		h:           -121,
		m:           -5810,
		expected:    "22:10",
	},
}

var addTestCases = []struct {
	description      string
	h, m, addedValue int
	expected         string
}{
	{
		description: "add minutes",
		h:           10,
		m:           0,
		addedValue:  3,
		expected:    "10:03",
	},
	{
		description: "add no minutes",
		h:           6,
		m:           41,
		addedValue:  0,
		expected:    "06:41",
	},
	{
		description: "add to next hour",
		h:           0,
		m:           45,
		addedValue:  40,
		expected:    "01:25",
	},
	{
		description: "add more than one hour",
		h:           10,
		m:           0,
		addedValue:  61,
		expected:    "11:01",
	},
	{
		description: "add more than two hours with carry",
		h:           0,
		m:           45,
		addedValue:  160,
		expected:    "03:25",
	},
	{
		description: "add across midnight",
		h:           23,
		m:           59,
		addedValue:  2,
		expected:    "00:01",
	},
	{
		description: "add more than one day (1500 min = 25 hrs)",
		h:           5,
		m:           32,
		addedValue:  1500,
		expected:    "06:32",
	},
	{
		description: "add more than two days",
		h:           1,
		m:           1,
		addedValue:  3500,
		expected:    "11:21",
	},
}

var subtractTestCases = []struct {
	description           string
	h, m, subtractedValue int
	expected              string
}{
	{
		description:     "subtract minutes",
		h:               10,
		m:               3,
		subtractedValue: 3,
		expected:        "10:00",
	},
	{
		description:     "subtract to previous hour",
		h:               10,
		m:               3,
		subtractedValue: 30,
		expected:        "09:33",
	},
	{
		description:     "subtract more than an hour",
		h:               10,
		m:               3,
		subtractedValue: 70,
		expected:        "08:53",
	},
	{
		description:     "subtract across midnight",
		h:               0,
		m:               3,
		subtractedValue: 4,
		expected:        "23:59",
	},
	{
		description:     "subtract more than two hours",
		h:               0,
		m:               0,
		subtractedValue: 160,
		expected:        "21:20",
	},
	{
		description:     "subtract more than two hours with borrow",
		h:               6,
		m:               15,
		subtractedValue: 160,
		expected:        "03:35",
	},
	{
		description:     "subtract more than one day (1500 min = 25 hrs)",
		h:               5,
		m:               32,
		subtractedValue: 1500,
		expected:        "04:32",
	},
	{
		description:     "subtract more than two days",
		h:               2,
		m:               20,
		subtractedValue: 3000,
		expected:        "00:20",
	},
}

// Compare two clocks for equality
type hm struct{ h, m int }

var equalTestCases = []struct {
	description string
	c1, c2      hm
	expected    bool
}{
	{
		description: "clocks with same time",
		c1:          hm{15, 37},
		c2:          hm{15, 37},
		expected:    true,
	},
	{
		description: "clocks a minute apart",
		c1:          hm{15, 36},
		c2:          hm{15, 37},
		expected:    false,
	},
	{
		description: "clocks an hour apart",
		c1:          hm{14, 37},
		c2:          hm{15, 37},
		expected:    false,
	},
	{
		description: "clocks with hour overflow",
		c1:          hm{10, 37},
		c2:          hm{34, 37},
		expected:    true,
	},
	{
		description: "clocks with hour overflow by several days",
		c1:          hm{3, 11},
		c2:          hm{99, 11},
		expected:    true,
	},
	{
		description: "clocks with negative hour",
		c1:          hm{22, 40},
		c2:          hm{-2, 40},
		expected:    true,
	},
	{
		description: "clocks with negative hour that wraps",
		c1:          hm{17, 3},
		c2:          hm{-31, 3},
		expected:    true,
	},
	{
		description: "clocks with negative hour that wraps multiple times",
		c1:          hm{13, 49},
		c2:          hm{-83, 49},
		expected:    true,
	},
	{
		description: "clocks with minute overflow",
		c1:          hm{0, 1},
		c2:          hm{0, 1441},
		expected:    true,
	},
	{
		description: "clocks with minute overflow by several days",
		c1:          hm{2, 2},
		c2:          hm{2, 4322},
		expected:    true,
	},
	{
		description: "clocks with negative minute",
		c1:          hm{2, 40},
		c2:          hm{3, -20},
		expected:    true,
	},
	{
		description: "clocks with negative minute that wraps",
		c1:          hm{4, 10},
		c2:          hm{5, -1490},
		expected:    true,
	},
	{
		description: "clocks with negative minute that wraps multiple times",
		c1:          hm{6, 15},
		c2:          hm{6, -4305},
		expected:    true,
	},
	{
		description: "clocks with negative hours and minutes",
		c1:          hm{7, 32},
		c2:          hm{-12, -268},
		expected:    true,
	},
	{
		description: "clocks with negative hours and minutes that wrap",
		c1:          hm{18, 7},
		c2:          hm{-54, -11513},
		expected:    true,
	},
	{
		description: "full clock and zeroed clock",
		c1:          hm{24, 0},
		c2:          hm{0, 0},
		expected:    true,
	},
}
