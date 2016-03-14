package clock

// Source: exercism/x-common
// Commit: 269f498 Merge pull request #48 from soniakeys/custom-set-json

// Test creating a new clock with an initial time.
var timeTests = []struct {
	h, m int
	want string
}{
	{8, 0, "08:00"},      // on the hour
	{9, 0, "09:00"},      // on the hour
	{11, 9, "11:09"},     // past the hour
	{11, 19, "11:19"},    // past the hour
	{24, 0, "00:00"},     // midnight is zero hours
	{25, 0, "01:00"},     // hour rolls over
	{1, 60, "02:00"},     // sixty minutes is next hour
	{0, 160, "02:40"},    // minutes roll over
	{25, 160, "03:40"},   // hour and minutes roll over
	{-1, 15, "23:15"},    // negative hour
	{-25, 0, "23:00"},    // negative hour rolls over
	{1, -40, "00:20"},    // negative minutes
	{1, -160, "22:20"},   // negative minutes roll over
	{-25, -160, "20:20"}, // negative hour and minutes both roll over
}

// Test adding and subtracting minutes.
var addTests = []struct {
	h, m, a int
	want    string
}{
	{10, 0, 3, "10:03"},     // add minutes
	{0, 45, 40, "01:25"},    // add to next hour
	{10, 0, 61, "11:01"},    // add more than one hour
	{23, 59, 2, "00:01"},    // add across midnight
	{5, 32, 1500, "06:32"},  // add more than one day (1500 min = 25 hrs)
	{0, 45, 160, "03:25"},   // add more than two hours with carry
	{10, 3, -3, "10:00"},    // subtract minutes
	{10, 3, -30, "09:33"},   // subtract to previous hour
	{10, 3, -70, "08:53"},   // subtract more than an hour
	{0, 3, -4, "23:59"},     // subtract across midnight
	{0, 0, -160, "21:20"},   // subtract more than two hours
	{5, 32, -1500, "04:32"}, // subtract more than one day (1500 min = 25 hrs)
	{6, 15, -160, "03:35"},  // subtract more than two hours with borrow
}

// Construct two separate clocks, set times, test if they are equal.
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
	// clocks set 24 hours apart
	{
		hm{10, 37},
		hm{34, 37},
		true,
	},
	// clocks with minute overflow
	{
		hm{0, 1},
		hm{0, 1441},
		true,
	},
}
