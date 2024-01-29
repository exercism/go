package dndcharacter

type CharacterSheet struct {
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
	panic("Please implement the 'Modifier' function")
}

// Ability should generate the score for a random ability
func Ability() int {
	panic("Please implement the 'Ability' function")
}

// Character should return a CharacterSheet with valid ability scores
func Character() CharacterSheet {
	panic("Please implement the 'Character' function")
}

// CalculateAbilityScore expects an array of 4 dice scores and returns the sum of the 3 highest numbers
func CalculateAbilityScore(scores []int) int {
	panic("Please implement the 'CalculateAbilityScore' function")
}
