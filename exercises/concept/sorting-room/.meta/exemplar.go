package sorting

import "fmt"

type NumberBox interface {
	Number() int
}

type BooleanBox interface {
	Boolean() bool
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// DescribeBoolean should return a string describing the boolean.
func DescribeBoolean(b bool) string {
	if b {
		return "This is true"
	}
	return "This is false"
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

// DescribeBooleanBox should return a string describing the BooleanBox.
func DescribeBooleanBox(bb BooleanBox) string {
	if bb.Boolean() {
		return "This box is true"
	}
	return "This box is false"
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		return DescribeNumber(float64(i))
	case float64:
		return DescribeNumber(i)
	case bool:
		return DescribeBoolean(i)
	case NumberBox:
		return DescribeNumberBox(i)
	case BooleanBox:
		return DescribeBooleanBox(i)
	default:
		return "Return to sender"
	}
}
