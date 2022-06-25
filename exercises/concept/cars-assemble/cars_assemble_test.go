package cars

import (
	"math"
	"testing"
)

const floatEqualityThreshold = 1e-9

func floatingPointEquals(got, want float64) bool {
	absoluteDifferenceBelowThreshold := math.Abs(got-want) <= floatEqualityThreshold
	relativeDifferenceBelowThreshold := math.Abs(got-want)/(math.Abs(got)+math.Abs(want)) <= floatEqualityThreshold
	return absoluteDifferenceBelowThreshold || relativeDifferenceBelowThreshold
}

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	tests := []struct {
		name           string
		productionRate int
		successRate    float64
		want           float64
	}{
		{
			name:           "calculate working cars per hour for production rate 0",
			productionRate: 0,
			successRate:    100,
			want:           0.0,
		},
		{
			name:           "calculate working cars per hour for 100% success rate",
			productionRate: 221,
			successRate:    100,
			want:           221.0,
		},
		{
			name:           "calculate working cars per hour for 80% success rate",
			productionRate: 426,
			successRate:    80,
			want:           340.8,
		},
		{
			name:           "calculate working cars per hour for 20.5% success rate",
			productionRate: 6824,
			successRate:    20.5,
			want:           1398.92,
		},
		{
			name:           "calculate working cars per hour for 0% success rate",
			productionRate: 8000,
			successRate:    0,
			want:           0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateWorkingCarsPerHour(tt.productionRate, tt.successRate)
			if !floatingPointEquals(got, tt.want) {
				t.Errorf(
					"CalculateWorkingCarsPerHour(%d, %f) = %f, want %f",
					tt.productionRate,
					tt.successRate,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalculateWorkingCarsPerMinute(t *testing.T) {
	tests := []struct {
		name           string
		productionRate int
		successRate    float64
		want           int
	}{
		{
			name:           "calculate working cars per minute for production rate 0",
			productionRate: 0,
			successRate:    100,
			want:           0,
		},
		{
			name:           "calculate working cars per minute for 100% success rate",
			productionRate: 221,
			successRate:    100,
			want:           3,
		},
		{
			name:           "calculate working cars per minute for 80% success rate",
			productionRate: 426,
			successRate:    80,
			want:           5,
		},
		{
			name:           "calculate working cars per minute for 20.5% success rate",
			productionRate: 6824,
			successRate:    20.5,
			want:           23,
		},
		{
			name:           "calculate working cars per minute for 0% success rate",
			productionRate: 8000,
			successRate:    0,
			want:           0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateWorkingCarsPerMinute(tt.productionRate, tt.successRate)
			if got != tt.want {
				t.Errorf(
					"CalculateWorkingCarsPerMinute(%d, %f) = %d, want %d",
					tt.productionRate,
					tt.successRate,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalculateCost(t *testing.T) {
	tests := []struct {
		name      string
		carsCount int
		want      uint
	}{
		{
			name:      "calculate cost to produce 0 cars",
			carsCount: 0,
			want:      0,
		},
		{
			name:      "calculate the cost of materials to produce 1 car",
			carsCount: 1,
			want:      10000,
		},
		{
			name:      "calculate cost to produce 2 cars",
			carsCount: 2,
			want:      20000,
		},
		{
			name:      "calculate cost to produce 9 cars",
			carsCount: 9,
			want:      90000,
		},
		{
			name:      "calculate cost to produce 10 cars",
			carsCount: 10,
			want:      95000,
		},
		{
			name:      "calculate cost to produce 100 cars",
			carsCount: 100,
			want:      950000,
		},
		{
			name:      "calculate cost to produce 21 cars",
			carsCount: 21,
			want:      200000,
		},
		{
			name:      "calculate cost to produce 37 cars",
			carsCount: 37,
			want:      355000,
		},
		{
			name:      "calculate cost to produce 56 cars",
			carsCount: 56,
			want:      535000,
		},
		{
			name:      "calculate cost to produce 148 cars",
			carsCount: 148,
			want:      1410000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateCost(tt.carsCount)
			if got != tt.want {
				t.Errorf(
					"CalculateCost(%d) = %d, want %d",
					tt.carsCount,
					got,
					tt.want,
				)
			}
		})
	}
}
