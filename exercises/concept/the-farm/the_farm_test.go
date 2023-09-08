package thefarm

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

const precision = 1e-5

var errDeterminingAmount = errors.New("amount could not be determined")
var errDeterminingFactor = errors.New("factor could not be determined")

type testFodderCalculator struct {
	amount    float64
	amountErr error
	factor    float64
	factorErr error
}

func (fc testFodderCalculator) FodderAmount(int) (float64, error) {
	return fc.amount, fc.amountErr
}

func (fc testFodderCalculator) FatteningFactor() (float64, error) {
	return fc.factor, fc.factorErr
}

// testRunnerTaskID=1
func TestDivideFood(t *testing.T) {
	tests := []struct {
		name             string
		fodderCalculator FodderCalculator
		cows             int
		wantAmount       float64
		wantErr          error
	}{
		{
			name: "success, simple inputs",
			fodderCalculator: testFodderCalculator{
				amount: 100,
				factor: 1,
			},
			cows:       10,
			wantAmount: 10,
			wantErr:    nil,
		},
		{
			name: "success, decimal inputs",
			fodderCalculator: testFodderCalculator{
				amount: 60.5,
				factor: 1.3,
			},
			cows:       5,
			wantAmount: 15.73,
			wantErr:    nil,
		},
		{
			name: "error when retrieving fodder amount",
			fodderCalculator: testFodderCalculator{
				amountErr: errDeterminingAmount,
			},
			cows:       10,
			wantAmount: 0,
			wantErr:    errDeterminingAmount,
		},
		{
			name: "error when retrieving fattening factor",
			fodderCalculator: testFodderCalculator{
				factorErr: errDeterminingFactor,
			},
			cows:       10,
			wantAmount: 0,
			wantErr:    errDeterminingFactor,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAmount, gotErr := DivideFood(tt.fodderCalculator, tt.cows)

			if gotErr != tt.wantErr {
				msg := "expected error %q but got %q\n" +
					"(if expected and actual look the same that means the message matches but the errors are not identical)"
				t.Fatalf(msg, tt.wantErr, gotErr)
			}

			if math.Abs(gotAmount-tt.wantAmount) > precision {
				t.Fatalf("expected amount %v but got %v", tt.wantAmount, gotAmount)
			}
		})
	}
}

// testRunnerTaskID=2
func TestValidateInputAndDivideFood(t *testing.T) {
	tests := []struct {
		name             string
		fodderCalculator FodderCalculator
		cows             int
		wantAmount       float64
		wantErr          error
	}{
		{
			name: "negative cows are invalid",
			fodderCalculator: testFodderCalculator{
				amount: 10,
				factor: 1,
			},
			cows:       -10,
			wantAmount: 0,
			wantErr:    errors.New("invalid number of cows"),
		},
		{
			name: "zero cows are invalid",
			fodderCalculator: testFodderCalculator{
				amount: 10,
				factor: 1,
			},
			cows:       0,
			wantAmount: 0,
			wantErr:    errors.New("invalid number of cows"),
		},
		{
			name: "success, simple inputs",
			fodderCalculator: testFodderCalculator{
				amount: 100,
				factor: 1,
			},
			cows:       10,
			wantAmount: 10,
			wantErr:    nil,
		},
		{
			name: "success, decimal inputs",
			fodderCalculator: testFodderCalculator{
				amount: 60.5,
				factor: 1.3,
			},
			cows:       5,
			wantAmount: 15.73,
			wantErr:    nil,
		},
		{
			name: "error when retrieving fodder amount",
			fodderCalculator: testFodderCalculator{
				amountErr: errDeterminingAmount,
			},
			cows:       10,
			wantAmount: 0,
			wantErr:    errDeterminingAmount,
		},
		{
			name: "error when retrieving fattening factor",
			fodderCalculator: testFodderCalculator{
				factorErr: errDeterminingFactor,
			},
			cows:       10,
			wantAmount: 0,
			wantErr:    errDeterminingFactor,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAmount, gotErr := ValidateInputAndDivideFood(tt.fodderCalculator, tt.cows)

			if tt.wantErr != nil && gotErr != nil && tt.wantErr.Error() != gotErr.Error() {
				t.Fatalf("expected error %q but got %q", tt.wantErr, gotErr)
			}

			if tt.wantErr == nil && gotErr != nil {
				t.Fatalf("expected nil but got error %q", gotErr)
			}

			if tt.wantErr != nil && gotErr == nil {
				t.Fatalf("expected error %q but got nil", tt.wantErr)
			}

			if math.Abs(gotAmount-tt.wantAmount) > precision {
				t.Fatalf("expected amount %v but got %v", tt.wantAmount, gotAmount)
			}
		})
	}
}

// testRunnerTaskID=3
func TestValidateNumberOfCows(t *testing.T) {
	tests := []struct {
		name          string
		cows          int
		errorExpected bool
		wantErrMsg    string
	}{
		{
			name:          "big positive number of cows",
			cows:          80,
			errorExpected: false,
		},
		{
			name:          "small positive number of cows",
			cows:          2,
			errorExpected: false,
		},
		{
			name:          "big negative number of cows",
			cows:          -20,
			errorExpected: true,
			wantErrMsg:    "-20 cows are invalid: there are no negative cows",
		},
		{
			name:          "small negative number of cows",
			cows:          -1,
			errorExpected: true,
			wantErrMsg:    "-1 cows are invalid: there are no negative cows",
		},
		{
			name:          "zero cows",
			cows:          0,
			errorExpected: true,
			wantErrMsg:    "0 cows are invalid: no cows don't need food",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := ValidateNumberOfCows(tt.cows)

			if tt.errorExpected && gotErr == nil {
				t.Fatalf("an error was expected but got nil")
			}

			if tt.errorExpected && tt.wantErrMsg != gotErr.Error() {
				t.Fatalf("want error %q but got %q", tt.wantErrMsg, gotErr)
			}

			if !tt.errorExpected && gotErr != nil {
				t.Fatalf("expected nil but got %q", gotErr)
			}
		})
	}
}

// testRunnerTaskID=3
func TestValidateNumberOfCows_PointerReturned(t *testing.T) {
	gotErr := ValidateNumberOfCows(-10)

	if reflect.ValueOf(gotErr).Kind() != reflect.Ptr {
		t.Fatalf("expected pointer but got %v", reflect.ValueOf(gotErr).Kind())
	}
}
