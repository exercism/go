# Performance

In this approach, we'll find out how to most efficiently calculate the value for Grains in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using `math.Pow()` and `big` exponentiation][approach-math-pow-big-exponentiation]
2. [Using bit shifting][approach-bit-shifting]


## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkSquare(b *testing.B) {

	for i := 0; i < b.N; i++ {

		for _, test := range squareTests {
			Square(test.input)
		}

	}
}

func BenchmarkTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Total()
	}
}
```

and received the following results:

```
math.Pow square
BenchmarkSquare-12    	 4761478	       241.9 ns/op	      48 B/op	       3 allocs/op

big exponentiation total
BenchmarkTotal-12    	 5606967	       238.7 ns/op	     128 B/op	       6 allocs/op

bit shift square
BenchmarkSquare-12    	98908706	        11.21 ns/op	       0 B/op	       0 allocs/op

bit shift total
BenchmarkTotal-12    	1000000000	         0.2689 ns/op	       0 B/op	       0 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The bit shifting approach is much faster than using `math.Pow()` and `big` exponentiation.

[approaches]: https://exercism.org/tracks/go/exercises/grains/approaches
[approach-math-pow-big-exponentiation]: https://exercism.org/tracks/go/exercises/grains/approaches/math-pow-big-exponentiation
[approach-bit-shifting]: https://exercism.org/tracks/go/exercises/grains/approaches/bit-shifting
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
