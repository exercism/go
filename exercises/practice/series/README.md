# Series

Given a string of digits, output all the contiguous substrings of length `n` in
that string in the order that they appear.

For example, the string "49142" has the following 3-digit series:

- "491"
- "914"
- "142"

And the following 4-digit series:

- "4914"
- "9142"

And if you ask for a 6-digit series from a 5-digit string, you deserve
whatever you get.

Note that these series are only required to occupy *adjacent positions*
in the input; the digits need not be *numerically consecutive*.

## Implementation

Define two functions: (Two? Yes, sometimes we ask more out of Go.)

All returns a list of all substrings of s with length n.

`All(n int, s string) []string`

UnsafeFirst returns the first substring of s with length n.

`UnsafeFirst(n int, s string) string`

At this point you could consider this exercise complete and move on.

But wait, maybe you ask a reasonable question: Why is the function
called **Unsafe** First?  If you are interested, read on for a bonus
exercise.

### Bonus exercise:

Once you get `go test` passing, try `go test -tags asktoomuch`.  This
uses a *build tag* to enable a test that wasn't enabled before. Build
tags allow for the selection of files to be used by the package. In
our case it will control which test files are used. You can read more
about those at
[the Go documentation](https://golang.org/pkg/go/build/#hdr-Build_Constraints).

You may notice that you can't make this *asktoomuch* test happy. We
need a way to signal that in some cases you can't take the first N
characters of the string. UnsafeFirst can't do that since it only
returns a string.

To fix that, let's add another return value to the function.  Define

`First(int, string) (first string, ok bool)`

and test with `go test -tags first`.

The `ok bool` second return argument is a common and idiomatic pattern
in Go. For example you see it in
[Map lookups](https://blog.golang.org/go-maps-in-action) and
[type assertions](https://tour.golang.org/methods/15).


## Coding the solution

Look for a stub file having the name series.go
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

A subset of the Problem 8 at Project Euler [http://projecteuler.net/problem=8](http://projecteuler.net/problem=8)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
