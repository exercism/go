//go:build step1 || (!step2 && !step3)

package robot

// This source file contains step 1 tests only.  For other tests see
// robot_simulator_step2_test.go and robot_simulator_step3_test.go.

import (
	"runtime"
	"testing"
)

func TestStep1(t *testing.T) {

	want := func(x, y int, dir Dir) {
		_, _, line, _ := runtime.Caller(1)
		if Step1Robot.X != x || Step1Robot.Y != y {
			t.Fatalf("(from line %d) robot at = %d, %d.  Want %d, %d.",
				line, Step1Robot.X, Step1Robot.Y, x, y)
		}
		if Step1Robot.Dir != dir {
			t.Fatalf("(from line %d) robot facing %v, want %v.",
				line, Step1Robot.Dir, dir)
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
