package sorting

import "testing"

func TestDescribeNumber(t *testing.T) {
	tests := []struct {
		description string
		input       float64
		want        string
	}{
		{
			description: "Describe 4.1",
			input:       4.1,
			want:        "This is the number 4.1",
		},
		{
			description: "Describe -3.2",
			input:       -3.2,
			want:        "This is the number -3.2",
		},
		{
			description: "Pads to single decimal place",
			input:       4.0,
			want:        "This is the number 4.0",
		},
		{
			description: "Truncates to single decimal place",
			input:       7.11,
			want:        "This is the number 7.1",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := DescribeNumber(test.input); got != test.want {
				t.Errorf("DescribeNumber(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestDescribeBoolean(t *testing.T) {
	tests := []struct {
		description string
		input       bool
		want        string
	}{
		{
			description: "Describe true",
			input:       true,
			want:        "This is true",
		},
		{
			description: "Describe false",
			input:       false,
			want:        "This is false",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := DescribeBoolean(test.input); got != test.want {
				t.Errorf("DescribeBoolean(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

type testNumberBox struct {
	n int
}

func (nb testNumberBox) Number() int {
	return nb.n
}

func TestDescribeNumberBox(t *testing.T) {
	tests := []struct {
		description string
		input       NumberBox
		want        string
	}{
		{
			description: "Describe NumberBox with 4",
			input:       testNumberBox{4},
			want:        "This is a box containing the number 4.0",
		},
		{
			description: "Describe NumberBox with -3",
			input:       testNumberBox{-3},
			want:        "This is a box containing the number -3.0",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := DescribeNumberBox(test.input); got != test.want {
				t.Errorf("DescribeNumberBox(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

type testBooleanBox struct {
	b bool
}

func (bb testBooleanBox) Boolean() bool {
	return bb.b
}

func TestDescribeBooleanBox(t *testing.T) {
	tests := []struct {
		description string
		input       BooleanBox
		want        string
	}{
		{
			description: "Describe BooleanBox with true",
			input:       testBooleanBox{true},
			want:        "This box contains true",
		},
		{
			description: "Describe false",
			input:       testBooleanBox{false},
			want:        "This box contains false",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := DescribeBooleanBox(test.input); got != test.want {
				t.Errorf("DescribeBooleanBox(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestDescribeAnything(t *testing.T) {
	tests := []struct {
		description string
		input       interface{}
		want        string
	}{
		{
			description: "Describe 7.2",
			input:       7.2,
			want:        "This is the number 7.2",
		},
		{
			description: "Describe false",
			input:       false,
			want:        "This is false",
		},
		{
			description: "Describe NumberBox with 16",
			input:       testNumberBox{16},
			want:        "This is a box containing the number 16.0",
		},
		{
			description: "BooleanBox with true",
			input:       testBooleanBox{true},
			want:        "This box contains true",
		},
		{
			description: "Something unknown is labelled return to sender",
			input:       "something we did not anticipate",
			want:        "Return to sender",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := DescribeAnything(test.input); got != test.want {
				t.Errorf("DescribeAnything(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}
