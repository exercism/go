# xGo

[![Build Status](https://travis-ci.org/exercism/xgo.png?branch=master)](https://travis-ci.org/exercism/xgo)

Exercism exercises in Go

## Issues

We welcome issues filed at https://github.com/exercism/xgo/issues for problems of any size.  Feel free to report
typographical errors or poor wording for example.  We are most interested in improving the quality of the test
suites.  You can greatly help us improve the quality of the exercises by filing reports of invalid solutions that
pass tests or of valid solutions that fail tests 

## Development setup

Beyond filing issues, if you would like to contribute directly to the Go code in xgo, you should follow some
standard Go development practices.  You should have a [recent version of Go](http://golang.org/doc/install)
installed.  Our continuous integration tests that our code base works with the current "release" version of Go,
the previous minor version release, and also with "tip."  You don't need to have all of these installed, any one
of them is fine.

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

## License

The MIT License (MIT)

Copyright (c) 2014 Katrina Owen, _@kytrinyx.com
