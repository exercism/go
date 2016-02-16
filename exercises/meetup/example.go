package meetup

import "time"

const testVersion = 3

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second              = 8
	Third               = 15
	Fourth              = 22
	Teenth              = 13
	Last                = -6
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	if wSched == Last {
		month++
	}
	t := time.Date(year, month, int(wSched), 12, 0, 0, 0, time.UTC)
	return t.Day() + int(wDay-t.Weekday()+7)%7
}
