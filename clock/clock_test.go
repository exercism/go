package clock

import "testing"

var newTests = []struct {
	h, m int
	disp string
}{
	{8, 0, "08:00"},
	{9, 0, "09:00"},
	{11, 9, "11:09"},
}

var addTests = []struct {
	h, m, a int
	disp    string
}{
	{10, 0, 3, "10:03"},
	{10, 0, 61, "11:01"},
	{23, 30, 60, "00:30"},
	{10, 0, -90, "08:30"},
	{0, 30, -60, "23:30"},
}

func TestClock(t *testing.T) {
	for _, n := range newTests {
		if c := New(n.h, n.m); c.String() != n.disp {
			t.Fatalf("New(%d, %d) = %q, want %q", n.h, n.m, c, n.disp)
		}
	}
	for _, a := range addTests {
		if c := New(a.h, a.m).Add(a.a); c.String() != a.disp {
			t.Fatalf("New(%d, %d).Add(%d) = %q, want %q", a.h, a.m, a.a, c, a.disp)
		}
	}
	clock0 := New(15, 37)
	clock1 := New(15, 37)
	clock2 := New(15, 36)
	clock3 := New(14, 37)
	if clock0 != clock1 {
		t.Fatal(clock0, "!=", clock1, "want ==")
	}
	if clock2 == clock1 {
		t.Fatal(clock2, "==", clock1, "want !=")
	}
	if clock3 == clock1 {
		t.Fatal(clock3, "==", clock1, "want !=")
	}
}
