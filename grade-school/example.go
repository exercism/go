package school

import "sort"

type School map[int][]string

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	(*s)[grade] = append((*s)[grade], student)
}

func (s *School) Grade(level int) []string {
	return (*s)[level]
}

func (s *School) Enrollment() []Grade {
	m := *s
	ls := make([]int, len(m))
	i := 0
	for l := range m {
		ls[i] = l
		i++
	}
	sort.Ints(ls)
	e := make([]Grade, len(m))
	for i, l := range ls {
		t := m[l]
		sort.Strings(t)
		e[i] = Grade{l, t}
	}
	return e
}

type Grade struct {
	Level    int
	Students []string
}
