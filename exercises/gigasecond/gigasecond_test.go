package gigasecond

// Write a function AddGigasecond that works with time.Time.

import (
	"os"
	"testing"
	"time"
)

const targetTestVersion = 4

// date formats used in test data
const (
	fmtD  = "2006-01-02"
	fmtDT = "2006-01-02T15:04:05"
)

func TestAddGigasecond(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
	for _, tc := range addCases {
		in := parse(tc.in, t)
		want := parse(tc.want, t)
		got := AddGigasecond(in)
		if !got.Equal(want) {
			t.Fatalf(`AddGigasecond(%s)
   = %s
want %s`, in, got, want)
		}
	}
	t.Log("Tested", len(addCases), "cases.")
}

func parse(s string, t *testing.T) time.Time {
	tt, err := time.Parse(fmtDT, s) // try full date time format first
	if err != nil {
		tt, err = time.Parse(fmtD, s) // also allow just date
	}
	if err != nil {
		// can't run tests if input won't parse.  if this seems to be a
		// development or ci environment, raise an error.  if this condition
		// makes it to the solver though, ask for a bug report.
		_, statErr := os.Stat("example_gen.go")
		if statErr == nil || os.Getenv("TRAVIS_GO_VERSION") > "" {
			t.Fatal(err)
		} else {
			t.Log(err)
			t.Skip("(Not your problem.  " +
				"please file issue at https://github.com/exercism/xgo.)")
		}
	}
	return tt
}

func BenchmarkAddGigasecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddGigasecond(time.Time{})
	}
}
