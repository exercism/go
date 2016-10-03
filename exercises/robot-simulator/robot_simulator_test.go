// +build step1 !step2,!step3

package robot

// Tests are separated into 3 steps.
//
// Run all tests with `go test` or run specific tests with the -tags option.
// Examples,
//
//    go test                      # run all tests
//    go test -tags step1          # run just step 1 tests.
//    go test -tags 'step1 step2'  # run step1 and step2 tests
//
// This source file contains step 1 tests only.  For other tests see
// robot_simulator_step2_test.go and robot_simulator_step3_test.go.
//
// You are given the source file defs.go which defines a number of things
// the test program requires.  It is organized into three sections by step.
//
// To complete step 1 you will define Right, Left, Advance, N, S, E, W,
// and Dir.String.  Complete step 1 before moving on to step 2.

import (
	"runtime"
	"testing"
)

const targetTestVersion = 2

func TestStep1(t *testing.T) {

	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	want := func(x, y int, dir Dir) {
		_, _, line, _ := runtime.Caller(1)
		if X != x || Y != y {
			t.Fatalf("(from line %d) robot at = %d, %d.  Want %d, %d.",
				line, X, Y, x, y)
		}
		if Facing != dir {
			t.Fatalf("(from line %d) robot facing %v, want %v.",
				line, Facing, dir)
		}
	}
	want(0, 0, N)

	Advance()
	want(0, 1, N)

	Right()
	want(0, 1, E)

	Advance()
	want(1, 1, E)

	Left()
	want(1, 1, N)

	Left()
	Left()
	Advance()
	want(1, 0, S)

	Right()
	Advance()
	want(0, 0, W)
}
