# Performance

In this approach, we'll find out how to most efficiently solve RNA Transcription in Go.

The [approaches page][approaches] lists three idiomatic approaches to this exercise:

1. [Using the `strings.Map()` function][approach-strings-map].
2. [Using a `map` object][approach-map-object].
3. [Using a `switch` statement][approach-switch-statement].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkRNATranscription(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range rnaTests {
			ToRNA(test.input)
		}
	}
}
```

and received the following results:

```
map function
BenchmarkRNATranscription-12    	 5396551	       217.5 ns/op	      37 B/op	       5 allocs/op

map object
BenchmarkRNATranscription-12    	 8124092	       154.8 ns/op	      16 B/op	       5 allocs/op

switch
BenchmarkRNATranscription-12    	11808524	        97.69 ns/op	      16 B/op	       5 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

Using `switch` statement was the fastest approach.
Using a `map` object was faster than using the `strings.Map()` function.

[approaches]: https://exercism.org/tracks/go/exercises/rna-transcription/approaches
[approach-strings-map]: https://exercism.org/tracks/go/exercises/rna-transcription/approaches/map-function
[approach-map-object]: https://exercism.org/tracks/go/exercises/rna-transcription/approaches/map-object
[approach-switch-statement]: https://exercism.org/tracks/go/exercises/rna-transcription/approaches/switch
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
