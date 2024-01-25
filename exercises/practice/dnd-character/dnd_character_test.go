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
	for _, tc := range abilityTests {
		t.Run(tc.description, func(t *testing.T) {
			actual := Ability()
			if actual != tc.expected {
				t.Fatalf("Ability(%#v) = %#v, want %#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestCharacter(t *testing.T) {
	for _, tc := range characterTests {
		t.Run(tc.description, func(t *testing.T) {
			actual := Character()
			if actual != tc.expected {
				t.Fatalf("Character(%#v) = %#v, want %#v", tc.input, actual, tc.expected)
			}
		})
	}
}
