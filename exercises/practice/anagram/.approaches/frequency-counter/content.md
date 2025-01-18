# Frequency Counter

```go
// Package anagram contains a solution to the anagram Exercism exercise.
package anagram

import (
	"strings"
	"unicode"
)

// Detect determines which words in cases are anagrams of the subject,
// ignoring words that are equal to subject (case-insensitive).
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
	if strings.ToLower(a) == strings.ToLower(b) {
		return false
	}

	freqCounter := make(map[rune]int)

	for _, r := range a {
		r = unicode.ToLower(r)
		freqCounter[r]++
	}

	for _, r := range b {
		r = unicode.ToLower(r)

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

This approach utilizes a frequency counter to determine if the strings are
anagrams of one another.

The [`strings.ToLower` function][strings.ToLower] is used to convert the strings
into their lowercase form where they are then checked for equality. Two strings
that are equal after converting to lowercase are not anagrams since they are the
same string.

A hash map of type `map[rune]int` is initialized as the frequency counter. The
first string is iterated over and the frequency counter is updated to hold the
number of occurances of each character in the string. Before a character is
counted in the frequency counter, it's converted to lowercase to account for
case-insensitive strings that should be anagrams (e.g., `foo` and `OOF`).

The second string is then iterated over and the frequency counter is checked to
see if the lowercase version of the character exists in the hash map. If it does
not then the strings cannot possibly be anagrams of one another and the
implementation returns early. Otherwise, the number of occurances of the current
character is decremented and, if the count of that character reaches 0, the
entry is removed from the hash map.

After iterating through both strings and updating the frequency counter the two
strings are anagrams if and only if the frequency counter is empty. That is, all
of the characters that were counted in the first string have been accounted for
when iterating through the second string. If the second string contained a
charater that the first string did not, then the implementation would return
early as described above. If the second string contained extra characters or
different characters than the first string then we'd have a non-empty hash map
left over at the end signifying a difference between the two strings.

This approach separates the anagram checking logic into its own function for
readability. That way, the `Detect` function can focus on iterating over the
candidates and building the resulting slice of anagrams.

[strings.ToLower]: https://pkg.go.dev/strings#ToLower
