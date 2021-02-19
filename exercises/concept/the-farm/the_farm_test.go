package thefarm

import (
	"errors"
	"fmt"
	"testing"
)

// Valid cases only here
func TestDivideFood(t *testing.T) {
	for _, test := range []struct {
		fodder WeightFodder
		cows   int
		out    float64
		err    bool
	}{
		{func() (float64, error) { return 100, nil }, 10, 10, false},
		{func() (float64, error) { return 10, nil }, 10, 1, false},
		{func() (float64, error) { return 10.5, nil }, 2, 5.25, false},
		{func() (float64, error) { return 5, nil }, 2, 2.5, false},
		{func() (float64, error) { return 0, nil }, 2, 0, false},
		{func() (float64, error) { return -1, nil }, 2, 0, true},
	} {
		if out, err := DivideFood(test.fodder, test.cows); out != test.out ||
			(err != nil) != test.err {
			expectedErr := "nil"
			if test.err {
				expectedErr = "<error>"
			}
			t.Errorf(
				"DivideFood(%v, %v) returned (%v, %v) while (%v, %v) was expected",
				test.fodder,
				test.cows,
				out,
				err,
				test.out,
				expectedErr,
			)
		}
	}
}

// Errors should be returned
func TestDivideFoodErrors(t *testing.T) {
	val1, err1 := DivideFood(func() (float64, error) { return 100, nil }, 0)
	val2, err2 := DivideFood(func() (float64, error) { return 100, nil }, -10)
	// 1
	if err1 == nil {
		t.Errorf("DivideFood(100, 0) should return an error.")
	}
	if val1 != 0 || val2 != 0 {
		t.Error("Returned value, in case of error, must be the default for the type: 0")
	}
	// 2
	if err2 == nil {
		t.Errorf("DivideFood(100, -10) should return an error.")
	}
	if fmt.Sprint(err1) == fmt.Sprint(err2) {
		t.Errorf("DivideFood() should return different errors for division by zero and negative cows.")
	}
}

// Checking errors from weight reader
func TestDivideFoodWeightErrors(t *testing.T) {
	errMsg := "Generic error"
	value, err := DivideFood(func() (float64, error) {
		return 100, errors.New(errMsg)
	}, 10)

	if err == nil || fmt.Sprint(err) != errMsg {
		t.Errorf("Expected the error from weightFodder func to be returned")
	}

	if value != 0 {
		t.Error("Returned value, in case of error, must be the default for the type: 0")
	}
}
