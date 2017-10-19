package bookstore

import (
	"math"
)

const bookPrice = 8

var discountTiers = [...]float64{0, 0.05, 0.1, 0.2, 0.25}

// CalculateBasketCost implements the book store exercise.
func CalculateBasketCost(books []int) float64 {
	return calculateBasketCost(books, 0.0)
}

func calculateBasketCost(books []int, priceSoFar float64) float64 {
	if len(books) == 0 {
		return priceSoFar
	}

	distinctBooks, remainingBooks := getDistinctBooks(books)
	minPrice := math.MaxFloat32

	for i := 1; i <= len(distinctBooks); i++ {
		newRemainingBooks := make([]int, len(remainingBooks))
		copy(newRemainingBooks, remainingBooks)
		newRemainingBooks = append(newRemainingBooks, distinctBooks[i:]...)

		price := calculateBasketCost(newRemainingBooks, priceSoFar+costOfGroupSize(i))
		minPrice = math.Min(minPrice, price)
	}

	return minPrice
}

func getDistinctBooks(books []int) (distinct []int, remaining []int) {
	distinctMap := make(map[int]bool)
	for _, book := range books {
		if exists := distinctMap[book]; exists {
			remaining = append(remaining, book)
		} else {
			distinct = append(distinct, book)
			distinctMap[book] = true
		}
	}

	return distinct, remaining
}

func costOfGroupSize(groupSize int) float64 {
	return float64(bookPrice*groupSize) * (1.00 - discountTiers[groupSize-1])
}
