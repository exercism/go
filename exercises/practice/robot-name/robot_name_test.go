package robotname

import (
	"regexp"
	"testing"
)

var namePat = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)
var seen = map[string]int{}

func New() *Robot { return new(Robot) }

// getName is a test helper function to facilitate optionally checking for seen
// robot names.
func (r *Robot) getName(t testing.TB, expectSeen bool) string {
	t.Helper()
	newName, err := r.Name()
	if err != nil {
		t.Fatalf("Name() returned unexpected error: %v", err)
	}
	if len(newName) != 5 {
		t.Fatalf("names should have 5 characters: name '%s' has %d character(s)", newName, len(newName))
	}

	_, chk := seen[newName]
	if !expectSeen && chk {
		t.Fatalf("Name %s reissued after %d robots.", newName, len(seen))
	}
	seen[newName] = 0
	return newName
}

func TestNameValid(t *testing.T) {
	n := New().getName(t, false)
	if !namePat.MatchString(n) {
		t.Errorf(`Invalid robot name %q, want form "AA###".`, n)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	n2 := r.getName(t, true)
	if n2 != n1 {
		t.Errorf(`Robot name changed.  Now %s, was %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1 := New().getName(t, false)
	n2 := New().getName(t, false)
	if n1 == n2 {
		t.Errorf(`Robots with same name.  Two %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	r.Reset()
	if r.getName(t, false) == n1 {
		t.Errorf(`Robot name not cleared on reset.  Still %s.`, n1)
	}
}

var maxNames = 26 * 26 * 10 * 10 * 10

func TestCollisions(t *testing.T) {
	var name string
	// Test uniqueness for new robots.
	for i := len(seen); i <= maxNames-600000; i++ {
		name = New().getName(t, false)
		if len(name) != 5 {
			t.Fatalf("names should have 5 characters: name '%s' has %d character(s)", name, len(name))
		}
	}

	// Test that names aren't recycled either.
	r := New()
	for i := len(seen); i < maxNames; i++ {
		r.Reset()
		name = r.getName(t, false)
		if len(name) != 5 {
			t.Fatalf("names should have 5 characters: name '%s' has %d character(s)", name, len(name))
		}
	}

	// Test that name exhaustion is handled more or less correctly.
	_, err := New().Name()
	if err == nil {
		t.Fatalf("should return error if namespace is exhausted")
	}
}
