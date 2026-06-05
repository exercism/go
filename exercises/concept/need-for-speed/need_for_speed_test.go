package speed

import "testing"

func TestNewCar(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected Car
	}{
		{
			name: "Create a new car.",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
			},
			expected: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewCar(tc.car.speed, tc.car.batteryDrain)

			if got != tc.expected {
				t.Errorf("NewCar(%+v,%+v) = %+v; expected %+v", tc.car.speed, tc.car.batteryDrain, got, tc.expected)
			}
		})
	}
}

func TestNewTrack(t *testing.T) {
	tests := []struct {
		name     string
		track    Track
		expected Track
	}{
		{
			name: "Create a long track.",
			track: Track{
				distance: 800,
			},
			expected: Track{
				distance: 800,
			},
		},
		{
			name: "Create a short track.",
			track: Track{
				distance: 360,
			},
			expected: Track{
				distance: 360,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewTrack(tc.track.distance)

			if got != tc.expected {
				t.Errorf("NewTrack(%+v) = %+v; expected %+v", tc.track.distance, got, tc.expected)
			}
		})
	}
}

func TestDrive(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		expected Car
	}{
		{
			name: "Drive the car with full battery.",
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
			name: "Drive the car with exactly enough battery.",
			car: Car{
				speed:        18,
				batteryDrain: 10,
				battery:      10,
				distance:     0,
			},
			expected: Car{
				speed:        18,
				batteryDrain: 10,
				battery:      0,
				distance:     18,
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
		{
			name: "Drive the car with battery drained.",
			car: Car{
				speed:        5,
				batteryDrain: 5,
				battery:      4,
				distance:     0,
			},
			expected: Car{
				speed:        5,
				batteryDrain: 5,
				battery:      4,
				distance:     0,
			},
		},
		{
			name: "Drive the car with positive initial distance and battery drained",
			car: Car{
				speed:        5,
				batteryDrain: 3,
				battery:      2,
				distance:     1,
			},
			expected: Car{
				speed:        5,
				batteryDrain: 3,
				battery:      2,
				distance:     1,
			},
		},
		{
			name: "Drive the car with positive initial distance and full battery",
			car: Car{
				speed:        5,
				batteryDrain: 3,
				battery:      100,
				distance:     1,
			},
			expected: Car{
				speed:        5,
				batteryDrain: 3,
				battery:      97,
				distance:     6,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Drive(tc.car)

			if got != tc.expected {
				t.Errorf("Drive(%+v) = %+v; expected %+v", tc.car, got, tc.expected)
			}
		})
	}
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name     string
		car      Car
		track    Track
		expected bool
	}{
		{
			name: "Car can easily finish the track distance.",
			car: Car{
				speed:        5,
				batteryDrain: 2,
				battery:      100,
				distance:     20,
			},
			track: Track{
				distance: 30,
			},
			expected: true,
		},
		{
			name: "Car can finish the track distance just as the battery is drained.",
			car: Car{
				speed:        5,
				batteryDrain: 5,
				battery:      100,
				distance:     20,
			},
			track: Track{
				distance: 100,
			},
			expected: true,
		},
		{
			name: "Car cannot finish the track distance.",
			car: Car{
				speed:        2,
				batteryDrain: 6,
				battery:      100,
				distance:     5,
			},
			track: Track{
				distance: 80,
			},
			expected: false,
		},
		{
			name: "Car can finish the track distance with initial battery less than 100%",
			car: Car{
				speed:        2,
				batteryDrain: 3,
				battery:      25,
				distance:     0,
			},
			track: Track{
				distance: 16,
			},
			expected: true,
		},
		{
			name: "Car cannot finish the track distance with initial battery less than 100%",
			car: Car{
				speed:        2,
				batteryDrain: 3,
				battery:      25,
				distance:     0,
			},
			track: Track{
				distance: 24,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CanFinish(tc.car, tc.track)

			if got != tc.expected {
				t.Errorf("CanFinish(%#v, %#v) = %+v; expected %+v", tc.car, tc.track, got, tc.expected)
			}
		})
	}
}
