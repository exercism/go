package poker

import (
	"reflect"
	"testing"
)

var invalidTestCases = []struct {
	description string
	hand        string
}{
	{
		description: "1 is an invalid card rank",
		hand:        "1♢ 2♡ 3♡ 4♡ 5♡",
	},
	{
		description: "11 is an invalid card rank",
		hand:        "11♢ 2♡ 3♡ 4♡ 5♡",
	},
	{
		description: "too few cards",
		hand:        "2♡ 3♡ 4♡ 5♡",
	},
	{
		description: "too many cards",
		hand:        "2♡ 3♡ 4♡ 5♡ 6♡ 7♡",
	},
	{
		description: "lack of rank",
		hand:        "11♢ 2♡ ♡ 4♡ 5♡",
	},
	{
		description: "lack of suit",
		hand:        "2♡ 3♡ 4 5♡ 7♡",
	},
	{
		description: "H is an invalid suit",
		hand:        "2♡ 3♡ 4H 5♡ 7♡",
	},
	{
		description: "♥ is an invalid suit",
		hand:        "2♡ 3♡ 4♥ 5♡ 7♡",
	},
	{
		description: "lack of spacing",
		hand:        "2♡ 3♡ 5♡7♡ 8♡",
	},
	{
		description: "double suits after rank",
		hand:        "2♡ 3♡ 5♡♡ 8♡ 9♡",
	},
}

func TestBestHandValid(t *testing.T) {
	for _, tc := range validTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := BestHand(tc.hands)
			if err != nil {
				t.Fatalf("BestHand(%v) returned error: %v, want: %v", tc.hands, err, tc.expected)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("BestHand(%v) = %v, want: %v", tc.hands, actual, tc.expected)
			}
		})
	}
}

func TestBestHandInvalid(t *testing.T) {
	for _, tc := range invalidTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := BestHand([]string{tc.hand})
			if err == nil {
				t.Fatalf("BestHand(%v) expected error, got: %v", []string{tc.hand}, actual)
			}
		})
	}
}

func BenchmarkBestHand(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range validTestCases {
			BestHand(tt.hands)
		}
	}
}
