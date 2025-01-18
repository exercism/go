package dominoes

import (
	"errors"
	"reflect"
	"sort"
	"testing"
)

func TestMakeChain(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			c, ok := MakeChain(tc.dominoes)
			if ok != tc.valid {
				t.Fatalf("MakeChain(%v)\nexpected 'ok' result: %t, actual 'ok': %t", tc.dominoes, tc.valid, ok)
			}
			if ok {
				// There can be a variety of "valid" chains. Verify the chain is valid.
				if err := verifyChain(tc.dominoes, c); err != nil {
					t.Fatalf("MakeChain(%v)\nverifying chain failed with error: %v\nchain: %v", tc.dominoes, err, c)
				}
			}
		})
	}
}

var (
	errWrongLengthChain          = errors.New("wrong length chain")
	errChainIsNotLegalAdj        = errors.New("chain is not legal - adjacent mismatch")
	errChainIsNotLegalEnd        = errors.New("chain is not legal - ends mismatch")
	errChainSetNotSameAsInputSet = errors.New("chain dominoes not same as input")
)

func verifyChain(input, chain []Domino) error {
	if len(input) != len(chain) {
		return errWrongLengthChain
	}

	switch len(input) {
	case 0:
		return nil
	case 1:
		if input[0] != chain[0] {
			return errChainSetNotSameAsInputSet
		}
		return nil
	}

	// Check adjacent pairs.
	for i := 0; i < len(chain)-1; i++ {
		if chain[i][1] != chain[i+1][0] {
			return errChainIsNotLegalAdj
		}
	}
	// Check end dominoes.
	if chain[0][0] != chain[len(chain)-1][1] {
		return errChainIsNotLegalEnd
	}

	// Make copies of input and chain.
	cinput := copyDominoes(input)
	cchain := copyDominoes(chain)

	sortDominoes(cinput)
	sortDominoes(cchain)

	// Compare for equality (same set in input and chain).
	if !reflect.DeepEqual(cinput, cchain) {
		return errChainSetNotSameAsInputSet
	}
	return nil
}

func copyDominoes(d []Domino) (c []Domino) {
	c = make([]Domino, len(d))
	// Put each domino in "canonical position" [a,b] where a <= b.
	for i := range d {
		c[i] = d[i]
		if c[i][0] > c[i][1] {
			c[i][0], c[i][1] = c[i][1], c[i][0]
		}
	}
	return c
}

func sortDominoes(d []Domino) {
	sort.Slice(d,
		func(i, j int) bool {
			if d[i][0] < d[j][0] {
				return true
			}
			if d[i][0] > d[j][0] {
				return false
			}
			return d[i][1] < d[j][1]
		})
}

func BenchmarkMakeChain(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			MakeChain(test.dominoes)
		}
	}
}
