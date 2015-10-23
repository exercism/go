// Clock stub file

// Build constraint.
// It lets our CI (Continuous Integration) testing test the test program.
// But it's served its purpose by now.  Feel free to delete.
// +build !example

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

// The value of TestVersion here must match `testVersion` in the file
// clock_test.go.
const TestVersion = 2

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock // Complete the type definition.  Pick a suitable data type.

func Time(hour, minute int) Clock {
}

func (Clock) String() string {
}

func (Clock) Add(minutes int) Clock {
}

// This comment and others that aren't of interest become nits if you
// don't delete them...
