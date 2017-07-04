package erratum

import "io"

// These are the support types and interface definitions used in the
// implementation if your Use function. See the test suite file at
// ./error_handling_test.go for information on the expected implementation.
//
// Because this are part of the package "erratum" if your solution file
// is also declared in the package you will automatically have access to
// these definitions (you do not have to re-declare them).

// TransientError is an error that may occur while opening a resource via
// ResourceOpener.
type TransientError struct {
	err error
}

func (e TransientError) Error() string {
	return e.err.Error()
}

// FrobError is a possible error from doing some frobbing, your implementation
// will require calling your Resource's Defrob(string) method.
// When this error occurs, the FrobError's defrobTag string will contain the
// string you must pass into Defrob.
type FrobError struct {
	defrobTag string
	inner     error
}

func (e FrobError) Error() string {
	return e.inner.Error()
}

type Resource interface {

	// Resource is using composition to inherit the requirements of the io.Closer
	// interface. What this means is that a Resource will have a .Close() method.
	io.Closer

	// Frob does something with the input string.
	// Because this is an incredibly badly designed system if there is an error
	// it panics.
	//
	// The paniced error may be a FrobError in which case Defrob should be called
	// with the defrobTag string.
	Frob(string)

	Defrob(string)
}

// ResourceOpener is a function that creates a resource.
//
// It may return a wrapped error of type TransientError. In this case the resource
// is temporarily unavailable and the caller should retry soon.
type ResourceOpener func() (Resource, error)
