# Say

Given a number from 0 to 999,999,999,999, spell out that number in English.

## Step 1

Handle the basic case of 0 through 99.

If the input to the program is `22`, then the output should be
`'twenty-two'`.

Your program should throw an error if given a number outside the
range of 0 through 99.

Some good test cases for this program are:

- 0
- 14
- 50
- 98
- 99 
- 100 (currently an error case)

### Extension

If you're on a Mac, shell out to Mac OS X's `say` program to talk out
loud.

## Step 2

Handle the the case of 0 through 999,999,999,999.

Implement breaking a number up into chunks of thousands.

So `1234567890` should yield a list like 1, 234, 567, and 890, while the
far simpler `1000` should yield just 1 and 0.

The program must also report any values that are larger than 999,999,999,999.

## Step 3

Now handle inserting the appropriate scale word between those chunks.

So `1234567890` should yield `'1 billion 234 million 567 thousand 890'`

The program must also report any values that are larger than 999,999,999,999.

## Step 4

Put it all together to get nothing but plain English.

`12345` should give `twelve thousand three hundred forty-five`.

The program must also report any values that are larger than 999,999,999,999.

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

## Source

A variation on JavaRanch CattleDrive, exercise 4a [http://www.javaranch.com/say.jsp](http://www.javaranch.com/say.jsp)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
