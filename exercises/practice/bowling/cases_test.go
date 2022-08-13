package bowling

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: daf84d6 bowling, transpose: conform array format to rest of file

var scoreTestCases = []struct {
	description   string
	previousRolls []int  // bowling rolls to do before the Score() test
	valid         bool   // true => no error, false => error expected
	score         int    // when .valid == true, the expected score value
	explainText   string // when .valid == false, error explanation text
}{
	{
		description:   "should be able to score a game with all zeros",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         true,
		score:         0,
		explainText:   "",
	},
	{
		description:   "should be able to score a game with no strikes or spares",
		previousRolls: []int{3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6},
		valid:         true,
		score:         90,
		explainText:   "",
	},
	{
		description:   "a spare followed by zeros is worth ten points",
		previousRolls: []int{6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         true,
		score:         10,
		explainText:   "",
	},
	{
		description:   "points scored in the roll after a spare are counted twice",
		previousRolls: []int{6, 4, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         true,
		score:         16,
		explainText:   "",
	},
	{
		description:   "consecutive spares each get a one roll bonus",
		previousRolls: []int{5, 5, 3, 7, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         true,
		score:         31,
		explainText:   "",
	},
	{
		description:   "a spare in the last frame gets a one roll bonus that is counted once",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 7},
		valid:         true,
		score:         17,
		explainText:   "",
	},
	{
		description:   "a strike earns ten points in a frame with a single roll",
		previousRolls: []int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         true,
		score:         10,
		explainText:   "",
	},
	{
		description:   "points scored in the two rolls after a strike are counted twice as a bonus",
		previousRolls: []int{10, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         true,
		score:         26,
		explainText:   "",
	},
	{
		description:   "consecutive strikes each get the two roll bonus",
		previousRolls: []int{10, 10, 10, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         true,
		score:         81,
		explainText:   "",
	},
	{
		description:   "a strike in the last frame gets a two roll bonus that is counted once",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 7, 1},
		valid:         true,
		score:         18,
		explainText:   "",
	},
	{
		description:   "rolling a spare with the two roll bonus does not get a bonus roll",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 7, 3},
		valid:         true,
		score:         20,
		explainText:   "",
	},
	{
		description:   "strikes with the two roll bonus do not get bonus rolls",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 10},
		valid:         true,
		score:         30,
		explainText:   "",
	},
	{
		description:   "last two strikes followed by only last bonus with non strike points",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 0, 1},
		valid:         true,
		score:         31,
		explainText:   "",
	},
	{
		description:   "a strike with the one roll bonus after a spare in the last frame does not get a bonus",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 10},
		valid:         true,
		score:         20,
		explainText:   "",
	},
	{
		description:   "all strikes is a perfect game",
		previousRolls: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		valid:         true,
		score:         300,
		explainText:   "",
	},
	{
		description:   "two bonus rolls after a strike in the last frame can score more than 10 points if one is a strike",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 6},
		valid:         true,
		score:         26,
		explainText:   "",
	},
	{
		description:   "an unstarted game cannot be scored",
		previousRolls: []int{},
		valid:         false,
		score:         0,
		explainText:   "Score cannot be taken until the end of the game",
	},
	{
		description:   "an incomplete game cannot be scored",
		previousRolls: []int{0, 0},
		valid:         false,
		score:         0,
		explainText:   "Score cannot be taken until the end of the game",
	},
	{
		description:   "bonus rolls for a strike in the last frame must be rolled before score can be calculated",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10},
		valid:         false,
		score:         0,
		explainText:   "Score cannot be taken until the end of the game",
	},
	{
		description:   "both bonus rolls for a strike in the last frame must be rolled before score can be calculated",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10},
		valid:         false,
		score:         0,
		explainText:   "Score cannot be taken until the end of the game",
	},
	{
		description:   "bonus roll for a spare in the last frame must be rolled before score can be calculated",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3},
		valid:         false,
		score:         0,
		explainText:   "Score cannot be taken until the end of the game",
	},
}

var rollTestCases = []struct {
	description   string
	previousRolls []int  // bowling rolls to do before the Roll(roll) test
	valid         bool   // true => no error, false => error expected
	roll          int    // pin count for the test roll
	explainText   string // when .valid == false, error explanation text
}{
	{
		description:   "rolls cannot score negative points",
		previousRolls: []int{},
		valid:         false,
		roll:          -1,
		explainText:   "Negative roll is invalid",
	},
	{
		description:   "a roll cannot score more than 10 points",
		previousRolls: []int{},
		valid:         false,
		roll:          11,
		explainText:   "Pin count exceeds pins on the lane",
	},
	{
		description:   "two rolls in a frame cannot score more than 10 points",
		previousRolls: []int{5},
		valid:         false,
		roll:          6,
		explainText:   "Pin count exceeds pins on the lane",
	},
	{
		description:   "bonus roll after a strike in the last frame cannot score more than 10 points",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10},
		valid:         false,
		roll:          11,
		explainText:   "Pin count exceeds pins on the lane",
	},
	{
		description:   "two bonus rolls after a strike in the last frame cannot score more than 10 points",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 5},
		valid:         false,
		roll:          6,
		explainText:   "Pin count exceeds pins on the lane",
	},
	{
		description:   "the second bonus rolls after a strike in the last frame cannot be a strike if the first one is not a strike",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 6},
		valid:         false,
		roll:          10,
		explainText:   "Pin count exceeds pins on the lane",
	},
	{
		description:   "second bonus roll after a strike in the last frame cannot score more than 10 points",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10},
		valid:         false,
		roll:          11,
		explainText:   "Pin count exceeds pins on the lane",
	},
	{
		description:   "cannot roll if game already has ten frames",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		valid:         false,
		roll:          0,
		explainText:   "Cannot roll after game is over",
	},
	{
		description:   "cannot roll after bonus roll for spare",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 2},
		valid:         false,
		roll:          2,
		explainText:   "Cannot roll after game is over",
	},
	{
		description:   "cannot roll after bonus rolls for strike",
		previousRolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 3, 2},
		valid:         false,
		roll:          2,
		explainText:   "Cannot roll after game is over",
	},
}
