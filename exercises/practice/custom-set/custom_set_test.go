package stringset

// Implement Set as a collection of unique string values.
//
// API:
//
// New() Set
// NewFromSlice([]string) Set
// (s Set) String() string
// (s Set) IsEmpty() bool
// (s Set) Has(string) bool
// Subset(s1, s2 Set) bool
// Disjoint(s1, s2 Set) bool
// Equal(s1, s2 Set) bool
// (s Set) Add(string)
// Intersection(s1, s2 Set) Set
// Difference(s1, s2 Set) Set
// Union(s1, s2 Set) Set
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements.  For example {"a", "b"}.
// Format the empty set as {}.

import (
	"math/rand"
	"strconv"
	"testing"
)

// A first set of tests uses Set.String() to judge correctness.

func TestNew(t *testing.T) {
	// New must return an empty set.
	want := "{}"
	if got := New().String(); got != want {
		t.Fatalf(`New().String() = %s, want %s.`, got, want)
	}
}

func TestNewFromSlice(t *testing.T) {
	// nil slice should give empty set
	want := "{}"
	if got := NewFromSlice(nil).String(); got != want {
		t.Fatalf(`NewFromSlice(nil) = %s, want %s.`, got, want)
	}

	// slice with one element:
	want = `{"a"}`
	if got := NewFromSlice([]string{"a"}).String(); got != want {
		t.Fatalf(`NewFromSlice([]string{"a"}) = %s, want %s.`, got, want)
	}

	// slice with repeated element:
	if got := NewFromSlice([]string{"a", "a"}).String(); got != want {
		t.Fatalf(`NewFromSlice([]string{"a", "a"}) = %s, want %s.`, got, want)
	}

	// slice with two elements:
	got := NewFromSlice([]string{"a", "b"}).String()
	want1 := `{"a", "b"}`
	want2 := `{"b", "a"}`
	if got != want1 && got != want2 { // order undefined
		t.Fatalf(`NewFromSlice([]string{"a", "b"}) = %s, want %s or (%s).`,
			got, want1, want2)
	}
}

// Trusting NewFromSlice now, remaining tests are table driven, taking data
// from cases_test.go and building sets with NewFromSlice.

// test case types used in cases_test.go

type (
	// unary function, bool result (IsEmpty)
	unaryBoolCase struct {
		set  []string
		want bool
	}
	// set-element function, bool result (Has)
	eleBoolCase struct {
		set  []string
		ele  string
		want bool
	}
	// binary function, bool result (Subset, Disjoint, Equal)
	binBoolCase struct {
		set1 []string
		set2 []string
		want bool
	}
	// set-element operator (Add)
	eleOpCase struct {
		set  []string
		ele  string
		want []string
	}
	// set-set operator (Intersection, Difference, Union)
	binOpCase struct {
		set1 []string
		set2 []string
		want []string
	}
)

func TestIsEmpty(t *testing.T) {
	for _, tc := range emptyCases {
		s := NewFromSlice(tc.set)
		got := s.IsEmpty()
		if got != tc.want {
			t.Fatalf("%v IsEmpty = %t, want %t", s, got, tc.want)
		}
	}
}

func TestHas(t *testing.T) {
	for _, tc := range containsCases {
		s := NewFromSlice(tc.set)
		got := s.Has(tc.ele)
		if got != tc.want {
			t.Fatalf("%v Has %q = %t, want %t", s, tc.ele, got, tc.want)
		}
	}
}

// helper for testing Subset, Disjoint, Equal
func testBinBool(name string, f func(Set, Set) bool, cases []binBoolCase, t *testing.T) {
	for _, tc := range cases {
		s1 := NewFromSlice(tc.set1)
		s2 := NewFromSlice(tc.set2)
		got := f(s1, s2)
		if got != tc.want {
			t.Fatalf("%s(%v, %v) = %t, want %t", name, s1, s2, got, tc.want)
		}
	}
}

func TestSubset(t *testing.T) {
	testBinBool("Subset", Subset, subsetCases, t)
}

func TestDisjoint(t *testing.T) {
	testBinBool("Disjoint", Disjoint, disjointCases, t)
}

func TestEqual(t *testing.T) {
	testBinBool("Equal", Equal, equalCases, t)
}

func TestAdd(t *testing.T) {
	for _, tc := range addCases {
		s := NewFromSlice(tc.set)
		s.Add(tc.ele)
		want := NewFromSlice(tc.want)
		if !Equal(s, want) {
			t.Fatalf("%v Add %q = %v, want %v", NewFromSlice(tc.set), tc.ele, s, want)
		}
	}
}

// helper for testing Intersection, Difference, Union
func testBinOp(name string, f func(Set, Set) Set, cases []binOpCase, t *testing.T) {
	for _, tc := range cases {
		s1 := NewFromSlice(tc.set1)
		s2 := NewFromSlice(tc.set2)
		want := NewFromSlice(tc.want)
		got := f(s1, s2)
		if !Equal(got, want) {
			t.Fatalf("%s(%v, %v) = %v, want %v", name, s1, s2, got, want)
		}
	}
}

func TestIntersection(t *testing.T) {
	testBinOp("Intersection", Intersection, intersectionCases, t)
}

func TestDifference(t *testing.T) {
	testBinOp("Difference", Difference, differenceCases, t)
}

func TestUnion(t *testing.T) {
	testBinOp("Union", Union, unionCases, t)
}

func BenchmarkNewFromSlice1e1(b *testing.B) { bench(1e1, b) }
func BenchmarkNewFromSlice1e2(b *testing.B) { bench(1e2, b) }
func BenchmarkNewFromSlice1e3(b *testing.B) { bench(1e3, b) }
func BenchmarkNewFromSlice1e4(b *testing.B) { bench(1e4, b) }

func bench(nAdd int, b *testing.B) {
	s := make([]string, nAdd)
	for i := range s {
		s[i] = strconv.Itoa(rand.Intn(len(s)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewFromSlice(s)
	}
}
