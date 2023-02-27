# Performance

In this approach, we'll find out how to most efficiently solve Clock in Go.

The [approaches page][approaches] lists two idiomatic approaches to this exercise:

1. [Typing `Clock` as an `int`][approach-clock-as-int]
2. [Typing `Clock` as a `struct`][approach-clock-as-struct]


## Benchmarks

To benchmark the approaches, we ran the following Benchmark code for each approach:

```go
func BenchmarkAddMinutes(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	c := New(12, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range addTests {
			c.Add(a.a)
		}
	}
}

func BenchmarkSubtractMinutes(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	c := New(12, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range subtractTests {
			c.Subtract(a.a)
		}
	}
}

func BenchmarkCreateClocks(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, n := range timeTests {
			New(n.h, n.m)
		}
	}
}
```

and received the following results:

```
// clock as int
BenchmarkCreateClocks-12    	182396848	         5.944 ns/op	       0 B/op	       0 allocs/op
BenchmarkSubtractMinutes-12    	372009392	         3.131 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMinutes-12    	386899329	         3.114 ns/op	       0 B/op	       0 allocs/op

// clock as struct
BenchmarkCreateClocks-12    	191981476	         5.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkSubtractMinutes-12    	367081141	         3.114 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMinutes-12    	361321388	         3.247 ns/op	       0 B/op	       0 allocs/op
```

The above are the fastest times benched for each approach from multiple runs.
Generally, the fewer bytes allocated per op (`B/op`) the faster (i.e. the fewer ns/op) the implementation.
More info on reading benchmarks can be found [here][benchmark].

The two approaches were very similar in performance.
The approach of `Clock` as a struct may have a bit of overhead with the indirection of having the minutes inside a struct,
but the approach of Clock as an `int` requires a lot of converting the `Clock` type to an `int` type,
even though the underlying type of `Clock` is an `int`.
That converting may also have some overhead.

[approaches]: https://exercism.org/tracks/go/exercises/clock/approaches
[approach-clock-as-int]: https://exercism.org/tracks/go/exercises/clock/approaches/clock-as-int
[approach-clock-as-struct]: https://exercism.org/tracks/go/exercises/clock/approaches/clock-as-struct
[benchmark]: https://www.mikenewswanger.com/posts/2018/benchmarking-in-go/
