package series

func All(n int, s string) (r []string) {
	for i := 0; n <= len(s); i++ {
		r = append(r, s[i:n])
		n++
	}
	return
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
