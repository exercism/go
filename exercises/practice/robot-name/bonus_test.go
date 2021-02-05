// +build bonus

package robotname

import "testing"

var maxNames = 26 * 26 * 10 * 10 * 10

func TestCollisions(t *testing.T) {
	// Test uniqueness for new robots.
	for i := len(seen); i <= maxNames-600000; i++ {
		_ = New().getName(t, false)
	}

	// Test that names aren't recycled either.
	r := New()
	for i := len(seen); i < maxNames; i++ {
		r.Reset()
		_ = r.getName(t, false)
	}

	// Test that name exhaustion is handled more or less correctly.
	_, err := New().Name()
	if err == nil {
		t.Fatalf("should return error if namespace is exhausted")
	}
}
