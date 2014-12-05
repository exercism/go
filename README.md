# xGo

[![Build Status](https://travis-ci.org/exercism/xgo.png?branch=master)](https://travis-ci.org/exercism/xgo)

Exercism exercises in Go

## Issues

We welcome issues filed at https://github.com/exercism/xgo/issues for problems of any size.  Feel free to report
typographical errors or poor wording.  We are most interested in improving the quality of the test suites.
You can greatly help us improve the quality of the exercises by filing reports of invalid solutions that
pass tests or of valid solutions that fail tests 

## Development setup

Beyond filing issues, if you would like to contribute directly to the Go code in xgo, you should follow some
standard Go development practices.  You should have a [recent version of Go](http://golang.org/doc/install)
installed.  Our continuous integration (CI) tests that our code base works with the current "release" version of
Go, the previous minor version release, and also with "tip."  You don't need to have all of these installed, any
one of them is fine.

You will need a github account and you will need to fork exercism/xgo to your account.
See [GitHub Help](https://help.github.com/articles/fork-a-repo/) if you are unfamiliar with the process.
Clone your fork into a Go workspace on your GOPATH.  Read [How to Write Go Code](http://golang.org/doc/code.html)
if you are unfamiliar with Go workspaces and GOPATH.  As an example, if you have `$HOME/dev` on your GOPATH, then
you should cd to `$HOME/dev/src/github.com/<you>` to type the clone command.
Type `git clone https://github.com/<you>/xgo` then, to clone into `$HOME/dev/src/github.com/<you>/xgo`.  Test your
clone by cding to the xgo directory and typing `go test ./...`.  You should see tests pass for all exercises.
(Oh wait, do you see the error message from paasio?  Try what it says, `go test -cpu 2 ./...`.  Tests should be
clean now.

## Contributing Guide

Please be familiar with the [contributing guide](https://github.com/exercism/x-api/blob/master/CONTRIBUTING.md#the-exercise-data)
in the x-api repository.  This describes how all the language tracks are put together, as well as details about
the common metadata, and high-level information about contributing to existing problems and adding new problems.

## Xgo style

Let's walk through the first problem in `config.json`, leap.  Cd down into the leap directory now, there are two
files there, example.go and leap_test.go.  Example.go is a reference solution.  It is a valid solution that CI can
run tests against.  Solvers generally will not see it though.  Files with "example" in the file name are skipped
by `fetch`.  Because of this, there is less need for this code to be an example of style, expression and
readability, or to use the best algorithm.  Examples can be plain, simple, concise, even naïve, as long as they
are correct.  The test program though, is fetched for the solver and deserves attention for consistency and
appearance.

Leap_test.go uses a data-driven test.  Test cases are defined as data, then a test function iterates over
the data.  Identifiers within the method appear in actual-expected order.  Here the identifier 'observed' is used
instead of actual.  That's fine.  The words 'got' and 'want' are common in go tests.  They are clear and short.
See [Useful Test Failures](http://code.google.com/p/go-wiki/wiki/CodeReviewComments#Useful_Test_Failures).
Really we like most of the advice on that page.

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

## Pull requests

Pull requests are welcome.  You forked, cloned, coded and tested and you have something good?  Awesome!  Use git
to add, commit, and push to your repository.  Checkout your repository on the web now.  You should see your commit
and the invitation to submit a pull request!

<img src="img/mars1.png">

Click on that big green button.  You have a chance to add more explanation to your pull request here, then send
it.  Looking at the exercism/xgo repository now instead of your own, you see this.

<img src="img/mars2.png">

That inconspicuous orange dot is important!  Hover over it (no, not on this page, on a real page) and you can see
it's indicating that the Travis CI build is in progress.  After a few minutes (usually) that dot will turn green
indicating that tests passed.  If there's a problem, it comes up red:

<img src="img/mars3.png">

This means you've still got work to do.  Click on "details" to go to the Travis site and look over the build log
for clues.  Usually error messages will be helpful and you can correct the problem.

## Direction

Directions are unlimited.  This code is fresh and evolving.  Explore the existing code and you will see some new
directions being tried.  Your fresh ideas and contributions are welcome.

## License

The MIT License (MIT)

Copyright (c) 2014 Katrina Owen, _@kytrinyx.com
