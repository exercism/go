# Instructions append

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
