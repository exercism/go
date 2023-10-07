# Introduction

There are several ways to solve Anagram. One approach is to convert both strings
to lowercase, sort them, and compare them. Another approach is to build a
frequency counter of the number of characters in each string and then ensure
both frequency counters have the same keys and values.

## General guidance

Consider the following when choosing an approach.

- Can you use additional memory?
  - Yes. The frequency counter approach has a faster running time at the expense
    of using additional memory.
  - No. The case-insensitive sorting approach has a slower running time but it
    uses no additional memory.

## Approach: Case-insensitive Sorting

```go
// Package anagram contains a solution to the anagram Exercism exercise.
package anagram

import (
	"sort"
	"strings"
)

// Detect determines which words in cases are anagrams of the subject.
func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0)

	subject = strings.ToLower(subject)

	for _, candidate := range candidates {
		c := strings.ToLower(candidate)

		if isAnagram(subject, c) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

// isAnagram determines whether a and b are anagrams of each other.
func isAnagram(a, b string) bool {
	return a != b && sortString(a) == sortString(b)
}

// sortString sorts a string lexicographically in non-decreasing order.
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
```

For more information, check the [case-insensitive sorting approach][approach-case-insensitive-sorting].

## Approach: Frequency Counter

```go
// Package anagram contains a solution to the anagram Exercism exercise.
package anagram

import (
	"strings"
)

// Detect determines which words in cases are anagrams of the subject.
func Detect(subject string, cases []string) []string {
	anagrams := make([]string, 0)

	for _, word := range cases {
		if isAnagram(subject, word) {
			anagrams = append(anagrams, word)
		}
	}

	return anagrams
}

// isAnagram determines whether a and b are anagrams of each other.
func isAnagram(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if a == b {
		return false
	}

	freqCounter := make(map[rune]int)

	for _, r := range a {
		freqCounter[r]++
	}

	for _, r := range b {
		if _, ok := freqCounter[r]; !ok {
			return false
		}

		freqCounter[r]--

		if freqCounter[r] == 0 {
			delete(freqCounter, r)
		}
	}

	return len(freqCounter) == 0
}

```

For more information, check the [frequency counter approach][approach-frequency-counter].

[approach-case-insensitive-sorting]: https://exercism.org/tracks/go/exercises/anagram/approaches/case-insensitive-sorting
[approach-frequency-counter]: https://exercism.org/tracks/go/exercises/anagram/approaches/frequency-counter
