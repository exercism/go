package jedlik

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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			car := tc.car
			car.Drive()

			if car != tc.expected {
				t.Errorf("method Drive() = %v; expected %v", car, tc.expected)
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.car.DisplayDistance()

			if got != tc.expected {
				t.Errorf("method DisplayDistance() = %v; Expected value %v", got, tc.expected)
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.car.DisplayBattery()

			if got != tc.expected {
				t.Errorf("method DisplayBattery() = %v; Expected value %v", got, tc.expected)
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
			name: "Car has 60% battery. Car cannot finish the race",
			car: Car{
				speed:        3,
				batteryDrain: 3,
				battery:      60,
			},
			trackDistance: 61,
			expected:      false,
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.car.CanFinish(tc.trackDistance)

			if got != tc.expected {
				t.Errorf("method CanFinish(%v) = %v. Expected value %v", tc.trackDistance, got, tc.expected)
			}
		})
	}
}
