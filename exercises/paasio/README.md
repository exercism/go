# Paasio

Report network IO statistics.

You are writing a [PaaS][], and you need a way to bill customers based
on network and filesystem usage.

Create a wrapper for network connections and files that can report IO
statistics. The wrapper must report:

- The total number of bytes read/written.
- The total number of read/write operations.

[PaaS]: http://en.wikipedia.org/wiki/Platform_as_a_service

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

Brian Matsuo [https://github.com/bmatsuo](https://github.com/bmatsuo)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
