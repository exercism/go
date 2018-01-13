package summultiples

// Source: exercism/problem-specifications
// Commit: fb5b0a1 sum-of-multiples: Apply new "input" policy
// Problem Specifications Version: 1.2.0

var varTests = []struct {
	divisors []int
	limit    int
	sum      int
}{
	{[]int{3, 5}, 1, 0},             // multiples of 3 or 5 up to 1
	{[]int{3, 5}, 4, 3},             // multiples of 3 or 5 up to 4
	{[]int{3}, 7, 9},                // multiples of 3 up to 7
	{[]int{3, 5}, 10, 23},           // multiples of 3 or 5 up to 10
	{[]int{3, 5}, 100, 2318},        // multiples of 3 or 5 up to 100
	{[]int{3, 5}, 1000, 233168},     // multiples of 3 or 5 up to 1000
	{[]int{7, 13, 17}, 20, 51},      // multiples of 7, 13 or 17 up to 20
	{[]int{4, 6}, 15, 30},           // multiples of 4 or 6 up to 15
	{[]int{5, 6, 8}, 150, 4419},     // multiples of 5, 6 or 8 up to 150
	{[]int{5, 25}, 51, 275},         // multiples of 5 or 25 up to 51
	{[]int{43, 47}, 10000, 2203160}, // multiples of 43 or 47 up to 10000
	{[]int{1}, 100, 4950},           // multiples of 1 up to 100
	{[]int{}, 10000, 0},             // multiples of an empty list up to 10000
}
