# Implementation

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
