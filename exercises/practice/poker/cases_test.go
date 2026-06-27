package poker

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 25c2ae5 build(deps): bump fast-uri from 3.1.0 to 3.1.2 (#2653)

type testCase struct {
	description string
	hands       []string
	expected    []string
}

var validTestCases = []testCase{
	{
		description: "single hand always wins",
		hands:       []string{"4♤ 5♤ 7♡ 8♢ J♧"},
		expected:    []string{"4♤ 5♤ 7♡ 8♢ J♧"},
	},
	{
		description: "highest card out of all hands wins",
		hands:       []string{"4♢ 5♤ 6♤ 8♢ 3♧", "2♤ 4♧ 7♤ 9♡ 10♡", "3♤ 4♤ 5♢ 6♡ J♡"},
		expected:    []string{"3♤ 4♤ 5♢ 6♡ J♡"},
	},
	{
		description: "a tie has multiple winners",
		hands:       []string{"4♢ 5♤ 6♤ 8♢ 3♧", "2♤ 4♧ 7♤ 9♡ 10♡", "3♤ 4♤ 5♢ 6♡ J♡", "3♡ 4♡ 5♧ 6♧ J♢"},
		expected:    []string{"3♤ 4♤ 5♢ 6♡ J♡", "3♡ 4♡ 5♧ 6♧ J♢"},
	},
	{
		description: "multiple hands with the same high cards, tie compares next highest ranked, down to last card",
		hands:       []string{"3♤ 5♡ 6♤ 8♢ 7♡", "2♤ 5♢ 6♢ 8♧ 7♤"},
		expected:    []string{"3♤ 5♡ 6♤ 8♢ 7♡"},
	},
	{
		description: "winning high card hand also has the lowest card",
		hands:       []string{"2♤ 5♡ 6♤ 8♢ 7♡", "3♤ 4♢ 6♢ 8♧ 7♤"},
		expected:    []string{"2♤ 5♡ 6♤ 8♢ 7♡"},
	},
	{
		description: "one pair beats high card",
		hands:       []string{"4♤ 5♡ 6♧ 8♢ K♡", "2♤ 4♡ 6♤ 4♢ J♡"},
		expected:    []string{"2♤ 4♡ 6♤ 4♢ J♡"},
	},
	{
		description: "highest pair wins",
		hands:       []string{"4♤ 2♡ 6♤ 2♢ J♡", "2♤ 4♡ 6♧ 4♢ J♢"},
		expected:    []string{"2♤ 4♡ 6♧ 4♢ J♢"},
	},
	{
		description: "both hands have the same pair, high card wins",
		hands:       []string{"4♡ 4♤ A♡ J♧ 3♢", "4♧ 4♢ A♤ 5♢ 6♧"},
		expected:    []string{"4♡ 4♤ A♡ J♧ 3♢"},
	},
	{
		description: "two pairs beats one pair",
		hands:       []string{"2♤ 8♡ 6♤ 8♢ J♡", "4♤ 5♡ 4♧ 8♧ 5♧"},
		expected:    []string{"4♤ 5♡ 4♧ 8♧ 5♧"},
	},
	{
		description: "both hands have two pairs, highest ranked pair wins",
		hands:       []string{"2♤ 8♡ 2♢ 8♢ 3♡", "4♤ 5♡ 4♧ 8♤ 5♢"},
		expected:    []string{"2♤ 8♡ 2♢ 8♢ 3♡"},
	},
	{
		description: "both hands have two pairs, with the same highest ranked pair, tie goes to low pair",
		hands:       []string{"2♤ Q♤ 2♧ Q♢ J♡", "J♢ Q♡ J♤ 8♢ Q♧"},
		expected:    []string{"J♢ Q♡ J♤ 8♢ Q♧"},
	},
	{
		description: "both hands have two identically ranked pairs, tie goes to remaining card (kicker)",
		hands:       []string{"J♢ Q♡ J♤ 8♢ Q♧", "J♤ Q♤ J♧ 2♢ Q♢"},
		expected:    []string{"J♢ Q♡ J♤ 8♢ Q♧"},
	},
	{
		description: "both hands have two pairs that add to the same value, win goes to highest pair",
		hands:       []string{"6♤ 6♡ 3♤ 3♡ A♤", "7♡ 7♤ 2♡ 2♤ A♧"},
		expected:    []string{"7♡ 7♤ 2♡ 2♤ A♧"},
	},
	{
		description: "two pairs first ranked by largest pair",
		hands:       []string{"5♧ 2♤ 5♤ 4♡ 4♧", "6♤ 2♤ 6♡ 7♧ 2♧"},
		expected:    []string{"6♤ 2♤ 6♡ 7♧ 2♧"},
	},
	{
		description: "three of a kind beats two pair",
		hands:       []string{"2♤ 8♡ 2♡ 8♢ J♡", "4♤ 5♡ 4♧ 8♤ 4♡"},
		expected:    []string{"4♤ 5♡ 4♧ 8♤ 4♡"},
	},
	{
		description: "both hands have three of a kind, tie goes to highest ranked triplet",
		hands:       []string{"2♤ 2♡ 2♧ 8♢ J♡", "4♤ A♡ A♤ 8♧ A♢"},
		expected:    []string{"4♤ A♡ A♤ 8♧ A♢"},
	},
	{
		description: "with multiple decks, two players can have same three of a kind, ties go to highest remaining cards",
		hands:       []string{"5♤ A♡ A♤ 7♧ A♢", "4♤ A♡ A♤ 8♧ A♢"},
		expected:    []string{"4♤ A♡ A♤ 8♧ A♢"},
	},
	{
		description: "a straight beats three of a kind",
		hands:       []string{"4♤ 5♡ 4♧ 8♢ 4♡", "3♤ 4♢ 2♤ 6♢ 5♧"},
		expected:    []string{"3♤ 4♢ 2♤ 6♢ 5♧"},
	},
	{
		description: "aces can end a straight (10 J Q K A)",
		hands:       []string{"4♤ 5♡ 4♧ 8♢ 4♡", "10♢ J♡ Q♤ K♢ A♧"},
		expected:    []string{"10♢ J♡ Q♤ K♢ A♧"},
	},
	{
		description: "aces can start a straight (A 2 3 4 5)",
		hands:       []string{"4♤ 5♡ 4♧ 8♢ 4♡", "4♢ A♡ 3♤ 2♢ 5♧"},
		expected:    []string{"4♢ A♡ 3♤ 2♢ 5♧"},
	},
	{
		description: "aces cannot be in the middle of a straight (Q K A 2 3)",
		hands:       []string{"2♧ 3♢ 7♡ 5♡ 2♤", "Q♤ K♡ A♧ 2♢ 3♤"},
		expected:    []string{"2♧ 3♢ 7♡ 5♡ 2♤"},
	},
	{
		description: "both hands with a straight, tie goes to highest ranked card",
		hands:       []string{"4♤ 6♧ 7♤ 8♢ 5♡", "5♤ 7♡ 8♤ 9♢ 6♡"},
		expected:    []string{"5♤ 7♡ 8♤ 9♢ 6♡"},
	},
	{
		description: "even though an ace is usually high, a 5-high straight is the lowest-scoring straight",
		hands:       []string{"2♡ 3♧ 4♢ 5♢ 6♡", "4♤ A♡ 3♤ 2♢ 5♡"},
		expected:    []string{"2♡ 3♧ 4♢ 5♢ 6♡"},
	},
	{
		description: "flush beats a straight",
		hands:       []string{"4♧ 6♡ 7♢ 8♢ 5♡", "2♤ 4♤ 5♤ 6♤ 7♤"},
		expected:    []string{"2♤ 4♤ 5♤ 6♤ 7♤"},
	},
	{
		description: "both hands have a flush, tie goes to high card, down to the last one if necessary",
		hands:       []string{"2♡ 7♡ 8♡ 9♡ 6♡", "3♤ 5♤ 6♤ 7♤ 8♤"},
		expected:    []string{"2♡ 7♡ 8♡ 9♡ 6♡"},
	},
	{
		description: "full house beats a flush",
		hands:       []string{"3♡ 6♡ 7♡ 8♡ 5♡", "4♤ 5♡ 4♧ 5♢ 4♡"},
		expected:    []string{"4♤ 5♡ 4♧ 5♢ 4♡"},
	},
	{
		description: "both hands have a full house, tie goes to highest-ranked triplet",
		hands:       []string{"4♡ 4♤ 4♢ 9♤ 9♢", "5♡ 5♤ 5♢ 8♤ 8♢"},
		expected:    []string{"5♡ 5♤ 5♢ 8♤ 8♢"},
	},
	{
		description: "with multiple decks, both hands have a full house with the same triplet, tie goes to the pair",
		hands:       []string{"5♡ 5♤ 5♢ 9♤ 9♢", "5♡ 5♤ 5♢ 8♤ 8♢"},
		expected:    []string{"5♡ 5♤ 5♢ 9♤ 9♢"},
	},
	{
		description: "four of a kind beats a full house",
		hands:       []string{"4♤ 5♡ 4♢ 5♢ 4♡", "3♤ 3♡ 2♤ 3♢ 3♧"},
		expected:    []string{"3♤ 3♡ 2♤ 3♢ 3♧"},
	},
	{
		description: "both hands have four of a kind, tie goes to high quad",
		hands:       []string{"2♤ 2♡ 2♧ 8♢ 2♢", "4♤ 5♡ 5♤ 5♢ 5♧"},
		expected:    []string{"4♤ 5♡ 5♤ 5♢ 5♧"},
	},
	{
		description: "with multiple decks, both hands with identical four of a kind, tie determined by kicker",
		hands:       []string{"3♤ 3♡ 2♤ 3♢ 3♧", "3♤ 3♡ 4♤ 3♢ 3♧"},
		expected:    []string{"3♤ 3♡ 4♤ 3♢ 3♧"},
	},
	{
		description: "straight flush beats four of a kind",
		hands:       []string{"4♤ 5♡ 5♤ 5♢ 5♧", "7♤ 8♤ 9♤ 6♤ 10♤"},
		expected:    []string{"7♤ 8♤ 9♤ 6♤ 10♤"},
	},
	{
		description: "aces can end a straight flush (10 J Q K A)",
		hands:       []string{"K♧ A♡ A♤ A♢ A♧", "10♧ J♧ Q♧ K♧ A♧"},
		expected:    []string{"10♧ J♧ Q♧ K♧ A♧"},
	},
	{
		description: "aces can start a straight flush (A 2 3 4 5)",
		hands:       []string{"K♤ A♡ A♤ A♢ A♧", "4♡ A♡ 3♡ 2♡ 5♡"},
		expected:    []string{"4♡ A♡ 3♡ 2♡ 5♡"},
	},
	{
		description: "aces cannot be in the middle of a straight flush (Q K A 2 3)",
		hands:       []string{"2♧ A♧ Q♧ 10♧ K♧", "Q♡ K♡ A♡ 2♡ 3♡"},
		expected:    []string{"2♧ A♧ Q♧ 10♧ K♧"},
	},
	{
		description: "both hands have a straight flush, tie goes to highest-ranked card",
		hands:       []string{"4♡ 6♡ 7♡ 8♡ 5♡", "5♤ 7♤ 8♤ 9♤ 6♤"},
		expected:    []string{"5♤ 7♤ 8♤ 9♤ 6♤"},
	},
	{
		description: "even though an ace is usually high, a 5-high straight flush is the lowest-scoring straight flush",
		hands:       []string{"2♡ 3♡ 4♡ 5♡ 6♡", "4♢ A♢ 3♢ 2♢ 5♢"},
		expected:    []string{"2♡ 3♡ 4♡ 5♡ 6♡"},
	},
}
