package thefarm

import (
	"errors"
	"testing"
)

func TestDivideFood(t *testing.T) {
	scaleError := errors.New("scale error")
	tests := []struct {
		description             string
		weightFodder            WeightFodder
		weightFodderDescription string
		cows                    int
		wantAmount              float64
		wantErr                 error
	}{
		{
			description:             "100 fodder for 10 cows",
			weightFodder:            func() (float64, error) { return 100, nil },
			weightFodderDescription: "100 fodder, no error",
			cows:                    10,
			wantAmount:              10,
			wantErr:                 nil,
		},
		{
			description:             "10 fodder for 10 cows",
			weightFodder:            func() (float64, error) { return 10, nil },
			weightFodderDescription: "10 fodder, no error",
			cows:                    10,
			wantAmount:              1,
			wantErr:                 nil,
		},
		{
			description:             "10.5 fodder for 2 cows",
			weightFodder:            func() (float64, error) { return 10.5, nil },
			weightFodderDescription: "10.5 fodder, no error",
			cows:                    2,
			wantAmount:              5.25,
			wantErr:                 nil,
		},
		{
			description:             "10 fodder for 10 cows",
			weightFodder:            func() (float64, error) { return 5, nil },
			weightFodderDescription: "5 fodder, no error",
			cows:                    2,
			wantAmount:              2.5,
			wantErr:                 nil,
		},
		{
			description:             "0 fodder for 2 cows",
			weightFodder:            func() (float64, error) { return 0, nil },
			weightFodderDescription: "0 fodder, no error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 nil,
		},
		{
			description:             "Generic scale error",
			weightFodder:            func() (float64, error) { return 10, scaleError },
			weightFodderDescription: "10 fodder, generic error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 scaleError,
		},
		{
			description:             "Scale returns 10 with ErrWeight for 2 cows",
			weightFodder:            func() (float64, error) { return 10, ErrWeight },
			weightFodderDescription: "10 fodder, ErrWeight",
			cows:                    2,
			wantAmount:              10,
			wantErr:                 nil,
		},
		{
			description:             "Scale returns 5 with ErrWeight for 10 cows",
			weightFodder:            func() (float64, error) { return 5, ErrWeight },
			weightFodderDescription: "5 fodder, ErrWeight",
			cows:                    10,
			wantAmount:              1,
			wantErr:                 nil,
		},
		{
			description:             "Negative fodder",
			weightFodder:            func() (float64, error) { return -1, nil },
			weightFodderDescription: "-1 fodder, no error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 errors.New("Negative fodder"),
		},
		{
			description:             "Zero cows",
			weightFodder:            func() (float64, error) { return 100, nil },
			weightFodderDescription: "100 fodder, no error",
			cows:                    0,
			wantAmount:              0,
			wantErr:                 errors.New("Division by zero"),
		},
		{
			description:             "Negative cows",
			weightFodder:            func() (float64, error) { return 100, nil },
			weightFodderDescription: "100 fodder, no error",
			cows:                    -10,
			wantAmount:              0,
			wantErr:                 SillyNephew,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			gotAmount, gotErr := DivideFood(test.weightFodder, test.cows)
			switch {
			case gotAmount != test.wantAmount:
				t.Errorf(
					"DivideFood(weightFodder(%v), %v) got (%v, %v) wanted (%v, %v)",
					test.weightFodderDescription,
					test.cows,
					gotAmount,
					gotErr,
					test.wantAmount,
					test.wantErr,
				)
			case gotErr != nil && test.wantErr == nil:
				t.Errorf(
					"DivideFood(weightFodder(%v), %v) got an unexpected error (%v)",
					test.weightFodderDescription,
					test.cows,
					gotErr,
				)
			case gotErr == nil && test.wantErr != nil:
				t.Errorf(
					"DivideFood(weightFodder(%v), %v) got no error, but wanted an error (%v)",
					test.weightFodderDescription,
					test.cows,
					test.wantErr,
				)
			case !errorsAreEqual(gotErr, test.wantErr):
				t.Errorf(
					"DivideFood(weightFodder(%v), %v) got error (%v), but wanted error (%v)",
					test.weightFodderDescription,
					test.cows,
					gotErr,
					test.wantErr,
				)
			}
		})
	}
}

func errorsAreEqual(got, want error) bool {
	if got == want {
		return true
	}
	// Only accept exact equality for the SillyNephew error
	if errors.Is(want, SillyNephew) {
		return false
	}
	return got.Error() == want.Error()
}
