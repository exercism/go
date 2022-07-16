package meetup

import "testing"

var weekName = map[WeekSchedule]string{
	First:  "first",
	Second: "second",
	Third:  "third",
	Fourth: "fourth",
	Teenth: "teenth",
	Last:   "last",
}

func TestDay(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Day(tc.week, tc.weekday, tc.month, tc.year)
			if actual != tc.expectedDay {
				t.Fatalf("Day(%q, %d, %d, %d) = %d, want: %d", weekName[tc.week], tc.weekday, tc.month, tc.year, actual, tc.expectedDay)
			}
		})
	}
}
