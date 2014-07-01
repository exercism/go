package stringset

// Implement Set as a collection of unique string values.
//
// API:
//
// New() Set
// NewFromSlice([]string) Set
// (s Set) Add(string)
// (s Set) Delete(string)
// (s Set) Has(string) bool
// (s Set) IsEmpty() bool
// (s Set) Len() int
// (s Set) Slice() []string
// (s Set) String() string
// Disjoint(s1, s2 Set) bool
// Equal(s1, s2 Set) bool
// Intersection(s1, s2 Set) Set
// Union(s1, s2 Set) Set
// Difference(s1, s2 Set) Set
// Subset(s1, s2 Set) bool
// SymetricDifference(s1 s2 Set) Set
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements.  For example {"a", "b"}.
// Format the empty set as {}.

import (
	"fmt"
	"reflect"
	"testing"
)

var _ fmt.Stringer = New()

var _empty = New()

// These first tests of New and Add also test String.  No way around the
// chicken and egg problem.
func TestNew(t *testing.T) {

	// New must return an empty set.
	want := "{}"
	if res := _empty.String(); res != want {
		t.Fatalf(`New().String() = %s, want %s.`, res, want)
	}
}

var (
	_a    = New()
	_b    = New()
	_c    = New()
	_ab   = New()
	_abc  = New()
	_aS   = `{"a"}`
	_bS   = `{"b"}`
	_abS  = `{"a", "b"}`
	_baS  = `{"b", "a"}`
	_abcS = `{"a", "b", "c"}`
)

func init() {
	_a.Add("a")
	_b.Add("b")
	_c.Add("c")
	_ab.Add("a")
	_ab.Add("b")
	_abc.Add("a")
	_abc.Add("b")
	_abc.Add("c")
}

func TestAdd(t *testing.T) {
	// some sets are created and elements added in init() above.
	// now test results:

	// single element
	if res := _a.String(); res != _aS {
		t.Fatalf(`Add("a") = %s, want %s.`, res, _aS)
	}

	// another set, with a different element element
	if res := _b.String(); res != _bS {
		t.Fatalf(`Add("b") = %s, want %s.`, res, _bS)
	}

	// a third set, with two elements added
	if res := _ab.String(); res != _abS && res != _baS { // order undefined
		t.Fatalf(`Add("a"); Add("b") = %s, want %s.`, res, _abS)
	}

	// test adding an element that already exists
	_b.Add("b")
	if res := _b.String(); res != _bS {
		t.Fatalf(`Add("b") = %s, want %s.`, res, _bS)
	}
}

func TestNewFromSlice(t *testing.T) {
	// nil slice should give empty set
	want := "{}"
	if res := NewFromSlice(nil).String(); res != want {
		t.Fatalf(`NewFromSlice(nil) = %s, want %s.`, res, want)
	}

	// slice with one element:
	if res := NewFromSlice([]string{"a"}).String(); res != _aS {
		t.Fatalf(`NewFromSlice([]string{"a"}) = %s, want %s.`, res, _aS)
	}

	// slice with two elements:
	res := NewFromSlice([]string{"a", "b"}).String()
	if res != _abS && res != _baS { // order undefined
		t.Fatalf(`NewFromSlice([]string{"a", "b"}) = %s, want %s.`, res, _abS)
	}

	// slice with repeated element:
	if res := NewFromSlice([]string{"a", "a"}).String(); res != _aS {
		t.Fatalf(`NewFromSlice([]string{"a", "a"}) = %s, want %s.`, res, _aS)
	}
}

func TestDeleteHasIsEmptyLen(t *testing.T) {
	// properties of empty set:
	s := New()
	t.Log("s := New()") // log messages provide context for failure messages
	if l := s.Len(); l != 0 {
		t.Fatalf(`s.Len() = %d, want 0.`, l)
	}
	// `if !s.IsEmpty() {` is readable, but then literal bools
	// in the fail message would be prone to being mistyped.
	// this more tedious form should be safer.
	if res, want := s.IsEmpty(), true; res != want {
		t.Fatalf(`s.IsEmpty() = %t, want %t.`, res, want)
	}

	// properties of set with one element:
	s.Add("a")
	t.Log(`s.Add("a")`)
	if l := s.Len(); l != 1 {
		t.Fatalf(`s.Len() = %d, want 1.`, l)
	}
	if res, want := s.IsEmpty(), false; res != want {
		t.Fatalf(`s.IsEmpty() = %t, want %t.`, res, want)
	}
	if res, want := s.Has("a"), true; res != want {
		t.Fatalf(`s.Has("a") = %t, want %t.`, res, want)
	}

	// delete sole element:
	s.Delete("a")
	t.Log(`s.Delete("a")`)
	if res, want := s.IsEmpty(), true; res != want {
		t.Fatalf(`s.IsEmpty = %t, want %t.`, res, want)
	}

	// add a different element
	s.Add("b")
	t.Log(`s.Add("b")`)
	if l := s.Len(); l != 1 {
		t.Fatalf(`s.Len() = %d, want 1.`, l)
	}
	if res, want := s.Has("b"), true; res != want {
		t.Fatalf(`s.Has("b") = %t, want %t.`, res, want)
	}
	if res, want := s.Has("a"), false; res != want {
		t.Fatalf(`s.Has("a") = %t, want %t.`, res, want)
	}

	// add a second element, then add it a second time.
	s.Add("a")
	t.Log(`s.Add("a")`)
	s.Add("a")
	t.Log(`s.Add("a")`)
	if l := s.Len(); l != 2 {
		t.Fatalf(`s.Len() = %d, want 2.`, l)
	}

	// delete the element that was added twice:
	s.Delete("a")
	t.Log(`s.Delete("a")`)
	if res, want := s.IsEmpty(), false; res != want {
		t.Fatalf(`s.IsEmpty() = %t, want %t.`, res, want)
	}
	if l := s.Len(); l != 1 {
		t.Fatalf(`s.Len() = %d, want 1.`, l)
	}
	if res, want := s.Has("b"), true; res != want {
		t.Fatalf(`s.Has("b") = %t, want %t.`, res, want)
	}
	if res := s.String(); res != _bS {
		t.Fatalf(`s = %s, want %s.`, res, _bS)
	}
}

func TestSlice(t *testing.T) {
	// empty set should produce empty slice
	s := New()
	t.Log(`s := New()`)
	if l := s.Slice(); len(l) != 0 {
		t.Fatalf(`s.Slice() = %q, want []`, l)
	}

	// one element:
	s.Add("a")
	t.Log(`s.Add("a")`)
	if l := s.Slice(); len(l) != 1 || l[0] != "a" {
		t.Fatalf(`s.Slice() = %q, want ["a"]`, l)
	}

	// add a second element:
	s.Add("b")
	t.Log(`s.Add("b")`)
	if l := s.Slice(); len(l) != 2 ||
		!(l[0] == "a" && l[1] == "b" ||
			l[0] == "b" && l[1] == "a") {
		t.Fatalf(`s.Slice() = %q, want ["a" "b"]`, l)
	}
}

func TestEqualSubsetDisjoint(t *testing.T) {
	// empty sets should be equal, subsets, disjoint
	s1 := New()
	t.Log(`s1 := New()`)
	s2 := New()
	t.Log(`s2 := New()`)
	if res, want := Equal(s1, s2), true; res != want {
		t.Fatalf(`Equal(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s1, s2), true; res != want {
		t.Fatalf(`Subset(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Disjoint(s1, s2), true; res != want {
		t.Fatalf(`Disjoint(s1, s2) = %t, want %t.`, res, want)
	}

	// compare set with single element against empty set
	s1.Add("a")
	t.Log(`s1.Add("a")`)
	if res, want := Equal(s1, s2), false; res != want {
		t.Fatalf(`Equal(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s1, s2), false; res != want {
		t.Fatalf(`Subset(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s2, s1), true; res != want {
		t.Fatalf(`Subset(s2, s1) = %t, want %t.`, res, want)
	}
	if res, want := Disjoint(s1, s2), true; res != want {
		t.Fatalf(`Disjoint(s1, s2) = %t, want %t.`, res, want)
	}

	// compare distinct Sets, both {a}
	s2.Add("a")
	t.Log(`s2.Add("a")`)
	if res, want := Equal(s1, s2), true; res != want {
		t.Fatalf(`Equal(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s1, s2), true; res != want {
		t.Fatalf(`Subset(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Disjoint(s1, s2), false; res != want {
		t.Fatalf(`Disjoint(s1, s2) = %t, want %t.`, res, want)
	}

	// compare {a} with {a, b}
	s2.Add("b")
	t.Log(`s2.Add("b")`)
	if res, want := Equal(s1, s2), false; res != want {
		t.Fatalf(`Equal(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s1, s2), true; res != want {
		t.Fatalf(`Subset(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s2, s1), false; res != want {
		t.Fatalf(`Subset(s2, s1) = %t, want %t.`, res, want)
	}
	if res, want := Disjoint(s1, s2), false; res != want {
		t.Fatalf(`Disjoint(s1, s2) = %t, want %t.`, res, want)
	}

	// compare {a} with {b}
	s2.Delete("a")
	t.Log(`s2.Delete("a")`)
	if res, want := Equal(s1, s2), false; res != want {
		t.Fatalf(`Equal(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s1, s2), false; res != want {
		t.Fatalf(`Subset(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s2, s1), false; res != want {
		t.Fatalf(`Subset(s2, s1) = %t, want %t.`, res, want)
	}
	if res, want := Disjoint(s1, s2), true; res != want {
		t.Fatalf(`Disjoint(s1, s2) = %t, want %t.`, res, want)
	}

	// compare {a, c} with {b, c}
	s1.Add("c")
	t.Log(`s1.Add("c")`)
	s2.Add("c")
	t.Log(`s2.Add("c")`)
	if res, want := Equal(s1, s2), false; res != want {
		t.Fatalf(`Equal(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s1, s2), false; res != want {
		t.Fatalf(`Subset(s1, s2) = %t, want %t.`, res, want)
	}
	if res, want := Subset(s2, s1), false; res != want {
		t.Fatalf(`Subset(s2, s1) = %t, want %t.`, res, want)
	}
	if res, want := Disjoint(s1, s2), false; res != want {
		t.Fatalf(`Disjoint(s1, s2) = %t, want %t.`, res, want)
	}
}

func TestUnionIntersectionDifferenceSymetricDifference(t *testing.T) {
	// operations on empty sets
	s1 := New()
	t.Log(`s1 := New()`)
	s2 := New()
	t.Log(`s2 := New()`)
	if res := Intersection(s1, s2); !res.IsEmpty() {
		t.Fatalf(`Intersection(s1, s2) = %v, want {}.`, res)
	}
	if res := Union(s1, s2); !res.IsEmpty() {
		t.Fatalf(`Union(s1, s2) = %v, want {}.`, res)
	}
	if res := Difference(s1, s2); !res.IsEmpty() {
		t.Fatalf(`Difference(s1, s2) = %v, want {}.`, res)
	}
	if res := SymetricDifference(s1, s2); !res.IsEmpty() {
		t.Fatalf(`SymetricDifference(s1, s2) = %v, want {}.`, res)
	}

	// {a} op empty set:
	s1.Add("a")
	t.Log(`s1.Add("a")`)
	if res := Intersection(s1, s2); !res.IsEmpty() {
		t.Fatalf(`Intersection(s1, s2) = %v, want {}.`, res)
	}
	if res := Union(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`Union(s1, s2) = %v, want %v.`, res, _a)
	}
	if res := Difference(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`Difference(s1, s2) = %v, want %v.`, res, _a)
	}
	if res := Difference(s2, s1); !res.IsEmpty() {
		t.Fatalf(`Difference(s2, s1) = %v, want {}.`, res)
	}
	if res := SymetricDifference(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`SymetricDifference(s1, s2) = %v, want %v.`, res, _a)
	}

	// {a} op {a}:
	s2.Add("a")
	t.Log(`s2.Add("a")`)
	if res := Intersection(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`Intersection(s1, s2) = %v, want %v.`, res, _a)
	}
	if res := Union(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`Union(s1, s2) = %v, want %v.`, res, _a)
	}
	if res := Difference(s1, s2); !res.IsEmpty() {
		t.Fatalf(`Difference(s1, s2) = %v, want {}.`, res)
	}
	if res := SymetricDifference(s1, s2); !res.IsEmpty() {
		t.Fatalf(`SymetricDifference(s1, s2) = %v, want {}.`, res)
	}

	// {a} op {a, b}
	s2.Add("b")
	t.Log(`s2.Add("b")`)
	if res := Intersection(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`Intersection(s1, s2) = %v, want %v.`, res, _a)
	}
	if res := Union(s1, s2); !reflect.DeepEqual(res, _ab) {
		t.Fatalf(`Union(s1, s2) = %v, want %s.`, res, _abS)
	}
	if res := Difference(s1, s2); !res.IsEmpty() {
		t.Fatalf(`Difference(s1, s2) = %v, want {}.`, res)
	}
	if res := Difference(s2, s1); !reflect.DeepEqual(res, _b) {
		t.Fatalf(`Difference(s2, s1) = %v, want %v.`, res, _b)
	}
	if res := SymetricDifference(s1, s2); !reflect.DeepEqual(res, _b) {
		t.Fatalf(`SymetricDifference(s1, s2) = %v, want %v.`, res, _b)
	}

	// {a} op {b}
	s2.Delete("a")
	t.Log(`s2.Delete("a")`)
	if res := Intersection(s1, s2); !res.IsEmpty() {
		t.Fatalf(`Intersection(s1, s2) = %v, want {}.`, res)
	}
	if res := Union(s1, s2); !reflect.DeepEqual(res, _ab) {
		t.Fatalf(`Union(s1, s2) = %v, want %s.`, res, _abS)
	}
	if res := Difference(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`Difference(s1, s2) = %v, want %v.`, res, _a)
	}
	if res := Difference(s2, s1); !reflect.DeepEqual(res, _b) {
		t.Fatalf(`Difference(s2, s1) = %v, want %v.`, res, _b)
	}
	if res := SymetricDifference(s1, s2); !reflect.DeepEqual(res, _ab) {
		t.Fatalf(`SymetricDifference(s1, s2) = %v, want %s.`, res, _abS)
	}

	// {a, c} op {b, c}
	s1.Add("c")
	t.Log(`s1.Add("c")`)
	s2.Add("c")
	t.Log(`s2.Add("c")`)
	if res := Intersection(s1, s2); !reflect.DeepEqual(res, _c) {
		t.Fatalf(`Intersection(s1, s2) = %v, want %v.`, res, _c)
	}
	if res := Union(s1, s2); !reflect.DeepEqual(res, _abc) {
		t.Fatalf(`Union(s1, s2) = %v, want %s.`, res, _abcS)
	}
	if res := Difference(s1, s2); !reflect.DeepEqual(res, _a) {
		t.Fatalf(`Difference(s1, s2) = %v, want %v.`, res, _a)
	}
	if res := Difference(s2, s1); !reflect.DeepEqual(res, _b) {
		t.Fatalf(`Difference(s2, s1) = %v, want %v.`, res, _b)
	}
	if res := SymetricDifference(s1, s2); !reflect.DeepEqual(res, _ab) {
		t.Fatalf(`SymetricDifference(s1, s2) = %v, want %s.`, res, _abS)
	}
}
