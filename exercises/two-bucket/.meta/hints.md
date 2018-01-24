## Implementation

In package twobucket, implement a Go func, Solve, with
the following signature:

```
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int)
```
Solve returns three values: the resulting goal bucket("one" or two"),
the number of moves/steps to achieve the goal amount,
and the liters left in the other bucket.
