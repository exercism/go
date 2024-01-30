package dndcharacter

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	panic("Please implement the Modifier() function")
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	panic("Please implement the Ability() function")
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	panic("Please implement the GenerateCharacter() function")
}
