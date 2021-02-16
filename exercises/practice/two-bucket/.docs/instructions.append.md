# Implementation

In package twobucket, implement a Go func, Solve, with
the following signature:

```
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error)
```
Solve returns four values: the resulting goal bucket("one" or two"),
the number of moves/steps to achieve the goal amount,
the liters left in the other bucket, and an error value.
The returned error value should be nil when the parameters are valid.
Return an error for any invalid parameter.
Solve should also return an error when a solution cannot be found,
but this error relates only to the bonus exercise below, so you may
ignore that error case for your initial solution.

## Bonus exercise

Once you get `go test` passing, try `go test -tags bonus`.  This uses a *build
tag* to enable tests that were not previously enabled. Build tags control which
files should be included in the package. You can read more about those at [the
Go documentation](https://golang.org/pkg/go/build/#hdr-Build_Constraints).

The exercise limits `go test` to only build the tests referenced in the
`two_bucket_test.go` file. The bonus test cases are found in the file
`bonus_test.go`. Enable those tests as described above.

To get the bonus tests to pass, the Solve func must detect when
a solution cannot be found and return an error.
A solution cannot be found when input test case bucket sizes
are not ones which allow the three operations to succeed in creating the goal amount,
which occurs when the two bucket sizes are not relatively prime to one another.
