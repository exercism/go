package stringset

// Implement Set as a collection of unique string values.
//
// API:
//
// New() Set
// NewFromSlice([]string) Set
// (s Set) Add(string)         // modify s
// (s Set) Delete(string)      // modify s
// (s Set) Has(string) bool
// (s Set) IsEmpty() bool
// (s Set) Len() int
// (s Set) Slice() []string
// (s Set) String() string
// Equal(s1, s2 Set) bool
// Subset(s1, s2 Set) bool     // return s1 ⊆ s2
// Disjoint(s1, s2 Set) bool
// Intersection(s1, s2 Set) Set
// Union(s1, s2 Set) Set
// Difference(s1, s2 Set) Set  // return s1 ∖ s2
// SymmetricDifference(s1, s2 Set) Set
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements.  For example {"a", "b"}.
// Format the empty set as {}.

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

const targetTestVersion = 3

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

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

func TestSlice(t *testing.T) {
	// empty set should produce empty slice
	s := New()
	if l := s.Slice(); len(l) != 0 {
		t.Fatalf(`s.Slice() = %q, want []`, l)
	}

	// one element:
	want := []string{"a"}
	s = NewFromSlice(want)
	got := s.Slice()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`%v Slice = %q, want %q`, s, got, want)
	}

	// two elements:
	w1 := []string{"a", "b"}
	w2 := []string{"b", "a"}
	s = NewFromSlice(w1)
	got = s.Slice()
	if !reflect.DeepEqual(got, w1) && !reflect.DeepEqual(got, w2) {
		t.Fatalf(`%v Slice = %q, want %q`, s, got, w1)
	}
}

// Trusting NewFromSlice now, remaining tests are table driven, taking data
// from cases_test.go and building sets with NewFromSlice.

// test case types used in cases_test.go
type (
	// binary function, bool result (Equal, Subset, Disjoint)
	binBoolCase struct {
		set1 []string
		set2 []string
		want bool
	}
	// unary function, bool result (IsEmpty)
	unaryBoolCase struct {
		set  []string
		want bool
	}
	// unary function, int result (Len)
	unaryIntCase struct {
		set  []string
		want int
	}
	// set-element function, bool result (Has)
	eleBoolCase struct {
		set  []string
		ele  string
		want bool
	}
	// set-element operator (Add, Delete)
	eleOpCase struct {
		set  []string
		ele  string
		want []string
	}
	// set-set operator (Union, Intersection, Difference, Symmetric-Difference)
	binOpCase struct {
		set1 []string
		set2 []string
		want []string
	}
)

// helper for testing Equal, Subset, Disjoint
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

func TestEqual(t *testing.T) {
	testBinBool("Equal", Equal, eqCases, t)
}

// With Equal tested, remaining tests use it to judge correctness.

// helper for testing Add, Delete
func testEleOp(name string, op func(Set, string), cases []eleOpCase, t *testing.T) {
	for _, tc := range cases {
		s := NewFromSlice(tc.set)
		op(s, tc.ele)
		want := NewFromSlice(tc.want)
		if !Equal(s, want) {
			t.Fatalf("%v %s %q = %v, want %v",
				NewFromSlice(tc.set), name, tc.ele, s, want)
		}
	}
}

func TestAdd(t *testing.T) {
	testEleOp("Add", Set.Add, addCases, t)
}

func TestDelete(t *testing.T) {
	testEleOp("Delete", Set.Delete, delCases, t)
}

func TestHas(t *testing.T) {
	for _, tc := range hasCases {
		s := NewFromSlice(tc.set)
		got := s.Has(tc.ele)
		if got != tc.want {
			t.Fatalf("%v Has %q = %t, want %t", s, tc.ele, got, tc.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tc := range emptyCases {
		s := NewFromSlice(tc.set)
		got := s.IsEmpty()
		if got != tc.want {
			t.Fatalf("%v IsEmpty = %t, want %t", s, got, tc.want)
		}
	}
}

func TestLen(t *testing.T) {
	for _, tc := range lenCases {
		s := NewFromSlice(tc.set)
		got := s.Len()
		if got != tc.want {
			t.Fatalf("%v Len = %d, want %d", s, got, tc.want)
		}
	}
}

func TestSubset(t *testing.T) {
	testBinBool("Subset", Subset, subsetCases, t)
}

func TestDisjoint(t *testing.T) {
	testBinBool("Disjoint", Disjoint, disjointCases, t)
}

// helper for testing Union, Intersection, Difference, SymmetricDifference
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

func TestUnion(t *testing.T) {
	testBinOp("Union", Union, unionCases, t)
}

func TestIntersection(t *testing.T) {
	testBinOp("Intersection", Intersection, intersectionCases, t)
}

func TestDifference(t *testing.T) {
	testBinOp("Difference", Difference, differenceCases, t)
}

func TestSymmetricDifference(t *testing.T) {
	testBinOp("SymmetricDifference", SymmetricDifference, symmetricDifferenceCases, t)
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
