# xGo

[![Build Status](https://travis-ci.org/exercism/xgo.png?branch=master)](https://travis-ci.org/exercism/xgo)

Exercism exercises in Go

## Issues

We welcome issues filed at https://github.com/exercism/xgo/issues for problems of any size.  Feel free to report
typographical errors or poor wording.  We are most interested in improving the quality of the test suites.
You can greatly help us improve the quality of the exercises by filing reports of invalid solutions that
pass tests or of valid solutions that fail tests.

## Development setup

Beyond filing issues, if you would like to contribute directly to the Go code in xgo, you should follow some
standard Go development practices.  You should have a [recent version of Go](http://golang.org/doc/install)
installed, ideally either the current release, the previous release, or tip.

You will need a github account and you will need to fork exercism/xgo to your account.
See [GitHub Help](https://help.github.com/articles/fork-a-repo/) if you are unfamiliar with the process.
Clone your fork into a Go workspace on your GOPATH.  Read [How to Write Go Code](http://golang.org/doc/code.html)
if you are unfamiliar with Go workspaces and GOPATH.  As an example, if you have `$HOME/dev` on your GOPATH,
you should cd to `$HOME/dev/src/github.com/<you>` to type the clone command:
`git clone https://github.com/<you>/xgo`.  Test your clone by cding to the xgo directory and typing
`go test ./...`. You should see tests pass for all exercises. (Oh wait, do you see the error message from paasio?
Try what it says, `go test -cpu 2 ./...`.  Tests should be clean now.)

## Contributing Guide

Please be familiar with the [contributing guide](https://github.com/exercism/x-api/blob/master/CONTRIBUTING.md#the-exercise-data)
in the x-api repository.  This describes how all the language tracks are put together, as well as details about
the common metadata, and high-level information about contributing to existing problems and adding new problems.

## Problem Versioning

Each problem defines a `const targetTestVersion` in the test program, and validates that the solution has defined a matching value `testVersion`.  Any xgo developer that changes the test program or test data increments `targetTestVersion`.

The benefit of all this is that nipickers can see which test version a posted solution was written for and be spared confusion over why an old posted solution might not pass current tests.

Notice that neither the `testVersion` nor the `targetTestVersion` is exported. This is so that golint will not complain about a missing comment. In general, adding tests for unexported names is considered an anti-pattern, but in this case the trade-off seems acceptable.

## Xgo style

Let's walk through the first problem in `config.json`, leap.  Cd down into the leap directory now, there are two
files there, example.go and leap_test.go.  Example.go is a reference solution.  It is a valid solution that CI can
run tests against.  Solvers generally will not see it though.  Files with "example" in the file name are skipped
by `fetch`.  Because of this, there is less need for this code to be a model of style, expression and
readability, or to use the best algorithm.  Examples can be plain, simple, concise, even naïve, as long as they
are correct.  The test program though, is fetched for the solver and deserves attention for consistency and
appearance.

Leap_test.go uses a data-driven test.  Test cases are defined as data, then a test function iterates over
the data.  Identifiers within the method appear in actual-expected order as described at
[Useful Test Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures).
Here the identifier 'observed' is used instead of actual.  That's fine.  More common are words 'got' and 'want'.
They are clear and short.  Note
[Useful Test Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures) is part of
[Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).  Really we like most of the
advice on that page.

Also here in leap_test.go is a benchmark.  We throw in benchmarks because they're interesting, and because it's
idiomatic in Go to think about performance.  There is no critical use for these though.  Usually, like this one in
leap, they will just bench the combined time to run over all the test data rather than attempt precise timings
on single function calls.  They are useful if they let the solver try a change and see a performance effect.

## Xgo compared

We do a few things differently than the other language tracks.  In Go we generally have all tests enabled and do
not ask the solver to edit the test program, to enable progressive tests for example.  `Testing.Fatal`, as seen
in leap_test.go, will stop tests at the first problem encountered so the solver is not faced with too many errors
all at once.

We like errors in Go.  It's not idiomatic Go to ignore invalid data or have undefined behavior.  Sometimes our
Go tests require an error return where other language tracks don't.

## Generating test cases

Some problems that are implemented in multiple tracks use the same inputs and outputs to define the test suites.
Where the [x-common](https://github.com/exercism/x-common) repository contains json files with these inputs and
outputs, we can generate the test cases programmatically.

See `leap/example_gen.go` for an example of how this can be done.

Whenever the shared JSON data changes, the test cases will need to be regenerated. To do this, make sure that the
x-common repository has been cloned in the same parent-directory as the xgo repository. Then `cd` to the `xgo`
directory and run `go run <problem>/example_gen.go`.

You should see that the `<problem>/test_cases.go` file has changed. Commit the change.

## Pull requests

Pull requests are welcome.  You forked, cloned, coded and tested and you have something good?  Awesome!  Use git
to add, commit, and push to your repository.  Checkout your repository on the web now.  You should see your commit
and the invitation to submit a pull request!

<img src="img/mars1.png">

Click on that big green button.  You have a chance to add more explanation to your pull request here, then send
it.  Looking at the exercism/xgo repository now instead of your own, you see this.

<img src="img/mars2.png">

That inconspicuous orange dot is important!  Hover over it (no, not on this image, on a real page) and you can see
it's indicating that a Travis CI build is in progress.  After a few minutes (usually) that dot will turn green
indicating that tests passed.  If there's a problem, it comes up red:

<img src="img/mars3.png">

This means you've still got work to do.  Click on "details" to go to the Travis site and look over the build log
for clues.  Usually error messages will be helpful and you can correct the problem.

## Direction

Directions are unlimited.  This code is fresh and evolving.  Explore the existing code and you will see some new
directions being tried.  Your fresh ideas and contributions are welcome.  :sparkles:

## License

The MIT License (MIT)

Copyright (c) 2014 Katrina Owen, _@kytrinyx.com
