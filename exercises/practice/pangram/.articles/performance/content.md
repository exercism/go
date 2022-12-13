# Performance

In this approach, we'll find out how to most efficiently determine if a string is an Pangram in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using `strings.ContainsRune()`][approach-strings-containsrune].
2. [Using a bit field][approach-bitfield].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkPangram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsPangram(test.input)
		}
	}
}
```

and received the following results:

```
bit field
BenchmarkPangram-12    	 6435968	       164.8 ns/op	       0 B/op	       0 allocs/op

strings.ContainsRune()
BenchmarkPangram-12    	  707517	      1723 ns/op	      96 B/op	       2 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The fastest is the `Bit feld` approach.

[approaches]: https://exercism.org/tracks/go/exercises/pangram/approaches
[approach-strings-containsrune]: https://exercism.org/tracks/go/exercises/pangram/approaches/strings-containsrune
[approach-bitfield]: https://exercism.org/tracks/go/exercises/pangram/approaches/bitfield
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
