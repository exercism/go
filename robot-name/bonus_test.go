// +build bonus

package robotname

import "testing"

func TestCollisions(t *testing.T) {
	m := map[string]bool{}
	// Test uniqueness for new robots.
	// 10k should be plenty to catch names generated
	// randomly without a uniqueness check.
	for i := 0; i < 10000; i++ {
		n := New().Name()
		if m[n] {
			t.Fatalf("Name %s reissued after %d robots.", n, i)
		}
		m[n] = true
	}
	// Test that names aren't recycled either.
	r := New()
	for i := 0; i < 10000; i++ {
		r.Reset()
		n := r.Name()
		if m[n] {
			t.Fatalf("Name %s reissued after Reset.", n)
		}
		m[n] = true
	}
}
