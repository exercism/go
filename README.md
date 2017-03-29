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
Clone your fork with the command: `git clone https://github.com/<you>/xgo`.
Test your clone by cding to the xgo directory and typing
`bin/test-without-stubs`. You should see tests pass for all exercises.

Note that unlike most other Go code, it is not necessary to clone this to your GOPATH.
This is because this repo only imports from the standard library and isn't expected to be imported by other packages.

There is a [misspelling tool](https://github.com/client9/misspell). You can install and occasionally run it to
find low hanging typo problems. [#570](https://github.com/exercism/xgo/pull/570) It's not added into CI since it could give false positives.

## Contributing Guide

Please be familiar with the [contributing guide](https://github.com/exercism/x-common/blob/master/CONTRIBUTING.md)
in the x-common repository.  This describes how all the language tracks are put together, as well as details about
the common metadata, and high-level information about contributing to existing problems and adding new problems.
In particular, please read the [Pull Request Guidelines](https://github.com/exercism/x-common/blob/master/CONTRIBUTING.md#pull-request-guidelines) before opening a pull request.

## xgo style

Let's walk through the `leap` exercise to see what is included in an exercise
implementation. Navigate into the *leap* directory, you'll see there are a
number of files there.

```sh
~/exercism/go/leap
$ tree
.
├── cases_test.go
├── example_gen.go
├── example.go
├── leap.go
└── leap_test.go

0 directories, 5 files
```

This list of files can vary across exercises so let's quickly run through
each file and briefly describe what it is.

**cases_test.go** - This file contains [generated test cases](#generating-test-cases),
and will only be present in some exercises.

**example_gen.go** - This file generates the *cases_test.go* file, and will only
be present in some exercises. See [generating test cases](#generating-test-cases)
for more information.

**example.go** - This is a reference solution for the exercise.

**leap.go** - This is a *stub file*, and will only be present in some exercises.

**leap_test.go** - This is the main test file for the exercise.

In some exercises there can be extra files, for instance the `series` exercise
contains extra test files.

### Example solutions

*example.go* is a reference solution. It is a valid solution that [Travis](https://travis-ci.org/exercism/xgo),
the CI (continuous integration) service, can run tests against. Solvers generally
will not see it though. Files with *"example"* in the file name are skipped by
the `exercism fetch` command. Because of this, there is less need for this code
to be a model of style, expression and readability, or to use the best algorithm.
Examples can be plain, simple, concise, even naïve, as long as they are correct.
The test file though, is fetched for the solver and deserves attention for consistency
and appearance.

### Tests

The `leap` exercise makes use of data-driven tests. Test cases are defined as
data, then a test function iterates over the data. In this exercise, as they are
generated, the test cases are defined in the *cases_test.go* file. The test function
that iterates over this data is defined in the *leap_test.go* file.

Identifiers within the test function appear in actual-expected order as described
at [Useful Test Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures).
Here the identifier `observed` is used instead of actual. That's fine. More
common are words `got` and `want`. They are clear and short. Note [Useful Test
Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures)
is part of [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
Really we like most of the advice on that page.

In Go we generally have all tests enabled and do not ask the solver to edit the
test program, to enable progressive tests for example. `t.Fatalf()`, as seen
in the *leap_test.go* file, will stop tests at the first failure encountered,
so the solver is not faced with too many failures all at once.

### Benchmarks

In most test files there will also be benchmark tests, as can be seen at the end
of the *leap_test.go* file. In Go, benchmarking is a first-class citizen of the
testing package. We throw in benchmarks because they're interesting, and because
it is idiomatic in Go to think about performance. There is no critical use for
these though. Usually they will just bench the combined time to run over all
the test data rather than attempt precise timings on single function calls. They
are useful if they let the solver try a change and see a performance effect.

### Testable examples

Some exercises can contain [Example tests](https://blog.golang.org/examples) that document the exercise API. These
examples are run alongside the standard exercise tests and will verify that the
exercise API is working as expected. They are not required by all exercises and
are most useful when clarifying the API of the exercise.

### Stub files

Stub files, such as *leap.go*, are a starting point for solutions. Not all exercises
need to do this; this is most helpful in the early exercises for newcomers to Go.
By convention, the stub file for an exercise with slug `exercise-slug`
must be named `exercise_slug.go`. This is because CI needs to delete stub files
to avoid conflicting definitions.

The track stops automatically providing stub files at the [twelve-days-exercise](exercises/twelve-days/).

### Problem Versioning

Each problem defines a `const targetTestVersion` in the test file, and validates
that the solution has defined a matching value `testVersion`. Any xgo developer
that changes the test file or test data must increment `targetTestVersion`.

The benefit of all this is that reviewers can see which test version a posted
solution was written for and be spared confusion over why an old posted solution
might not pass current tests.

Notice that neither the `testVersion` nor the `targetTestVersion` is exported.
This is so that golint will not complain about a missing comment. In general,
adding tests for unexported names is considered an anti-pattern, but in this
case the trade-off seems acceptable.

### Errors

We like errors in Go. It's not idiomatic Go to ignore invalid data or have undefined
behavior. Sometimes our Go tests require an error return where other language
tracks don't.

## Generating test cases

Some problems that are implemented in multiple tracks use the same inputs and
outputs to define the test suites. Where the [x-common](https://github.com/exercism/x-common)
repository contains json files with these inputs and outputs, we can generate
the test cases programmatically.

See the *example_gen.go* file in the `leap` exercise for an example of how this
can be done.

Whenever the shared JSON data changes, the test cases will need to be regenerated.
To do this, make sure that the **x-common** repository has been cloned in the same
parent-directory as the **xgo** repository. Then navigate into the **xgo**
directory and run `go run exercises/<problem>/example_gen.go`. You should see
that the `<problem>/test_cases.go` file has changed. Commit the change.

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

### Go icon
The Go logo was designed by Renée French, and has been released under the Creative Commons 3.0 Attributions license.
