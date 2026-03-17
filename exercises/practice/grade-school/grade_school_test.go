//nolint:gosec // In the context of this exercise, it is fine to use math.Rand instead of crypto.Rand.
package gradeschool

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	for _, tc := range testCasesAdd {
		t.Run(tc.description, func(t *testing.T) {
			var calls strings.Builder

			s := New()
			calls.WriteString("s := New()")

			for i, student := range tc.students {
				got := s.Add(student.name, student.grade)
				calls.WriteString(fmt.Sprintf("; s.Add(%q, %d)", student.name, student.grade))
				if got != tc.expected[i] {
					t.Fatalf("%s = %t, expected %t", calls.String(), got, tc.expected[i])
				}
			}
		})
	}
}

func TestGrade(t *testing.T) {
	for _, tc := range testCasesGrade {
		t.Run(tc.description, func(t *testing.T) {
			var calls strings.Builder

			s := New()
			calls.WriteString("s := New(); ")

			for _, student := range tc.students {
				s.Add(student.name, student.grade)
				calls.WriteString(fmt.Sprintf("s.Add(%q, %d); ", student.name, student.grade))
			}
			if got := s.Grade(tc.grade); !slices.Equal(got, tc.expected) {
				t.Fatalf("%s s.Grade(%d) = %#v, expected %#v", calls.String(), tc.grade, got, tc.expected)
			}
		})
	}
}

func TestEnrollment(t *testing.T) {
	for _, tc := range testCasesEnrollment {
		t.Run(tc.description, func(t *testing.T) {
			var calls strings.Builder

			s := New()
			calls.WriteString("s := New(); ")

			for _, student := range tc.students {
				s.Add(student.name, student.grade)
				calls.WriteString(fmt.Sprintf("s.Add(%q, %d); ", student.name, student.grade))
			}
			if got := s.Enrollment(); !slices.Equal(got, tc.expected) {
				t.Fatalf("%s s.Enrollment() = %#v, expected %#v", calls.String(), got, tc.expected)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for range b.N {
		s := New()
		for _, tc := range testCasesAdd {
			for _, student := range tc.students {
				s.Add(student.name, student.grade)
			}
		}
	}
}

func BenchmarkGrade(b *testing.B) {
	for _, tc := range testCasesGrade {
		b.StopTimer()
		s := New()
		for _, student := range tc.students {
			s.Add(student.name, student.grade)
		}
		b.StartTimer()
		for range b.N {
			s.Grade(tc.grade)
		}
	}
}

func BenchmarkEnrollment(b *testing.B) {
	for _, tc := range testCasesEnrollment {
		b.StopTimer()
		s := New()
		for _, student := range tc.students {
			s.Add(student.name, student.grade)
		}
		b.StartTimer()
		for range b.N {
			s.Enrollment()
		}
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
	for range b.N {
		// bench combined time to create a school and add
		// a number of students, drawn from a pool of students
		s := New()
		for t := 0; t < enrollment; t++ {
			s.Add(names[p], levels[p])
			p = (p + 1) % pool
		}
	}
}

func BenchmarkRandomEnrollment(b *testing.B) {
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
	for range b.N {
		// bench time to get enrollment of a full school,
		// averaged over a pool of schools.
		ss[p].Enrollment()
		p = (p + 1) % pool
	}
}
