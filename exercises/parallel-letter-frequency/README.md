# Parallel Letter Frequency

Count the frequency of letters in texts using parallel computation.

Parallelism is about doing things in parallel that can also be done
sequentially. A common example is counting the frequency of letters.
Create a function that returns the total frequency of each letter in a
list of texts and that employs parallelism.

## Concurrency vs Parallelism

Go supports concurrency via "[goroutines](https://golangbot.com/goroutines/)"
which are started with the `go` keyword. It is a simple, lightweight and elegant
way to provide concurrency support and is one of the greatest strengths of the
language.

You may notice that while this exercise is called _Parallel_ letter frequency
you don't see the term "Parallel" used very often in Go. Gophers prefer to use
the term **Concurrent** to describe the management of multiple independent
goroutines ("processes" or "threads" in other language contexts). Although
these terms are often used interchangeably Gophers like to be technically
correct and use "concurrent" when discussing the seemingly simultaneous
executions of goroutines.

While we can plan for our programs to run in parallel, and at times they may
appear to run in parallel, without strict knowledge of the execution context of
our code all we can guarantee is that processes will run concurrently. In other
words they may be executing sequentially faster than we can distinguish but not
strictly simultaneously.

For more take a look at The Go Blog's post: [Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism).

## Concurrency Resources

If you are new to the concurrency features in Go here are some resources to get
you started. We recommend looking over these before starting this exercise:

* [Concurrency in the Golang Book](https://www.golang-book.com/books/intro/10)
* [A Tour of Go's concurrency section](https://tour.golang.org/concurrency/1)
* [Go's sync.Map](https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c)

For a really deep dive you can try the book [Concurrency in Go](http://shop.oreilly.com/product/0636920046189.do) by [@kat-co](https://github.com/kat-co).


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
