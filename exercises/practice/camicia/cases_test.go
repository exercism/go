package camicia

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 7fc6778 Add new exercise: camicia (#2581)

var testCases = []struct {
	description string
	playerA     []string
	playerB     []string
	expected    Outcome
}{
	{
		description: "two cards, one trick",
		playerA:     []string{"2"},
		playerB:     []string{"3"},
		expected: Outcome{
			finishes: true,
			cards:    2,
			tricks:   1,
		},
	},
	{
		description: "three cards, one trick",
		playerA:     []string{"2", "4"},
		playerB:     []string{"3"},
		expected: Outcome{
			finishes: true,
			cards:    3,
			tricks:   1,
		},
	},
	{
		description: "four cards, one trick",
		playerA:     []string{"2", "4"},
		playerB:     []string{"3", "5", "6"},
		expected: Outcome{
			finishes: true,
			cards:    4,
			tricks:   1,
		},
	},
	{
		description: "the ace reigns supreme",
		playerA:     []string{"2", "A"},
		playerB:     []string{"3", "4", "5", "6", "7"},
		expected: Outcome{
			finishes: true,
			cards:    7,
			tricks:   1,
		},
	},
	{
		description: "the king beats ace",
		playerA:     []string{"2", "A"},
		playerB:     []string{"3", "4", "5", "6", "K"},
		expected: Outcome{
			finishes: true,
			cards:    7,
			tricks:   1,
		},
	},
	{
		description: "the queen seduces the king",
		playerA:     []string{"2", "A", "7", "8", "Q"},
		playerB:     []string{"3", "4", "5", "6", "K"},
		expected: Outcome{
			finishes: true,
			cards:    10,
			tricks:   1,
		},
	},
	{
		description: "the jack betrays the queen",
		playerA:     []string{"2", "A", "7", "8", "Q"},
		playerB:     []string{"3", "4", "5", "6", "K", "9", "J"},
		expected: Outcome{
			finishes: true,
			cards:    12,
			tricks:   1,
		},
	},
	{
		description: "the 10 just wants to put on a show",
		playerA:     []string{"2", "A", "7", "8", "Q", "10"},
		playerB:     []string{"3", "4", "5", "6", "K", "9", "J"},
		expected: Outcome{
			finishes: true,
			cards:    13,
			tricks:   1,
		},
	},
	{
		description: "simple loop with decks of 3 cards",
		playerA:     []string{"J", "2", "3"},
		playerB:     []string{"4", "J", "5"},
		expected: Outcome{
			finishes: false,
			cards:    8,
			tricks:   3,
		},
	},
	{
		description: "the story is starting to get a bit complicated",
		playerA:     []string{"2", "6", "6", "J", "4", "K", "Q", "10", "K", "J", "Q", "2", "3", "K", "5", "6", "Q", "Q", "A", "A", "6", "9", "K", "A", "8", "K", "2", "A", "9", "A", "Q", "4", "K", "K", "K", "3", "5", "K", "8", "Q", "3", "Q", "7", "J", "K", "J", "9", "J", "3", "3", "K", "K", "Q", "A", "K", "7", "10", "A", "Q", "7", "10", "J", "4", "5", "J", "9", "10", "Q", "J", "J", "K", "6", "10", "J", "6", "Q", "J", "5", "J", "Q", "Q", "8", "3", "8", "A", "2", "6", "9", "K", "7", "J", "K", "K", "8", "K", "Q", "6", "10", "J", "10", "J", "Q", "J", "10", "3", "8", "K", "A", "6", "9", "K", "2", "A", "A", "10", "J", "6", "A", "4", "J", "A", "J", "J", "6", "2", "J", "3", "K", "2", "5", "9", "J", "9", "6", "K", "A", "5", "Q", "J", "2", "Q", "K", "A", "3", "K", "J", "K", "2", "5", "6", "Q", "J", "Q", "Q", "J", "2", "J", "9", "Q", "7", "7", "A", "Q", "7", "Q", "J", "K", "J", "A", "7", "7", "8", "Q", "10", "J", "10", "J", "J", "9", "2", "A", "2"},
		playerB:     []string{"7", "2", "10", "K", "8", "2", "J", "9", "A", "5", "6", "J", "Q", "6", "K", "6", "5", "A", "4", "Q", "7", "J", "7", "10", "2", "Q", "8", "2", "2", "K", "J", "A", "5", "5", "A", "4", "Q", "6", "Q", "K", "10", "8", "Q", "2", "10", "J", "A", "Q", "8", "Q", "Q", "J", "J", "A", "A", "9", "10", "J", "K", "4", "Q", "10", "10", "J", "K", "10", "2", "J", "7", "A", "K", "K", "J", "A", "J", "10", "8", "K", "A", "7", "Q", "Q", "J", "3", "Q", "4", "A", "3", "A", "Q", "Q", "Q", "5", "4", "K", "J", "10", "A", "Q", "J", "6", "J", "A", "10", "A", "5", "8", "3", "K", "5", "9", "Q", "8", "7", "7", "J", "7", "Q", "Q", "Q", "A", "7", "8", "9", "A", "Q", "A", "K", "8", "A", "A", "J", "8", "4", "8", "K", "J", "A", "10", "Q", "8", "J", "8", "6", "10", "Q", "J", "J", "A", "A", "J", "5", "Q", "6", "J", "K", "Q", "8", "K", "4", "Q", "Q", "6", "J", "K", "4", "7", "J", "J", "9", "9", "A", "Q", "Q", "K", "A", "6", "5", "K"},
		expected: Outcome{
			finishes: true,
			cards:    361,
			tricks:   1,
		},
	},
	{
		description: "two tricks",
		playerA:     []string{"J"},
		playerB:     []string{"3", "J"},
		expected: Outcome{
			finishes: true,
			cards:    5,
			tricks:   2,
		},
	},
	{
		description: "more tricks",
		playerA:     []string{"J", "2", "4"},
		playerB:     []string{"3", "J", "A"},
		expected: Outcome{
			finishes: true,
			cards:    12,
			tricks:   4,
		},
	},
	{
		description: "simple loop with decks of 4 cards",
		playerA:     []string{"2", "3", "J", "6"},
		playerB:     []string{"K", "5", "J", "7"},
		expected: Outcome{
			finishes: false,
			cards:    16,
			tricks:   4,
		},
	},
	{
		description: "easy card combination",
		playerA:     []string{"4", "8", "7", "5", "4", "10", "3", "9", "7", "3", "10", "10", "6", "8", "2", "8", "5", "4", "5", "9", "6", "5", "2", "8", "10", "9"},
		playerB:     []string{"6", "9", "4", "7", "2", "2", "3", "6", "7", "3", "A", "A", "A", "A", "K", "K", "K", "K", "Q", "Q", "Q", "Q", "J", "J", "J", "J"},
		expected: Outcome{
			finishes: true,
			cards:    40,
			tricks:   4,
		},
	},
	{
		description: "easy card combination, inverted decks",
		playerA:     []string{"3", "3", "5", "7", "3", "2", "10", "7", "6", "7", "A", "A", "A", "A", "K", "K", "K", "K", "Q", "Q", "Q", "Q", "J", "J", "J", "J"},
		playerB:     []string{"5", "10", "8", "2", "6", "7", "2", "4", "9", "2", "6", "10", "10", "5", "4", "8", "4", "8", "6", "9", "8", "5", "9", "3", "4", "9"},
		expected: Outcome{
			finishes: true,
			cards:    40,
			tricks:   4,
		},
	},
	{
		description: "mirrored decks",
		playerA:     []string{"2", "A", "3", "A", "3", "K", "4", "K", "2", "Q", "2", "Q", "10", "J", "5", "J", "6", "10", "2", "9", "10", "7", "3", "9", "6", "9"},
		playerB:     []string{"6", "A", "4", "A", "7", "K", "4", "K", "7", "Q", "7", "Q", "5", "J", "8", "J", "4", "5", "8", "9", "10", "6", "8", "3", "8", "5"},
		expected: Outcome{
			finishes: true,
			cards:    59,
			tricks:   4,
		},
	},
	{
		description: "opposite decks",
		playerA:     []string{"4", "A", "9", "A", "4", "K", "9", "K", "6", "Q", "8", "Q", "8", "J", "10", "J", "9", "8", "4", "6", "3", "6", "5", "2", "4", "3"},
		playerB:     []string{"10", "7", "3", "2", "9", "2", "7", "8", "7", "5", "J", "7", "J", "10", "Q", "10", "Q", "3", "K", "5", "K", "6", "A", "2", "A", "5"},
		expected: Outcome{
			finishes: true,
			cards:    151,
			tricks:   21,
		},
	},
	{
		description: "random decks #1",
		playerA:     []string{"K", "10", "9", "8", "J", "8", "6", "9", "7", "A", "K", "5", "4", "4", "J", "5", "J", "4", "3", "5", "8", "6", "7", "7", "4", "9"},
		playerB:     []string{"6", "3", "K", "A", "Q", "10", "A", "2", "Q", "8", "2", "10", "10", "2", "Q", "3", "K", "9", "7", "A", "3", "Q", "5", "J", "2", "6"},
		expected: Outcome{
			finishes: true,
			cards:    542,
			tricks:   76,
		},
	},
	{
		description: "random decks #2",
		playerA:     []string{"8", "A", "4", "8", "5", "Q", "J", "2", "6", "2", "9", "7", "K", "A", "8", "10", "K", "8", "10", "9", "K", "6", "7", "3", "K", "9"},
		playerB:     []string{"10", "5", "2", "6", "Q", "J", "A", "9", "5", "5", "3", "7", "3", "J", "A", "2", "Q", "3", "J", "Q", "4", "10", "4", "7", "4", "6"},
		expected: Outcome{
			finishes: true,
			cards:    327,
			tricks:   42,
		},
	},
	{
		description: "Kleber 1999",
		playerA:     []string{"4", "8", "9", "J", "Q", "8", "5", "5", "K", "2", "A", "9", "8", "5", "10", "A", "4", "J", "3", "K", "6", "9", "2", "Q", "K", "7"},
		playerB:     []string{"10", "J", "3", "2", "4", "10", "4", "7", "5", "3", "6", "6", "7", "A", "J", "Q", "A", "7", "2", "10", "3", "K", "9", "6", "8", "Q"},
		expected: Outcome{
			finishes: true,
			cards:    5790,
			tricks:   805,
		},
	},
	{
		description: "Collins 2006",
		playerA:     []string{"A", "8", "Q", "K", "9", "10", "3", "7", "4", "2", "Q", "3", "2", "10", "9", "K", "A", "8", "7", "7", "4", "5", "J", "9", "2", "10"},
		playerB:     []string{"4", "J", "A", "K", "8", "5", "6", "6", "A", "6", "5", "Q", "4", "6", "10", "8", "J", "2", "5", "7", "Q", "J", "3", "3", "K", "9"},
		expected: Outcome{
			finishes: true,
			cards:    6913,
			tricks:   960,
		},
	},
	{
		description: "Mann and Wu 2007",
		playerA:     []string{"K", "2", "K", "K", "3", "3", "6", "10", "K", "6", "A", "2", "5", "5", "7", "9", "J", "A", "A", "3", "4", "Q", "4", "8", "J", "6"},
		playerB:     []string{"4", "5", "2", "Q", "7", "9", "9", "Q", "7", "J", "9", "8", "10", "3", "10", "J", "4", "10", "8", "6", "8", "7", "A", "Q", "5", "2"},
		expected: Outcome{
			finishes: true,
			cards:    7157,
			tricks:   1007,
		},
	},
	{
		description: "Nessler 2012",
		playerA:     []string{"10", "3", "6", "7", "Q", "2", "9", "8", "2", "8", "4", "A", "10", "6", "K", "2", "10", "A", "5", "A", "2", "4", "Q", "J", "K", "4"},
		playerB:     []string{"10", "Q", "4", "6", "J", "9", "3", "J", "9", "3", "3", "Q", "K", "5", "9", "5", "K", "6", "5", "7", "8", "J", "A", "7", "8", "7"},
		expected: Outcome{
			finishes: true,
			cards:    7207,
			tricks:   1015,
		},
	},
	{
		description: "Anderson 2013",
		playerA:     []string{"6", "7", "A", "3", "Q", "3", "5", "J", "3", "2", "J", "7", "4", "5", "Q", "10", "5", "A", "J", "2", "K", "8", "9", "9", "K", "3"},
		playerB:     []string{"4", "J", "6", "9", "8", "5", "10", "7", "9", "Q", "2", "7", "10", "8", "4", "10", "A", "6", "4", "A", "6", "8", "Q", "K", "K", "2"},
		expected: Outcome{
			finishes: true,
			cards:    7225,
			tricks:   1016,
		},
	},
	{
		description: "Rucklidge 2014",
		playerA:     []string{"8", "J", "2", "9", "4", "4", "5", "8", "Q", "3", "9", "3", "6", "2", "8", "A", "A", "A", "9", "4", "7", "2", "5", "Q", "Q", "3"},
		playerB:     []string{"K", "7", "10", "6", "3", "J", "A", "7", "6", "5", "5", "8", "10", "9", "10", "4", "2", "7", "K", "Q", "10", "K", "6", "J", "J", "K"},
		expected: Outcome{
			finishes: true,
			cards:    7959,
			tricks:   1122,
		},
	},
	{
		description: "Nessler 2021",
		playerA:     []string{"7", "2", "3", "4", "K", "9", "6", "10", "A", "8", "9", "Q", "7", "A", "4", "8", "J", "J", "A", "4", "3", "2", "5", "6", "6", "J"},
		playerB:     []string{"3", "10", "8", "9", "8", "K", "K", "2", "5", "5", "7", "6", "4", "3", "5", "7", "A", "9", "J", "K", "2", "Q", "10", "Q", "10", "Q"},
		expected: Outcome{
			finishes: true,
			cards:    7972,
			tricks:   1106,
		},
	},
	{
		description: "Nessler 2022",
		playerA:     []string{"2", "10", "10", "A", "J", "3", "8", "Q", "2", "5", "5", "5", "9", "2", "4", "3", "10", "Q", "A", "K", "Q", "J", "J", "9", "Q", "K"},
		playerB:     []string{"10", "7", "6", "3", "6", "A", "8", "9", "4", "3", "K", "J", "6", "K", "4", "9", "7", "8", "5", "7", "8", "2", "A", "7", "4", "6"},
		expected: Outcome{
			finishes: true,
			cards:    8344,
			tricks:   1164,
		},
	},
	{
		description: "Casella 2024, first infinite game found",
		playerA:     []string{"2", "8", "4", "K", "5", "2", "3", "Q", "6", "K", "Q", "A", "J", "3", "5", "9", "8", "3", "A", "A", "J", "4", "4", "J", "7", "5"},
		playerB:     []string{"7", "7", "8", "6", "10", "10", "6", "10", "7", "2", "Q", "6", "3", "2", "4", "K", "Q", "10", "J", "5", "9", "8", "9", "9", "K", "A"},
		expected: Outcome{
			finishes: false,
			cards:    474,
			tricks:   66,
		},
	},
}
