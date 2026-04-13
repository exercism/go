package savings

import (
	"testing"
)

type interestRateTestCase struct {
	description string
	want        float32
}

type daysPerYearTestCase struct {
	description string
	want        int
}

type getMonthTestCase struct {
	description string
	month       int
	want        int
}

type getAccountNumberTestCase struct {
	description string
	want        AccNo
}

func TestGetFixedInterestRate(t *testing.T) {
	tests := []interestRateTestCase{
		{description: "GetFixedInterestRate 1", want: 0.05},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := GetFixedInterestRate(); got != tc.want {
				t.Errorf("GetFixedInterestRate() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetDaysPerYear(t *testing.T) {
	tests := []daysPerYearTestCase{
		{description: "GetDaysPerYear 1", want: 365},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := GetDaysPerYear(); got != tc.want {
				t.Errorf("GetDaysPerYear() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetMonth(t *testing.T) {
	tests := []getMonthTestCase{
		{description: "GetMonth(Jan)", month: Jan, want: 1},
		{description: "GetMonth(Mar)", month: Mar, want: 3},
		{description: "GetMonth(Oct)", month: Oct, want: 10},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := GetMonth(tc.month); got != tc.want {
				t.Errorf("GetMonth(%v) = %v, want %v", tc.month, got, tc.want)
			}
		})
	}
}

func TestGetAccountNumber(t *testing.T) {
	tests := []getAccountNumberTestCase{
		{description: "GetAccountNumber 1", want: "XF348IJ"},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if got := GetAccountNumber(); got != tc.want {
				t.Errorf("GetAccountNumber() = %v, want %v", got, tc.want)
			}
		})
	}
}
