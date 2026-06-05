package lasagna

import "testing"

type lasagnaTests struct {
	name                   string
	layers, time, expected int
}

func TestOvenTime(t *testing.T) {
	tests := []lasagnaTests{
		{
			name:     "Calculates how many minutes the lasagna should be in the oven",
			layers:   0,
			time:     40,
			expected: 40,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := OvenTime; got != tc.expected {
				t.Errorf("OvenTime()  = %d; want %d", got, tc.expected)
			}
		})
	}
}

func TestRemainingOvenTime(t *testing.T) {
	tests := []lasagnaTests{
		{
			name:     "Remaining minutes in oven after 15 min",
			layers:   0,
			time:     15,
			expected: 25,
		},
		{
			name:     "Remaining minutes in oven after 30 min",
			layers:   0,
			time:     30,
			expected: 10,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := RemainingOvenTime(tc.time); got != tc.expected {
				t.Errorf("RemainingOvenTime(%d) = %d; want %d", tc.time, got, tc.expected)
			}
		})
	}
}

func TestPreparationTime(t *testing.T) {
	tests := []lasagnaTests{
		{
			name:     "Preparation time in minutes for one layer",
			layers:   1,
			time:     0,
			expected: 2,
		},
		{
			name:     "Preparation time in minutes for multiple layers",
			layers:   4,
			time:     0,
			expected: 8,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := PreparationTime(tc.layers); got != tc.expected {
				t.Errorf("PreparationTime(%d) = %d; want %d", tc.layers, got, tc.expected)
			}
		})
	}
}

func TestElapsedTime(t *testing.T) {
	tests := []lasagnaTests{
		{
			name:     "Total time in minutes for one layer",
			layers:   1,
			time:     30,
			expected: 32,
		},
		{
			name:     "Total time in minutes for multiple layers",
			layers:   4,
			time:     8,
			expected: 16,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ElapsedTime(tc.layers, tc.time); got != tc.expected {
				t.Errorf("ElapsedTime(%d, %d) = %d; want %d", tc.layers, tc.time, got, tc.expected)
			}
		})
	}
}
