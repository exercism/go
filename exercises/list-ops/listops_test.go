package listops

import "reflect"
import "testing"

var foldTestCases = []struct {
	name     string
	property string
	fn       binFunc
	initial  int
	list     IntSlice
	want     int
}{
	{
		name:     "empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldl",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     0,
		list:     []int{2, 5},
	},
	{
		name:     "empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "foldr",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     2,
		list:     []int{2, 5},
	},
}

func TestFold(t *testing.T) {
	var got int
	for _, tt := range foldTestCases {
		if tt.property == "foldr" {
			got = tt.list.Foldr(tt.fn, tt.initial)
		} else {
			got = tt.list.Foldl(tt.fn, tt.initial)
		}
		if got != tt.want {
			t.Fatalf("Build for test case %q for property %s returned %d but was expected to return %d",
				tt.name, tt.property, got, tt.want)
		}
	}

}

var filterTestCases = []struct {
	name     string
	property string
	fn       predFunc
	list     []int
	want     []int
}{
	{
		name:     "empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "filter",
		fn:       func(n int) bool { return n%2 == 1 },
		list:     []int{1, 2, 3, 4, 5},
		want:     []int{1, 3, 5},
	},
}

func TestFilterMethod(t *testing.T) {
	for _, tt := range filterTestCases {
		in := IntSlice(tt.list)
		got := in.Filter(tt.fn)
		if !reflect.DeepEqual(IntSlice(tt.want), got) {
			t.Fatalf("Build for test case %q for property %s returned %v but was expected to return %v",
				tt.name, tt.property, got, tt.want)
		}
	}
}

var lengthTestCases = []struct {
	name     string
	property string
	list     IntSlice
	want     int
}{
	{
		name:     "empty list",
		property: "length",
		list:     []int{},
		want:     0,
	},
	{
		name:     "non-empty list",
		property: "length",
		list:     []int{1, 2, 3, 4},
		want:     4,
	},
}

func TestLengthMethod(t *testing.T) {
	for _, tt := range lengthTestCases {
		got := tt.list.Length()
		if tt.want != got {
			t.Fatalf("Build for test case %q for property %s returned %d but was expected to return %d",
				tt.name, tt.property, got, tt.want)
		}
	}
}

var mapTestCases = []struct {
	name     string
	property string
	list     IntSlice
	fn       unaryFunc
	want     IntSlice
}{
	{
		name:     "empty list",
		property: "map",
		list:     []int{},
		fn:       func(x int) int { return x + 1 },
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "map",
		list:     []int{1, 3, 5, 7},
		fn:       func(x int) int { return x + 1 },
		want:     []int{2, 4, 6, 8},
	},
}

func TestMapMethod(t *testing.T) {
	for _, tt := range mapTestCases {
		got := tt.list.Map(tt.fn)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("Build for test case %q for property %s returned %v but was expected to return %v",
				tt.name, tt.property, got, tt.want)
		}
	}
}

var reverseTestCases = []struct {
	name     string
	property string
	list     IntSlice
	want     IntSlice
}{
	{
		name:     "empty list",
		property: "reverse",
		list:     []int{},
		want:     []int{},
	},
	{
		name:     "non-empty list",
		property: "reverse",
		list:     []int{1, 3, 5, 7},
		want:     []int{7, 5, 3, 1},
	},
}

func TestReverseMethod(t *testing.T) {
	for _, tt := range reverseTestCases {
		got := tt.list.Reverse()
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("Build for test case %q for property %s returned %v but was expected to return %v",
				tt.name, tt.property, got, tt.want)
		}
	}
}
