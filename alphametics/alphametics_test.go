package alphametics

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Define a function Solve(a, b, sum string) map[rune]int.
//
// Also define an exported TestVersion with a value that matches
// the internal testVersion here.

const testVersion = 1

var testCases = []struct {
	a, b, sum string
	solution  bool
}{
	{
		// The classic
		a: "send", b: "more", sum: "money",
		solution: true,
	},
	{
		a: "stone", b: "rock", sum: "cavern",
		solution: true,
	},
	{
		a: "math", b: "code", sum: "same",
		solution: true,
	},
	{
		a: "yes", b: "no", sum: "maybe",
		solution: false,
	},
	{
		a: "zeroes", b: "one", sum: "binary",
		solution: false,
	},
	{
		a: "one", b: "zeroes", sum: "binary",
		solution: false,
	},
}

// Without this you'd get output like `map[101:3 109:0 115:8]`
func prettyMap(rm map[rune]int) string {
	runes := make([]int, 0, len(rm))
	for k := range rm {
		runes = append(runes, int(k))
	}
	// Some type juggling is needed to use this
	sort.Ints(runes)
	var buf bytes.Buffer
	buf.WriteRune('[')
	for _, r := range runes {
		r := rune(r)
		buf.WriteRune(r)
		buf.WriteRune(':')
		buf.WriteString(strconv.Itoa(rm[r]))
		buf.WriteRune(' ')
	}
	if len(runes) > 0 {
		buf.Truncate(buf.Len() - 1) // Remove last ' '
	}
	buf.WriteRune(']')
	return buf.String()
}

func prettyEquation(a, b, sum string, digits map[rune]int) string {
	s := fmt.Sprintf("% 10s\n% 10s +\n----------\n% 10s", a, b, sum)
	for k, v := range digits {
		s = strings.Replace(s, string(k), strconv.Itoa(v), -1)
	}
	return s
}

func isCorrect(a, b, sum string, digits map[rune]int) bool {
	for k, v := range digits {
		a = strings.Replace(a, string(k), strconv.Itoa(v), -1)
		b = strings.Replace(b, string(k), strconv.Itoa(v), -1)
		sum = strings.Replace(sum, string(k), strconv.Itoa(v), -1)
	}
	// Starting a number with a zero isn't allowed.
	if a[0] == '0' || b[0] == '0' || sum[0] == '0' {
		return false
	}
	aInt, err := strconv.ParseInt(a, 10, 32)
	if err != nil {
		return false
	}
	bInt, err := strconv.ParseInt(b, 10, 32)
	if err != nil {
		return false
	}
	sumInt, err := strconv.ParseInt(sum, 10, 32)
	if err != nil {
		return false
	}
	return aInt+bInt == sumInt
}

func TestSolve(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, tt := range testCases {
		actual := Solve(tt.a, tt.b, tt.sum)
		if tt.solution {
			if actual == nil {
				t.Fatalf("Got no solution for solvable %s+%s==%s", tt.a, tt.b, tt.sum)
			}
			if !isCorrect(tt.a, tt.b, tt.sum, actual) {
				t.Fatalf("Incoherent answer for %s+%s==%q: got %v; in equation form:\n%s",
					tt.a, tt.b, tt.sum, prettyMap(actual),
					prettyEquation(tt.a, tt.b, tt.sum, actual))
			}
		} else {
			if actual != nil {
				t.Fatalf("Got solution for unsolvable %s+%s==%q: got %v; in equation form:\n%s",
					tt.a, tt.b, tt.sum, prettyMap(actual),
					prettyEquation(tt.a, tt.b, tt.sum, actual))
			}
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Solve(tt.a, tt.b, tt.sum)
		}
	}
}
