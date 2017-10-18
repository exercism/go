package collatz

var testCases = []struct {
	num         int
	expected    int
	description string
}{
	{1, 0, "zero steps for one"},
	{16, 4, "divide if even"},
	{12, 9, "even and odd steps"},
	{1000000, 152, "Large number of even and odd steps"},
}

var testThrowErrorCases = []struct {
	num         int
	description string
}{
	{0, "zero is an error"},
	{-15, "Only positive numbers are allowed"},
}
