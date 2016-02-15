package meetup

import "testing"

/* API
package meetup
type Weekschedule
WeekSchedule First
WeekSchedule Second
WeekSchedule Third
WeekSchedule Fourth
WeekSchedule Last
WeekSchedule Teenth
func Day(WeekSchedule, time.Weekday, time.Month, int) int
*/

const targetTestVersion = 3

var weekName = map[WeekSchedule]string{
	First:  "first",
	Second: "second",
	Third:  "third",
	Fourth: "fourth",
	Teenth: "teenth",
	Last:   "last",
}

func TestDay(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, test := range testCases {
		res := Day(test.week, test.weekday, test.month, test.year)
		if res != test.expDay {
			t.Fatalf("For %s %s of %s 2013 got date of %d, want %d",
				weekName[test.week], test.weekday, test.month, res, test.expDay)
		}
	}
}
