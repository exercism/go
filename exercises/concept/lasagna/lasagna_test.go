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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OvenTime; got != tt.expected {
				t.Errorf("OvenTime(%d)  = %d; want %d", tt.expected, got, tt.expected)
			}
		})
	}
}

func TestRemainingOvenTime(t *testing.T) {
	tests := []lasagnaTests{
		{
			name:     "Remaining minutes in oven",
			layers:   0,
			time:     15,
			expected: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemainingOvenTime(tt.time); got != tt.expected {
				t.Errorf("RemainingOvenTime(%d) = %d; want %d", tt.time, got, tt.expected)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreparationTime(tt.layers); got != tt.expected {
				t.Errorf("PreparationTime(%d) = %d; want %d", tt.layers, got, tt.expected)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElapsedTime(tt.layers, tt.time); got != tt.expected {
				t.Errorf("ElapsedTime(%d, %d) = %d; want %d", tt.layers, tt.time, got, tt.expected)
			}
		})

	}
}
