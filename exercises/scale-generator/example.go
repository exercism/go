// Package scale provides a sample implementation
package scale

import "strings"

var chromaticScale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A",
	"A#", "B"}
var flatChromaticScale = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab",
	"A", "Bb", "B"}
var flatKeys = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb",
	"eb"}

// Scale returns a type of scale based on the inputs
func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = strings.Repeat("m", 12)
	}
	ft := formatTonic(tonic)
	// chromatic scale
	if flatKey(tonic, flatKeys) {
		start := findStart(ft, flatChromaticScale)
		return printScale(ft, interval, start, flatChromaticScale)
	}
	start := findStart(ft, chromaticScale)
	return printScale(ft, interval, start, chromaticScale)

}

func printScale(tonic, interval string, start int, arr []string) []string {
	var res []string
	res = append(res, tonic)
	for _, e := range interval[0 : len(interval)-1] {
		if e == 'm' {
			start++
		} else if e == 'M' {
			start += 2
		} else if e == 'A' {
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
	for i := range arr {
		if arr[i] == tonic {
			return true
		}
	}
	return false
}

func formatTonic(tonic string) string {
	if len(tonic) == 1 {
		return strings.ToUpper(tonic)
	}
	return strings.ToUpper(tonic[0:1]) + tonic[1:]
}
