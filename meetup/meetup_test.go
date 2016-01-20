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

const testVersion = 2

// Retired testVersions
// (none) fcef587e5fc4f22d82eea9366fedd7a5363147d1
// 1      0be0802be138413671fbc43294ef495f407cdcb2

var weekName = map[WeekSchedule]string{
	First:  "first",
	Second: "second",
	Third:  "third",
	Fourth: "fourth",
	Teenth: "teenth",
	Last:   "last",
}

func TestDay(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, test := range testCases {
		res := Day(test.week, test.weekday, test.month, test.year)
		if res != test.expDay {
			t.Fatalf("For %s %s of %s 2013 got date of %d, want %d",
				weekName[test.week], test.weekday, test.month, res, test.expDay)
		}
	}
}
