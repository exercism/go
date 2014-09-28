package clock

import (
	"reflect"
	"testing"
)

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
	{10, 0, 3, "10:03"},   // simple
	{10, 0, 61, "11:01"},  // add > 1 hour
	{23, 30, 60, "00:30"}, // add across midnight
	{10, 0, -90, "08:30"}, // subtract > 1 hour
	{0, 30, -60, "23:30"}, // subtract across midnight

	{0, 45, 40, "01:25"},   // hour carry
	{1, 25, -40, "00:45"},  // hour borrow
	{23, 45, 40, "00:25"},  // carry across midnight
	{0, 25, -40, "23:45"},  // borrow across midnight
	{0, 0, 160, "02:40"},   // add > 2 hrs
	{0, 45, 160, "03:25"},  // add > 2 hrs with carry
	{0, 0, -160, "21:20"},  // subtract > 2 hrs
	{6, 15, -160, "03:35"}, // subtract > 2 hrs with borrow

	{0, 160, 0, "02:40"},  // initial minutes roll over
	{25, 0, 0, "01:00"},   // initial hour rolls over
	{25, 160, 0, "03:40"}, // both rollover
	{0, -160, 0, "21:20"}, // same cases, negative
	{-25, 0, 0, "23:00"},
	{-25, -160, 0, "20:20"},
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
	if !reflect.DeepEqual(clock0, clock1) {
		t.Fatal(clock0, "!=", clock1, "want ==")
	}
	if reflect.DeepEqual(clock1, clock2) {
		t.Fatal(clock2, "==", clock1, "want !=")
	}
	if reflect.DeepEqual(clock1, clock3) {
		t.Fatal(clock3, "==", clock1, "want !=")
	}

}
