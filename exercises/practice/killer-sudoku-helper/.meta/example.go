package killersudokuhelper

import "slices"

type cachedCombos struct {
	seen map[[3]int]struct {
		result [][]int
		ok     bool
	}
}

func (cc *cachedCombos) getCombos(numbers []int, target, size int) ([][]int, bool) {
	if numbers[0] > target {
		return nil, false
	}
	fp := [3]int{numbers[0], target, size}
	if r, ok := cc.seen[fp]; ok {
		return r.result, r.ok
	}
	if size == 1 {
		if slices.Contains(numbers, target) {
			result := [][]int{{target}}
			cc.seen[fp] = struct {
				result [][]int
				ok     bool
			}{result, true}
			return result, true
		}
		cc.seen[fp] = struct {
			result [][]int
			ok     bool
		}{nil, false}
		return nil, false
	}
	var result [][]int
	for i := range len(numbers) - size {
		n := numbers[i]
		if n > target {
			break
		}
		if subCombos, ok := cc.getCombos(numbers[i+1:], target-n, size-1); ok {
			for _, combo := range subCombos {
				if !slices.ContainsFunc(result, func(r []int) bool { return slices.Equal(r, combo) }) {
					result = append(result, append(combo, n))
				}
			}
		}
		if subCombos, ok := cc.getCombos(numbers[i+1:], target, size); ok {
			for _, combo := range subCombos {
				if !slices.ContainsFunc(result, func(r []int) bool { return slices.Equal(r, combo) }) {
					result = append(result, combo)
				}
			}
		}
	}
	cc.seen[fp] = struct {
		result [][]int
		ok     bool
	}{result, true}
	return result, true
}

func Combinations(sum, size int, exclude []int) [][]int {
	var includes []int
	for i := 1; i <= sum; i++ {
		if !slices.Contains(exclude, i) {
			includes = append(includes, i)
		}
	}
	cc := cachedCombos{make(map[[3]int]struct {
		result [][]int
		ok     bool
	})}
	combos, _ := cc.getCombos(includes, sum, size)

	var unique [][]int
	for _, combo := range combos {
		slices.Sort(combo)
		if !slices.ContainsFunc(unique, func(r []int) bool { return slices.Equal(r, combo) }) {
			unique = append(unique, combo)
		}
	}
	return unique
}
