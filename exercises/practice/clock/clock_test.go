package clock

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestCreateClock(t *testing.T) {
	for _, tc := range timeTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := New(tc.h, tc.m); tc.expected != actual.String() {
				t.Errorf("New(%d, %d) = %q, want %q", tc.h, tc.m, actual, tc.expected)
			}
		})
	}
}

func TestAddMinutes(t *testing.T) {
	for _, tc := range addTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := New(tc.h, tc.m).Add(tc.addedValue); tc.expected != actual.String() {
				t.Errorf("New(%d, %d).Add(%d) = %q, want %q", tc.h, tc.m, tc.addedValue, actual, tc.expected)
			}
		})
	}
}

func TestSubtractMinutes(t *testing.T) {
	for _, tc := range subtractTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := New(tc.h, tc.m).Subtract(tc.subtractedValue); tc.expected != actual.String() {
				t.Errorf("New(%d, %d).Subtract(%d) = %q, want %q", tc.h, tc.m, tc.subtractedValue, actual, tc.expected)
			}
		})
	}
}

func TestAddMinutesStringless(t *testing.T) {
	for _, tc := range addTestCases {
		t.Run(tc.description, func(t *testing.T) {
			split := strings.SplitN(tc.expected, ":", 2)
			if len(split) != 2 {
				t.Fatalf("error in test setup: expected time in format hh:mm, got: %s", tc.expected)
			}
			wantHour, _ := strconv.Atoi(split[0])
			wantMin, _ := strconv.Atoi(split[1])

			expected := New(wantHour, wantMin)
			if actual := New(tc.h, tc.m).Add(tc.addedValue); !reflect.DeepEqual(actual, expected) {
				t.Errorf("New(%d, %d).Add(%d)\n\t Got: %q (%#v)\n\tWant: %q (%#v)",
					tc.h, tc.m, tc.addedValue, actual, actual, expected, expected)
			}
		})
	}
}

func TestSubtractMinutesStringless(t *testing.T) {
	for _, tc := range subtractTestCases {
		t.Run(tc.description, func(t *testing.T) {
			split := strings.SplitN(tc.expected, ":", 2)
			if len(split) != 2 {
				t.Fatalf("error in test setup: expected time in format hh:mm, got: %s", tc.expected)
			}
			wantHour, _ := strconv.Atoi(split[0])
			wantMin, _ := strconv.Atoi(split[1])

			expected := New(wantHour, wantMin)
			if actual := New(tc.h, tc.m).Subtract(tc.subtractedValue); !reflect.DeepEqual(actual, expected) {
				t.Errorf("New(%d, %d).Subtract(%d)\n\t Got: %q (%#v)\n\tWant: %q (%#v)",
					tc.h, tc.m, tc.subtractedValue, actual, actual, expected, expected)
			}
		})
	}
}

func TestCompareClocks(t *testing.T) {
	for _, tc := range equalTestCases {
		t.Run(tc.description, func(t *testing.T) {
			clock1 := New(tc.c1.h, tc.c1.m)
			clock2 := New(tc.c2.h, tc.c2.m)
			actual := clock1 == clock2
			if actual != tc.expected {
				t.Errorf("Clock1 == Clock2 is %t, want %t\n\tClock1: %q (%#v)\n\tClock2: %q (%#v)",
					actual, tc.expected, clock1, clock1, clock2, clock2)
				if reflect.DeepEqual(clock1, clock2) {
					t.Log("(Hint: see comments in clock_test.go.)")
				}
			}
		})
	}
}

func BenchmarkAddMinutes(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	c := New(12, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range addTestCases {
			c.Add(a.addedValue)
		}
	}
}

func BenchmarkSubtractMinutes(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	c := New(12, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range subtractTestCases {
			c.Subtract(a.subtractedValue)
		}
	}
}

func BenchmarkCreateClocks(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, n := range timeTestCases {
			New(n.h, n.m)
		}
	}
}
