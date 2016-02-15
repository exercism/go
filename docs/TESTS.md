Go exercises within your exercism project directory can be run by changing to the exercise directory, and running `go test`.

```bash
$ cd exercism/project/directory/go/leap
$ go test
```

### Running benchmarks

Most exercises contain benchmarks, that you can use to determine how changes to your solution affect its performance. To run the benchmarks for an exercise use the command `go test -bench .` inside the exercise directory.

### The `testVersion` constant

Some exercises require you to define a testVersion constant in your package.
There are sometimes problems where tests change and solutions which worked for
the old version no longer work. The testVersion system makes it obvious when
this happens.

To avoid compilation warnings due to testVersion not being defined
just add `const testVersion = VERSION` (where VERSION should be replaced by the
value of `targetTestVersion` in the test file) to your submission file.

Should the tests change in a way that old submissions will or could no longer
work the test writer should increase the `targetTestVersion` value which will cause
the tests to fail with a clear message when an old submission is run with them.

## Go fmt

Please run [`go fmt`](http://blog.golang.org/go-fmt-your-code) on your code before submitting it.

## Hints and tips

### Coding the exercise

The README.md for each exercise gives a general description but the Go test program will be very specific.  Open the test program and give it a quick look - if it seems like cheating, do it anyway.  Look for helpful comments, test data, and just the names of the test functions.  Try running `go test` before you have written anything and see if the error messages give you an idea of where to start.

Your first goal it to get something to compile, even though it fails tests.  For this, you should "stub" functions.  That means leave the body empty, except for whatever it must return.  Write `return 0` or `return ""` or whatever just to get it to compile.  Sometimes to figure out function signatures you will have to go back to the test program and read in more detail.  Once you have figured out all the required function signatures, the test program will complain that 0 is the wrong answer.  Now start filling in function bodies.

### Fixing

Tests will often stop at the first error.  Fix that one, repeat until no errors.  `go test` has useful options.  `go test -v` will give more verbose output. `go test -run <regexp>` will run just tests with names that match <regexp>.

Panics produce lots of output.  It might seem scary at first, but be brave.  Look at the very top of the panic output for the panic message.  That should give a general description of the problem.  Then read through the rest of it looking for the first reference to your code.  That will often point right to the problem.

To find problems with concurrent programs, try `go test -race`.

### Cleaning up your code

Let's say that again, run `go fmt`.
[Golint](https://github.com/golang/lint) can inform you about style issues in your code that `go fmt` won't catch.
You might also look at `go vet`, which warns about potential problems in your code.
Read and heed [frequent comments from code reviews](https://code.google.com/p/go-wiki/wiki/CodeReviewComments).
For an exercise that involves concurrency, try `go test -race` even if passes its tests.
Keep readability in mind even if you optimize for benchmark performance.  Sometimes optimization can help you discover solutions that are concise and readable, but there can be a temptation to abandon readability in the pursuit of performance.

### Bonus tests

A few exercises have bonus tests that do not run with just `go test`.  If you are solving the bonus exercise in addition to the regular exercise, you will have to type `go test -tags bonus`, or some similar command as given in the test program documentation.
