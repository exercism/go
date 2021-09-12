package cars

import (
	"testing"
)

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
			if got != tt.want {
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
					"CalculateWorkingItemsRate(%d) = %d, want %d",
					tt.speed,
					got,
					tt.want,
				)
			}
		})
	}
}
