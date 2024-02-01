package knapsack

import "testing"

func TestKnapsack(t *testing.T) {
	for _, tc := range maximumValueTests {
		actual := Knapsack(tc.input.MaximumWeight, tc.input.Items)
		expected := tc.expected

		if actual != expected {
			t.Fatalf("Knapsack(%#v) = %d, want %d", tc.input, actual, tc.expected)
		}
	}
}
