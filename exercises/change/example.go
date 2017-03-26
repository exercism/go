// Package change demonstrates making change with fewest number of coins.
package change

import (
	"errors"
	"sort"
)

const testVersion = 1

var (
	ErrNoSolutionFound = errors.New("No solution")
	ErrNegativeTarget  = errors.New("Negative Target")
)

// Change returns a list of coins in increasing value order to make
// change for target amount from the coins given; an infinite supply of coins is assumed.
func Change(coins []int, target int) (result []int, err error) {
	var solution []int
	if target == 0 {
		// No change needed for 0.
		return []int{}, nil
	}
	if target < 0 {
		// No change possible for negative.
		return nil, ErrNegativeTarget
	}
	findSolution(coins, target, &solution)
	if len(solution) == 0 {
		// No solution possible with coins given.
		return nil, ErrNoSolutionFound
	}
	return solution, nil
}

// findSolution attempts to find a change solution for target amount
// from the coins given.
func findSolution(coins []int, target int, s *[]int) {
	// Find the coins to consider, less than the target change amount.
	last := len(coins) - 1
	for ; last >= 0; last-- {
		if coins[last] == target {
			// Perfect single coin solution, so done!
			*s = []int{target}
			return
		}
		if coins[last] > target {
			// Cannot use this too large coin.
			continue
		}
		// Now, 0 .. last are the coins to consider.
		break
	}
	if last < 0 {
		// All coins larger than target.
		return
	}
	findBestMixedSolution(coins[0:last+1], target, s)
}

// findBestMixedSolution attempts to find a best mixed coin amount solution
// making correct change for the amount target.
func findBestMixedSolution(coins []int, target int, s *[]int) {
	// Consider all combinations of the coins.
	// Eliminate the combination if it sums to value greater than target.
	n := len(coins)
	bestsize := target + 1 // an initial "worst" bestcase
comboLoop:
	for m := uint64((1 << uint(n)) - 1); m >= 1; m-- {
		// Make a list, c, of the coins matching combination m,
		// as long as the sum of coins in c doesn't exceed target.
		sum := 0
		c := make([]int, 0, n)
		for b := 0; b < n; b++ {
			// smaller b is larger coin
			if (m & (1 << uint(b))) != 0 {
				sum += coins[n-1-b]
				if sum > target {
					// Already sum exceeds target for this combination,
					// so consider it no more.
					continue comboLoop
				}
				c = append(c, coins[n-1-b])
			}
		}
		// Now the combination of coins in c can be considered.
		// The sum of coins in c doesn't exceed target.
		// When the sum matches the target, those coins are a solution.
		if sum == target {
			updateSolution(c, s, &bestsize)
			continue
		}
		// Recursively search for a best solution.
		partialSolution(nil, c, target, s, &bestsize)
	}
}

// updateSolution updates solution with newsolution if no prior solution
// or when length of newsolution is better.
func updateSolution(newsolution []int, solution *[]int, bestsize *int) {
	if len(*solution) == 0 || len(newsolution) < len(*solution) {
		*bestsize = len(newsolution)
		*solution = make([]int, len(newsolution))
		copy(*solution, newsolution)
		sort.Ints(*solution)
	}
}

// partialSolution takes a partial solution, and coin list c looking to meet target
// while updating solution s to include partial plus the coins meeting target.
func partialSolution(partial []int, c []int, target int, s *[]int, bestsize *int) {
	// We want to consider coins in c in largest to smallest order.
	sort.Sort(sort.Reverse(sort.IntSlice(c)))

	for e := range c {
		// Find n, which is highest number of c[e] that could be used in change.
		n := target / c[e]

		// Try n, n-1 ... 1
		for k := n; k >= 1; k-- {
			if (k * c[e]) == target {
				// k coins of c[e] completes a solution.
				if (len(partial) + k) < *bestsize {
					list := append(partial, nCoins(k, c[e])...)
					updateSolution(list, s, bestsize)
				}
				// No need to consider smaller k for c[e], so exit this loop.
				break
			}
			if len(c) > 1 {
				// k coins of *c[e] might be potential partial solution.
				newtarget := target - (k * c[e])
				// Use a heuristic estimate of whether a partial solution
				// might be worth it: if the larget coin in c not yet used
				// in an already found solution allows the new solution
				// to be smaller, then it would be worth pursuing.

				// Find the largest coin in c not yet used in *s solution
				// omitting c[e], since it is used in this partial solution.
				largest := largestNotUsed(c, *s, c[e])
				// It will take at least the partial, plus k coins plus 1 more.
				if (len(partial) + k + 1) < *bestsize {
					// Calculate the minimum potential size of a solution
					// which doesn't include coin c[e].
					potentialSize := len(partial) + k + (newtarget / largest)
					if (newtarget % largest) != 0 {
						potentialSize++
					}
					if potentialSize < *bestsize {
						// Put k instances of c[e] in a new partial solution.
						newPartial := append(partial, nCoins(k, c[e])...)
						d := make([]int, len(c))
						copy(d, c)
						// Remove e'th coin since it is in the partial solution.
						d = append(d[:e], d[e+1:]...)
						// Now look for solution meeting *new* target.
						partialSolution(newPartial, d, newtarget, s, bestsize)
					}
				}
			}
		}
	}
}

// Find the largest coin in coins not yet used in solution, not including coin 'omit'.
// The coins are in largest to smallest order.
func largestNotUsed(coins []int, solution []int, omit int) int {
coinLoop:
	for _, coin := range coins {
		if coin == omit {
			continue
		}
		for _, used := range solution {
			if coin == used {
				continue coinLoop
			}
		}
		return coin
	}
	// Should not get here. But unity is safe return value.
	return 1
}

// nCoins returns a list(slice) of length n of a specific coin value.
func nCoins(n int, coin int) (r []int) {
	r = make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = coin
	}
	return r
}
