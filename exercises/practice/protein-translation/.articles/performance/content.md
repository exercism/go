# Performance

In this approach, we'll find out how to most efficiently solve Protein Translation in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using a `switch` statement][approach-switch-statement].
2. [Using a `map` object][approach-map-object].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkProtein(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, test := range proteinTestCases {
		for i := 0; i < b.N; i++ {
			FromRNA(test.input)
		}
	}
}
```

and received the following results:

```
switch statement
BenchmarkProtein-12    	 1207272	       952.3 ns/op	     400 B/op	      12 allocs/op

map object
BenchmarkProtein-12    	 1254066	       953.3 ns/op	     400 B/op	      12 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

Both approaches are about the same performance.

[approaches]: https://exercism.org/tracks/go/exercises/protein-translation/approaches
[approach-switch-statement]: https://exercism.org/tracks/go/exercises/protein-translation/approaches/switch
[approach-map-object]: https://exercism.org/tracks/go/exercises/protein-translation/approaches/map
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
