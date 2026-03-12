package listops

import (
	"slices"
	"testing"
)

func TestFoldl(t *testing.T) {
	for _, tc := range testCasesFoldl {
		t.Run(tc.description, func(t *testing.T) {
			if got := IntList(tc.list).Foldl(tc.function, tc.initial); got != tc.expected {
				t.Fatalf("Foldl(%#v, %d) = %d, want: %d", tc.list, tc.initial, got, tc.expected)
			}
		})
	}
}

func TestFoldr(t *testing.T) {
	for _, tc := range testCasesFoldr {
		t.Run(tc.description, func(t *testing.T) {
			if got := IntList(tc.list).Foldr(tc.function, tc.initial); got != tc.expected {
				t.Fatalf("Foldr(%#v, %d) = %d, want: %d", tc.list, tc.initial, got, tc.expected)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	for _, tc := range testCasesFilter {
		t.Run(tc.description, func(t *testing.T) {
			if got := IntList(tc.list).Filter(tc.function); !slices.Equal(got, tc.expected) {
				t.Fatalf("Filter(%#v, %s) = %#v, want: %#v", tc.list, tc.functionStr, got, tc.expected)
			}
		})
	}
}

func TestLength(t *testing.T) {
	for _, tc := range testCasesLength {
		t.Run(tc.description, func(t *testing.T) {
			if got := IntList(tc.list).Length(); got != tc.expected {
				t.Fatalf("Length(%#v) = %d, want: %d", tc.list, got, tc.expected)
			}
		})
	}
}

func TestMap(t *testing.T) {
	for _, tc := range testCasesMap {
		t.Run(tc.description, func(t *testing.T) {
			if got := IntList(tc.list).Map(tc.function); !slices.Equal(got, tc.expected) {
				t.Fatalf("Map(%#v, %s) = %#v, want: %#v", tc.list, tc.functionStr, got, tc.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range testCasesReverse {
		t.Run(tc.description, func(t *testing.T) {
			if got := IntList(tc.list).Reverse(); !slices.Equal(got, tc.expected) {
				t.Fatalf("Reverse(%#v) = %#v, want: %#v", tc.list, got, tc.expected)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	for _, tc := range testCasesAppend {
		t.Run(tc.description, func(t *testing.T) {
			if got := IntList(tc.list1).Append(tc.list2); !slices.Equal(got, tc.expected) {
				t.Fatalf("Append(%#v, %#v) = %#v, want: %#v", tc.list1, tc.list2, got, tc.expected)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	for _, tc := range testCasesConcat {
		t.Run(tc.description, func(t *testing.T) {
			var lists []IntList
			for _, list := range tc.lists {
				lists = append(lists, IntList(list))
			}
			if got := (IntList{}).Concat([]IntList(lists)); !slices.Equal(got, IntList(tc.expected)) {
				t.Fatalf("Append(%#v) = %#v, want: %#v", tc.lists, got, tc.expected)
			}
		})
	}
}
