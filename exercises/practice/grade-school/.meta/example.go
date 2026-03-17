package gradeschool

import (
	"maps"
	"slices"
)

type Grade struct {
	Level    int
	Students []string
}

type School map[int][]string

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) bool {
	if slices.Contains(s.Enrollment(), student) {
		return false
	}
	(*s)[grade] = append((*s)[grade], student)
	return true
}

func (s *School) Grade(level int) []string {
	g := (*s)[level]
	slices.Sort(g)
	return g
}

func (s *School) Enrollment() []string {
	var students []string
	grades := slices.Sorted(maps.Keys(*s))
	for _, grade := range grades {
		for _, student := range s.Grade(grade) {
			students = append(students, student)
		}
	}
	return students
}
