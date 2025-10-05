package listops

import (
	"reflect"
	"testing"
)

var foldTestCases = []struct {
	name     string
	property string
	fn       func(int, int) int
	initial  int
	list     IntList
	want     int
}{
	{
		name:     "empty list",
		property: "Foldl",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "Foldl",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "Foldl",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     0,
		list:     []int{2, 5},
	},
	{
		name:     "empty list",
		property: "Foldr",
		fn:       func(x, y int) int { return x * y },
		initial:  2,
		want:     2,
		list:     []int{},
	},
	{
		name:     "direction independent function applied to non-empty list",
		property: "Foldr",
		fn:       func(x, y int) int { return x + y },
		initial:  5,
		want:     15,
		list:     []int{1, 2, 3, 4},
	},
	{
		name:     "direction dependent function applied to non-empty list",
		property: "Foldr",
		fn:       func(x, y int) int { return x / y },
		initial:  5,
		want:     2,
		list:     []int{2, 5},
	},
}

func TestFold(t *testing.T) {
	for _, tc := range foldTestCases {
		t.Run(tc.name, func(t *testing.T) {
			var got int
			if tc.property == "Foldr" {
				got = tc.list.Foldr(tc.fn, tc.initial)
			} else {
				got = tc.list.Foldl(tc.fn, tc.initial)
			}
			if got != tc.want {
				t.Fatalf("%s() = %d, want: %d\ntestcase name: %s", tc.property, got, tc.want, tc.name)
			}
		})
	}
}

var filterTestCases = []struct {
	name     string
	property string
	fn       func(int) bool
	list     IntList
	want     IntList
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
	for _, tc := range filterTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Filter(tc.fn)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("IntList(%v).Filter()\n got: %v\nwant: %v\ntestcase name: %s", tc.list, got, tc.want, tc.name)
			}
		})
	}
}

var lengthTestCases = []struct {
	name     string
	property string
	list     IntList
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
	for _, tc := range lengthTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.list.Length(); tc.want != got {
				t.Fatalf("IntList(%v).Length() = %d, want: %d", tc.list, got, tc.want)
			}
		})
	}
}

var mapTestCases = []struct {
	name     string
	property string
	list     IntList
	fn       func(int) int
	want     IntList
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
	for _, tc := range mapTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.list.Map(tc.fn); !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("IntList(%v).Map()\n got: %v\nwant: %v\ntestcase name: %s", tc.list, got, tc.want, tc.name)
			}
		})
	}
}

var reverseTestCases = []struct {
	name     string
	property string
	list     IntList
	want     IntList
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
	for _, tc := range reverseTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Reverse()
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("IntList(%v).Reverse()\n got: %v\nwant: %v", tc.list, got, tc.want)
			}
		})
	}
}

var appendTestCases = []struct {
	name       string
	property   string
	list       IntList
	appendThis IntList
	want       IntList
}{
	{
		name:       "empty list",
		property:   "append",
		list:       []int{},
		appendThis: []int{},
		want:       []int{},
	},
	{
		name:       "empty list to list",
		property:   "append",
		list:       []int{},
		appendThis: []int{1, 2, 3, 4},
		want:       []int{1, 2, 3, 4},
	},
	{
		name:       "non-empty lists",
		property:   "append",
		list:       []int{1, 2},
		appendThis: []int{2, 3, 4, 5},
		want:       []int{1, 2, 2, 3, 4, 5},
	},
}

func TestAppendMethod(t *testing.T) {
	for _, tc := range appendTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Append(tc.appendThis)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("IntList(%v).Append()\n got: %v\nwant: %v", tc.list, got, tc.want)
			}
		})
	}
}

var concatTestCases = []struct {
	name     string
	property string
	list     IntList
	args     []IntList
	want     IntList
}{
	{
		name:     "empty list",
		property: "concat",
		list:     []int{},
		args:     []IntList{},
		want:     []int{},
	},
	{
		name:     "list of lists",
		property: "concat",
		list:     []int{1, 2},
		args:     []IntList{[]int{3}, []int{}, []int{4, 5, 6}},
		want:     []int{1, 2, 3, 4, 5, 6},
	},
}

func TestConcatMethod(t *testing.T) {
	for _, tc := range concatTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Concat(tc.args)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("IntList(%v).Concat(%v)\n got: %v\nwant: %v", tc.list, tc.args, got, tc.want)
			}
		})
	}
}
