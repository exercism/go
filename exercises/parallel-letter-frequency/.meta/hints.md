## Concurrency vs Parallelism

You may notice that while this exercise is called _Parallel_ letter frequency
you don't see the term "Parallel" used very often in Go. Gophers prefer to use
the term **Concurrent** to describe the management of multiple independent
processes. Although these terms are often used interchangeably Gophers like to
be technically correct and use "concurrent" when discussing the seeming
simultaneous executions of processes. While we can plan for our programs to run
in parallel, and at times they may appear to run in parallel, without strictly
knowing the execution context of our code all we can guarantee is that processes
will run concurrently. In other words they may be executing sequentially faster
than we can distinguish but not strictly simultaneously.

For more take a look at The Go Blog's post: [Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
