package clock

import (
	"reflect"
	"testing"
)

// Clock type API:
//
// Time(hour, minute int) Clock    // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
//
// The Add method should also handle subtraction by accepting negative values.
// To satisfy the readme requirement about clocks being equal, values of
// your Clock type need to work with the == operator.
//
// It might help to study the time.Time type in the standard library
// (https://golang.org/pkg/time/#Time) as a model.  See how constructors there
// (Date and Now) return Time values rather than pointers.  Note also how
// most time.Time methods have value receivers rather than pointer receivers.
// For more background on this read
// https://github.com/golang/go/wiki/CodeReviewComments#receiver-type.

const testVersion = 2

// Retired testVersions
// (none) 79937f6d58e25ebafe12d1cb4a9f88f4de70cfd6
// 1      8d0cb8b617be2e36b2ca5ad2034e5f80f2372924

func TestCreateClock(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, n := range timeTests {
		if got := Time(n.h, n.m); got.String() != n.want {
			t.Fatalf("Time(%d, %d) = %q, want %q", n.h, n.m, got, n.want)
		}
	}
	t.Log(len(timeTests), "test cases")
}

func TestAddMinutes(t *testing.T) {
	for _, a := range addTests {
		if got := Time(a.h, a.m).Add(a.a); got.String() != a.want {
			t.Fatalf("Time(%d, %d).Add(%d) = %q, want %q",
				a.h, a.m, a.a, got, a.want)
		}
	}
	t.Log(len(addTests), "test cases")
}

func TestCompareClocks(t *testing.T) {
	for _, e := range eqTests {
		clock1 := Time(e.c1.h, e.c1.m)
		clock2 := Time(e.c2.h, e.c2.m)
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
