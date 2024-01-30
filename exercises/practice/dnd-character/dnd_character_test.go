package dndcharacter

import "testing"

func TestModifier(t *testing.T) {
	for _, tc := range modifierTests {
		t.Run(tc.description, func(t *testing.T) {
			actual := Modifier(tc.input.Score)
			if actual != tc.expected {
				t.Fatalf("Modifier(%d) = %d, want %d", tc.input.Score, actual, tc.expected)
			}
		})
	}
}

func TestAbility(t *testing.T) {
	t.Run("should generate ability score within accepted range", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			got := Ability()
			if !inAcceptedRange(got) {
				t.Fatalf("Ability() returned a score for an ability outside the accepted range. Got %d, expected a value between 3 and 18 inclusive.", got)
			}
		}
	})
}

func TestGenerateCharacter(t *testing.T) {
	t.Run("should generate a character with random ability scores", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			character := GenerateCharacter()

			assertCharacterAbilityScoreInRange(t, "Charisma", character.Charisma)
			assertCharacterAbilityScoreInRange(t, "Strength", character.Strength)
			assertCharacterAbilityScoreInRange(t, "Dexterity", character.Dexterity)
			assertCharacterAbilityScoreInRange(t, "Wisdom", character.Wisdom)
			assertCharacterAbilityScoreInRange(t, "Intelligence", character.Intelligence)
			assertCharacterAbilityScoreInRange(t, "Constitution", character.Constitution)

			expectedHitpoints := 10 + Modifier(character.Constitution)
			if character.Hitpoints != expectedHitpoints {
				t.Fatalf("Got %d hitpoints for a character with %d constitution, expected %d hitpoints", character.Hitpoints, character.Constitution, expectedHitpoints)
			}
		}
	})
}

func inAcceptedRange(score int) bool {
	return score >= 3 && score <= 18
}

func assertCharacterAbilityScoreInRange(t *testing.T, ability string, score int) {
	t.Helper()

	if !inAcceptedRange(score) {
		t.Fatalf("GenerateCharacter() created a character with a %s score of %d, but the score for an ability is expected to be between 3 and 18 inclusive", ability, score)
	}
}

func BenchmarkModifier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Modifier(i)
	}
}

func BenchmarkAbility(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ability()
	}
}

func BenchmarkCharacter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateCharacter()
	}
}
