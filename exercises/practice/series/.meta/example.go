package series

func All(n int, s string) (r []string) {
	if n <= 0 {
		return nil
	}
	for i := 0; n <= len(s); i++ {
		r = append(r, s[i:n])
		n++
	}
	return r
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
