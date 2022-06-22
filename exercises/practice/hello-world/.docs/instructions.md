# Instructions

The classical introductory exercise. Just say "Hello, World!".

["Hello, World!"](http://en.wikipedia.org/wiki/%22Hello,_world!%22_program) is
the traditional first program for beginning programming in a new language
or environment.

The objectives are simple:

- Write a function that returns the string "Hello, World!".
- Run the test suite and make sure that it succeeds.
- Submit your solution and check it at the website.

If everything goes well, you will be ready to fetch your first real exercise.

You may find Go's [testable documentation examples](https://go.dev/blog/examples) useful for testing, experimenting an debugging your solution.

Just treat the inside of the example function as a main function and run it with `go test -run Example…` instead of `go run main.go`.

1. Before trying to solve the hello world exercise, create a `hello_world_examples_test.go` with these contents.

```go
package greeting

import (
	"fmt"
)

func ExampleHelloWorld() {
	fmt.Println(HelloWorld())
	// Output:
	// Goodbye, Mars!
}
```

2. Run the example.

```bash
$ go test -v -run ExampleHelloWorld
=== RUN   ExampleHelloWorld
--- PASS: ExampleHelloWorld (0.00s)
PASS
ok      greeting        0.012s
```

3. You can also run all the example tests with the regular tests with `go test -v`. If your solution is incomplete, it's easier to run specific examples instead of all the test.
