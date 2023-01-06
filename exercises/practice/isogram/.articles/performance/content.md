# Performance

In this approach, we'll find out how to most efficiently determine if a string is an Isogram in Go.

The [approaches page][approaches] lists three idiomatic approaches to this exercise:

1. [Using `map`][approach-map].
2. [Using `strings.Count()`][approach-strings-count].
3. [Using a bit field][approach-bitfield].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkIsIsogram(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range testCases {
			IsIsogram(c.input)
		}

	}
}
```

and received the following results:

```
map
BenchmarkIsIsogram-12    	  237565	      4960 ns/op	    1201 B/op	      13 allocs/op

strings.count
BenchmarkIsIsogram-12    	  896049	      1409 ns/op	      40 B/op	       3 allocs/op

bit field
BenchmarkIsIsogram-12    	 6214525	       185.2 ns/op	       0 B/op	       0 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The fastest is the `Bitfeld` approach, but it depends on all of the letters being [ASCII][ascii].
The fastest approach that supports Unicode is the `strings.Count()` approach.

[approaches]: https://exercism.org/tracks/go/exercises/isogram/approaches
[approach-map]: https://exercism.org/tracks/go/exercises/isogram/approaches/map
[approach-strings-count]: https://exercism.org/tracks/go/exercises/isogram/approaches/strings-count
[approach-bitfield]: https://exercism.org/tracks/go/exercises/isogram/approaches/bitfield
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
[ascii]: https://www.asciitable.com/
