# Introduction

There are several variants of a general approach for solving Sieve.
The general approach is to keep track of composite numbers with a slice of `bool` values.

## General guidance

Something to consider is to keep the number of allocations at a minimum to get the best performance.

## Approach: Slice of `bool`s

```go
// Package sieve is a small library for finding prime numbers
package sieve

// Sieve calculates and returns prime numbers up to the limit
func Sieve(limit int) []int {
	composite := make([]bool, limit+1)
	primes := make([]int, limit/2)
	prime_index := 0

	for number := 2; number <= limit; number++ {
		if !composite[number] {
			primes[prime_index] = number
			prime_index++
			for idx := number + number; idx <= limit; idx += number {
				composite[idx] = true
			}
		}
	}
	return primes[:prime_index]
}
```

For more information, check the [Slice of `bool`s approach][approach-slice-of-bools].

[approach-slice-of-bools]: https://exercism.org/tracks/go/exercises/sieve/approaches/slice-of-bools
