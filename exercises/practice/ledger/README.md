# Ledger

Refactor a ledger printer.

The ledger exercise is a refactoring exercise. There is code that prints a
nicely formatted ledger, given a locale (American or Dutch) and a currency (US
dollar or euro). The code however is rather badly written, though (somewhat
surprisingly) it consistently passes the test suite.

Rewrite this code. Remember that in refactoring the trick is to make small steps
that keep the tests passing. That way you can always quickly go back to a
working version.  Version control tools like git can help here as well.

Please keep a log of what changes you've made and make a comment on the exercise
containing that log, this will help reviewers.

## Coding the solution

Look for a stub file having the name ledger.go
and place your solution code in that file.

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

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
