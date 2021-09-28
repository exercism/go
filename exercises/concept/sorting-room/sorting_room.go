package sorting

type NumberBox interface {
	Number() int
}

type BooleanBox interface {
	Boolean() bool
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	panic("Please implement DescribeNumber")
}

// DescribeBoolean should return a string describing the boolean.
func DescribeBoolean(b bool) string {
	panic("Please implement DescribeBoolean")
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	panic("Please implement DescribeNumberBox")
}

// DescribeBooleanBox should return a string describing the BooleanBox.
func DescribeBooleanBox(nb NumberBox) string {
	panic("Please implement DescribeBooleanBox")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	panic("Please implement DescribeAnything")
}
