# Hamming

Calculate the Hamming Distance between two DNA strands.

Your body is made up of cells that contain DNA. Those cells regularly wear out and need replacing, which they achieve by dividing into daughter cells. In fact, the average human body experiences about 10 quadrillion cell divisions in a lifetime!

When cells divide, their DNA replicates too. Sometimes during this process mistakes happen and single pieces of DNA get encoded with the incorrect information. If we compare two strands of DNA and count the differences between them we can see how many mistakes occurred. This is known as the "Hamming Distance".

We read DNA using the letters C,A,G and T. Two strands might look like this:

    GAGCCTACTAACGGGAT
    CATCGTAATGACGGCCT
    ^ ^ ^  ^ ^    ^^

They have 7 differences, and therefore the Hamming Distance is 7.

The Hamming Distance is useful for lots of things in science, not just biology, so it's a nice phrase to be familiar with :)

# Implementation notes

The Hamming distance is only defined for sequences of equal length, so
an attempt to calculate it between sequences of different lengths should
not work, returning an error. 

With this exercise, you are introduced to the use of 
[multiple return](https://tour.golang.org/basics/6) feature and here there 
is an example:

    func Foo(a, b string) (string, error) {
        if len(a) > len(b) {
            return "", errors.New("An error occured!")
        }
    
        // No errors:
        return fmt.Sprintf("%s - %s", a, b), nil
    }
    

The caller can now easily check if everything was OK:

    foo, err := Foo(a, b)
    if err != nil {
        fmt.Println(err)
        
        // other logic, if needed
        return
    }
    
    fmt.Println(foo)

There are multiple ways to create an error:

    errors.New("Some description")
    fmt.Errorf("Error %d occured", 503)
    
    // worth then mention but should be avoided:
    errors.New(fmt.Sprintf("Error %d occured", 503)
    
Did you know? The entire errors package in Go is only 4 
[lines of code](https://go.googlesource.com/go/+/go1.15.7/src/errors/errors.go).

You may be wondering about the `cases_test.go` file. We explain it in the
[leap exercise][leap-exercise-readme].

[leap-exercise-readme]: https://github.com/exercism/go/blob/master/exercises/leap/README.md 


## Coding the solution

Look for a stub file having the name hamming.go
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

The Calculating Point Mutations problem at Rosalind [http://rosalind.info/problems/hamm/](http://rosalind.info/problems/hamm/)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
