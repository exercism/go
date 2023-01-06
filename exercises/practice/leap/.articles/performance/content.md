# Performance

In this approach, we'll find out how to most efficiently calculate if a year is a leap year in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Using the boolean chain][approach-boolean-chain]
2. [Using a maximum of two checks][approach-max-two-checks]

For our performance investigation, we'll also include two other approaches:

- [Using the `time.Add()` approach][approach-time-add]
- [Using the `time.YearDay()` approach][approach-time-yearday]

## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
// Benchmark 400 year interval to get fair weighting of different years.
func Benchmark400(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for y := 1600; y < 2000; y++ {
			IsLeapYear(y)
		}
	}
}
```

and received the following results:

```
boolean chain
Benchmark400-12    	 4936674	       219.4 ns/op	       0 B/op	       0 allocs/op

max two checks
Benchmark400-12    	10733403	       110.8 ns/op	       0 B/op	       0 allocs/op

time add
Benchmark400-12    	   89797	     12741 ns/op	       0 B/op	       0 allocs/op

time yearday
Benchmark400-12    	  118605	      9700 ns/op	       0 B/op	       0 allocs/op
```

Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

You can see that the maximum of two checks approach was about twice as fast as the boolean chain,
even though the maximum of two checks approach starts from a less likely condition.

[approaches]: https://exercism.org/tracks/go/exercises/leap/approaches
[approach-boolean-chain]: https://exercism.org/tracks/go/exercises/leap/approaches/boolean-chain
[approach-max-two-checks]: https://exercism.org/tracks/go/exercises/leap/approaches/max-two-checks
[approach-time-add]: https://exercism.org/tracks/go/exercises/leap/approaches/time-addition
[approach-time-yearday]: https://exercism.org/tracks/go/exercises/leap/approaches/time-yearday
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
