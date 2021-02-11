package bowling

// Source: exercism/problem-specifications
// Commit: 1806718 bowling: add tests for rolling after bonus rolls
// Problem Specifications Version: 1.2.0

var scoreTestCases = []struct {
	description   string
	previousRolls []int  // bowling rolls to do before the Score() test
	valid         bool   // true => no error, false => error expected
	score         int    // when .valid == true, the expected score value
	explainText   string // when .valid == false, error explanation text
}{
	{
		"should be able to score a game with all zeros",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		true,
		0,
		"",
	},
	{
		"should be able to score a game with no strikes or spares",
		[]int{3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6},
		true,
		90,
		"",
	},
	{
		"a spare followed by zeros is worth ten points",
		[]int{6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		true,
		10,
		"",
	},
	{
		"points scored in the roll after a spare are counted twice",
		[]int{6, 4, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		true,
		16,
		"",
	},
	{
		"consecutive spares each get a one roll bonus",
		[]int{5, 5, 3, 7, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		true,
		31,
		"",
	},
	{
		"a spare in the last frame gets a one roll bonus that is counted once",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 7},
		true,
		17,
		"",
	},
	{
		"a strike earns ten points in a frame with a single roll",
		[]int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		true,
		10,
		"",
	},
	{
		"points scored in the two rolls after a strike are counted twice as a bonus",
		[]int{10, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		true,
		26,
		"",
	},
	{
		"consecutive strikes each get the two roll bonus",
		[]int{10, 10, 10, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		true,
		81,
		"",
	},
	{
		"a strike in the last frame gets a two roll bonus that is counted once",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 7, 1},
		true,
		18,
		"",
	},
	{
		"rolling a spare with the two roll bonus does not get a bonus roll",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 7, 3},
		true,
		20,
		"",
	},
	{
		"strikes with the two roll bonus do not get bonus rolls",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 10},
		true,
		30,
		"",
	},
	{
		"a strike with the one roll bonus after a spare in the last frame does not get a bonus",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 10},
		true,
		20,
		"",
	},
	{
		"all strikes is a perfect game",
		[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		true,
		300,
		"",
	},

	{
		"two bonus rolls after a strike in the last frame can score more than 10 points if one is a strike",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 6},
		true,
		26,
		"",
	},

	{
		"an unstarted game cannot be scored",
		[]int{},
		false,
		0,
		"Score cannot be taken until the end of the game",
	},
	{
		"an incomplete game cannot be scored",
		[]int{0, 0},
		false,
		0,
		"Score cannot be taken until the end of the game",
	},

	{
		"bonus rolls for a strike in the last frame must be rolled before score can be calculated",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10},
		false,
		0,
		"Score cannot be taken until the end of the game",
	},
	{
		"both bonus rolls for a strike in the last frame must be rolled before score can be calculated",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10},
		false,
		0,
		"Score cannot be taken until the end of the game",
	},
	{
		"bonus roll for a spare in the last frame must be rolled before score can be calculated",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3},
		false,
		0,
		"Score cannot be taken until the end of the game",
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
		"rolls cannot score negative points",
		[]int{},
		false,
		-1,
		"Negative roll is invalid",
	},
	{
		"a roll cannot score more than 10 points",
		[]int{},
		false,
		11,
		"Pin count exceeds pins on the lane",
	},
	{
		"two rolls in a frame cannot score more than 10 points",
		[]int{5},
		false,
		6,
		"Pin count exceeds pins on the lane",
	},
	{
		"bonus roll after a strike in the last frame cannot score more than 10 points",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10},
		false,
		11,
		"Pin count exceeds pins on the lane",
	},
	{
		"two bonus rolls after a strike in the last frame cannot score more than 10 points",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 5},
		false,
		6,
		"Pin count exceeds pins on the lane",
	},

	{
		"the second bonus rolls after a strike in the last frame cannot be a strike if the first one is not a strike",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 6},
		false,
		10,
		"Pin count exceeds pins on the lane",
	},
	{
		"second bonus roll after a strike in the last frame cannot score more than 10 points",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10},
		false,
		11,
		"Pin count exceeds pins on the lane",
	},

	{
		"cannot roll if game already has ten frames",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		false,
		0,
		"Cannot roll after game is over",
	},

	{
		"cannot roll after bonus roll for spare",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 2},
		false,
		2,
		"Cannot roll after game is over",
	},
	{
		"cannot roll after bonus rolls for strike",
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 3, 2},
		false,
		2,
		"Cannot roll after game is over",
	},
}
