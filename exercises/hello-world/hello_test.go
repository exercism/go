package greeting

import "testing"

// Define a function named HelloWorld that takes no arguments,
// and returns a string.
// In other words, define a function with the following signature:
// HelloWorld() string
//
// In any exercise solution using the following tests you must also define a
// "testVersion" constant with a value that matches the targetTestVersion
// here. See how it is done in the the ./hello_world.go file provided.
// This is a common convention for ensuring test consistency throughout the
// Exercism Go language track.
//
// See TestTestVersion function below for how this and the testVersion
// constants are used.
const targetTestVersion = 4

// TestTestVersion is used to ensure the exercise solution is considered
// compatible with the rest of the tests. It should be the initial testing
// function in any test suite. If there is an incompatibility the tests can
// be considered unreliable therefore this function will abort the testing.
func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if observed := HelloWorld(); observed != expected {
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
	}
}
