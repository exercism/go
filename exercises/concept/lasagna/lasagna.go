package main

import (
	"fmt"
)

const ovenTime = 40

func main() {
	fmt.Println(RemainingOvenTime(30))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(3, 20))

}

func RemainingOvenTime(actual int) int {
	return ovenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return actualMinutesInOven + PreparationTime(numberOfLayers)
}
