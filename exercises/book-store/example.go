package bookstore

import (
	"math"
	"sort"
)

const bookPrice = 800 // 800 cents = $8.00

var discountTiers = [...]int{0, 5, 10, 20, 25}

// Cost implements the book store exercise.
func Cost(books []int) int {
	organize(books)
	return cost(books, 0)
}

// rework the input array so all the repetitions
// are together at first, ordering by most repeated
// e.g. [1,1,2,3,4,4,5,4,2] -> [4,4,4,1,1,2,2,3,5]
func organize(books []int) {
	// used for sorting
	type kv struct {
		Key   int
		Value int
	}

	// calc book frequency: how many 1's ,2's and so on
	freq := make(map[int]int)
	for i := range books {
		freq[books[i]]++
	}
	// sort frequency in descending order
	ss := make([]kv, len(freq))
	for k, v := range freq {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	// transform the frequencies back to repetitions
	// e.g. 4*1, 3*2 -> 1,1,1,1,2,2,2
	p := 0
	for _, kv := range ss {
		for i := 0; i < kv.Value; i++ {
			books[p] = kv.Key
			p++
		}
	}
	// give back the modified array
}

func cost(books []int, priceSoFar int) int {
	if len(books) == 0 {
		return priceSoFar
	}

	distinctBooks, remainingBooks := getDistinctBooks(books)
	minPrice := math.MaxInt32

	for i := 1; i <= len(distinctBooks); i++ {
		newRemainingBooks := make([]int, len(remainingBooks))
		copy(newRemainingBooks, remainingBooks)
		newRemainingBooks = append(newRemainingBooks, distinctBooks[i:]...)

		price := cost(newRemainingBooks, priceSoFar+groupCost(i))
		if price < minPrice {
			minPrice = price
		}
	}

	return minPrice
}

func getDistinctBooks(books []int) (distinct, remaining []int) {
	exists := make(map[int]bool)
	for _, book := range books {
		if exists[book] {
			remaining = append(remaining, book)
		} else {
			distinct = append(distinct, book)
			exists[book] = true
		}
	}

	return
}

func groupCost(groupSize int) int {
	normalPrice := bookPrice * groupSize
	discount := (normalPrice * discountTiers[groupSize-1]) / 100
	return normalPrice - discount
}
