package stringset

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
	"testing"
)

var _ fmt.Stringer = New()

var _empty = New()

// Early tests of New and Add also test String.  No way around the chicken
// and egg problem.
func TestNew(t *testing.T) {
	if res := _empty.String(); res != "{}" {
		t.Fatalf(`New().String() = %s, want "{}".`, res)
	}
}

var (
	_a    = New()
	_b    = New()
	_ab   = New()
	_aS   = `{"a"}`
	_bS   = `{"b"}`
	_abS1 = `{"a", "b"}`
	_abS2 = `{"b", "a"}`
)

func init() {
	_a.Add("a")
	_b.Add("b")
	_ab.Add("a")
	_ab.Add("b")
}

func TestAdd(t *testing.T) {
	if res := _a.String(); res != _aS {
		t.Fatalf(`Add("a") = %s, want %s.`, res, _aS)
	}
	if res := _b.String(); res != _bS {
		t.Fatalf(`Add("b") = %s, want %s.`, res, _bS)
	}
	if res := _ab.String(); res != _abS1 && res != _abS2 {
		t.Fatalf(`Add("a"); Add("b") = %s, want %s.`, res, _abS1)
	}
	// another b should not change b
	_b.Add("b")
	if res := _b.String(); res != _bS {
		t.Fatalf(`Add("b") = %s, want %s.`, res, _bS)
	}
}

func TestNewFromSlice(t *testing.T) {
	if res := NewFromSlice(nil).String(); res != "{}" {
		t.Fatalf(`NewFromSlice(nil) = %s, want "{}".`, res)
	}
	if res := NewFromSlice([]string{"a"}).String(); res != _aS {
		t.Fatalf(`NewFromSlice([]string{"a"}) = %s, want %s.`,
			res, _aS)
	}
	res := NewFromSlice([]string{"a", "b"}).String()
	if res != _abS1 && res != _abS2 {
		t.Fatalf(`NewFromSlice([]string{"a", "b"}) = %s, want %s.`,
			res, _abS1)
	}
}

func TestDeleteHasIsEmptyLen(t *testing.T) {
	s := New()
	t.Log("s := New()")
	if l := s.Len(); l != 0 {
		t.Fatalf(`s.Len() = %d, want 0.`, l)
	}
	if !s.IsEmpty() {
		t.Fatal(`s.IsEmpty() = false, want true.`)
	}

	s.Add("a")
	t.Log(`s.Add("a")`)
	if l := s.Len(); l != 1 {
		t.Fatalf(`s.Len() = %d, want 1.`, l)
	}
	if s.IsEmpty() {
		t.Fatal(`s.IsEmpty() = true, want false.`)
	}
	if !s.Has("a") {
		t.Fatal(`s.Has("a") = false, want true.`)
	}

	s.Delete("a")
	t.Log(`s.Delete("a")`)
	if !s.IsEmpty() {
		t.Fatalf(`s.IsEmpty = false, want true.`)
	}

	s.Add("b")
	t.Log(`s.Add("b")`)
	if l := s.Len(); l != 1 {
		t.Fatalf(`s.Len() = %d, want 1.`, l)
	}
	if !s.Has("b") {
		t.Fatal(`s.Has("b") = false, want true.`)
	}
	if s.Has("a") {
		t.Fatal(`s.Has("a") = true, want false.`)
	}

	s.Add("a")
	t.Log(`s.Add("a")`)
	s.Add("a")
	t.Log(`s.Add("a")`)
	if l := s.Len(); l != 2 {
		t.Fatalf(`s.Len() = %d, want 2.`, l)
	}

	s.Delete("a")
	t.Log(`s.Delete("a")`)
	if s.IsEmpty() {
		t.Fatal(`s.IsEmpty() = true, want false.`)
	}
	if l := s.Len(); l != 1 {
		t.Fatalf(`s.Len() = %d, want 1.`, l)
	}
	if !s.Has("b") {
		t.Fatal(`s.Has("b") = false, want true.`)
	}
	if res := s.String(); res != _bS {
		t.Fatalf(`s = %s, want %s.`, res, _bS)
	}
}

func TestSlice(t *testing.T) {
	s := New()
	t.Log(`s := New()`)
	if l := s.Slice(); len(l) != 0 {
		t.Fatalf(`s.Slice() = %q, want []`, l)
	}
	s.Add("a")
	t.Log(`s.Add("a")`)
	if l := s.Slice(); len(l) != 1 || l[0] != "a" {
		t.Fatalf(`s.Slice() = %q, want ["a"]`, l)
	}
	s.Add("b")
	t.Log(`s.Add("b")`)
	if l := s.Slice(); len(l) != 2 || !(l[0] == "a" && l[1] == "b" ||
		l[0] == "b" && l[1] == "a") {
		t.Fatalf(`s.Slice() = %q, want ["a" "b"]`, l)
	}
}

func TestUnionIntersectionDifferenceSymetricDifference(t *testing.T) {
}

func TestEqualSubsetDisjoint(t *testing.T) {
	s1 := New()
	t.Log(`s1 := New()`)
	s2 := New()
	t.Log(`s2 := New()`)
	if !Equal(s1, s2) {
		t.Fatalf(`Equal(s1, s2) = false, want true.`)
	}
	if !Subset(s1, s2) {
		t.Fatalf(`Subset(s1, s2) = false, want true.`)
	}

	s2.Add("b")
	t.Log(`s2.Add("b")`)
	if Equal(s1, s2) {
		t.Fatalf(`Equal(s1, s2) = true, want false.`)
	}
	if !Subset(s1, s2) {
		t.Fatalf(`Subset(s1, s2) = false, want true.`)
	}
	if Subset(s2, s1) {
		t.Fatalf(`Subset(s2, s1) = true, want false.`)
	}

	s1.Add("a")
	t.Log(`s1.Add("a")`)
	s1.Add("b")
	t.Log(`s1.Add("b")`)
	s2.Add("b")
	t.Log(`s2.Add("b")`)
	if Subset(s1, s2) {
		t.Fatalf(`Subset(s1, s2) = false, want true.`)
	}
	if !Subset(s2, s1) {
		t.Fatalf(`Subset(s2, s1) = false, want true.`)
	}
}
