# Performance

In this approach, we'll find out how to most efficiently solve ISBN Verifier in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Iterating characters as bytes][approach-iterate-bytes].
2. [Iterating characters as runes][approach-iterate-runes].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkIsValidISBN(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, n := range testCases {
			IsValidISBN(n.isbn)
		}
	}
}
```

and received the following results:

```
iterate as bytes
BenchmarkIsValidISBN-12    	 4561827	       253.0 ns/op	       0 B/op	       0 allocs/op

iterate as runes
BenchmarkIsValidISBN-12    	 3094387	       354.9 ns/op	       0 B/op	       0 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The fastest is the iterate as bytes approach.
As of the time of this writing we can iterate bytes, since all of the characters are [ASCII][ascii].

[approaches]: https://exercism.org/tracks/go/exercises/isbn-verifier/approaches
[approach-iterate-bytes]: https://exercism.org/tracks/go/exercises/isbn-verifier/approaches/iterate-bytes
[approach-iterate-runes]: https://exercism.org/tracks/go/exercises/isbn-verifier/approaches/iterate-runes
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
[ascii]: https://www.asciitable.com/
