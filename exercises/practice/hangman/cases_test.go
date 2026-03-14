package hangman

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: cfbb6ab Add UUIDs

var testCases = []struct {
	description       string
	word              string
	guesses           []rune
	state             string
	maskedWord        string
	remainingFailures int
	error             string
}{
	{
		description:       "Initially 9 failures are allowed and no letters are guessed",
		word:              "loot",
		guesses:           []rune{},
		state:             "Ongoing",
		maskedWord:        "____",
		remainingFailures: 9,
		error:             "",
	},
	{
		description:       "After 10 failures the game is over",
		word:              "loot",
		guesses:           []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'},
		state:             "Lose",
		maskedWord:        "____",
		remainingFailures: 0,
		error:             "",
	},
	{
		description:       "Losing with several correct guesses",
		word:              "loot",
		guesses:           []rune{'t', 'o', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'},
		state:             "Lose",
		maskedWord:        "_oot",
		remainingFailures: 0,
		error:             "",
	},
	{
		description:       "Feeding a correct letter removes underscores",
		word:              "loot",
		guesses:           []rune{'t'},
		state:             "Ongoing",
		maskedWord:        "___t",
		remainingFailures: 9,
		error:             "",
	},
	{
		description:       "Feeding a correct letter twice counts as a failure",
		word:              "loot",
		guesses:           []rune{'t', 't'},
		state:             "Ongoing",
		maskedWord:        "___t",
		remainingFailures: 8,
		error:             "",
	},
	{
		description:       "Guessing a repeated letter reveals all instances",
		word:              "loot",
		guesses:           []rune{'t', 't', 'o'},
		state:             "Ongoing",
		maskedWord:        "_oot",
		remainingFailures: 8,
		error:             "",
	},
	{
		description:       "Getting all the letters right makes for a win",
		word:              "loot",
		guesses:           []rune{'t', 't', 'o', 'l'},
		state:             "Win",
		maskedWord:        "loot",
		remainingFailures: 8,
		error:             "",
	},
	{
		description:       "Winning on the last guess is still a win",
		word:              "loot",
		guesses:           []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 't', 'o', 'l'},
		state:             "Win",
		maskedWord:        "loot",
		remainingFailures: 0,
		error:             "",
	},
	{
		description:       "Guessing after a lose is error",
		word:              "loot",
		guesses:           []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'},
		state:             "",
		maskedWord:        "",
		remainingFailures: 0,
		error:             "cannot guess after the game is lost",
	},
	{
		description:       "Guessing after a win is error",
		word:              "loot",
		guesses:           []rune{'t', 'o', 'l', 'l'},
		state:             "",
		maskedWord:        "",
		remainingFailures: 0,
		error:             "cannot guess after the game is won",
	},
}
