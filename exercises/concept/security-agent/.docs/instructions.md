# Instructions

You work for an agency that is responsible for information security.
Your job includes writing functions that will be useful for verifying and improving security.

## 1. Generate important number

In the agency, _important_ numbers are frequently used.
An _important_ number is defined as a number obtained in a special way: given two `int`, `num` and `power`, we multiply `num` by 2 raised to the power of `power` to obtain the _important_ number.
These numbers are often used for various data manipulations, encoding and decoding.
Your task is to write a function to generate such _important_ number.

```go
ImportantNumber(3, 7) // => 3 * (2 ^ 7) = 384
```

~~~~exercism/note
Do not use the `math` package or implement this through direct multiplication.
~~~~

## 2. Integrity Check

During data transmission, the agency transfers `int` type numbers, which are stored in a special slice.
If integrity is preserved, then each number in the slice should have a matching pair.
Write a function that checks this.
The function should return `true` if integrity is preserved, and `false` otherwise.

```go
nums1 := []int{1, 1, 2, 2, 7, -5, 7, -5}
nums2 := []int{4, 5, 4, 6, 7, 6, 7}
nums3 := []int{4, 4, 4, 5, 5, 4}

CheckIntegrity(nums1) // => true
CheckIntegrity(nums2) // => false
CheckIntegrity(nums3) // => true (There may be more than one pair of a particular number)
```

~~~~exercism/note
Having an even number of elements does not imply pairs, e.g., `{1, 2, 3, 4}`.
~~~~

## 3. Setting Flags

In programming, flags are sometimes used to set specific parameters.
To combine several flags, bits of a number can be used.
Your management wants you to write a function that sets the second and fourth bits of a number to `0` (`false`).
Bits are counted from the _right_.
You don't know what these bits are used for, because this information is classified from you.

```go
var a uint8 = 255 // 1111 1111
var b uint8 = 151 // 1001 0111

SetFlags(a) // => 1111 0101
SetFlags(b) // => 1001 0101
```

## 4. Generate Key

In information security, [hash][hash-link] functions are used to store sensitive data.
These are functions that cannot be reversed to retrieve the original input.
Your agency wants to have its own such function.
It should generate a hash based on a given number `num` of type `uint8` and return `uint32`.
The following algorithm was devised to create this hash (bytes are counted from _left to right_):

- First byte: `num / 13`
- Second byte: `num + 113`
- Third byte: `num * 7`
- Fourth byte: `num - 79`

As you may have noticed, people in information security have a fondness for simple numbers.

```go
var num uint32 = 9

GenerateHash(num) // => 8011706 (Yes, num is large)
```

~~~~exercism/note
Be careful with the types of numbers. You need to apply the operation and then get the desired byte.
~~~~

[hash-link]: https://en.wikipedia.org/wiki/Hash_function
