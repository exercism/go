package pascal

func Triangle(n int) (t [][]int) {
	if n < 1 {
		return
	}
	t = make([][]int, n)
	r := []int{1}
	t[0] = r
	for i := 1; i < n; i++ {
		last := r
		r = make([]int, i+1)
		r[0] = 1
		r[i] = 1
		for j := 1; j < i; j++ {
			r[j] = last[j-1] + last[j]
		}
		t[i] = r
	}
	return t
}
