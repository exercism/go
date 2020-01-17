package alphametics

import (
	"errors"
	"strings"
	"unicode"
)

var decDigits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

type problem struct {
	vDigits      [][]rune
	maxDigits    int
	letterValues [26]int
	lettersUsed  []rune
	nLetters     int
}

func Solve(puzzle string) (map[string]int, error) {
	p := parsePuzzle(puzzle)
	if p == nil {
		return nil, errors.New("invalid puzzle")
	}
	return p.solvePuzzle()
}

// parsePuzzle parses the puzzle input into a problem for solving.
func parsePuzzle(puzzle string) (p *problem) {
	var valueStrings []string
	p = new(problem)
	fields := strings.Fields(puzzle)
	for _, field := range fields {
		if field == "+" || field == "==" {
			continue
		} else {
			valueStrings = append(valueStrings, field)
			if len(field) > p.maxDigits {
				p.maxDigits = len(field)
			}
			for _, r := range field {
				if !unicode.IsUpper(r) {
					return nil
				}
				v := r - 'A'
				p.letterValues[v] = -1
			}
		}
	}
	// Count the letters used.
	for v := 0; v < len(p.letterValues); v++ {
		if p.letterValues[v] == -1 {
			p.nLetters++
		}
	}
	p.vDigits = make([][]rune, len(valueStrings))
	for i := range valueStrings {
		p.vDigits[i] = make([]rune, p.maxDigits)
		for d, r := range valueStrings[i] {
			j := len(valueStrings[i]) - 1 - d
			// Each vDigits value is 0 for nothing in this position,
			// or 1..26 for 'A'..'Z'.
			p.vDigits[i][j] = r - 'A' + 1
		}
	}
	// Make a list of the letters used, where 0 == 'A'
	p.lettersUsed = make([]rune, p.nLetters)
	for n, v := 0, 0; v < len(p.letterValues); v++ {
		if p.letterValues[v] == -1 {
			p.lettersUsed[n] = rune(v) // 0 == 'A'
			n++
		}
	}
	return p
}

func (p *problem) solvePuzzle() (map[string]int, error) {
	for _, digValues := range permutations(decDigits, p.nLetters) {
		if p.isPuzzleSolution(digValues) {
			// Check leading digit of answer for 0 (invalid solution).
			r := p.vDigits[len(p.vDigits)-1][p.maxDigits-1]
			if p.letterValues[r-1] == 0 {
				continue
			}
			return p.puzzleMap(), nil
		}
	}
	return nil, errors.New("no solution")
}

// isPuzzleSolution returns true if the values work out as a solution.
func (p *problem) isPuzzleSolution(values []int) bool {
	// Put the candidate values into the letterValues for the lettersUsed.
	for i, r := range p.lettersUsed {
		p.letterValues[r] = values[i]
	}
	// For each column up to maxDigits
	// check that the sum of the digits corresponding to values in vDigits
	// add up to the digits (modulo 10) of values in last row.
	carry := 0
	for d := 0; d < p.maxDigits; d++ {
		// Get initial sum for the digits for the this column of vDigits
		r := p.vDigits[0][d]
		sum := carry
		if r != 0 {
			// There's a character in this position.
			sum += p.letterValues[r-1]
		}
		// Sum remaining rows for this column.
		for n := 1; n < len(p.vDigits)-1; n++ {
			r = p.vDigits[n][d]
			if r != 0 {
				// There's a character in this position.
				sum += p.letterValues[r-1]
			}
		}
		carry = sum / 10
		sum %= 10
		// Check the result sum against the answer row digit.
		r = p.vDigits[len(p.vDigits)-1][d]
		if r == 0 || sum != p.letterValues[r-1] {
			return false
		}
	}
	return true
}

// puzzleMap creates a "by letter" map from the letterValues used.
func (p *problem) puzzleMap() map[string]int {
	pm := make(map[string]int, p.nLetters)
	for _, v := range p.lettersUsed {
		r := v + 'A'
		s := string(r)
		pm[s] = p.letterValues[v]
	}
	return pm
}

// permutations returns a slice containing the r length permutations of the elements in iterable.
// The implementation is modeled after the Python itertools.permutations().
func permutations(iterable []int, r int) (perms [][]int) {
	pool := iterable
	n := len(pool)
	nperm := 1
	for i := n; i > 1; i-- {
		nperm *= i
	}
	if r < n {
		d := 1
		for i := (n - r); i > 1; i-- {
			d *= i
		}
		nperm /= d
	}
	perms = make([][]int, 0, nperm)

	if r > n {
		return
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	p := make([]int, len(result))
	copy(p, result)
	perms = append(perms, p)

	for n > 0 {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				for k := i; k < r; k++ {
					result[k] = pool[indices[k]]
				}

				p = make([]int, len(result))
				copy(p, result)
				perms = append(perms, p)

				break
			}
		}

		if i < 0 {
			return
		}

	}
	return
}
