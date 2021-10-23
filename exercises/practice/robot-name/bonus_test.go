//go:build bonus
// +build bonus

package robotname

import "testing"

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
