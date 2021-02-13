package meetup

import "time"

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = -6
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	if wSched == Last {
		month++
	}
	t := time.Date(year, month, int(wSched), 12, 0, 0, 0, time.UTC)
	return t.Day() + int(wDay-t.Weekday()+7)%7
}
