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
	newName, err := r.Name()
	if err != nil {
		t.Fatalf("Name() returned unexpected error: %v", err)
	}
	_, chk := seen[newName]
	if !expectSeen && chk {
		t.Fatalf("Name %s reissued after %d robots.", newName, len(seen))
	}
	seen[newName] = 0
	return newName
}
