# Performance

In this approach, we'll find out how to most efficiently validate a number with the Luhn algorithm in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using the scrub and validate first appproach][approach-scrub-and-validate-first]
2. [Uisng the validate as you go approach][approach-validate-as-you-go]

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Valid(test.input)
		}
	}
}
```

and received the following results:

```
validate as you go
BenchmarkValid-12    	 4668248	       248.4 ns/op	       0 B/op	       0 allocs/op

scrub and validate first
BenchmarkValid-12    	  969022	      1121 ns/op	     352 B/op	      14 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The fastest is the validate as you go approach.

[approaches]: https://exercism.org/tracks/go/exercises/luhn/approaches
[approach-scrub-and-validate-first]: https://exercism.org/tracks/go/exercises/luhn/approaches/scrub-and-validate-first
[approach-validate-as-you-go]: https://exercism.org/tracks/go/exercises/luhn/approaches/validate-as-you-go
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
