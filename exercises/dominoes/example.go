package dominoes

type Dominoe [2]int

// MakeChain returns a chain and true if the dominoes in input
// can be arranged in a legal chain; otherwise it returns nil, false.
func MakeChain(input []Dominoe) (chain []Dominoe, ok bool) {
	switch len(input) {
	case 0:
		return []Dominoe{}, true
	case 1:
		// A single is legal if the ends match.
		if input[0][0] == input[0][1] {
			return input, true
		}
		return nil, false
	}
	chain, ok = buildDominoeChain(input)
	if !ok {
		return nil, false
	}
	return chain, true
}

// buildDominoeChain looks for a matching dominoe chain using the 2 or more dominoes in d.
func buildDominoeChain(d []Dominoe) ([]Dominoe, bool) {
	permuteList := dominoePermutations(d, len(d))
	for _, p := range permuteList {
		chain, ok := arrangeChain(p)
		if ok {
			return chain, true
		}
	}
	return nil, false
}

// arrangeChain uses the dominoes positionally in d[] attempting to make a matching chain;
// reversing of sides of any dominoe is permitted, but the order of dominoes in d remains
// the same left to right.
func arrangeChain(d []Dominoe) (chain []Dominoe, ok bool) {
	// A good chain will match length of d, so preallocate the chain.
	chain = make([]Dominoe, len(d))
	n := 0
	// Put first dominoe into the chain for starting point.
	chain[n] = d[0]
	// Loop through the remaining dominoes in d, attempting to add them into the chain,
	// allowing a reverse of either single one in chain or the new dominoe t at each step.
	for i := 1; i < len(d); i++ {
		t := d[i]
		last := chain[n]
		if n == 0 && last[0] == t[0] {
			// reverse first and only item in chain to match t, and add t to chain.
			chain[n] = reverseDominoe(last)
			chain[n+1] = t
		} else if n == 0 && last[0] == t[1] {
			// reverse only item in chain and t to match, and add reversed t to chain.
			chain[n] = reverseDominoe(last)
			chain[n+1] = reverseDominoe(t)
		} else if last[1] == t[0] {
			// add t as-is into chain.
			chain[n+1] = t
		} else if last[1] == t[1] {
			// reverse t to match last one in chain, and add reversed t to chain.
			chain[n+1] = reverseDominoe(t)
		} else {
			// no match for this chain configuration of d, even with swapping.
			return nil, false
		}
		// chain is now filled in with one more dominoe.
		n++
	}
	// Both ends of the chain must also match.
	if chain[0][0] != chain[len(chain)-1][1] {
		return nil, false
	}
	return chain, true
}

// reverseDominoe returns given Dominoe x, with its sides reversed/rotated.
func reverseDominoe(x Dominoe) Dominoe {
	return Dominoe{x[1], x[0]}
}

// dominoePermutations returns a slice containing the r length permutations of the elements in iterable.
// The implementation is modeled after the Python itertools.permutations().
func dominoePermutations(iterable []Dominoe, r int) (perms [][]Dominoe) {
	pool := iterable
	n := len(pool)

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

	result := make([]Dominoe, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	p := make([]Dominoe, len(result))
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

				p = make([]Dominoe, len(result))
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
