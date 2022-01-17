# Instructions append

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

- [Concurrency in the Golang Book](https://www.golang-book.com/books/intro/10)
- [A Tour of Go's concurrency section](https://tour.golang.org/concurrency/1)
- [Go's sync.Map](https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c)

For a really deep dive you can try the book [Concurrency in Go](http://shop.oreilly.com/product/0636920046189.do) by [@kat-co](https://github.com/kat-co).

## Testing

For this exercise, the unit tests cannot determine whether you wrote a good concurrent solution.
Instead, it is best to solve this exercise [locally][cli] and execute the benchmarks with `go test -bench .`.
For a good solution, you should see that the concurrent version shows a lower number of nano seconds per operation (`ns/op`) than the sequential version.

[cli]: https://exercism.org/docs/using/solving-exercises/working-locally
