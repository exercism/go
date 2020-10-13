# Resistor Color

Resistors have color coded bands, where each color maps to a number. The first 2 bands of a resistor have a simple encoding scheme: each color maps to a single number.

These colors are encoded as follows:

- Black: 0
- Brown: 1
- Red: 2
- Orange: 3
- Yellow: 4
- Green: 5
- Blue: 6
- Violet: 7
- Grey: 8
- White: 9

The goal of this exercise is to create a way:
- to look up the numerical value associated with a particular color band
- to list the different band colors

Mnemonics map the colors to the numbers, that, when stored as an array, happen to map to their index in the array: Better Be Right Or Your Great Big Values Go Wrong.

More information on the color encoding of resistors can be found in the [Electronic color code Wikipedia article](https://en.wikipedia.org/wiki/Electronic_color_code)

## Coding the solution

Look for a stub file having the name acronym.go
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

Julien Vanier [https://github.com/monkbroc](https://github.com/monkbroc)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
