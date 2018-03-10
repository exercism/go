//Package yacht is a sample solution to the yacht exercise for go
package yacht

import (
	"sort"
)

//Ones Twos represent enumerate the scoring categories
const (
	Ones = iota
	Twos
	Threes
	Fours
	Fives
	Sixes
	FullHouse
	FourOfAKind
	LittleStraight
	BigStraight
	Choice
	Yacht
)

//scoreN calculates the score of the lower categories
//It returns the total of those dice that equal n
func scoreN(dice []int, n int) int {
	score := 0
	for _, d := range dice {
		if d == n {
			score += n
		}
	}
	return score
}

//scoreAll adds up and returns the total of the five dice
func scoreAll(dice []int) int {
	return dice[0] + dice[1] + dice[2] + dice[3] + dice[4]
}

//Score returns the total score of a hand of Yacht dice in a given category.
func Score(dice []int, category int) int {
	sort.Ints(dice)
	switch category {
	case Ones, Twos, Threes, Fours, Fives, Sixes:
		//The category value of "ones" is 0, so adding 1 to the category is needed
		return scoreN(dice, category+1)
	case FullHouse:
		//Note, a yacht scores zero as a full house
		if dice[0] == dice[1] && dice[2] == dice[3] && dice[3] == dice[4] && dice[1] != dice[2] {
			return scoreAll(dice)
		}
		if dice[0] == dice[1] && dice[1] == dice[2] && dice[3] == dice[4] && dice[2] != dice[3] {
			return scoreAll(dice)
		}
		return 0
	case FourOfAKind:
		//Note, a yacht can score as four of a kind, but only 4 dice are counted
		if dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] {
			return dice[0] * 4
		}
		if dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4] {
			return dice[1] * 4
		}
		return 0
	case LittleStraight:
		if dice[0] == 1 && dice[1] == 2 && dice[2] == 3 && dice[3] == 4 && dice[4] == 5 {
			return 30
		}
		return 0
	case BigStraight:
		if dice[0] == 2 && dice[1] == 3 && dice[2] == 4 && dice[3] == 5 && dice[4] == 6 {
			return 30
		}
		return 0
	case Choice:
		return scoreAll(dice)
	case Yacht:
		if dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4] {
			return 50
		}
		return 0
	default:
		//This can't happen, so panic is appropriate
		panic("Unknown category")
	}
}
