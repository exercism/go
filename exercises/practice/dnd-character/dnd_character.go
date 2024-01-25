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

func Modifier(score int) int {
	panic("implement me")
}

func Ability() string {
	panic("implement me")
}

func Character() CharacterSheet {
	panic("implement me")
}
