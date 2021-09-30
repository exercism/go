package sorting

type NumberBox interface {
	Number() int
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	panic("Please implement DescribeNumber")
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	panic("Please implement DescribeNumberBox")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	panic("Please implement DescribeAnything")
}
