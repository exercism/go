package elon

import (
	"testing"
)

func TestDrive(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected Car
	}{
		{
			name: "Drive the car once.",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
				distance:     0,
			},
			expected: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      98,
				distance:     5,
			},
		},
		{
			name: "Drive when battery percentage is below battery drain",
			car: Car{
				speed:        5,
				batteryDrain: 7,
				battery:      3,
				distance:     0,
			},
			expected: Car{
				speed:        5,
				batteryDrain: 7,
				battery:      3,
				distance:     0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			car := tt.car
			car.Drive()

			if car != tt.expected {
				t.Errorf("method Drive() = %v; expected %v", car, tt.expected)
			}
		})
	}
}

func TestDisplayDistance(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected string
	}{
		{
			name: "Car displays driven distance",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
				distance:     0,
			},
			expected: "Driven 0 meters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.car.DisplayDistance()

			if got != tt.expected {
				t.Errorf("method DisplayDistance() = %v; Expected value %v", got, tt.expected)
			}
		})
	}
}

func TestDisplayBattery(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected string
	}{
		{
			name: "Display battery percentage",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
				distance:     0,
			},
			expected: "Battery at 100%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.car.DisplayBattery()

			if got != tt.expected {
				t.Errorf("method DisplayBattery() = %v; Expected value %v", got, tt.expected)
			}
		})
	}
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		car           Car
		trackDistance int
		expected      bool
	}{
		{
			name: "Car has 100% battery. Car can finish the race",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
			},
			trackDistance: 100,
			expected:      true,
		},
		{
			name: "Car has 40% battery. Car can finish the race",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      40,
			},
			trackDistance: 100,
			expected:      true,
		},
		{
			name: "Car has 30% battery. Car cannot finish the race",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      30,
			},
			trackDistance: 100,
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.car.CanFinish(tt.trackDistance)

			if got != tt.expected {
				t.Errorf("method CanFinish(%v) = %v. Expected value %v", tt.trackDistance, got, tt.expected)
			}
		})
	}
}
