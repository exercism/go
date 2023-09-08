# Introduction

There are many variants of a general way to solve Binary Search.
A general approach is to loop and use a switch to iterate toward a solution.

## Approach: Loop with a `switch`

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
		default:
			left = mid + 1
		}
	}
	return -1
}
```

For more information, check the [Loop with `switch` approach][approach-loop-with-switch].

[approach-loop-with-switch]:  https://exercism.org/tracks/go/exercises/binary-search/approaches/loop-with-switch
