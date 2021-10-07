package cars

import (
	"math"
	"testing"
)

const FloatEqualityThreshold = 1e-9

func TestSuccessRate(t *testing.T) {
	tests := []struct {
		name  string
		speed int
		want  float64
	}{
		{
			name:  "calculate success rate for speed zero",
			speed: 0,
			want:  0.0,
		},
		{
			name:  "calculate success rate for speed one",
			speed: 1,
			want:  1.0,
		},
		{
			name:  "calculate success rate for speed four",
			speed: 4,
			want:  1.0,
		},
		{
			name:  "calculate success rate for speed seven",
			speed: 7,
			want:  0.9,
		},
		{
			name:  "calculate success rate for speed nine",
			speed: 9,
			want:  0.77,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SuccessRate(tt.speed)
			if math.Abs(got-tt.want) > FloatEqualityThreshold {
				t.Errorf(
					"SuccessRate(%d) = %f, want %f",
					tt.speed,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalculateProductionRatePerHour(t *testing.T) {
	tests := []struct {
		name  string
		speed int
		want  float64
	}{
		{
			name:  "calculate production rate per hour for speed zero",
			speed: 0,
			want:  0.0,
		},
		{
			name:  "calculate production rate per hour for speed one",
			speed: 1,
			want:  221.0,
		},
		{
			name:  "calculate production rate per hour for speed four",
			speed: 4,
			want:  884.0,
		},
		{
			name:  "calculate production rate per hour for speed seven",
			speed: 7,
			want:  1392.3,
		},
		{
			name:  "calculate production rate per hour for speed nine",
			speed: 9,
			want:  1531.53,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateProductionRatePerHour(tt.speed)
			if math.Abs(got-tt.want) > FloatEqualityThreshold {
				t.Errorf(
					"CalculateProductionRatePerHour(%d) = %f, want %f",
					tt.speed,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalculateProductionRatePerMinute(t *testing.T) {
	tests := []struct {
		name  string
		speed int
		want  int
	}{
		{
			name:  "calculate production rate per minute for speed zero",
			speed: 0,
			want:  0,
		},
		{
			name:  "calculate production rate per minute for speed one",
			speed: 1,
			want:  3,
		},
		{
			name:  "calculate production rate per minute for speed five",
			speed: 5,
			want:  16,
		},
		{
			name:  "calculate production rate per minute for speed eight",
			speed: 8,
			want:  26,
		},
		{
			name:  "calculate production rate per minute for speed ten",
			speed: 10,
			want:  28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateProductionRatePerMinute(tt.speed)
			if got != tt.want {
				t.Errorf(
					"CalculateProductionRatePerMinute(%d) = %d, want %d",
					tt.speed,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestCalculateLimitedProductionRatePerHour(t *testing.T) {
	tests := []struct {
		name  string
		speed int
		limit float64
		want  float64
	}{
		{
			name:  "calculate limited production rate per hour for speed zero",
			speed: 0,
			limit: 500.0,
			want:  0.0,
		},
		{
			name:  "calculate limited production rate per hour below limit",
			speed: 1,
			limit: 500.0,
			want:  221.0,
		},
		{
			name:  "calculate limited production rate per hour above limit",
			speed: 9,
			limit: 500.0,
			want:  500.0,
		},
		{
			name:  "calculate limited production rate per hour at limit",
			speed: 3,
			limit: 663.0,
			want:  663.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateLimitedProductionRatePerHour(tt.speed, tt.limit)
			if math.Abs(got-tt.want) > FloatEqualityThreshold {
				t.Errorf(
					"CalculateLimitedProductionRatePerHour(%d, %f) = %f, want %f",
					tt.speed,
					tt.limit,
					got,
					tt.want,
				)
			}
		})
	}
}
