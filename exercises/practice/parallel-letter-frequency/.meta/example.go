package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency is a (not necessarily ideal) example of how to call Frequency() concurrently.
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
