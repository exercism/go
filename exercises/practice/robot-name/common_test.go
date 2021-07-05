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
