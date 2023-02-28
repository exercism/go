# Slice of `bool`s

```go
// Package sieve is a small library for finding prime numbers
package sieve

// Sieve calculates and returns prime numbers up to the limit
func Sieve(limit int) []int {
	composite := make([]bool, limit+1)
	primes := make([]int, limit/2)
	primeIndex := 0

	for number := 2; number <= limit; number++ {
		if !composite[number] {
			primes[primeIndex] = number
			primeIndex++
			for idx := number + number; idx <= limit; idx += number {
				composite[idx] = true
			}
		}
	}
	return primes[:primeIndex]
}
```

This approach starts by using the [make][make] function to define the slice which will keep track of the composite numbers.
The length of `limit+1` is used so that the slice does not have to be reallocated.
The `bool` values are initialized to their zero `false` value.

Similarly, the slice to keep track of primes is also defined with a length, which is set to half of the number being calculated up to, since
at least half of the numbers will be even and so not prime.
An index into the slice of primes is set for `0`.

Since values less than `2` are not prime, the iteration of the outer [for loop][for-loop] begins at `2`.
If the number being iterated is not a composite number, then the element in the slice of primes at the prime index is set to that number, and the
prime index is incremented.

Since any number evenly divisible by that prime is not prime, the inner `for` loop iterates from the prime plus itself, setting each of those
numbers to `true` for being a composite number.

After the outer loop is done, a slice of the primes slice up to but not including the prime index is returned.

For example if `limit` is `2`, then the slice  of primes would be made with a length of `1`.
The element at index `0` would be set to `2` and the prime index incremented to `1`.
The slice of primes would be sliced from the beginning up to but not including the element at index `1`
(good, since index `1` is beyond the length of the slice),
and `[]int{2}` would be returned.

if `limit` is `10`, then the slice  of primes would be made with a length of `5`.
The last prime would be `7`.
The element at index `3` would be set to `7` and the prime index incremented to `4`.
The slice of primes would be sliced from the beginning up to but not including the element at index `4`,
and `[]int{2, 3, 5, 7}` would be returned.

[make]: https://go.dev/tour/moretypes/13
[for-loop]: https://go.dev/tour/flowcontrol/1
