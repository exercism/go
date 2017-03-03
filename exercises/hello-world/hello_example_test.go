package greeting

import "fmt"

// ExampleHelloWorld() is an Example function. Examples are testable snippets of
// Go code that are used for documenting, and verifying, the package API.
// They may be present in some exercises to demonstrate the expected use of the
// exercise API and can be run as part of a package's test suite
//
// When an Example test is run the data that is written to standard output is
// compared to the data that comes after the "Output: " comment.
//
// So, below, the result of HelloWorld() is passed to standard output
// using fmt.Println, and this is compared against the expected output; Hello, World!
// If they are equal, the test passes.
func ExampleHelloWorld() {
	h := HelloWorld()
	fmt.Println(h)
	// Output: Hello, World!
}
