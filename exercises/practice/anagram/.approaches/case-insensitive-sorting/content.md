# Case-insensitive Sorting

```go
// Package anagram contains a solution to the anagram Exercism exercise.
package anagram

import (
	"sort"
	"strings"
)

// Detect determines which words in candidates are anagrams of the subject.
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

This approach normalizes both strings to lowercase, sorts them, and compares the
resulting sorted strings to determine if they are anagrams.

The [`strings.ToLower` function][strings.ToLower] is used to convert the strings
into their lowercase form. This normalizes the strings so that the solution is
case-insensitive (e.g., `foo` and `OOF` normalize to `foo` and `oof`). At this
point if the two strings normalize to the same string, they cannot be anagrams
since a word cannot be an anagram of itself.

The normalized strings are then sorted using the [`sort` package][sort]. It
doesn't matter if the strings are sorted in non-decreasing or non-increasing
order as long as both strings are sorted in the same way. Sorting the strings
allows us to determine if the strings are anagrams (e.g., `foo` and `oof` sort
to `foo` and `foo`). In Go, sorting operates on slices, not strings, so the
string is first split into a slice using `strings.Split`, sorted using
`sort.Strings`, and then joined back into a string using `strings.Join`.

Now that the strings are normalized and sorted, they can be compared directly to
determine if they are anagrams. If the strings are the same, they are anagrams.
Otherwise, they are not anagrams.

This approach separates the sorting and anagram checking logic into their own
functions for readability. That way, the `Detect` function can focus on
iterating over the candidates and building the resulting slice of anagrams.

[strings.ToLower]: https://pkg.go.dev/strings#ToLower
[sort]: https://pkg.go.dev/sort
