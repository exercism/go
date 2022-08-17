package gigasecond

// Write a function AddGigasecond that works with time.Time.

import (
	"testing"
	"time"
)

// date formats used in test data
const (
	fmtD  = "2006-01-02"
	fmtDT = "2006-01-02T15:04:05"
)

func TestAddGigasecond(t *testing.T) {
	for _, tc := range addCases {
		t.Run(tc.description, func(t *testing.T) {
			in := parse(tc.in, t)
			want := parse(tc.want, t)
			got := AddGigasecond(in)
			if !got.Equal(want) {
				t.Fatalf("AddGigasecond(%v) = %v, want: %v", in, got, want)
			}
		})
	}
}

func parse(s string, t *testing.T) time.Time {
	tt, err := time.Parse(fmtDT, s) // try full date time format first
	if err != nil {
		tt, err = time.Parse(fmtD, s) // also allow just date
	}
	if err != nil {
		t.Fatalf("error in test setup: TestAddGigasecond requires datetime in one of the following formats: \nformat 1:%q\nformat 2:%q\ngot:%q", fmtD, fmtDT, s)
	}
	return tt
}

func BenchmarkAddGigasecond(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		AddGigasecond(time.Time{})
	}
}
