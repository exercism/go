package bookstore

import (
	"math"
)

const bookPrice = 8

var discountTiers = [...]float64{0, 0.05, 0.1, 0.2, 0.25}

// Cost implements the book store exercise.
func Cost(books []int) float64 {
	return cost(books, 0)
}

func cost(books []int, priceSoFar float64) float64 {
	if len(books) == 0 {
		return priceSoFar
	}

	distinctBooks, remainingBooks := getDistinctBooks(books)
	minPrice := math.MaxFloat32

	for i := 1; i <= len(distinctBooks); i++ {
		newRemainingBooks := make([]int, len(remainingBooks))
		copy(newRemainingBooks, remainingBooks)
		newRemainingBooks = append(newRemainingBooks, distinctBooks[i:]...)

		price := cost(newRemainingBooks, priceSoFar+groupCost(i))
		minPrice = math.Min(minPrice, price)
	}

	return minPrice
}

func getDistinctBooks(books []int) (distinct []int, remaining []int) {
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

func groupCost(groupSize int) float64 {
	return float64(bookPrice*groupSize) * (1.00 - discountTiers[groupSize-1])
}
