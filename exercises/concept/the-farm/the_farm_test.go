package thefarm

import (
	"errors"
	"testing"
)

type testWeightFodder struct {
	fodder float64
	err    error
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
			weightFodder:            testWeightFodder{fodder: 100, err: nil},
			weightFodderDescription: "100 fodder, no error",
			cows:                    10,
			wantAmount:              10,
			wantErr:                 nil,
		},
		{
			description:             "10 fodder for 10 cows",
			weightFodder:            testWeightFodder{fodder: 10, err: nil},
			weightFodderDescription: "10 fodder, no error",
			cows:                    10,
			wantAmount:              1,
			wantErr:                 nil,
		},
		{
			description:             "10.5 fodder for 2 cows",
			weightFodder:            testWeightFodder{fodder: 10.5, err: nil},
			weightFodderDescription: "10.5 fodder, no error",
			cows:                    2,
			wantAmount:              5.25,
			wantErr:                 nil,
		},
		{
			description:             "5 fodder for 2 cows",
			weightFodder:            testWeightFodder{fodder: 5, err: nil},
			weightFodderDescription: "5 fodder, no error",
			cows:                    2,
			wantAmount:              2.5,
			wantErr:                 nil,
		},
		{
			description:             "0 fodder for 2 cows",
			weightFodder:            testWeightFodder{fodder: 0, err: nil},
			weightFodderDescription: "0 fodder, no error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 nil,
		},
		{
			description:             "Generic error from the scale is returned",
			weightFodder:            testWeightFodder{fodder: 10, err: nonScaleError},
			weightFodderDescription: "10 fodder, generic error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 nonScaleError,
		},
		{
			description:             "Negative fodder with generic error from the scale is returned",
			weightFodder:            testWeightFodder{fodder: -10, err: nonScaleError},
			weightFodderDescription: "-10 fodder, generic error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 nonScaleError,
		},
		{
			description:             "Scale returns 10 with ErrScaleMalfunction for 2 cows",
			weightFodder:            testWeightFodder{fodder: 10, err: ErrScaleMalfunction},
			weightFodderDescription: "10 fodder, ErrScaleMalfunction",
			cows:                    2,
			wantAmount:              10,
			wantErr:                 nil,
		},
		{
			description:             "Scale returns 1 with ErrScaleMalfunction for 10 cows",
			weightFodder:            testWeightFodder{fodder: 5, err: ErrScaleMalfunction},
			weightFodderDescription: "5 fodder, ErrScaleMalfunction",
			cows:                    10,
			wantAmount:              1,
			wantErr:                 nil,
		},
		{
			description:             "Negative fodder",
			weightFodder:            testWeightFodder{fodder: -1, err: nil},
			weightFodderDescription: "-1 fodder, no error",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 errors.New("negative fodder"),
		},
		{
			description:             "Negative fodder with ScaleError",
			weightFodder:            testWeightFodder{fodder: -1, err: ErrScaleMalfunction},
			weightFodderDescription: "-1 fodder, ScaleError",
			cows:                    2,
			wantAmount:              0,
			wantErr:                 errors.New("negative fodder"),
		},
		{
			description:             "Zero cows",
			weightFodder:            testWeightFodder{fodder: 100, err: nil},
			weightFodderDescription: "100 fodder, no error",
			cows:                    0,
			wantAmount:              0,
			wantErr:                 errors.New("division by zero"),
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
			case !(gotErr == test.wantErr || gotErr.Error() == test.wantErr.Error()):
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

func TestDivideFoodSillyNephewError(t *testing.T) {
	tests := []struct {
		description string
		cows        int
		wantErrMsg  string
	}{
		{
			description: "Negative ten cows",
			cows:        -10,
			wantErrMsg:  "silly nephew, there cannot be -10 cows",
		},
		{
			description: "Negative seven cows",
			cows:        -7,
			wantErrMsg:  "silly nephew, there cannot be -7 cows",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			weightFodder := testWeightFodder{fodder: 100, err: nil}
			gotAmount, gotErr := DivideFood(weightFodder, test.cows)
			if gotAmount != 0 {
				t.Errorf(
					"DivideFood(weightFodder(%v), %v) got amount %v, but wanted amount 0",
					"100 fodder, no error",
					test.cows,
					gotAmount,
				)
			}
			if gotErr.Error() != test.wantErrMsg {
				t.Errorf(
					"DivideFood(weightFodder(%v), %v) got error msg %q, but wanted error msg %q",
					"100 fodder, no error",
					test.cows,
					gotErr.Error(),
					test.wantErrMsg,
				)
			}
		})
	}
}
