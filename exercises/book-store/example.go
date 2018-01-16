package bookstore

import "math"

const bookPrice = 800 // 800 cents = $8.00

var discountTiers = [...]int{0, 5, 10, 20, 25}

// Cost implements the book store exercise.
func Cost(books []int) int {
	return cost(books, 0)
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

func groupCost(groupSize int) int {
	normalPrice := bookPrice * groupSize
	discount := (normalPrice * discountTiers[groupSize-1]) / 100
	return normalPrice - discount
}
