package stateoftictactoe

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 4326086 exercise state-of-tic-tac-toe: set proper value for 'reimplements' key of reimplemented test case (#2070)

var testCases = []struct {
	description string
	board       []string
	expected    State
	wantErr     bool
}{
	{
		description: "Finished game where X won via left column victory",
		board:       []string{"XOO", "X  ", "X  "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via middle column victory",
		board:       []string{"OXO", " X ", " X "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via right column victory",
		board:       []string{"OOX", "  X", "  X"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via left column victory",
		board:       []string{"OXX", "OX ", "O  "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via middle column victory",
		board:       []string{"XOX", " OX", " O "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via right column victory",
		board:       []string{"XXO", " XO", "  O"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via top row victory",
		board:       []string{"XXX", "XOO", "O  "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via middle row victory",
		board:       []string{"O  ", "XXX", " O "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via bottom row victory",
		board:       []string{" OO", "O X", "XXX"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via top row victory",
		board:       []string{"OOO", "XXO", "XX "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via middle row victory",
		board:       []string{"XX ", "OOO", "X  "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via bottom row victory",
		board:       []string{"XOX", " XX", "OOO"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via falling diagonal victory",
		board:       []string{"XOO", " X ", "  X"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via rising diagonal victory",
		board:       []string{"O X", "OX ", "X  "},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via falling diagonal victory",
		board:       []string{"OXX", "OOX", "X O"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where O won via rising diagonal victory",
		board:       []string{"  O", " OX", "OXX"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via a row and a column victory",
		board:       []string{"XXX", "XOO", "XOO"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Finished game where X won via two diagonal victories",
		board:       []string{"XOX", "OXO", "XOX"},
		expected:    Win,
		wantErr:     false,
	},
	{
		description: "Draw",
		board:       []string{"XOX", "XXO", "OXO"},
		expected:    Draw,
		wantErr:     false,
	},
	{
		description: "Another draw",
		board:       []string{"XXO", "OXX", "XOO"},
		expected:    Draw,
		wantErr:     false,
	},
	{
		description: "Ongoing game: one move in",
		board:       []string{"   ", "X  ", "   "},
		expected:    Ongoing,
		wantErr:     false,
	},
	{
		description: "Ongoing game: two moves in",
		board:       []string{"O  ", " X ", "   "},
		expected:    Ongoing,
		wantErr:     false,
	},
	{
		description: "Ongoing game: five moves in",
		board:       []string{"X  ", " XO", "OX "},
		expected:    Ongoing,
		wantErr:     false,
	},
	{
		description: "Invalid board: X went twice",
		board:       []string{"XX ", "   ", "   "},
		expected:    "",
		wantErr:     true,
	},
	{
		description: "Invalid board: O started",
		board:       []string{"OOX", "   ", "   "},
		expected:    "",
		wantErr:     true,
	},
	{
		description: "Invalid board: X won and O kept playing",
		board:       []string{"XXX", "OOO", "   "},
		expected:    "",
		wantErr:     true,
	},
	{
		description: "Invalid board: players kept playing after a win",
		board:       []string{"XXX", "OOO", "XOX"},
		expected:    "",
		wantErr:     true,
	},
}
