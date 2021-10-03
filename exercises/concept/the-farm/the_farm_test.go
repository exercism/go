package thefarm

import (
	"errors"
	"testing"
)

type testWeightFodder struct {
	fodder float64
	err error
}

func (wf testWeightFodder) FodderAmount() (float64, error) {
	return wf.fodder, wf.err
}

func TestDivideFood(t *testing.T) {
	nonScaleError := errors.New("non-scale error")
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
			weightFodder:            testWeightFodder{ 100, nil },
			weightFodderDescription: "100 fodder, no error",
			cows:                    10,
			wantAmount:              10,
			wantErr:                 nil,
		},
		{
			description:             "10 fodder for 10 cows",
			weightFodder:            testWeightFodder{ 10, nil },
			weightFodderDescription: "10 fodder, no error",
			cows:                    10,
			wantAmount:              1,
			wantErr:                 nil,
		},
		{
			description:             "10.5 fodder for 2 cows",
			weightFodder:            testWeightFodder{ 10.5, nil },
			weightFodderDescription: "10.5 fodder, no error",
			cows:                    2,
			wantAmount:              5.25,
			wantErr:                 nil,
		},
		{
			description:             "5 fodder for 2 cows",
			weightFodder:            testWeightFodder{ 5, nil },
			weightFodderDescription: "5 fodder, no error",
			cows:                    2,
			wantAmount:              2.5,
			wantErr:                 nil,
		},
		{
			description:             "0 fodder for 2 cows",
			weightFodder:            testWeightFodder{ 0, nil },
			weightFodderDescription: "0 fodder, no error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 nil,
		},
		{
			description:             "Generic error from the scale is returned",
			weightFodder:            testWeightFodder{ 10, nonScaleError },
			weightFodderDescription: "10 fodder, generic error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 nonScaleError,
		},
		{
			description:             "Scale returns 10 with ScaleError for 2 cows",
			weightFodder:            testWeightFodder{ 10, ScaleError{} },
			weightFodderDescription: "10 fodder, ScaleError",
			cows:                    2,
			wantAmount:              10,
			wantErr:                 nil,
		},
		{
			description:             "Scale returns 5 with ScaleError for 10 cows",
			weightFodder:            testWeightFodder{ 5, ScaleError{} },
			weightFodderDescription: "5 fodder, ScaleError",
			cows:                    10,
			wantAmount:              1,
			wantErr:                 nil,
		},
		{
			description:             "Negative fodder",
			weightFodder:            testWeightFodder{ -1, nil },
			weightFodderDescription: "-1 fodder, no error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 errors.New("Negative fodder"),
		},
		{
			description:             "Zero cows",
			weightFodder:            testWeightFodder{ 100, nil },
			weightFodderDescription: "100 fodder, no error",
			cows:                    0,
			wantAmount:              0,
			wantErr:                 errors.New("Division by zero"),
		},
		{
			description:             "Negative ten cows",
			weightFodder:            testWeightFodder{ 100, nil },
			weightFodderDescription: "100 fodder, no error",
			cows:                    -10,
			wantAmount:              0,
			wantErr:                 SillyNephewError{-10},
		},
		{
			description:             "Negative seven cows",
			weightFodder:            testWeightFodder{ 100, nil },
			weightFodderDescription: "100 fodder, no error",
			cows:                    -7,
			wantAmount:              0,
			wantErr:                 SillyNephewError{-7},
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
	var wantSne *SillyNephewError
	var gotSne *SillyNephewError
	if errors.As(want, &wantSne) {
		if errors.As(got, &gotSne) {
			return wantSne.Cows == gotSne.Cows
		}
		return false
	}
	// Got a SillyNephewError that we did not want
	if errors.As(got, &gotSne) {
		return false
	}
	// Otherwise, just comparer the error strings
	return got.Error() == want.Error()
}
