package binarysearch

// SearchInts search given key in the given slice
// upon finding returns its index, otherwise -1
func SearchInts(s []int, k int) int {
	for i, j := 0, len(s)-1; i <= j; {
		if h := (i + j) / 2; s[h] < k {
			i = h + 1
		} else if s[h] > k {
			j = h - 1
		} else {
			return h
		}
	}
	return -1
}
