//+build slice

package stringset

// Example based on slices rather than maps, to validate that such a solution
// can pass the tests.
//
// The API is a little awkward for this because it has value receivers but
// a pointer in a struct works okay.

import "fmt"

// Set represents some properties of a mathematical set.
// Sets are finite and all elements are unique string values.
type Set struct {
	s *[]string
}

// New constructs a new Set representing an empty set.
func New() Set {
	return Set{new([]string)}
}

// NewFromSlice constructs a new Set populated from elements of a slice.
func NewFromSlice(l []string) Set {
	s := New()
	for _, e := range l {
		s.Add(e)
	}
	return s
}

// String returns a printable representation of s.
func (s Set) String() string {
	r := "{"
	i := 0
	for _, v := range *s.s {
		if i > 0 {
			r += ", "
		}
		r += fmt.Sprintf("%q", v)
		i++
	}
	return r + "}"
}

// IsEmpty returns true if s represents the empty set.
func (s Set) IsEmpty() bool {
	return len(*s.s) == 0
}

// Has returns true if e is an element of s.
func (s Set) Has(v string) bool {
	for _, e := range *s.s {
		if e == v {
			return true
		}
	}
	return false
}

// Add adds element e to Set s.
// If e is already in s, s is unchanged.
func (s Set) Add(v string) {
	for _, e := range *s.s {
		if e == v {
			return
		}
	}
	*s.s = append(*s.s, v)
}

// Subset returns true if all elements of s1 are also in s2.
func Subset(s1, s2 Set) bool {
	for _, v := range *s1.s {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 have no elements in common.
func Disjoint(s1, s2 Set) bool {
	for _, v := range *s1.s {
		if s2.Has(v) {
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
	var i []string
	for _, v := range *s1.s {
		if s2.Has(v) {
			i = append(i, v)
		}
	}
	return Set{&i}
}

// Difference returns a new Set populated with elements that appear in s1
// but not in s2.
func Difference(s1, s2 Set) Set {
	var d []string
	for _, v := range *s1.s {
		if !s2.Has(v) {
			d = append(d, v)
		}
	}
	return Set{&d}
}

// Union constructs a new Set populated with elements that appear in s1, s2,
// or both.
func Union(s1, s2 Set) Set {
	s := append([]string{}, *s1.s...)
	u := Set{&s}
	for _, v := range *s2.s {
		u.Add(v)
	}
	return u
}
