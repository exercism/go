package sorting

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	panic("Please implement DescribeNumber")
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	panic("Please implement DescribeNumberBox")
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	panic("Please implement ExtractFancyNumber")
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	panic("Please implement DescribeFancyNumberBox")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	panic("Please implement DescribeAnything")
}
