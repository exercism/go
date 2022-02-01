package interest

import (
	"math"
	"testing"
)

const floatEqualityThreshold = 1e-5

func floatingPointEquals(got, want float64) bool {
	absoluteDifferenceBelowThreshold := math.Abs(got-want) <= floatEqualityThreshold
	relativeDifferenceBelowThreshold := math.Abs(got-want)/(math.Abs(got)+math.Abs(want)) <= floatEqualityThreshold
	return absoluteDifferenceBelowThreshold || relativeDifferenceBelowThreshold
}

func TestInterestRate(t *testing.T) {
	tests := []struct {
		name    string
		balance float64
		want    float32
	}{
		{
			name:    "Minimal first interest rate",
			balance: 0,
			want:    float32(0.5),
		},
		{
			name:    "Tiny first interest rate",
			balance: 0.000001,
			want:    float32(0.5),
		},
		{
			name:    "Maximum first interest rate",
			balance: 999.9999,
			want:    float32(0.5),
		},
		{
			name:    "Minimal second interest rate",
			balance: 1000.0,
			want:    float32(1.621),
		},
		{
			name:    "Tiny second interest rate",
			balance: 1000.0001,
			want:    float32(1.621),
		},
		{
			name:    "Maximum second interest rate",
			balance: 4999.9990,
			want:    float32(1.621),
		},
		{
			name:    "Minimal third interest rate",
			balance: 5000.0000,
			want:    float32(2.475),
		},
		{
			name:    "Tiny third interest rate",
			balance: 5000.0001,
			want:    float32(2.475),
		},
		{
			name:    "Large third interest rate",
			balance: 5639998.742909,
			want:    float32(2.475),
		},
		{
			name:    "Rate on minimal negative balance",
			balance: -0.000001,
			want:    float32(3.213),
		},
		{
			name:    "Rate on small negative balance",
			balance: -0.123,
			want:    float32(3.213),
		},
		{
			name:    "Rate on regular negative balance",
			balance: -300.0,
			want:    float32(3.213),
		},
		{
			name:    "Rate on large negative balance",
			balance: -152964.231,
			want:    float32(3.213),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InterestRate(tt.balance)
			if !floatingPointEquals(float64(got), float64(tt.want)) {
				t.Errorf(
					"InterestRate(%f) = %f, want %f",
					tt.balance,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestInterest(t *testing.T) {
	tests := []struct {
		name    string
		balance float64
		want    float64
	}{
		{
			name:    "Interest on negative balance",
			balance: -10000.0,
			want:    -321.3,
		},
		{
			name:    "Interest on small balance",
			balance: 555.55,
			want:    2.77775,
		},
		{
			name:    "Interest on medium balance",
			balance: 4999.99,
			want:    81.0498379,
		},
		{
			name:    "Interest on large balance",
			balance: 34600.80,
			want:    856.369767,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Interest(tt.balance)
			if !floatingPointEquals(got, tt.want) {
				t.Errorf(
					"Interest(%f) = %f, want %f",
					tt.balance,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestAnnualBalanceUpdate(t *testing.T) {
	tests := []struct {
		name    string
		balance float64
		want    float64
	}{
		{
			name:    "Annual balance update for empty start balance",
			balance: 0.0,
			want:    0.0000,
		},
		{
			name:    "Annual balance update for small positive start balance",
			balance: 0.000001,
			want:    0.000001005,
		},
		{
			name:    "Annual balance update for average positive start balance",
			balance: 1000.0,
			want:    1016.210000,
		},
		{
			name:    "Annual balance update for large positive start balance",
			balance: 1000.0001,
			want:    1016.210101621,
		},
		{
			name:    "Annual balance update for huge positive start balance",
			balance: 898124017.826243404425,
			want:    920352586.410925,
		},
		{
			name:    "Annual balance update for small negative start balance",
			balance: -0.123,
			want:    -0.12695199,
		},
		{
			name:    "Annual balance update for large negative start balance",
			balance: -152964.231,
			want:    -157878.971832,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnnualBalanceUpdate(tt.balance)
			if !floatingPointEquals(got, tt.want) {
				t.Errorf(
					"AnnualBalanceUpdate(%f) = %f, want %f",
					tt.balance,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestYearsBeforeDesiredBalance(t *testing.T) {
	tests := []struct {
		name          string
		balance       float64
		targetBalance float64
		want          int
	}{
		{
			name:          "Years before desired balance for small start balance",
			balance:       100.0,
			targetBalance: 125.80,
			want:          47,
		},
		{
			name:          "Years before desired balance for average start balance",
			balance:       1000.0,
			targetBalance: 1100.0,
			want:          6,
		},
		{
			name:          "Years before desired balance for large start balance",
			balance:       8080.80,
			targetBalance: 9090.90,
			want:          5,
		},
		{
			name:          "Years before large difference between start and target balance",
			balance:       2345.67,
			targetBalance: 12345.6789,
			want:          85,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := YearsBeforeDesiredBalance(tt.balance, tt.targetBalance)
			if got != tt.want {
				t.Errorf(
					"YearsBeforeDesiredBalance(%f, %f) = %d, want %d",
					tt.balance,
					tt.targetBalance,
					got,
					tt.want,
				)
			}
		})
	}
}
