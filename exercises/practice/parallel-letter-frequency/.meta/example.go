package parallelletterfrequency

import (
	"slices"
	"strings"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

var Ignore = []rune{' ', '\t', '\n', '\r', '!', '?', '.', ',', ':', '-', ';', '"', '\'', '(', ')'}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range strings.ToLower(text) {
		if r >= '0' && r <= '9' || slices.Contains(Ignore, r) {
			continue
		}
		frequencies[r]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) FreqMap {
	res := make(FreqMap)
	ch := make(chan FreqMap, 10)

	var wg sync.WaitGroup
	wg.Add(len(texts))

	for _, text := range texts {
		go func(t string) {
			ch <- Frequency(t)
			wg.Done()
		}(text)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for freqmap := range ch {
		for letter, freq := range freqmap {
			res[letter] += freq
		}
	}

	return res
}
