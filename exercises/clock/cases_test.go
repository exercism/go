package clock

// Source: exercism/problem-specifications
// Commit: 8b96944 clock: remove obsolete sentence from comment
// Problem Specifications Version: 2.2.1

// Create a new clock with an initial time
var timeTests = []struct {
	h, m int
	want string
}{
	{8, 0, "08:00"},        // on the hour
	{11, 9, "11:09"},       // past the hour
	{24, 0, "00:00"},       // midnight is zero hours
	{25, 0, "01:00"},       // hour rolls over
	{100, 0, "04:00"},      // hour rolls over continuously
	{1, 60, "02:00"},       // sixty minutes is next hour
	{0, 160, "02:40"},      // minutes roll over
	{0, 1723, "04:43"},     // minutes roll over continuously
	{25, 160, "03:40"},     // hour and minutes roll over
	{201, 3001, "11:01"},   // hour and minutes roll over continuously
	{72, 8640, "00:00"},    // hour and minutes roll over to exactly midnight
	{-1, 15, "23:15"},      // negative hour
	{-25, 0, "23:00"},      // negative hour rolls over
	{-91, 0, "05:00"},      // negative hour rolls over continuously
	{1, -40, "00:20"},      // negative minutes
	{1, -160, "22:20"},     // negative minutes roll over
	{1, -4820, "16:40"},    // negative minutes roll over continuously
	{-25, -160, "20:20"},   // negative hour and minutes both roll over
	{-121, -5810, "22:10"}, // negative hour and minutes both roll over continuously
}

// Add minutes
var addTests = []struct {
	h, m, a int
	want    string
}{
	{10, 0, 3, "10:03"},    // add minutes
	{6, 41, 0, "06:41"},    // add no minutes
	{0, 45, 40, "01:25"},   // add to next hour
	{10, 0, 61, "11:01"},   // add more than one hour
	{0, 45, 160, "03:25"},  // add more than two hours with carry
	{23, 59, 2, "00:01"},   // add across midnight
	{5, 32, 1500, "06:32"}, // add more than one day (1500 min = 25 hrs)
	{1, 1, 3500, "11:21"},  // add more than two days
}

// Subtract minutes
var subtractTests = []struct {
	h, m, a int
	want    string
}{
	{10, 3, 3, "10:00"},    // subtract minutes
	{10, 3, 30, "09:33"},   // subtract to previous hour
	{10, 3, 70, "08:53"},   // subtract more than an hour
	{0, 3, 4, "23:59"},     // subtract across midnight
	{0, 0, 160, "21:20"},   // subtract more than two hours
	{6, 15, 160, "03:35"},  // subtract more than two hours with borrow
	{5, 32, 1500, "04:32"}, // subtract more than one day (1500 min = 25 hrs)
	{2, 20, 3000, "00:20"}, // subtract more than two days
}

// Compare two clocks for equality
type hm struct{ h, m int }

var eqTests = []struct {
	c1, c2 hm
	want   bool
}{
	// clocks with same time
	{
		hm{15, 37},
		hm{15, 37},
		true,
	},
	// clocks a minute apart
	{
		hm{15, 36},
		hm{15, 37},
		false,
	},
	// clocks an hour apart
	{
		hm{14, 37},
		hm{15, 37},
		false,
	},
	// clocks with hour overflow
	{
		hm{10, 37},
		hm{34, 37},
		true,
	},
	// clocks with hour overflow by several days
	{
		hm{3, 11},
		hm{99, 11},
		true,
	},
	// clocks with negative hour
	{
		hm{22, 40},
		hm{-2, 40},
		true,
	},
	// clocks with negative hour that wraps
	{
		hm{17, 3},
		hm{-31, 3},
		true,
	},
	// clocks with negative hour that wraps multiple times
	{
		hm{13, 49},
		hm{-83, 49},
		true,
	},
	// clocks with minute overflow
	{
		hm{0, 1},
		hm{0, 1441},
		true,
	},
	// clocks with minute overflow by several days
	{
		hm{2, 2},
		hm{2, 4322},
		true,
	},
	// clocks with negative minute
	{
		hm{2, 40},
		hm{3, -20},
		true,
	},
	// clocks with negative minute that wraps
	{
		hm{4, 10},
		hm{5, -1490},
		true,
	},
	// clocks with negative minute that wraps multiple times
	{
		hm{6, 15},
		hm{6, -4305},
		true,
	},
	// clocks with negative hours and minutes
	{
		hm{7, 32},
		hm{-12, -268},
		true,
	},
	// clocks with negative hours and minutes that wrap
	{
		hm{18, 7},
		hm{-54, -11513},
		true,
	},
}
