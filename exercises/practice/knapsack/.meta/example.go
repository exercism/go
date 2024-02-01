package knapsack

import "math"

type item struct {
	Weight, Value int
}

func Knapsack(maximumWeight int, items []item) int {
	amountOfItems := len(items)
	knapsack := make([][]int, amountOfItems+1)

	for i := range knapsack {
		knapsack[i] = make([]int, maximumWeight+1)
	}

	for currentItem := 1; currentItem <= amountOfItems; currentItem++ {
		for currentItemWeight := 0; currentItemWeight <= maximumWeight; currentItemWeight++ {
			if items[currentItem-1].Weight > currentItemWeight {
				knapsack[currentItem][currentItemWeight] = knapsack[currentItem-1][currentItemWeight]
			} else {
				knapsack[currentItem][currentItemWeight] = int(math.Max(float64(knapsack[currentItem-1][currentItemWeight]), float64(items[currentItem-1].Value+knapsack[currentItem-1][currentItemWeight-items[currentItem-1].Weight])))
			}
		}
	}

	return knapsack[amountOfItems][maximumWeight]
}
