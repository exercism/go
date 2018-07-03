# Two Bucket

Given two buckets of different size, demonstrate how to measure an exact number of liters by strategically transferring liters of fluid between the buckets.

Since this mathematical problem is fairly subject to interpretation / individual approach, the tests have been written specifically to expect one overarching solution.

To help, the tests provide you with which bucket to fill first. That means, when starting with the larger bucket full, you are NOT allowed at any point to have the smaller bucket full and the larger bucket empty (aka, the opposite starting point); that would defeat the purpose of comparing both approaches!

Your program will take as input:
- the size of bucket one
- the size of bucket two
- the desired number of liters to reach
- which bucket to fill first, either bucket one or bucket two

Your program should determine:
- the total number of "moves" it should take to reach the desired number of liters, including the first fill
- which bucket should end up with the desired number of liters (let's say this is bucket A) - either bucket one or bucket two
- how many liters are left in the other bucket (bucket B)

Note: any time a change is made to either or both buckets counts as one (1) move.

Example:
Bucket one can hold up to 7 liters, and bucket two can hold up to 11 liters. Let's say bucket one, at a given step, is holding 7 liters, and bucket two is holding 8 liters (7,8). If you empty bucket one and make no change to bucket two, leaving you with 0 liters and 8 liters respectively (0,8), that counts as one "move". Instead, if you had poured from bucket one into bucket two until bucket two was full, leaving you with 4 liters in bucket one and 11 liters in bucket two (4,11), that would count as only one "move" as well.

To conclude, the only valid moves are:
- pouring from one bucket to another
- emptying one bucket and doing nothing to the other
- filling one bucket and doing nothing to the other

Written with <3 at [Fullstack Academy](http://www.fullstackacademy.com/) by Lindsay Levine.

## Implementation

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

### Bonus exercise

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


## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/about).

## Source

Water Pouring Problem [http://demonstrations.wolfram.com/WaterPouringProblem/](http://demonstrations.wolfram.com/WaterPouringProblem/)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
