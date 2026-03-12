package bafflingbirthdays

import (
	"testing"
	"time"
)

func TestSharedBirthday(t *testing.T) {
	for _, tc := range sharedTestCase {
		t.Run(tc.description, func(t *testing.T) {
			var dates []time.Time
			for _, d := range tc.input {
				if date, err := time.Parse(time.DateOnly, d); err != nil {
					t.Fatalf("Error parsing test data, %q; please report on https://forum.exercism.org/", d)
				} else {
					dates = append(dates, date)
				}
			}
			if actual := SharedBirthday(dates); actual != tc.expected {
				t.Fatalf("SharedBirthday(%#v) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestEstimatedProbabilityOfSharedBirthday(t *testing.T) {
	for _, tc := range probabilityTestCase {
		t.Run(tc.description, func(t *testing.T) {
			actual := EstimatedProbability(tc.input)
			diff := actual - tc.expected
			if diff < 0 {
				diff = -diff
			}
			if diff > 2 {
				t.Fatalf("EstimatedProbability(%v) = %v, want: %v ± 2%%", tc.input, actual, tc.expected)
			}
		})
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func TestRandomBirthdays(t *testing.T) {
	size := 1024 * 32
	birthdates := RandomBirthdates(size)
	// "uuid": "70b38cea-d234-4697-b146-7d130cd4ee12"
	t.Run("generate requested number of birthdates", func(t *testing.T) {
		if len(birthdates) != size {
			t.Fatalf("RandomBirthdates(%d) gave %d birthdates; want %d birthdates", size, len(birthdates), size)
		}
	})
	// "uuid": "d9d5b7d3-5fea-4752-b9c1-3fcd176d1b03"
	t.Run("years are not leap years", func(t *testing.T) {
		for _, date := range birthdates {
			if isLeapYear(date.Year()) {
				t.Fatalf("RandomBirthdates(%d) gave a leap year, %v; want non-leap years old", size, date)
			}
		}
	})
	// "uuid": "d1074327-f68c-4c8a-b0ff-e3730d0f0521"
	t.Run("months are random", func(t *testing.T) {
		avg := float64(size / 12)
		errBand := 0.2
		months := make(map[time.Month]int)
		for _, date := range birthdates {
			months[date.Month()]++
		}
		for month, count := range months {
			if float64(count) < (1-errBand)*avg || float64(count) > (1+errBand)*avg {
				t.Fatalf("RandomBirthdates(%d) has %d dates in %v; expected a uniform distribution (about %v per month)", size, count, month, avg)
			}
		}
	})
	// "uuid": "7df706b3-c3f5-471d-9563-23a4d0577940"
	t.Run("days are random", func(t *testing.T) {
		days := make(map[int]int)
		for _, date := range birthdates {
			days[date.Day()]++
		}
		for day, count := range days {
			want := int(0.80 * float64(size) / 31)
			if day == 31 {
				want /= 2
			}
			if count < want {
				t.Fatalf("RandomBirthdates(%d), day-of-month %d appears %d times; expected a uniform distribution (about %d per day)", size, day, count, want)
			}
		}
	})
}

func BenchmarkSharedBirthday(b *testing.B) {
	for _, tc := range sharedTestCase {
		var dates []time.Time
		for _, d := range tc.input {
			date, _ := time.Parse(time.DateOnly, d)
			dates = append(dates, date)
		}
		for range b.N {
			SharedBirthday(dates)
		}
	}
}

func BenchmarkRandomBirthdates(b *testing.B) {
	size := 1024 * 16
	for range b.N {
		RandomBirthdates(size)
	}
}
