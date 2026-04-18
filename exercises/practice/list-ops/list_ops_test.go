package listops

import (
	"fmt"
	"slices"
	"strings"
	"strconv"
	"testing"
)

func sliceIntToStr(ints []int) string {
	parts := make([]string, len(ints))
	for i, v := range ints {
		parts[i] = strconv.Itoa(v)
	}
	return fmt.Sprintf("%s", strings.Join(parts, ", "))
}

func (l IntList) String() string {
	return fmt.Sprintf("IntList{%s}", sliceIntToStr(l))
}

func sliceString(l []IntList) string {
	parts := make([]string, len(l))
	for i, p := range l {
		parts[i] = fmt.Sprintf("{%s}", sliceIntToStr(p))
	}
	return fmt.Sprintf("[]IntList{%s}", strings.Join(parts, ", "))
}

func TestFoldl(t *testing.T) {
	for _, tc := range testCasesFoldl {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.list.Foldl(tc.function, tc.initial); got != tc.expected {
				t.Fatalf("%s.Foldl(%s, %d) = %d, want: %d", tc.list, tc.functionStr, tc.initial, got, tc.expected)
			}
		})
	}
}

func TestFoldr(t *testing.T) {
	for _, tc := range testCasesFoldr {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.list.Foldr(tc.function, tc.initial); got != tc.expected {
				t.Fatalf("%s.Foldr(%s, %d) = %d, want: %d", tc.list, tc.functionStr, tc.initial, got, tc.expected)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	for _, tc := range testCasesFilter {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.list.Filter(tc.function); !slices.Equal(got, tc.expected) {
				t.Fatalf("%s.Filter(%s) = %s, want: %s", tc.list, tc.functionStr, got, tc.expected)
			}
		})
	}
}

func TestLength(t *testing.T) {
	for _, tc := range testCasesLength {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.list.Length(); got != tc.expected {
				t.Fatalf("%s.Length() = %d, want: %d", tc.list, got, tc.expected)
			}
		})
	}
}

func TestMap(t *testing.T) {
	for _, tc := range testCasesMap {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.list.Map(tc.function); !slices.Equal(got, tc.expected) {
				t.Fatalf("%s.Map(%s) = %s, want: %s", tc.list, tc.functionStr, got, tc.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range testCasesReverse {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.list.Reverse(); !slices.Equal(got, tc.expected) {
				t.Fatalf("%s.Reverse() = %s, want: %s", tc.list, got, tc.expected)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	for _, tc := range testCasesAppend {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.list1.Append(tc.list2); !slices.Equal(got, tc.expected) {
				t.Fatalf("%s.Append(%s) = %s, want: %s", tc.list1, tc.list2, got, tc.expected)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	for _, tc := range testCasesConcat {
		t.Run(tc.description, func(t *testing.T) {
			var initial IntList
			if got := initial.Concat(tc.lists); !slices.Equal(got, tc.expected) {
				t.Fatalf("%s.Concat(%s) = %s, want: %s", initial, sliceString(tc.lists), got, tc.expected)
			}
		})
	}
}
