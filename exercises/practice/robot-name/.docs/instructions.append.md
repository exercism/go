# Instructions append

## Bonus exercise

Once you get `go test` passing, try `go test -tags bonus`.  This uses a *build
tag* to enable a test that wasn't previously enabled. Build tags control which
files should be included in the package. You can read more about those at [the
Go documentation](https://golang.org/pkg/go/build/#hdr-Build_Constraints).

For this exercise it limits `go test` to only build the tests in the
`robot_name_test.go` file. We can then include the bonus test in the
`bonus_test.go` file with the tags flag as described above.

To get the bonus test to pass you'll have to ensure a robot's name has not been
used before.
