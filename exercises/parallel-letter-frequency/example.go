package letter

func ConcurrentFrequency(l []string) FreqMap {
	switch len(l) {
	case 0:
		return FreqMap{}
	case 1:
		return Frequency(l[0])
	}
	ch := make(chan FreqMap)
	f := func(l []string) {
		ch <- ConcurrentFrequency(l)
	}
	half := len(l) / 2
	go f(l[:half])
	go f(l[half:])
	m := <-ch
	for r, n := range <-ch {
		m[r] += n
	}
	return m
}
