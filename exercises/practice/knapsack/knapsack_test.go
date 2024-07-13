package knapsack

import "testing"

func TestKnapsack(t *testing.T) {
	for _, tc := range maximumValueTests {
		actual := Knapsack(tc.input.MaximumWeight, tc.input.Items)
		expected := tc.expected

		if actual != expected {
			t.Fatalf("Knapsack(%d, %+v) = %d, want %d", tc.input.MaximumWeight, tc.input.Items, actual, tc.expected)
		}
	}
}

func BenchmarkKnapsack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range maximumValueTests {
			Knapsack(tc.input.MaximumWeight, tc.input.Items)
		}
	}
}
