# Instructions append

It is possible to run out of robot names - once every possible robot name has been allocated then an error should be returned.

## Bonus test 

Once you get `go test` passing, you will find that there is a test in the test suite named `TestCollisions` which is skipped by default.

To get the test to run, remove this line:
```go
t.Skip("skipping test as it can take a long time to run if solution is sub-optimal.")
```

To get this test to pass you'll have to ensure a robot's name has not been used before and that names can be allocated efficiently.
