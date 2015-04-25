## Running tests

Go exercises within your exercism project directory can be run by changing to the exercise directory, and running `go test`.

```bash
$ cd exercism/project/directory/go/leap
$ go test
```

### Running benchmarks

Most exercises contain benchmarks, that you can use to determine how changes to your solution affect its performance. To run the benchmarks for an exercise use the command `go test -bench .` inside the exercise directory.

### The `TestVersion` constant

Some exercises require you to define a TestVersion constant in your package.
There are sometimes problems where tests change and solutions which worked for
the old version no longer work. The TestVersion system makes it obvious when
this happens.

To avoid compilation warnings due to TestVersion not being defined
just add `const TestVersion = VERSION` (where VERSION should be replaced by the
value of `testVersion` in the test file) to your submission file.

Should the tests change in a way that old submissions will or could no longer
work the test writer should increase the `testVersion` value which will cause
the tests to fail with a clear message when an old submission is run with them.

## Go fmt

Please run [`go fmt`](http://blog.golang.org/go-fmt-your-code) on your code before submitting it.

