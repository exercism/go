//+build slice

package stringset

// Example based on slices rather than maps, to validate that such a solution
// can pass the tests.
//
// The API is a little akward for this because it has value receivers but
// a pointer in a struct works okay.

import "fmt"

const TestVersion = 2

type Set struct{ s *[]string }

func New() Set {
	return Set{new([]string)}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, e := range l {
		s.Add(e)
	}
	return s
}

func (s Set) Add(v string) {
	for _, e := range *s.s {
		if e == v {
			return
		}
	}
	*s.s = append(*s.s, v)
}

func (s Set) Delete(v string) {
	ss := *s.s
	for i, e := range ss {
		if e == v {
			last := len(ss) - 1
			ss[i] = ss[last]
			*s.s = ss[:last]
			return
		}
	}
}

func (s Set) Has(v string) bool {
	for _, e := range *s.s {
		if e == v {
			return true
		}
	}
	return false
}

func (s Set) IsEmpty() bool {
	return len(*s.s) == 0
}

func (s Set) Len() int {
	return len(*s.s)
}

func (s Set) Slice() []string {
	return append([]string{}, *s.s...)
}

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

func Disjoint(s1, s2 Set) bool {
	for _, v := range *s1.s {
		if s2.Has(v) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return s1.Len() == s2.Len() && Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	var i []string
	for _, v := range *s1.s {
		if s2.Has(v) {
			i = append(i, v)
		}
	}
	return Set{&i}
}

func Union(s1, s2 Set) Set {
	s := s1.Slice()
	u := Set{&s}
	for _, v := range *s2.s {
		u.Add(v)
	}
	return u
}

func Difference(s1, s2 Set) Set {
	var d []string
	for _, v := range *s1.s {
		if !s2.Has(v) {
			d = append(d, v)
		}
	}
	return Set{&d}
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

// SymmetricDifference constructs a new set populated with elements that are
// members of s1 or s2 but not both.
func SymmetricDifference(s1, s2 Set) Set {
	var s []string
	for _, v := range *s1.s {
		if !s2.Has(v) {
			s = append(s, v)
		}
	}
	for _, v := range *s2.s {
		if !s1.Has(v) {
			s = append(s, v)
		}
	}
	return Set{&s}
}
