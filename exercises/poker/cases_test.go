package poker

// Source: exercism/x-common
// Commit: d078ba8 poker: Add canonical-data (#793)
// x-common version: 1.0.0

type validTestCase struct {
	name  string
	hands []string
	best  []string
}

var validTestCases = []validTestCase{
	{
		name: "single hand always wins",
		hands: []string{
			"4♤ 5♤ 7♡ 8♢ J♧",
		},
		best: []string{"4♤ 5♤ 7♡ 8♢ J♧"},
	},
	{
		name: "highest card out of all hands wins",
		hands: []string{
			"4♢ 5♤ 6♤ 8♢ 3♧",
			"2♤ 4♧ 7♤ 9♡ 10♡",
			"3♤ 4♤ 5♢ 6♡ J♡",
		},
		best: []string{"3♤ 4♤ 5♢ 6♡ J♡"},
	},
	{
		name: "a tie has multiple winners",
		hands: []string{
			"4♢ 5♤ 6♤ 8♢ 3♧",
			"2♤ 4♧ 7♤ 9♡ 10♡",
			"3♤ 4♤ 5♢ 6♡ J♡",
			"3♡ 4♡ 5♧ 6♧ J♢",
		},
		best: []string{"3♤ 4♤ 5♢ 6♡ J♡", "3♡ 4♡ 5♧ 6♧ J♢"},
	},
	{
		name: "multiple hands with the same high cards, tie compares next highest ranked, down to last card",
		hands: []string{
			"3♤ 5♡ 6♤ 8♢ 7♡",
			"2♤ 5♢ 6♢ 8♧ 7♤",
		},
		best: []string{"3♤ 5♡ 6♤ 8♢ 7♡"},
	},
	{
		name: "one pair beats high card",
		hands: []string{
			"4♤ 5♡ 6♧ 8♢ K♡",
			"2♤ 4♡ 6♤ 4♢ J♡",
		},
		best: []string{"2♤ 4♡ 6♤ 4♢ J♡"},
	},
	{
		name: "highest pair wins",
		hands: []string{
			"4♤ 2♡ 6♤ 2♢ J♡",
			"2♤ 4♡ 6♧ 4♢ J♢",
		},
		best: []string{"2♤ 4♡ 6♧ 4♢ J♢"},
	},
	{
		name: "two pairs beats one pair",
		hands: []string{
			"2♤ 8♡ 6♤ 8♢ J♡",
			"4♤ 5♡ 4♧ 8♧ 5♧",
		},
		best: []string{"4♤ 5♡ 4♧ 8♧ 5♧"},
	},
	{
		name: "both hands have two pairs, highest ranked pair wins",
		hands: []string{
			"2♤ 8♡ 2♢ 8♢ 3♡",
			"4♤ 5♡ 4♧ 8♤ 5♢",
		},
		best: []string{"2♤ 8♡ 2♢ 8♢ 3♡"},
	},
	{
		name: "both hands have two pairs, with the same highest ranked pair, tie goes to low pair",
		hands: []string{
			"2♤ Q♤ 2♧ Q♢ J♡",
			"J♢ Q♡ J♤ 8♢ Q♧",
		},
		best: []string{"J♢ Q♡ J♤ 8♢ Q♧"},
	},
	{
		name: "both hands have two identically ranked pairs, tie goes to remaining card (kicker)",
		hands: []string{
			"J♢ Q♡ J♤ 8♢ Q♧",
			"J♤ Q♤ J♧ 2♢ Q♢",
		},
		best: []string{"J♢ Q♡ J♤ 8♢ Q♧"},
	},
	{
		name: "three of a kind beats two pair",
		hands: []string{
			"2♤ 8♡ 2♡ 8♢ J♡",
			"4♤ 5♡ 4♧ 8♤ 4♡",
		},
		best: []string{"4♤ 5♡ 4♧ 8♤ 4♡"},
	},
	{
		name: "both hands have three of a kind, tie goes to highest ranked triplet",
		hands: []string{
			"2♤ 2♡ 2♧ 8♢ J♡",
			"4♤ A♡ A♤ 8♧ A♢",
		},
		best: []string{"4♤ A♡ A♤ 8♧ A♢"},
	},
	{
		name: "with multiple decks, two players can have same three of a kind, ties go to highest remaining cards",
		hands: []string{
			"4♤ A♡ A♤ 7♧ A♢",
			"4♤ A♡ A♤ 8♧ A♢",
		},
		best: []string{"4♤ A♡ A♤ 8♧ A♢"},
	},
	{
		name: "a straight beats three of a kind",
		hands: []string{
			"4♤ 5♡ 4♧ 8♢ 4♡",
			"3♤ 4♢ 2♤ 6♢ 5♧",
		},
		best: []string{"3♤ 4♢ 2♤ 6♢ 5♧"},
	},
	{
		name: "aces can end a straight (10 J Q K A)",
		hands: []string{
			"4♤ 5♡ 4♧ 8♢ 4♡",
			"10♢ J♡ Q♤ K♢ A♧",
		},
		best: []string{"10♢ J♡ Q♤ K♢ A♧"},
	},
	{
		name: "aces can start a straight (A 2 3 4 5)",
		hands: []string{
			"4♤ 5♡ 4♧ 8♢ 4♡",
			"4♢ A♡ 3♤ 2♢ 5♧",
		},
		best: []string{"4♢ A♡ 3♤ 2♢ 5♧"},
	},
	{
		name: "both hands with a straight, tie goes to highest ranked card",
		hands: []string{
			"4♤ 6♧ 7♤ 8♢ 5♡",
			"5♤ 7♡ 8♤ 9♢ 6♡",
		},
		best: []string{"5♤ 7♡ 8♤ 9♢ 6♡"},
	},
	{
		name: "even though an ace is usually high, a 5-high straight is the lowest-scoring straight",
		hands: []string{
			"2♡ 3♧ 4♢ 5♢ 6♡",
			"4♤ A♡ 3♤ 2♢ 5♡",
		},
		best: []string{"2♡ 3♧ 4♢ 5♢ 6♡"},
	},
	{
		name: "flush beats a straight",
		hands: []string{
			"4♧ 6♡ 7♢ 8♢ 5♡",
			"2♤ 4♤ 5♤ 6♤ 7♤",
		},
		best: []string{"2♤ 4♤ 5♤ 6♤ 7♤"},
	},
	{
		name: "both hands have a flush, tie goes to high card, down to the last one if necessary",
		hands: []string{
			"4♡ 7♡ 8♡ 9♡ 6♡",
			"2♤ 4♤ 5♤ 6♤ 7♤",
		},
		best: []string{"4♡ 7♡ 8♡ 9♡ 6♡"},
	},
	{
		name: "full house beats a flush",
		hands: []string{
			"3♡ 6♡ 7♡ 8♡ 5♡",
			"4♤ 5♡ 4♧ 5♢ 4♡",
		},
		best: []string{"4♤ 5♡ 4♧ 5♢ 4♡"},
	},
	{
		name: "both hands have a full house, tie goes to highest-ranked triplet",
		hands: []string{
			"4♡ 4♤ 4♢ 9♤ 9♢",
			"5♡ 5♤ 5♢ 8♤ 8♢",
		},
		best: []string{"5♡ 5♤ 5♢ 8♤ 8♢"},
	},
	{
		name: "with multiple decks, both hands have a full house with the same triplet, tie goes to the pair",
		hands: []string{
			"5♡ 5♤ 5♢ 9♤ 9♢",
			"5♡ 5♤ 5♢ 8♤ 8♢",
		},
		best: []string{"5♡ 5♤ 5♢ 9♤ 9♢"},
	},
	{
		name: "four of a kind beats a full house",
		hands: []string{
			"4♤ 5♡ 4♢ 5♢ 4♡",
			"3♤ 3♡ 2♤ 3♢ 3♧",
		},
		best: []string{"3♤ 3♡ 2♤ 3♢ 3♧"},
	},
	{
		name: "both hands have four of a kind, tie goes to high quad",
		hands: []string{
			"2♤ 2♡ 2♧ 8♢ 2♢",
			"4♤ 5♡ 5♤ 5♢ 5♧",
		},
		best: []string{"4♤ 5♡ 5♤ 5♢ 5♧"},
	},
	{
		name: "with multiple decks, both hands with identical four of a kind, tie determined by kicker",
		hands: []string{
			"3♤ 3♡ 2♤ 3♢ 3♧",
			"3♤ 3♡ 4♤ 3♢ 3♧",
		},
		best: []string{"3♤ 3♡ 4♤ 3♢ 3♧"},
	},
	{
		name: "straight flush beats four of a kind",
		hands: []string{
			"4♤ 5♡ 5♤ 5♢ 5♧",
			"7♤ 8♤ 9♤ 6♤ 10♤",
		},
		best: []string{"7♤ 8♤ 9♤ 6♤ 10♤"},
	},
	{
		name: "both hands have straight flush, tie goes to highest-ranked card",
		hands: []string{
			"4♡ 6♡ 7♡ 8♡ 5♡",
			"5♤ 7♤ 8♤ 9♤ 6♤",
		},
		best: []string{"5♤ 7♤ 8♤ 9♤ 6♤"},
	},
}
