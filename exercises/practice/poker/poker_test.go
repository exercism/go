package poker

import (
	"reflect"
	"testing"
)

// Define a function BestHand([]string) ([]string, error).

var invalidTestCases = []struct {
	name string
	hand string
}{
	{
		name: "1 is an invalid card rank",
		hand: "1♢ 2♡ 3♡ 4♡ 5♡",
	},
	{
		name: "11 is an invalid card rank",
		hand: "11♢ 2♡ 3♡ 4♡ 5♡",
	},
	{
		name: "too few cards",
		hand: "2♡ 3♡ 4♡ 5♡",
	},
	{
		name: "too many cards",
		hand: "2♡ 3♡ 4♡ 5♡ 6♡ 7♡",
	},
	{
		name: "lack of rank",
		hand: "11♢ 2♡ ♡ 4♡ 5♡",
	},
	{
		name: "lack of suit",
		hand: "2♡ 3♡ 4 5♡ 7♡",
	},
	{
		name: "H is an invalid suit",
		hand: "2♡ 3♡ 4H 5♡ 7♡",
	},
	{
		name: "♥ is an invalid suit",
		hand: "2♡ 3♡ 4♥ 5♡ 7♡",
	},
	{
		name: "lack of spacing",
		hand: "2♡ 3♡ 5♡7♡ 8♡",
	},
	{
		name: "double suits after rank",
		hand: "2♡ 3♡ 5♡♡ 8♡ 9♡",
	},
}

func TestBestHandValid(t *testing.T) {
	for _, tt := range validTestCases {
		actual, err := BestHand(tt.hands)
		if err != nil {
			var _ error = err
			t.Fatalf("Got unexpected error in valid case %q: %v", tt.name, err)
		}
		if !reflect.DeepEqual(actual, tt.best) {
			t.Fatalf("Mismatch in result of valid case %q: got %#v, want %#v",
				tt.name, actual, tt.best)
		}
	}
}

func TestBestHandInvalid(t *testing.T) {
	for _, tt := range invalidTestCases {
		_, err := BestHand([]string{tt.hand})
		if err == nil {
			t.Fatalf("Did not get an error for invalid case %q", tt.name)
		}
	}
}

func BenchmarkBestHand(b *testing.B) {
	if testing.Short() {
		t.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range validTestCases {
			BestHand(tt.hands)
		}
	}
}
