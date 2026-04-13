package booking

import (
	"testing"
	"time"
)

type scheduleTestCase struct {
	description string
	in          string
	want        time.Time
}

func TestSchedule(t *testing.T) {
	tests := []scheduleTestCase{
		{description: "Schedule 1", in: "7/13/2020 20:32:00", want: time.Date(2020, time.July, 13, 20, 32, 0, 0, time.UTC)},
		{description: "Schedule 2", in: "11/28/1984 2:02:02", want: time.Date(1984, time.November, 28, 2, 2, 2, 0, time.UTC)},
		{description: "Schedule 3", in: "2/29/2112 11:59:59", want: time.Date(2112, time.February, 29, 11, 59, 59, 0, time.UTC)},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := Schedule(tc.in); !got.Equal(tc.want) {
				t.Errorf("Schedule(%s) = '%v', want '%v'", tc.in, got, tc.want)
			}
		})
	}
}

type hasPassedTestCase struct {
	description string
	in          string
	want        bool
}

func TestHasPassed(t *testing.T) {
	tests := []hasPassedTestCase{
		{description: "HasPassed 1", in: "October 3, 2019 20:32:00", want: true},
		{description: "HasPassed 2", in: "January 28, 1974 2:02:02", want: true},
		{description: "HasPassed 3", in: "December 9, 2112 11:59:59", want: false},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := HasPassed(tc.in); got != tc.want {
				t.Errorf("HasPassed(%s) = '%v', want '%v'", tc.in, got, tc.want)
			}
		})
	}
}

type isAfternoonAppointmentTestCase struct {
	description string
	in          string
	want        bool
}

func TestIsAfternoonAppointment(t *testing.T) {
	tests := []isAfternoonAppointmentTestCase{
		{description: "IsAfternoonAppointment 1", in: "Thursday, May 13, 2010 20:32:00", want: false},
		{description: "IsAfternoonAppointment 2", in: "Friday, March 8, 1974 12:02:02", want: true},
		{description: "IsAfternoonAppointment 3", in: "Friday, September 9, 2112 11:59:59", want: false},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := IsAfternoonAppointment(tc.in); got != tc.want {
				t.Errorf("IsAfternoonAppointment(%s) = '%v', want '%v'", tc.in, got, tc.want)
			}
		})
	}
}

type descriptionTestCase struct {
	description string
	in          string
	want        string
}

func TestDescription(t *testing.T) {
	tests := []descriptionTestCase{
		{description: "Description 1", in: "6/6/2005 10:30:00", want: "You have an appointment on Monday, June 6, 2005, at 10:30."},
		{description: "Description 2", in: "9/19/1994 12:15:00", want: "You have an appointment on Monday, September 19, 1994, at 12:15."},
		{description: "Description 3", in: "4/4/2012 16:45:00", want: "You have an appointment on Wednesday, April 4, 2012, at 16:45."},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := Description(tc.in); got != tc.want {
				t.Errorf("Description(%s) = '%v', want '%v'", tc.in, got, tc.want)
			}
		})
	}
}

type anniversaryDateTestCase struct {
	description string
	want        time.Time
}

func TestAnniversaryDate(t *testing.T) {
	tests := []anniversaryDateTestCase{
		{description: "AnniversaryDate 1", want: time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := AnniversaryDate(); !got.Equal(tc.want) {
				t.Errorf("AnniversaryDate() = '%v', want '%v'", got, tc.want)
			}
		})
	}
}
