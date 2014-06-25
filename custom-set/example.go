package stringset

import (
	"fmt"
	"reflect"
)

type Set map[string]bool

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	s := Set{}
	for _, e := range l {
		s[e] = true
	}
	return s
}

func (s Set) Add(e string) {
	s[e] = true
}

func (s Set) Delete(e string) {
	delete(s, e)
}

func (s Set) Has(e string) bool {
	return s[e]
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Slice() []string {
	l := make([]string, len(s))
	i := 0
	for e := range s {
		l[i] = e
		i++
	}
	return l
}

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

func Disjoint(s1, s2 Set) bool {
	return len(Intersection(s1, s2)) == 0
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	r := Set{}
	for e := range s1 {
		if s2[e] {
			r[e] = true
		}
	}
	return r
}

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

func Difference(s1, s2 Set) Set {
	r := Set{}
	for e := range s1 {
		if !s2[e] {
			r[e] = true
		}
	}
	return r
}

func Subset(s1, s2 Set) bool {
	for e := range s1 {
		if !s2[e] {
			return false
		}
	}
	return true
}

func SymetricDifference(s1, s2 Set) Set {
	r := Set{}
	for e := range s1 {
		if !s2[e] {
			r[e] = true
		}
	}
	for e := range s2 {
		if !s1[e] {
			r[e] = true
		}
	}
	return r
}
