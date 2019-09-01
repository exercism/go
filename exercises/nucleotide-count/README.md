# Nucleotide Count

Given a single stranded DNA string, compute how many times each nucleotide occurs in the string.

The genetic language of every living thing on the planet is DNA.
DNA is a large molecule that is built from an extremely long sequence of individual elements called nucleotides.
4 types exist in DNA and these differ only slightly and can be represented as the following symbols: 'A' for adenine, 'C' for cytosine, 'G' for guanine, and 'T' thymine.

Here is an analogy:
- twigs are to birds nests as
- nucleotides are to DNA as
- legos are to lego houses as
- words are to sentences as...

## Implementation

You should define a custom type 'DNA' with a function 'Counts' that outputs two values: 

- a frequency count for the given DNA strand
- an error (if there are invalid nucleotides)

Which is a good type for a DNA strand ? 

Which is the best Go types to represent the output values ?

Take a look at the test cases to get a hint about what could be the possible inputs.


## note about the tests
You may be wondering about the `cases_test.go` file. We explain it in the
[leap exercise][leap-exercise-readme].

[leap-exercise-readme]: https://github.com/exercism/go/blob/master/exercises/leap/README.md



## Coding the solution

Look for a stub file having the name nucleotide_count.go
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

The Calculating DNA Nucleotides_problem at Rosalind [http://rosalind.info/problems/dna/](http://rosalind.info/problems/dna/)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
