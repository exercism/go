package erratum

import (
	"errors"
	"testing"
)

// Define a function `Use(o ResourceOpener, input string) error` that
// opens a resource, calls Frob(input) and closes the resource
// (in all cases). The function should properly handle errors,
// as defined by the expectations of this test suite.
//
// Also define a testVersion with a value that matches
// the targetTestVersion here.

const targetTestVersion = 2

// Little helper to let us customize behaviour of the resource on a per-test
// basis.
type mockResource struct {
	close  func() error
	frob   func(string)
	defrob func(string)
}

func (mr mockResource) Close() error      { return mr.close() }
func (mr mockResource) Frob(input string) { mr.frob(input) }
func (mr mockResource) Defrob(tag string) { mr.defrob(tag) }

// If this test fails and you've properly defined testVersion the requirements
// of the tests have changed since you wrote your submission.
func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

// Use should not return an error on the "happy" path.
func TestNoErrors(t *testing.T) {
	var frobInput string
	var closeCalled bool
	mr := mockResource{
		close: func() error { closeCalled = true; return nil },
		frob:  func(input string) { frobInput = input },
	}
	opener := func() (Resource, error) { return mr, nil }
	inp := "hello"
	err := Use(opener, inp)
	if err != nil {
		t.Fatalf("Unexpected error from Use: %v", err)
	}
	if frobInput != inp {
		t.Fatalf("Wrong string passed to Frob: got %v, expected %v", frobInput, inp)
	}
	if !closeCalled {
		t.Fatalf("Close was not called")
	}
}

// Use should keep trying if a transient error is returned on open.
func TestKeepTryOpenOnTransient(t *testing.T) {
	var frobInput string
	mr := mockResource{
		close: func() error { return nil },
		frob:  func(input string) { frobInput = input },
	}
	nthCall := 0
	opener := func() (Resource, error) {
		if nthCall < 3 {
			nthCall++
			return mockResource{}, TransientError{errors.New("some error")}
		} else {
			return mr, nil
		}
	}
	inp := "hello"
	err := Use(opener, inp)
	if err != nil {
		t.Fatalf("Unexpected error from Use: %v", err)
	}
	if frobInput != inp {
		t.Fatalf("Wrong string passed to Frob: got %v, expected %v", frobInput, inp)
	}
}

// Use should fail if a non-transient error is returned on open.
func TestFailOpenOnNonTransient(t *testing.T) {
	nthCall := 0
	opener := func() (Resource, error) {
		if nthCall < 3 {
			nthCall++
			return mockResource{}, TransientError{errors.New("some error")}
		} else {
			return nil, errors.New("too awesome")
		}
	}
	inp := "hello"
	err := Use(opener, inp)
	if err == nil {
		t.Fatalf("Unexpected lack of error from Use")
	}
	if err.Error() != "too awesome" {
		t.Fatalf("Invalid error returned from Use")
	}
}

// Use should call Defrob and Close on FrobError panic from Frob
// and return the error.
func TestCallDefrobAndCloseOnFrobError(t *testing.T) {
	tag := "moo"
	var closeCalled bool
	var defrobTag string
	mr := mockResource{
		close: func() error { closeCalled = true; return nil },
		frob:  func(input string) { panic(FrobError{tag, errors.New("meh")}) },
		defrob: func(tag string) {
			if closeCalled {
				t.Fatalf("Close was called before Defrob")
			}
			defrobTag = tag
		},
	}
	opener := func() (Resource, error) { return mr, nil }
	inp := "hello"
	err := Use(opener, inp)
	if err == nil {
		t.Fatalf("Unexpected lack of error from Use")
	}
	if err.Error() != "meh" {
		t.Fatalf("Invalid error returned from Use")
	}
	if defrobTag != tag {
		t.Fatalf("Wrong string passed to Defrob: got %v, expected %v", defrobTag, tag)
	}
	if !closeCalled {
		t.Fatalf("Close was not called")
	}
}

// Use should call Close but not Defrob on non-FrobError panic from Frob
// and return the error.
func TestCallCloseNonOnFrobError(t *testing.T) {
	var closeCalled bool
	var defrobCalled bool
	mr := mockResource{
		close:  func() error { closeCalled = true; return nil },
		frob:   func(input string) { panic(errors.New("meh")) },
		defrob: func(tag string) { defrobCalled = true },
	}
	opener := func() (Resource, error) { return mr, nil }
	inp := "hello"
	err := Use(opener, inp)
	if err == nil {
		t.Fatalf("Unexpected lack of error from Use")
	}
	if err.Error() != "meh" {
		t.Fatalf("Invalid error returned from Use")
	}
	if defrobCalled {
		t.Fatalf("Defrob was called")
	}
	if !closeCalled {
		t.Fatalf("Close was not called")
	}
}
