//go:build !slice
// +build !slice

package stringset

import "fmt"

// Set represents some properties of a mathematical set.
// Sets are finite and all elements are unique string values.
type Set map[string]bool

// New constructs a new Set representing an empty set.
func New() Set {
	return Set{}
}

// NewFromSlice constructs a new Set populated from elements of a slice.
func NewFromSlice(l []string) Set {
	s := Set{}
	for _, e := range l {
		s[e] = true
	}
	return s
}

// String returns a printable representation of s.
func (s Set) String() string {
	r := "{"
	i := 0
	for e := range s {
		if i > 0 {
			r += ", "
		}
		r += fmt.Sprintf("%q", e)
		i++
	}
	return r + "}"
}

// IsEmpty returns true if s represents the empty set.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns true if e is an element of s.
func (s Set) Has(e string) bool {
	return s[e]
}

// Add adds element e to Set s.
// If e is already in s, s is unchanged.
func (s Set) Add(e string) {
	s[e] = true
}

// Subset returns true if all elements of s1 are also in s2.
func Subset(s1, s2 Set) bool {
	for e := range s1 {
		if !s2[e] {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 have no elements in common.
func Disjoint(s1, s2 Set) bool {
	for e := range s1 {
		if s2[e] {
			return false
		}
	}
	return true
}

// Equal returns true if s1 and s2 contain the same elements.
func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

// Intersection constructs a new Set populated with the elements common to
// both s1 and s2
func Intersection(s1, s2 Set) Set {
	r := Set{}
	for e := range s1 {
		if s2[e] {
			r[e] = true
		}
	}
	return r
}

// Difference returns a new Set populated with elements that appear in s1
// but not in s2.
func Difference(s1, s2 Set) Set {
	r := Set{}
	for e := range s1 {
		if !s2[e] {
			r[e] = true
		}
	}
	return r
}

// Union constructs a new Set populated with elements that appear in s1, s2,
// or both.
func Union(s1, s2 Set) Set {
	r := Set{}
	for e := range s1 {
		r[e] = true
	}
	for e := range s2 {
		r[e] = true
	}
	return r
}
