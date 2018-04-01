// +build bonus

package robotname

import "testing"

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
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code should panic if namespace is exhausted.")
		}
	}()

	_ = New().getName(t, false)
}
