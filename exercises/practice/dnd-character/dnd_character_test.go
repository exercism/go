// TODO: implement test cases
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

}

func TestCharacter(t *testing.T) {

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
