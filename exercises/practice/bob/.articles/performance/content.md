# Performance

In this approach, we'll find out how to most efficiently determine the response for Bob in Go.

The [approaches page][approaches] lists three idiomatic approaches to this exercise:

1. [Using `if` statements][approach-if].
2. [Using a `switch` statement][approach-switch].
3. [Using an answer array][approach-answer-array].

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go

func BenchmarkHey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Hey(tt.input)
		}
	}
}
```

and received the following results:

```
if statements
BenchmarkHey-12    	  624798	      1746 ns/op	     432 B/op	      12 allocs/op

switch statement
BenchmarkHey-12    	  617388	      1791 ns/op	     432 B/op	      12 allocs/op

answer array
BenchmarkHey-12    	  632600	      1819 ns/op	     432 B/op	      12 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

Benchmarks can vary between one run and another by as much as a hundred nanoseconds,
so these values represent the fastest results from several runs.

The `if` statements approach gave the fastest result.
The `switch` statement approach was the second-fastest.
The answer array approach came in third.
Since each approach sometimes gave results slower than the other approaches, which to use could be a matter of stylistic choice. 

[approaches]: https://exercism.org/tracks/go/exercises/bob/approaches
[approach-if]: https://exercism.org/tracks/go/exercises/bob/approaches/if-statements
[approach-switch]: https://exercism.org/tracks/go/exercises/bob/approaches/switch-statement
[approach-answer-array]: https://exercism.org/tracks/go/exercises/bob/approaches/answer-array
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
