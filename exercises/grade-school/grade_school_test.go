// API:
//
// type Grade struct {
// 	int, []string
// }
//
// type School
// func New() *School
// func (s *School) Add(string, int)
// func (s *School) Grade(int) []string
// func (s *School) Enrollment() []Grade

package school

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestNewSchoolIsEmpty(t *testing.T) {
	if len(New().Enrollment()) != 0 {
		t.Error("New school not empty")
	}
}

func list(e []Grade) (s string) {
	for _, l := range e {
		s += fmt.Sprintln(l)
	}
	return s
}

func TestAddStudent(t *testing.T) {
	exp := list([]Grade{{2, []string{"Aimee"}}})
	s := New()
	s.Add("Aimee", 2)
	got := list(s.Enrollment())
	if got != exp {
		t.Errorf(`Add Aimee level 2, got
%sexpected:
%s`, got, exp)
	}
}

func TestAddMoreSameGrade(t *testing.T) {
	exp := list([]Grade{{2, []string{"Blair James Paul"}}})
	s := New()
	s.Add("Blair", 2)
	s.Add("James", 2)
	s.Add("Paul", 2)
	got := list(s.Enrollment())
	if got != exp {
		t.Errorf(`Add more same grade, got
%sexpected:
%s`, got, exp)
	}
}

func TestAddDifferentGrades(t *testing.T) {
	exp := list([]Grade{
		{3, []string{"Chelsea"}},
		{7, []string{"Logan"}},
	})
	s := New()
	s.Add("Chelsea", 3)
	s.Add("Logan", 7)
	got := list(s.Enrollment())
	if got != exp {
		t.Errorf(`Add different grades, got
%sexpected:
%s`, got, exp)
	}
}

func TestGetGrade(t *testing.T) {
	exp := []string{"Bradley", "Franklin"}
	s := New()
	s.Add("Bradley", 5)
	s.Add("Franklin", 5)
	s.Add("Jeff", 1)
	got := s.Grade(5)
	if len(got) == len(exp) {
		if got[0] == exp[0] && got[1] == exp[1] ||
			got[0] == exp[1] && got[1] == exp[0] { // accept out of order
			return
		}
	}
	t.Errorf(`Get grade, got
%q
expected
%q`, got, exp)
}

func TestNonExistantGrade(t *testing.T) {
	s := New()
	got := s.Grade(1)
	if len(got) != 0 {
		t.Errorf(`Get non-existent grade, got
%q
expected
[]`, got)
	}
}

func TestSortedEnrollment(t *testing.T) {
	exp := list([]Grade{
		{3, []string{"Kyle"}},
		{4, []string{"Christopher Jennifer"}},
		{6, []string{"Kareem"}},
	})
	s := New()
	s.Add("Jennifer", 4)
	s.Add("Kareem", 6)
	s.Add("Christopher", 4)
	s.Add("Kyle", 3)
	got := list(s.Enrollment())
	if got != exp {
		t.Errorf(`Sorted enrollment, got
%sexpected:
%s`, got, exp)
	}
}

const (
	minLevel   = 1
	maxLevel   = 9
	enrollment = 400
)

func BenchmarkAddStudents(b *testing.B) {
	const pool = 1e6 // pool of students
	names := make([]string, pool)
	levels := make([]int, pool)
	for i := range names {
		names[i] = strconv.Itoa(rand.Intn(1e5))
		levels[i] = minLevel + rand.Intn(maxLevel-minLevel+1)
	}
	p := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// bench combined time to create a school and add
		// a number of students, drawn from a pool of students
		s := New()
		for t := 0; t < enrollment; t++ {
			s.Add(names[p], levels[p])
			p = (p + 1) % pool
		}
	}
}

func BenchmarkEnrollment(b *testing.B) {
	const pool = 1000 // pool of schools
	ss := make([]*School, pool)
	for i := range ss {
		s := New()
		for t := 0; t < enrollment; t++ {
			s.Add(
				strconv.Itoa(rand.Intn(1e5)),
				minLevel+rand.Intn(maxLevel-minLevel+1))
		}
		ss[i] = s
	}
	p := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// bench time to get enrollment of a full school,
		// averaged over a pool of schools.
		ss[p].Enrollment()
		p = (p + 1) % pool
	}
}
