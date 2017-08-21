package greeting

import "testing"

// Define a function named HelloWorld that takes no arguments,
// and returns a string.
// In other words, define a function with the following signature:
// HelloWorld() string

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
