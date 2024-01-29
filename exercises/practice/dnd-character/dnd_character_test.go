package dndcharacter

import "testing"

func TestModifier(t *testing.T) {
	for _, tc := range modifierTests {
		t.Run(tc.description, func(t *testing.T) {
			actual := Modifier(tc.input.Score)
			if actual != tc.expected {
				t.Fatalf("Modifier(%#v) = %#v, want %#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestAbility(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			got := Ability()
			if !inAcceptedRange(got) {
				t.Fatal("Ability score is not within accepted range (3-18)")
			}
		})
	}
}

func TestCharacter(t *testing.T) {
	t.Run("should generate a character sheet with random ability scores", func(t *testing.T) {
		character := Character()

		assertAbilityScoreInRange(t, "Charisma", character.Charisma)
		assertAbilityScoreInRange(t, "Strength", character.Strength)
		assertAbilityScoreInRange(t, "Dexterity", character.Dexterity)
		assertAbilityScoreInRange(t, "Wisdom", character.Wisdom)
		assertAbilityScoreInRange(t, "Intelligence", character.Intelligence)
		assertAbilityScoreInRange(t, "Constitution", character.Constitution)

		if character.Hitpoints < 10 {
			t.Fatalf("Character's base hitpoints are incorrect. Got %d", character.Hitpoints)
		}
	})
}

func BenchmarkAbility(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ability()
	}
}

func BenchmarkCharacter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Character()
	}
}

func BenchmarkModifier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Modifier(i)
	}
}

func inAcceptedRange(score int) bool {
	return score >= 3 && score <= 18
}

func assertAbilityScoreInRange(t testing.TB, ability string, score int) {
	t.Helper()

	if !inAcceptedRange(score) {
		t.Fatalf("%s score is not withing accepted range. Got %d", ability, score)
	}
}
