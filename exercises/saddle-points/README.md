# Saddle Points

Detect saddle points in a matrix.

So say you have a matrix like so:

```text
    0  1  2
  |---------
0 | 9  8  7
1 | 5  3  2     <--- saddle point at (1,0)
2 | 6  6  7
```

It has a saddle point at (1, 0).

It's called a "saddle point" because it is greater than or equal to
every element in its row and less than or equal to every element in
its column.

A matrix may have zero or more saddle points.

Your code should be able to provide the (possibly empty) list of all the
saddle points for any given matrix.

Note that you may find other definitions of matrix saddle points online,
but the tests for this exercise follow the above unambiguous definition.

## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/resources).

## Source

J Dalbey's Programming Practice problems [http://users.csc.calpoly.edu/~jdalbey/103/Projects/ProgrammingPractice.html](http://users.csc.calpoly.edu/~jdalbey/103/Projects/ProgrammingPractice.html)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
