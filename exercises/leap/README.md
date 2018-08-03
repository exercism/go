# Leap

Given a year, report if it is a leap year.

The tricky thing here is that a leap year in the Gregorian calendar occurs:

```text
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
```

For example, 1997 is not a leap year, but 1996 is.  1900 is not a leap
year, but 2000 is.

If your language provides a method in the standard library that does
this look-up, pretend it doesn't exist and implement it yourself.

## Notes

Though our exercise adopts some very simple rules, there is more to
learn!

For a delightful, four minute explanation of the whole leap year
phenomenon, go watch [this youtube video][video].

[video]: http://www.youtube.com/watch?v=xX96xng7sAE

You will see a `cases_test.go` file in this exercise. This holds the test
cases used in the `leap_test.go`. You can mostly ignore this file.

However, if you are interested... we sometimes generate the test data from a
[cross language repository][problem-specifications-leap]. In that repo
exercises may have a [.json file][problem-specifications-leap-json] that
contains common test data. Some of our local exercises have an
[intermediary program][local-leap-gen] that takes the problem specification
JSON and turns in into Go structs that are fed into the `<exercise>_test.go`
file. The Go specific transformation of that data lives in the `cases_test.go` file.

[problem-specifications-leap]: https://github.com/exercism/problem-specifications/tree/master/exercises/leap
[problem-specifications-leap-json]: https://github.com/exercism/problem-specifications/blob/master/exercises/leap/canonical-data.json
[local-leap-gen]: https://github.com/exercism/go/blob/master/exercises/leap/.meta/gen.go


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

JavaRanch Cattle Drive, exercise 3 [http://www.javaranch.com/leap.jsp](http://www.javaranch.com/leap.jsp)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
