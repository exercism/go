// Package change demonstrates making change with fewest number of coins.
package change

// Change returns a list of coins in increasing value order to make
// change for target amount from the coins given; an infinite supply of coins is assumed.
func Change(coins []int, target int) (result []int) {
	solutions := make([][]int, 0)
	if target == 0 {
		result = []int{}
	} else {
		if target > 0 {
			findSolutions(coins, target, &solutions)
		}
		if len(solutions) == 0 {
			result = []int{-1}
		} else if len(solutions) == 1 {
			result = solutions[0]
		} else {
			result = chooseBestSolution(solutions)
		}
	}
	return result
}

// chooseBestSolution returns the best(fewest coins) for a list of solutions.
func chooseBestSolution(solutions [][]int) []int {
	bestIndex := 0
	bestSize := len(solutions[0])
	for i := 1; i < len(solutions); i++ {
		if len(solutions[i]) < bestSize {
			bestSize = len(solutions[i])
			bestIndex = i
		}
	}
	return solutions[bestIndex]
}

// findSolutions attempts to find change solutions for target amount
// from the coins given.
func findSolutions(coins []int, target int, s *[][]int) {
	// Find the coins to consider, less than the target change amount.
	i := len(coins) - 1
	for ; i > 0; i-- {
		if coins[i] == target {
			// Perfect single coin solution, so done!
			*s = append(*s, []int{target})
			return
		}
		if coins[i] > target {
			// Cannot use this too large coin.
			continue
		}
		// Now, 0 .. i are the coins to consider.
		break
	}
	last := i
	// Consider the largest coin usable: coins[last]
	// Will a multiple of this coin meet the target?
	if target%coins[last] == 0 {
		// Yes. Best solution since coins[last] is largest coin.
		p := nCoins(target/coins[last], coins[last])
		*s = append(*s, p)
	} else {
		// Look for solution with mixture of coins.
		findMixedSolutions(coins[0:last+1], target, s)
	}
}

// findMixedSolutions attempts to find mixed coin amount solutions
// making correct change for the amount target.
func findMixedSolutions(coins []int, target int, s *[][]int) {
	// Consider all the coins in this outer loop.
	for i := len(coins) - 1; i >= 0; i-- {
		// Now loop to find a solution using only the last i+1 coins,
		// using the largest coins first.
		sum := 0
		coinlist := make([]int, 0)
		for j := i; j >= 0; j-- {
			// Compute remaining amount of the target for which we must make change.
			remaining := target - sum
			// Find n, which is how many of coin[j] can be used in change.
			n := remaining / coins[j]
			if remaining%coins[j] == 0 {
				// Using n coins[j] will complete the solution.
				coinlist = append(coinlist, nCoins(n, coins[j])...)
				// Reverse the coinlist to put into lowest to highest coin order.
				coinlist = reverseInts(coinlist)
				*s = append(*s, coinlist)
				break
			}
			// Use n coins as part of the coinlist solution and keep looping to use other lessor coins.
			coinlist = append(coinlist, nCoins(n, coins[j])...)
			// Update sum to track amount used by coinlist.
			sum += n * coins[j]
		}
	}
}

// nCoins returns a list(slice) of length n of a specific coin value.
func nCoins(n int, coin int) (r []int) {
	r = make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = coin
	}
	return r
}

// reverseInts reverses and returns the items in list a.
func reverseInts(a []int) []int {
	b := make([]int, len(a))
	j := 0
	for i := len(a) - 1; i >= 0; i-- {
		b[j] = a[i]
		j++
	}
	return b
}

const testVersion = 1
