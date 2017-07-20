package erratum

import (
	"errors"
	"testing"
)

// Because this exercise is generally unique to each language and how it
// handles errors, most of the definition of your expected solution is provided
// here instead of the README.
// You should read this carefully (more than once) before implementation.

// Define a function `Use(o ResourceOpener, input string) error` that opens a
// resource, calls Frob(input) and closes the resource (in all cases). Your
// function should properly handle errors, as defined by the expectations of
// this test suite. ResourceOpener will be a function you may invoke directly
// `o()` in an attempt to "open" the resource. It returns a Resource and error
// value in the idiomatic Go fashion:
// https://blog.golang.org/error-handling-and-go
//
// See the ./common.go file for the definitions of Resource, ResourceOpener,
// FrobError and TransientError.
//
// There will be a few places in your Use function where errors may occur:
//
// - Invoking the ResourceOpener function passed into Use as the first
//   parameter, it may fail with a TransientError, if so keep trying to open it.
//   If it is some other sort of error, return it.
//
// - Calling the Frob function on the Resource returned from the ResourceOpener
//   function, it may panic with a FrobError (or another type of error). If
//   it is indeed a FrobError you will have to call the Resource's Defrob
//   function using the FrobError's defrobTag variable as input. Either way
//   return the error.
//
// Also note: if the Resource was opened successfully make sure to call its
// Close function no matter what (even if errors occur).
//
// If you are new to Go errors or panics here is a good place to start:
// https://blog.golang.org/defer-panic-and-recover
//
// You may also need to look at named return values as a helpful way to
// return error information from panic recovery:
// https://tour.golang.org/basics/7

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
		}
		return mr, nil
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
		}
		return nil, errors.New("too awesome")
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
