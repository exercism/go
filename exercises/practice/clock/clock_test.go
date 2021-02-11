package clock

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// Clock API:
//
// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock
//
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

func TestSubtractMinutes(t *testing.T) {
	for _, a := range subtractTests {
		if got := New(a.h, a.m).Subtract(a.a); got.String() != a.want {
			t.Fatalf("New(%d, %d).Subtract(%d) = %q, want %q",
				a.h, a.m, a.a, got, a.want)
		}
	}
	t.Log(len(subtractTests), "test cases")
}

func TestAddMinutesStringless(t *testing.T) {
	for _, a := range addTests {
		var wantHour, wantMin int
		split := strings.SplitN(a.want, ":", 2)
		if len(split) > 0 {
			wantHour, _ = strconv.Atoi(split[0])
		}
		if len(split) > 1 {
			wantMin, _ = strconv.Atoi(split[1])
		}
		want := New(wantHour, wantMin)
		if got := New(a.h, a.m).Add(a.a); !reflect.DeepEqual(got, want) {
			t.Fatalf("New(%d, %d).Add(%d) = %v, want %v",
				a.h, a.m, a.a, got, want)
		}
	}
	t.Log(len(addTests), "test cases")
}

func TestSubtractMinutesStringless(t *testing.T) {
	for _, a := range subtractTests {
		var wantHour, wantMin int
		split := strings.SplitN(a.want, ":", 2)
		if len(split) > 0 {
			wantHour, _ = strconv.Atoi(split[0])
		}
		if len(split) > 1 {
			wantMin, _ = strconv.Atoi(split[1])
		}
		want := New(wantHour, wantMin)
		if got := New(a.h, a.m).Subtract(a.a); !reflect.DeepEqual(got, want) {
			t.Fatalf("New(%d, %d).Subtract(%d) = %v, want %v",
				a.h, a.m, a.a, got, want)
		}
	}
	t.Log(len(subtractTests), "test cases")
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

func TestAddAndCompare(t *testing.T) {
	clock1 := New(15, 45).Add(16)
	clock2 := New(16, 1)
	if !reflect.DeepEqual(clock1, clock2) {
		t.Errorf("clock.New(15,45).Add(16) differs from clock.New(16,1)")
	}
}

func TestSubtractAndCompare(t *testing.T) {
	clock1 := New(16, 1).Subtract(16)
	clock2 := New(15, 45)
	if !reflect.DeepEqual(clock1, clock2) {
		t.Errorf("clock.New(16,1).Subtract(16) differs from clock.New(15,45)")
	}
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

func BenchmarkSubtractMinutes(b *testing.B) {
	c := New(12, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range subtractTests {
			c.Subtract(a.a)
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
