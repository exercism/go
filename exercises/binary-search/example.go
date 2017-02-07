package binarysearch

import "fmt"

const testVersion = 1

func SearchInts(s []int, k int) (i int) {
	for j := len(s); i < j; {
		if h := (i + j) / 2; s[h] < k {
			i = h + 1
		} else {
			j = h
		}
	}
	return
}

func Message(s []int, k int) string {
	x := SearchInts(s, k)
	switch {
	case x == len(s):
		if x == 0 {
			return "slice has no values"
		}
		return fmt.Sprintf("%d > all %d values", k, x)
	case s[x] != k:
		if x == 0 {
			return fmt.Sprintf("%d < all values", k)
		}
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d",
			k, s[x-1], x-1, s[x], x)
	}
	switch x {
	case 0:
		return fmt.Sprintf("%d found at beginning of slice", k)
	case len(s) - 1:
		return fmt.Sprintf("%d found at end of slice", k)
	}
	return fmt.Sprintf("%d found at index %d", k, x)
}
