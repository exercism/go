# Loop with a `switch`

```go
// Package binarysearch is a small library for finding an element in an ordered sequence
package binarysearch

// SearchInts searches for an element in an ordered sequence and returns its index or -1 if not found
func SearchInts(list []int, key int) int {
	for left, right := 0, len(list); left != right; {
		mid := (left + right) / 2

		switch {
		case list[mid] == key:
			return mid
		case key < list[mid]:
			right = mid
		case key > list[mid]:
			left = mid + 1
		}
	}
	return -1
}
```

This approach starts by initializing a [`for` loop][for-loop].
The `left` is initialized to `0` and the `right` is initialized to the length of the slice of `int`s passed in.
The loop keeps iterating while `left` does not equal `right`.

The middle value is set by `left` plus `right` divided by `2`.
For example, if `left` is `0` and `right` is `10`, then the middle is calculated to `5`.
if `left` is `6` and `right` is `10`, then the middle is calculated to `8`.

A [`switch` with no condition][switch-no-condition] is used to check the value of the element whose index in the slice of `int`s is the middle value.

- If the element at the index of the middle value matches the value being searched for, then the middle value is returned.

- If the value being searched for is less than the element at the index of the middle value, then `right` is set to the middle value
so that the next iteration will look at lower numbers.

- Otherwise, the value being searched is greater than the element at the index of the middle value, so `left` is set to the middle value
plus one so that the next iteration will look for higher numbers.

If `left` and `right` are changed during the iterations so that they equal other, then the value being searched for is not in the slice of `int`s.
The loop exits and `-1` is returned from the function.

[for-loop]: https://go.dev/tour/flowcontrol/1
[switch-no-condition]: https://go.dev/tour/flowcontrol/11
