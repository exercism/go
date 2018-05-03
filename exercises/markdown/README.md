# Markdown

Refactor a Markdown parser.

The markdown exercise is a refactoring exercise. There is code that parses a
given string with [Markdown
syntax](https://guides.github.com/features/mastering-markdown/) and returns the
associated HTML for that string. Even though this code is confusingly written
and hard to follow, somehow it works and all the tests are passing! Your
challenge is to re-write this code to make it easier to read and maintain
while still making sure that all the tests keep passing.

It would be helpful if you made notes of what you did in your refactoring in
comments so reviewers can see that, but it isn't strictly necessary. The most
important thing is to make the code better!

## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `-bench`
flag:

    go test -bench .

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/about).

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
