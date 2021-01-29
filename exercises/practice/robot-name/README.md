# Robot Name

Manage robot factory settings.

When robots come off the factory floor, they have no name.

The first time you boot them up, a random name is generated in the format
of two uppercase letters followed by three digits, such as RX837 or BC811.

Every once in a while we need to reset a robot to its factory settings,
which means that their name gets wiped. The next time you ask, it will
respond with a new random name.

The names must be random: they should not follow a predictable sequence.
Random names means a risk of collisions. Your solution must ensure that
every existing robot has a unique name.

### Bonus exercise

Once you get `go test` passing, try `go test -tags bonus`.  This uses a *build
tag* to enable a test that wasn't previously enabled. Build tags control which
files should be included in the package. You can read more about those at [the
Go documentation](https://golang.org/pkg/go/build/#hdr-Build_Constraints).

For this exercise it limits `go test` to only build the tests in the
`robot_name_test.go` file. We can then include the bonus test in the
`bonus_test.go` file with the tags flag as described above.

To get the bonus test to pass you'll have to ensure a robot's name has not been
used before.


## Coding the solution

Look for a stub file having the name robot_name.go
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

## Source

A debugging session with Paul Blackwell at gSchool. [http://gschool.it](http://gschool.it)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
