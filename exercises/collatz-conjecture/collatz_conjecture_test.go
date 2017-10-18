package collatz

import (
	"errors"
	"fmt"
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestSteps(t *testing.T) {
	for _, test := range testCases {
		observed := Steps(test.num)
		if observed != test.expected {
			t.Fatalf("Steps(%d) = %d, want %d (%s)",
				test.num, observed, test.expected, test.description)
		}
	}
}

func TestExit(t *testing.T) {
	for _, test := range testThrowErrorCases {
		err := testExit(1, test.num, Steps)
		if err != nil {
			t.Fatalf("%s", test.description)
		}
	}
}

type ExitError int

func (e ExitError) Error() string {
	return fmt.Sprintf("exited with code %d", int(e))
}

func init() {
	exit = func(n int) {
		panic(ExitError(n))
	}
}

func testExit(code, elem int, f func(int) int) (err error) {
	defer func() {

		e := recover()

		switch t := e.(type) {
		case ExitError:

			if int(t) == code {
				err = nil
			} else {
				err = fmt.Errorf("expected exit with %v but %v", code, e)
			}
		default:

			err = fmt.Errorf("expected exit with %v but %v", code, e)
		}
	}()

	f(elem)
	return errors.New("expected exited but not")
}
