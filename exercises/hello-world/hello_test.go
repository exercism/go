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

// BenchmarkHelloWorld() is a benchmarking function. These functions follow the
// form `func BenchmarkXxx(*testing.B)` and can be used to test the performance
// of your implementation. They may not be present in every exercise, but when
// they are you can run them by including the `-bench` flag with the `go test`
// command, like so: `go test -bench .`
//
// You will see output similar to the following:
//
// BenchmarkHelloWorld   	2000000000	         0.46 ns/op
//
// This means that the loop ran 2000000000 times at a speed of 0.46 ns per loop.
//
// While benchmarking can be useful to compare different iterations of the same
// exercise, keep in mind that others will run the same benchmarks on different
// machines, with different specs, so the results from these benchmark tests may
// vary.
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld()
	}
}
