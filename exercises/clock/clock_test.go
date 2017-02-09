package clock

import (
	"reflect"
	"testing"
)

// Clock type API:
//
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
//
// The Add method should also handle subtraction by accepting negative values.

// To satisfy the README requirement about clocks being equal, values of
// your Clock type need to work with the == operator. This means that if your
// New function returns a pointer rather than a value, your clocks will
// probably not work with ==.
//
// While the time.Time type in the standard library (https://golang.org/pkg/time/#Time)
// doesn't necessarily need to be used as a basis for your Clock type, it might
// help to look at how constructors there (Date and Now) return values rather
// than pointers. Note also how most time.Time methods have value receivers
// rather than pointer receivers.
//
// For some useful guidelines on when to use a value receiver or a pointer
// receiver see: https://github.com/golang/go/wiki/CodeReviewComments#receiver-type

const targetTestVersion = 4

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestCreateClock(t *testing.T) {
	for _, n := range timeTests {
		if got := New(n.h, n.m); got.String() != n.want {
			t.Fatalf("New(%d, %d) = %q, want %q", n.h, n.m, got, n.want)
		}
	}
	t.Log(len(timeTests), "test cases")
}

func TestAddMinutes(t *testing.T) {
	for _, a := range addTests {
		if got := New(a.h, a.m).Add(a.a); got.String() != a.want {
			t.Fatalf("New(%d, %d).Add(%d) = %q, want %q",
				a.h, a.m, a.a, got, a.want)
		}
	}
	t.Log(len(addTests), "test cases")
}

func TestCompareClocks(t *testing.T) {
	for _, e := range eqTests {
		clock1 := New(e.c1.h, e.c1.m)
		clock2 := New(e.c2.h, e.c2.m)
		got := clock1 == clock2
		if got != e.want {
			t.Log("Clock1:", clock1)
			t.Log("Clock2:", clock2)
			t.Logf("Clock1 == Clock2 is %t, want %t", got, e.want)
			if reflect.DeepEqual(clock1, clock2) {
				t.Log("(Hint: see comments in clock_test.go.)")
			}
			t.FailNow()
		}
	}
	t.Log(len(eqTests), "test cases")
}

func BenchmarkAddMinutes(b *testing.B) {
	c := New(12, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range addTests {
			c.Add(a.a)
		}
	}
}

func BenchmarkCreateClocks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range timeTests {
			New(n.h, n.m)
		}
	}
}
