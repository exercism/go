// Package change demonstrates making change with fewest number of coins.
package change

import (
	"errors"
	"sort"
)

var (
	ErrNoSolutionFound = errors.New("No solution")
	ErrNegativeTarget  = errors.New("Negative Target")
)

// Change returns a list of coins in increasing value order to make
// change for target amount from the coins given; an infinite supply of coins is assumed.
func Change(coins []int, target int) ([]int, error) {
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

	bestsize := target + 1 // an initial "worst" bestsize
	findUsingPartialSolution(nil, coins[0:last+1], target, s, &bestsize)
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

// findUsingPartialSolution takes a partial solution, and coin list c looking to meet target
// while updating solution s to include partial plus the coins meeting target.
func findUsingPartialSolution(partial, c []int, target int, s *[]int, bestsize *int) {
	// We want to consider coins in c in largest to smallest order.
	sort.Sort(sort.Reverse(sort.IntSlice(c)))

	for e, coin := range c {
		// Find n, which is highest number of coin that could be used in change.
		n := target / coin

		// Try n, n-1 ... 1
		for k := n; k >= 1; k-- {
			if (k * coin) == target {
				// k coin(s) completes a solution.
				if (len(partial) + k) < *bestsize {
					list := partial
					list = append(list, nCoins(k, coin)...)
					updateSolution(list, s, bestsize)
				}
				// No need to consider smaller k for coin, so exit this loop.
				break
			}
			if len(c) > 1 {
				// k coin(s) might be potential partial solution.
				newtarget := target - (k * coin)
				// Use a heuristic estimate of whether a partial solution
				// might be worth it: if the larget coin in c not yet used
				// in an already found solution allows the new solution
				// to be smaller, then it would be worth pursuing.

				// Find the largest coin in c not yet used in *s solution
				// omitting coin, since it is used in this partial solution.
				largest := largestNotUsed(c, *s, coin)
				// It will take at least the partial, plus k coins plus 1 more.
				if (len(partial) + k + 1) < *bestsize {
					// Calculate the minimum potential size of a solution
					// which doesn't include this coin.
					potentialSize := len(partial) + k + (newtarget / largest)
					if (newtarget % largest) != 0 {
						potentialSize++
					}
					if potentialSize < *bestsize {
						// Put k instances of coin in a new partial solution.
						newPartial := partial
						newPartial = append(newPartial, nCoins(k, coin)...)
						d := append(append([]int{}, c[:e]...), c[e+1:]...)
						// Now look for solution meeting *new* target.
						findUsingPartialSolution(newPartial, d, newtarget, s, bestsize)
					}
				}
			}
		}
	}
}

// Find the largest coin in coins not yet used in solution, not including coin 'omit'.
// The coins are in largest to smallest order.
func largestNotUsed(coins, solution []int, omit int) int {
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
func nCoins(n, coin int) []int {
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		coins[i] = coin
	}
	return coins
}
