package strain

import (
	"reflect"
	"slices"
	"testing"
)

func TestInt(t *testing.T) {
	for _, tc := range intTestCases {
		t.Run(tc.description, func(t *testing.T) {
			fn, name := Discard[int], "Discard"
			if tc.keep {
				fn, name = Keep[int], "Keep"
			}
			data := slices.Clone(tc.input)
			got := fn(data, tc.filter)

			if !slices.Equal(data, tc.input) {
				t.Fatalf("%s(%#v, %q) modified the input data; input data should not change", name, tc.input, tc.filterStr)
			} else if len(got) == 0 && len(tc.expected) == 0 { // treat nil and empty the same
			} else if !slices.Equal(got, tc.expected) {
				t.Fatalf("%s(%#v, %q) = %#v, want %#v",  name, tc.input, tc.filterStr, got, tc.expected)
			}
		})
	}
}

func TestString(t *testing.T) {
	for _, tc := range stringTestCases {
		t.Run(tc.description, func(t *testing.T) {
			fn, name := Discard[string], "Discard"
			if tc.keep {
				fn, name = Keep[string], "Keep"
			}
			data := slices.Clone(tc.input)
			got := fn(data, tc.filter)

			if !slices.Equal(data, tc.input) {
				t.Fatalf("%s(%#v, %q) modified the input data; input data should not change", name, tc.input, tc.filterStr)
			} else if len(got) == 0 && len(tc.expected) == 0 { // treat nil and empty the same
			} else if !slices.Equal(got, tc.expected) {
				t.Fatalf("%s(%#v, %q) = %#v, want %#v",  name, tc.input, tc.filterStr, got, tc.expected)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	for _, tc := range sliceTestCases {
		t.Run(tc.description, func(t *testing.T) {
			fn, name := Discard[[]int], "Discard"
			if tc.keep {
				fn, name = Keep[[]int], "Keep"
			}
			data := slices.Clone(tc.input)
			got := fn(data, tc.filter)

			if !reflect.DeepEqual(data, tc.input) {
				t.Fatalf("%s(%#v, %q) modified the input data; input data should not change", name, tc.input, tc.filterStr)
			} else if len(got) == 0 && len(tc.expected) == 0 { // treat nil and empty the same
			} else if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("%s(%#v, %q) = %#v, want %#v",  name, tc.input, tc.filterStr, got, tc.expected)
			}
		})
	}
}

func BenchmarkInt(b *testing.B) {
	for range b.N {
		for _, tc := range intTestCases {
			fn := Discard[int]
			if tc.keep {
				fn = Keep[int]
			}
			fn(tc.input, tc.filter)
		}
	}
}

func BenchmarkString(b *testing.B) {
	for range b.N {
		for _, tc := range stringTestCases {
			fn := Discard[string]
			if tc.keep {
				fn = Keep[string]
			}
			fn(tc.input, tc.filter)
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	for range b.N {
		for _, tc := range sliceTestCases {
			fn := Discard[[]int]
			if tc.keep {
				fn = Keep[[]int]
			}
			fn(tc.input, tc.filter)
		}
	}
}
