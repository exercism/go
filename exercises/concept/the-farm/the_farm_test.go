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

var NonScaleError = errors.New("non-scale error")

func TestDivideFood(t *testing.T) {
	tests := []struct {
		description      string
		fodderAmount     WeightFodder
		cows             int
		wantFodderAmount float64
		wantFodderErr    error
	}{
		{
			description:      "100 fodder for 10 cows with no error",
			fodderAmount:     testWeightFodder{fodder: 100, err: nil},
			cows:             10,
			wantFodderAmount: 10,
			wantFodderErr:    nil,
		},
		{
			description:      "10 fodder for 10 cows with no error",
			fodderAmount:     testWeightFodder{fodder: 10, err: nil},
			cows:             10,
			wantFodderAmount: 1,
			wantFodderErr:    nil,
		},
		{
			description:      "10.5 fodder for 2 cows with no error",
			fodderAmount:     testWeightFodder{fodder: 10.5, err: nil},
			cows:             2,
			wantFodderAmount: 5.25,
			wantFodderErr:    nil,
		},
		{
			description:      "5 fodder for 2 cows with no error",
			fodderAmount:     testWeightFodder{fodder: 5, err: nil},
			cows:             2,
			wantFodderAmount: 2.5,
			wantFodderErr:    nil,
		},
		{
			description:      "0 fodder for 2 cows with no error",
			fodderAmount:     testWeightFodder{fodder: 0, err: nil},
			cows:             2,
			wantFodderAmount: 0,
			wantFodderErr:    nil,
		},
		{
			description:      "10 fodder for 2 cows with generic error",
			fodderAmount:     testWeightFodder{fodder: 10, err: NonScaleError},
			cows:             2,
			wantFodderAmount: 0,
			wantFodderErr:    NonScaleError,
		},
		{
			description:      "Negative fodder for 2 cows with generic error",
			fodderAmount:     testWeightFodder{fodder: -10, err: NonScaleError},
			cows:             2,
			wantFodderAmount: 0,
			wantFodderErr:    NonScaleError,
		},
		{
			description:      "10 fodder for 2 cows with malfunction error",
			fodderAmount:     testWeightFodder{fodder: 10, err: ErrScaleMalfunction},
			cows:             2,
			wantFodderAmount: 10,
			wantFodderErr:    nil,
		},
		{
			description:      "5 fodder for 10 cows with malfunction error",
			fodderAmount:     testWeightFodder{fodder: 5, err: ErrScaleMalfunction},
			cows:             10,
			wantFodderAmount: 1,
			wantFodderErr:    nil,
		},
		{
			description:      "Negative fodder for 2 cows with no error",
			fodderAmount:     testWeightFodder{fodder: -1, err: nil},
			cows:             2,
			wantFodderAmount: 0,
			wantFodderErr:    errors.New("negative fodder"),
		},
		{
			description:      "Negative fodder for 2 cows with malfunction error",
			fodderAmount:     testWeightFodder{fodder: -1, err: ErrScaleMalfunction},
			cows:             2,
			wantFodderAmount: 0,
			wantFodderErr:    errors.New("negative fodder"),
		},
		{
			description:      "100 fodder for zero cows with no error",
			fodderAmount:     testWeightFodder{fodder: 100, err: nil},
			cows:             0,
			wantFodderAmount: 0,
			wantFodderErr:    errors.New("division by zero"),
		},
	}

	for _, test := range tests {
		t.Logf("Testing condition %s", test.description)

		gotFodderAmount, gotErr := DivideFood(test.fodderAmount, test.cows)

		if gotFodderAmount != test.wantFodderAmount && gotErr != test.wantFodderErr {
			t.Errorf("Wrong food division ... \ngot %v and %v, \nwant %v and %v",
				gotFodderAmount, gotErr, test.wantFodderAmount, test.wantFodderErr)
		}
	}
}

func TestDivideFoodSillyNephewError(t *testing.T) {
	tests := []struct {
		description string
		cows        int
		wantErrMsg  error
	}{
		{
			description: "negative ten cows",
			cows:        -10,
			wantErrMsg:  errors.New("silly nephew, there cannot be -10 cows"),
		},
		{
			description: "negative seven cows",
			cows:        -7,
			wantErrMsg:  errors.New("silly nephew, there cannot be -7 cows"),
		},
	}

	for _, test := range tests {
		t.Logf("Testing for %s", test.description)

		fodderWeight := testWeightFodder{fodder: 100, err: nil}
		gotFodderAmount, gotErrMsg := DivideFood(fodderWeight, test.cows)

		if gotFodderAmount != 0 && gotErrMsg != test.wantErrMsg {
			t.Errorf("Nepher error ... \ngot %v and %v, \nwant 0 and %v",
				gotFodderAmount, gotErrMsg, test.wantErrMsg)
		}
	}
}
