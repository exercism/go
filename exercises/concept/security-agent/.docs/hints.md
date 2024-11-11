# Hints

## 1. Generate important number

- Remember that there’s an operation that can replace multiplication by 2
- Consider how `this` can help multiply a number by 2 several times in one step

## 2. Integrity Check

- Imagine that you have two identical numbers
- Think about using an operation that has a specific behavior with identical numbers
- If you start with `0` and apply this operation with number `a`, you get `a`: `0 _ a = a`
- Applying the operation again with the same number returns `0`:
`0 _ a _ a = 0`
- This operation doesn’t require a strict order of numbers — if the numbers are swapped, it still works
- `0 _ a _ b _ a = b`, `0 _ b _ a _ a _ b = 0`
- Apply this operation across [all elements][range-over-link] of the slice to determine if any number lacks a pair

## 3. Setting Flags

- Recall what you can use when you know which bits need to be changed
- If you need to set bits to `0`, consider an operation that always returns `0` regardless of the bit value in the other operand

## 4. Generate Key

- Start by performing the specified operation on `num`
- If you need to end up with a `uint32`, remember to [handle][type-conversions-link] that
- Place each bit in its _correct_ position

[range-over-link]: https://gobyexample.com/range-over-built-in-types
[type-conversions-link]: https://go.dev/tour/basics/13
