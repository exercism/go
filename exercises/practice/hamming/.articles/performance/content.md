# Performance

In this approach, we'll find out how to most efficiently solve Hamming in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Iterating characters as bytes][approach-iterate-bytes].
2. [Iterating characters as runes][approach-iterate-runes].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			// ignoring errors and results because we're just timing function execution
			_, _ = Distance(tc.s1, tc.s2)
		}
	}
}
```

and received the following results:

```
iterate as bytes
BenchmarkHamming-12    	43623610	        26.92 ns/op	       0 B/op	       0 allocs/o

iterate as runes
BenchmarkHamming-12    	 3330171	       361.5 ns/op	      64 B/op	       4 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The fastest is the iterate as bytes approach.
As of the time of this writing we can iterate bytes, since all of the characters are [ASCII][ascii].

[approaches]: https://exercism.org/tracks/go/exercises/hamming/approaches
[approach-iterate-bytes]: https://exercism.org/tracks/go/exercises/hamming/approaches/iterate-bytes
[approach-iterate-runes]: https://exercism.org/tracks/go/exercises/hamming/approaches/iterate-runes
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
[ascii]: https://www.asciitable.com/
