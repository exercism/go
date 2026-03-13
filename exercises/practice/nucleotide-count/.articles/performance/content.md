# Performance

In this approach, we'll find out how to most efficiently solve Nucleotide Count in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using a `switch` statement][approach-switch-statement].
2. [Using `if` with a short statement][approach-if-with-short-statement].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkCounts(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			dna := DNA(test.strand)
			dna.Counts()
		}
	}
}
```

and received the following results:

```text
switch statement
BenchmarkCounts-12    	  593500	      1942 ns/op	     840 B/op	      12 allocs/op

if with short statement
BenchmarkCounts-12    	  527168	      2208 ns/op	     840 B/op	      12 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
See [more info on reading benchmarks][benchmark].

The fastest is the `switch` statement approach.

[approaches]: https://exercism.org/tracks/go/exercises/nucleotide-count/approaches
[approach-switch-statement]: https://exercism.org/tracks/go/exercises/nucelotide-count/approaches/switch-statement
[approach-if-with-short-statement]: https://exercism.org/tracks/go/exercises/nucelotide-count/approaches/if-with-short-statement
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
