# Performance

In this approach, we'll find out how to most efficiently determine the Scrabble Score in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using `map`][approach-map].
2. [Using a `switch` on runes][approach-switch-on-runes].

For our performance investigation, we'll also include a third approach that [uses a `switch` on strings][approach-switch-on-strings].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}
```

and received the following results:

```
map
BenchmarkScore-12    	 1544814	       711.7 ns/op	       0 B/op	       0 allocs/op

switch on runes
BenchmarkScore-12    	 4171696	       242.1 ns/op	       0 B/op	       0 allocs/op

switch on strings
BenchmarkScore-12    	  413484	      2527 ns/op	    1088 B/op	      12 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

Using `switch` on runes was about three times faster than the next fastest approach, which was using a `map`.
Using a `switch` on each letter being a separate string was over ten times slower than switching on runes.

[approaches]: https://exercism.org/tracks/go/exercises/scrabble-score/approaches
[approach-map]: https://exercism.org/tracks/go/exercises/scrabble-score/approaches/map
[approach-switch-on-runes]: https://exercism.org/tracks/go/exercises/scrabble-score/approaches/switch-on-runes
[approach-switch-on-strings]: https://exercism.org/tracks/go/exercises/scrabble-score/approaches/switch-on-strings
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
