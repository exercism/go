# Implementation

In package twobucket, implement a Go func, Solve, with the following signature:

```go
func Solve(sizeBucketOne,
	sizeBucketTwo, goalAmount int, startBucket string,
) (goalBucket string, numSteps, otherBucketLevel int, e error)
```

Solve returns four values: the resulting goal bucket("one" or two"), the number of moves/steps to achieve the goal amount, the liters left in the other bucket, and an error value.
Return an error for any invalid parameter or no solution is possible.
