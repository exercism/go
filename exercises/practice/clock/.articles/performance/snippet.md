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
