package thefarm

import (
	"errors"
	"fmt"
	"testing"
)

func TestDivideFood(t *testing.T) {
	tests := []struct {
		description string
		fodder      WeightFodder
		cows        int
		wantAmount  float64
		wantErr     error
	}{
		{
			description:  "100 fodder for 10 cows",
			weightFodder: func() (float64, error) { return 100, nil },
			cows:         10,
			wantAmount:   10,
			wantErr:      nil,
		},
		{
			description:  "10 fodder for 10 cows",
			weightFodder: func() (float64, error) { return 10, nil },
			cows:         10,
			wantAmount:   1,
			wantErr:      nil,
		},
		{
			description:  "10.5 fodder for 2 cows",
			weightFodder: func() (float64, error) { return 10.5, nil },
			cows:         2,
			wantAmount:   5.25,
			wantErr:      nil,
		},
		{
			description:  "10 fodder for 10 cows",
			weightFodder: func() (float64, error) { return 5, nil },
			cows:         2,
			wantAmount:   2.5,
			wantErr:      nil,
		},
		{
			description:  "0 fodder for 2 cows",
			weightFodder: func() (float64, error) { return 0, nil },
			cows:         2,
			wantAmount:   0,
			wantErr:      nil,
		},
		{
			description:  "Generic scale error",
			weightFodder: func() (float64, error) { return 10, errors.New("scale error") },
			cows:         2,
			wantAmount:   0,
			wantErr:      errors.New("scale error"),
		},
		{
			description:  "Scale returns 10 with ErrWeight for 2 cows",
			weightFodder: func() (float64, error) { return 10, ErrWeight },
			cows:         2,
			wantAmount:   10,
			wantErr:      nil,
		},
		{
			description:  "Scale returns 5 with ErrWeight for 10 cows",
			weightFodder: func() (float64, error) { return 5, ErrWeight },
			cows:         10,
			wantAmount:   1,
			wantErr:      nil,
		},
		{
			description:  "Negative fodder",
			weightFodder: func() (float64, error) { return -1, nil },
			cows:         2,
			wantAmount:   1,
			wantErr:      errors.New("Negative fodder"),
		},
		{
			description:  "Zero cows",
			weightFodder: func() (float64, error) { return 100, nil },
			cows:         0,
			wantAmount:   0,
			wantErr:      errors.New("Division by zero"),
		},
		{
			description:  "Negative cows",
			weightFodder: func() (float64, error) { return 100, nil },
			cows:         -10,
			wantAmount:   0,
			wantErr:      SillyNephew,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			gotAmount, gotErr := DivideFood(test.fodder, test.cows)
			if gotAmount != test.wantAmount || gotErr != test.wantErr {
				t.Errorf(
					"DivideFood(%v, %v) got (%v, %v) wanted (%v, %v)",
					test.fodder,
					test.cows,
					gotAmount,
					gotErr,
					test.wantAmount,
					test.wantErr,
				)
			}
		})
	}
}
