package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier should calculate the correct modifier for a given score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability should generate the score for a random ability
func Ability() int {
	var scores []int

	for i := 0; i < 4; i++ {
		roll := RollDice()
		scores = append(scores, roll)
	}

	return CalculateAbilityScore(scores)
}

// GenerateCharacter should return a Character with valid ability scores
func GenerateCharacter() Character {
	character := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	character.Hitpoints = 10 + character.Constitution

	return character
}

// CalculateAbilityScore expects an array of 4 dice scores and returns the sum of the 3 highest numbers
func CalculateAbilityScore(scores []int) int {
	var sum int

	for _, score := range scores {
		sum += score
	}

	return sum - slices.Min(scores)
}

func RollDice() int {
	return rand.Intn(6) + 1
}
