package binarysearch

// SearchInts search given key in the given slice
// upon finding returns its index, otherwise -1
func SearchInts(s []int, k int) int {
	for i, j := 0, len(s)-1; i <= j; {
		h := (i + j) / 2
		switch {
		case s[h] < k:
			i = h + 1
		case s[h] > k:
			j = h - 1
		default:
			return h
		}
	}
	return -1
}
