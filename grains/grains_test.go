package grains

import "testing"

var squareTests = []struct {
	input    int
	expected uint64
}{
	{1, 1},
	{2, 2},
	{3, 4},
	{4, 8},
	{16, 32768},
	{32, 2147483648},
	{64, 9223372036854775808},
}

func TestSquare(t *testing.T) {
	for _, test := range squareTests {
		if actual := Square(test.input); actual != test.expected {
			t.Errorf("Square(%d) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func TestTotal(t *testing.T) {
	var expected uint64 = 18446744073709551615
	if actual := Total(); actual != expected {
		t.Errorf("Total() expected %d, Actual %d", expected, actual)
	}
}
