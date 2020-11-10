package constants

import (
	"testing"
)

func TestGetFixedInterestRate(t *testing.T) {
	tests := map[string]struct {
		want float32
	}{
		"GetFixedInterestRate 1": {want: 0.05},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := GetFixedInterestRate(); got != tc.want {
				t.Errorf("GetFixedInterestRate() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetDaysPerYear(t *testing.T) {
	tests := map[string]struct {
		want int
	}{
		"GetDaysPerYear 1": {want: 365},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := GetDaysPerYear(); got != tc.want {
				t.Errorf("GetDaysPerYear() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetMonth(t *testing.T) {
	tests := map[string]struct {
		arg, want int
	}{
		"GetMonth(Jan)": {arg: Jan, want: 1},
		"GetMonth(Mar)": {arg: Mar, want: 3},
		"GetMonth(Oct)": {arg: Oct, want: 10},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := GetMonth(tc.arg); got != tc.want {
				t.Errorf("GetMonth(%v) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestGetAccountNumber(t *testing.T) {
	tests := map[string]struct {
		want AccNo
	}{
		"GetAccountNumber 1": {want: "XF348IJ"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := GetAccountNumber(); got != tc.want {
				t.Errorf("GetAccountNumber() = %v, want %v", got, tc.want)
			}
		})
	}
}
