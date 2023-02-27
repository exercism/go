# Performance

In this approach, we'll find out how to most efficiently solve Pig Latin in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using map lookups][approach-map-lookups]
2. [Using map lookups with generics][approach-map-lookups-with-generics]


## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkSentence(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Sentence(test.input)
		}
	}
}
```

and received the following results:

```
// map lookups
BenchmarkSentence-12    	  310228	      3845 ns/op	     616 B/op	      46 allocs/op

// map lookups with generics
BenchmarkSentence-12    	  306960	      3809 ns/op	     616 B/op	      46 allocs/op
```

The above are the fastest times benched for each approach from multiple runs.
Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The fastest benchmarks were from map lookups with generics, but not consistently so.
Benchmarks runs for each approach varied by as much as 100 nanoseconds to 200 nanoseconds or so, but the fastest benchmarks from several runs were for map lookups with generics.

[approaches]: https://exercism.org/tracks/go/exercises/pig-latin/approaches
[approach-map-lookups]: https://exercism.org/tracks/go/exercises/pig-latin/approaches/map-lookups
[approach-map-lookups-with-generics]: https://exercism.org/tracks/go/exercises/pig-latin/approaches/map-lookups-with-generics
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
