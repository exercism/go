package meetup

import (
	"testing"
	"time"
)

/* API
package meetup
type Weekschedule
WeekSchedule First
WeekSchedule Second
WeekSchedule Third
WeekSchedule Fourth
WeekSchedule Last
WeekSchedule Teenth
func MeetupDay(WeekSchedule, time.Weekday, time.Month, int) int
*/

type GivenYearTestCase struct {
	expDay  int
	month   time.Month
	weekday time.Weekday
	week    WeekSchedule
}

var test2013 = []GivenYearTestCase{
	{13, time.May, time.Monday, Teenth},
	{19, time.August, time.Monday, Teenth},
	{16, time.September, time.Monday, Teenth},
	{19, time.March, time.Tuesday, Teenth},
	{16, time.April, time.Tuesday, Teenth},
	{13, time.August, time.Tuesday, Teenth},
	{16, time.January, time.Wednesday, Teenth},
	{13, time.February, time.Wednesday, Teenth},
	{19, time.June, time.Wednesday, Teenth},
	{16, time.May, time.Thursday, Teenth},
	{13, time.June, time.Thursday, Teenth},
	{19, time.September, time.Thursday, Teenth},
	{19, time.April, time.Friday, Teenth},
	{16, time.August, time.Friday, Teenth},
	{13, time.September, time.Friday, Teenth},
	{16, time.February, time.Saturday, Teenth},
	{13, time.April, time.Saturday, Teenth},
	{19, time.October, time.Saturday, Teenth},
	{19, time.May, time.Sunday, Teenth},
	{16, time.June, time.Sunday, Teenth},
	{13, time.October, time.Sunday, Teenth},

	{4, time.March, time.Monday, First},
	{1, time.April, time.Monday, First},
	{7, time.May, time.Tuesday, First},
	{4, time.June, time.Tuesday, First},
	{3, time.July, time.Wednesday, First},
	{7, time.August, time.Wednesday, First},
	{5, time.September, time.Thursday, First},
	{3, time.October, time.Thursday, First},
	{1, time.November, time.Friday, First},
	{6, time.December, time.Friday, First},
	{5, time.January, time.Saturday, First},
	{2, time.February, time.Saturday, First},
	{3, time.March, time.Sunday, First},
	{7, time.April, time.Sunday, First},

	{11, time.March, time.Monday, Second},
	{8, time.April, time.Monday, Second},
	{14, time.May, time.Tuesday, Second},
	{11, time.June, time.Tuesday, Second},
	{10, time.July, time.Wednesday, Second},
	{14, time.August, time.Wednesday, Second},
	{12, time.September, time.Thursday, Second},
	{10, time.October, time.Thursday, Second},
	{8, time.November, time.Friday, Second},
	{13, time.December, time.Friday, Second},
	{12, time.January, time.Saturday, Second},
	{9, time.February, time.Saturday, Second},
	{10, time.March, time.Sunday, Second},
	{14, time.April, time.Sunday, Second},

	{18, time.March, time.Monday, Third},
	{15, time.April, time.Monday, Third},
	{21, time.May, time.Tuesday, Third},
	{18, time.June, time.Tuesday, Third},
	{17, time.July, time.Wednesday, Third},
	{21, time.August, time.Wednesday, Third},
	{19, time.September, time.Thursday, Third},
	{17, time.October, time.Thursday, Third},
	{15, time.November, time.Friday, Third},
	{20, time.December, time.Friday, Third},
	{19, time.January, time.Saturday, Third},
	{16, time.February, time.Saturday, Third},
	{17, time.March, time.Sunday, Third},
	{21, time.April, time.Sunday, Third},

	{25, time.March, time.Monday, Fourth},
	{22, time.April, time.Monday, Fourth},
	{28, time.May, time.Tuesday, Fourth},
	{25, time.June, time.Tuesday, Fourth},
	{24, time.July, time.Wednesday, Fourth},
	{28, time.August, time.Wednesday, Fourth},
	{26, time.September, time.Thursday, Fourth},
	{24, time.October, time.Thursday, Fourth},
	{22, time.November, time.Friday, Fourth},
	{27, time.December, time.Friday, Fourth},
	{26, time.January, time.Saturday, Fourth},
	{23, time.February, time.Saturday, Fourth},
	{24, time.March, time.Sunday, Fourth},
	{28, time.April, time.Sunday, Fourth},

	{25, time.March, time.Monday, Last},
	{29, time.April, time.Monday, Last},
	{28, time.May, time.Tuesday, Last},
	{25, time.June, time.Tuesday, Last},
	{31, time.July, time.Wednesday, Last},
	{28, time.August, time.Wednesday, Last},
	{26, time.September, time.Thursday, Last},
	{31, time.October, time.Thursday, Last},
	{29, time.November, time.Friday, Last},
	{27, time.December, time.Friday, Last},
	{26, time.January, time.Saturday, Last},
	{23, time.February, time.Saturday, Last},
	{31, time.March, time.Sunday, Last},
	{28, time.April, time.Sunday, Last},
}

var weekName = map[WeekSchedule]string{
	First:  "first",
	Second: "second",
	Third:  "third",
	Fourth: "fourth",
	Teenth: "teenth",
	Last:   "last",
}

func TestMeetup_Day2013(t *testing.T) {
	for _, test := range test2013 {
		res := MeetupDay(test.week, test.weekday, test.month, 2013)
		if res != test.expDay {
			t.Fatalf("For %s %s of %s 2013 got date of %d, want %d",
				weekName[test.week], test.weekday, test.month, res, test.expDay)
		}
	}
}
