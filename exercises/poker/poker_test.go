package poker

import (
	"reflect"
	"testing"
)

// Define a function BestHand([]string) ([]string, error).
//
// Also define a testVersion with a value that matches
// the internal targetTestVersion here.

const targetTestVersion = 2

var validTestCases = []struct {
	name  string
	hands []string
	best  []string
}{
	{
		name:  "single hand is always best",
		hands: []string{"3♡ 10♢ 7♧ 8♤ A♢"},
		best:  []string{"3♡ 10♢ 7♧ 8♤ A♢"},
	},
	{
		name:  "highest card",
		hands: []string{"3♢ 2♢ 5♤ 6♤ 9♡", "3♡ 2♡ 5♧ 6♢ 10♡"},
		best:  []string{"3♡ 2♡ 5♧ 6♢ 10♡"},
	},
	{
		name:  "highest card with mostly same cards",
		hands: []string{"3♢ 2♢ 5♤ 6♤ 9♡", "3♡ 3♤ 5♧ 6♢ 9♢"},
		best:  []string{"3♡ 3♤ 5♧ 6♢ 9♢"},
	},
	{
		name:  "pair beats lower",
		hands: []string{"4♢ 3♤ 4♤ J♤ K♤", "A♡ K♡ J♢ 10♧ 9♡"},
		best:  []string{"4♢ 3♤ 4♤ J♤ K♤"},
	},
	{
		name:  "best pair",
		hands: []string{"4♡ 2♡ 5♧ 4♢ 10♡", "3♢ 3♡ 5♤ 6♤ 9♡"},
		best:  []string{"4♡ 2♡ 5♧ 4♢ 10♡"},
	},
	{
		name:  "best pair with same pair and highest cards",
		hands: []string{"4♡ 2♡ 5♧ 4♢ 10♡", "4♤ 4♧ 5♡ 10♢ 3♡"},
		best:  []string{"4♤ 4♧ 5♡ 10♢ 3♡"},
	},
	{
		name: "two pair beats lower",
		hands: []string{
			"4♢ 3♤ 4♤ J♤ K♤",
			"A♡ K♡ J♢ 10♧ 9♡",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"2♢ 8♡ 5♢ 2♡ 8♧"},
	},
	{
		name: "best two pair",
		hands: []string{
			"4♢ J♧ 4♤ J♤ K♤",
			"A♡ K♡ J♢ 10♧ 9♡",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"4♢ J♧ 4♤ J♤ K♤"},
	},
	{
		name: "best two pair with equal highest pair",
		hands: []string{
			"4♢ J♧ 4♤ J♤ K♤",
			"A♡ K♡ J♢ 10♧ 9♡",
			"3♢ J♡ 5♢ 3♡ J♢",
		},
		best: []string{"4♢ J♧ 4♤ J♤ K♤"},
	},
	{
		name: "best two pair with equal pairs",
		hands: []string{
			"4♢ J♧ 4♤ J♤ 2♤",
			"A♡ K♡ J♢ 10♧ 9♡",
			"4♧ J♡ 5♢ 4♡ J♢",
		},
		best: []string{"4♧ J♡ 5♢ 4♡ J♢"},
	},
	{
		name: "three of a kind beats lower",
		hands: []string{
			"4♢ 3♤ 4♤ J♤ K♤",
			"A♡ K♡ J♢ 10♧ 9♡",
			"3♢ 8♡ 3♢ 3♧ 8♧",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"3♢ 8♡ 3♢ 3♧ 8♧"},
	},
	{
		name: "best three of a kind",
		hands: []string{
			"4♢ 3♤ 4♤ J♤ 4♡",
			"A♡ K♡ J♢ 10♧ 9♡",
			"3♢ 8♡ 3♢ 3♧ 9♧",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"4♢ 3♤ 4♤ J♤ 4♡"},
	},
	{
		name: "straight beats lower",
		hands: []string{
			"4♢ 3♤ 4♤ J♤ K♤",
			"Q♡ K♡ J♢ 10♧ 9♡",
			"3♢ 8♡ 3♢ 3♧ 9♧",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"Q♡ K♡ J♢ 10♧ 9♡"},
	},
	{
		name: "straight includes ace as one",
		hands: []string{
			"4♢ 3♤ 4♤ J♤ K♤",
			"2♤ 3♡ A♤ 5♤ 4♤",
			"3♢ 8♡ 3♢ 3♧ 9♧",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"2♤ 3♡ A♤ 5♤ 4♤"},
	},
	{
		name: "best straight",
		hands: []string{
			"4♢ 3♤ 4♤ J♤ K♤",
			"Q♡ K♡ J♢ 10♧ 9♡",
			"A♢ K♧ 10♢ J♢ Q♢",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"A♢ K♧ 10♢ J♢ Q♢"},
	},
	{
		name: "flush beats lower",
		hands: []string{
			"4♤ 3♤ 8♤ J♤ K♤",
			"Q♡ K♡ J♢ 10♧ 9♡",
			"3♢ 8♡ 3♢ 3♧ 9♧",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"4♤ 3♤ 8♤ J♤ K♤"},
	},
	{
		name: "best flush",
		hands: []string{
			"4♤ 3♤ 8♤ J♤ K♤",
			"Q♡ K♡ J♢ 10♧ 9♡",
			"3♢ 8♢ A♢ 3♢ 7♢",
			"2♢ 8♡ 5♢ 2♡ 8♧",
		},
		best: []string{"3♢ 8♢ A♢ 3♢ 7♢"},
	},
	{
		name: "full house beats lower",
		hands: []string{
			"4♤ 3♤ 8♤ J♤ K♤",
			"2♢ 8♡ 8♢ 2♡ 8♧",
			"Q♡ K♡ J♢ 10♧ 9♡",
			"3♢ A♡ 3♢ 3♧ A♧",
		},
		best: []string{"2♢ 8♡ 8♢ 2♡ 8♧"},
	},
	{
		name: "best full house",
		hands: []string{
			"4♤ 3♤ 8♤ J♤ K♤",
			"2♢ 8♡ 8♢ 2♡ 8♧",
			"5♡ 5♢ A♢ 5♧ A♢",
			"3♢ A♡ 3♢ 3♧ A♧",
		},
		best: []string{"2♢ 8♡ 8♢ 2♡ 8♧"},
	},
	{
		name: "four of a kind beats lower",
		hands: []string{
			"4♤ 5♤ 8♤ J♤ K♤",
			"2♢ 8♡ 8♢ 2♡ 8♧",
			"Q♡ K♡ J♢ 10♧ 9♡",
			"3♢ 3♡ 3♤ 3♧ A♧",
		},
		best: []string{"3♢ 3♡ 3♤ 3♧ A♧"},
	},
	{
		name: "best four of a kind",
		hands: []string{
			"4♤ 5♤ 8♤ J♤ K♤",
			"2♢ 2♧ 8♢ 2♡ 2♤",
			"Q♡ K♡ J♢ 10♧ 9♡",
			"3♢ 3♡ 3♤ 3♧ A♧",
		},
		best: []string{"3♢ 3♡ 3♤ 3♧ A♧"},
	},
	{
		name: "straight flush beats lower",
		hands: []string{
			"4♤ 5♢ 8♤ J♤ K♤",
			"2♢ 8♡ 8♢ 2♡ 8♧",
			"Q♡ K♡ 8♡ 10♡ 9♡",
			"2♤ 3♤ A♤ 5♤ 4♤",
		},
		best: []string{"2♤ 3♤ A♤ 5♤ 4♤"},
	},
	{
		name: "best straight flush is royal flush",
		hands: []string{
			"4♤ 5♤ 8♤ J♤ K♤",
			"2♢ 8♡ 8♢ 2♡ 8♧",
			"Q♡ K♡ J♡ 10♡ 9♡",
			"Q♢ K♢ J♢ 10♢ A♢",
		},
		best: []string{"Q♢ K♢ J♢ 10♢ A♢"},
	},
	{
		name:  "tie for best pair",
		hands: []string{"4♡ 2♡ 5♧ 4♢ 10♡", "4♧ 10♢ 5♤ 2♤ 4♧"},
		best:  []string{"4♡ 2♡ 5♧ 4♢ 10♡", "4♧ 10♢ 5♤ 2♤ 4♧"},
	},
	{
		name: "tie of three",
		hands: []string{
			"A♡ 2♡ 3♡ 4♡ 5♡",
			"A♤ 2♤ 3♤ 4♤ 5♤",
			"5♧ 4♧ 3♧ 2♧ A♧",
			"A♢ 2♢ 6♢ 4♢ 5♢",
		},
		best: []string{
			"A♡ 2♡ 3♡ 4♡ 5♡",
			"A♤ 2♤ 3♤ 4♤ 5♤",
			"5♧ 4♧ 3♧ 2♧ A♧",
		},
	},
}

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
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, tt := range validTestCases {
		actual, err := BestHand(tt.hands)
		if err != nil {
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
	for i := 0; i < b.N; i++ {
		for _, tt := range validTestCases {
			BestHand(tt.hands)
		}
	}
}
