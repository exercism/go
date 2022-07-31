// Package scale provides a sample implementation
package scale

import (
	"strings"
	"unicode"
)

// Title is a copy of strings.Title function of the stdlib.
// The copy is here because strings.Title is deprecated but we still
// want to use this function as the alternative would require us to support
// external dependencies which we don't yet (tracking issue https://github.com/exercism/go/issues/2379).
// Students should still be able to use strings.Title if they want.
func Title(s string) string {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

// Copy of strings.isSeparator function of the stdlib.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}

var chromaticScale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var flatChromaticScale = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var flatKeys = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

// Scale returns a type of scale based on the inputs
func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = strings.Repeat("m", 11)
	}
	ft := formatTonic(tonic)
	scale := chromaticScale
	if flatKey(tonic, flatKeys) {
		scale = flatChromaticScale
	}
	start := findStart(ft, scale)
	return printScale(ft, interval, start, scale)
}

func printScale(tonic, interval string, start int, arr []string) []string {
	res := []string{tonic}
	for _, e := range interval {
		switch e {
		case 'm':
			start++
		case 'M':
			start += 2
		case 'A':
			start += 3
		}
		res = append(res, arr[start%12])
	}
	return res
}

func findStart(tonic string, arr []string) int {
	for i := range arr {
		if arr[i] == tonic {
			return i
		}
	}
	return -1
}

func flatKey(tonic string, arr []string) bool {
	return findStart(tonic, arr) > -1
}

func formatTonic(tonic string) string {
	if len(tonic) == 1 {
		return strings.ToUpper(tonic)
	}
	return Title(tonic)
}
