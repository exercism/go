// TODO:
// include a comment in each function briefly explaining what the function should do. This is a bit of quality of life for
// students, to help them know which part of the instructions that function should implement.

// One thing that slightly annoys me in this exercise is that the instructions mention that to get the value for an ability,
// you should roll 4 dice, discard the worst result and sum the 3 best dice rolls. But with the functions we have currently
// (which are the ones most tracks have), that logic is never enforced or tested anywhere. I think it would make sense to
// make students write a function (which would be included in the skeleton given to them) that should take an array of 4
// elements, or receive 4 arguments, and return the sum of the 3 biggest ones.
// We could then also write manual tests for this function.

package dndcharacter

type CharacterSheet struct {
	Strength     string
	Dexterity    string
	Constitution string
	Intelligence string
	Wisdom       string
	Charisma     string
	Hitpoints    string
}

// Modifier
func Modifier(score int) int {
	panic("Please implement the 'Modifier' function")
}

// Ability
func Ability() string {
	panic("Please implement the 'Ability' function")
}

// Character
func Character() CharacterSheet {
	panic("Please implement the 'Character' function")
}
