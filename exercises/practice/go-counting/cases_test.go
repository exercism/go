package gocounting

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 2e820e1 Auto-format portions of some JSON files (#1967)

var oneTerritoryTestCases = []struct {
	description string
	board       []string
	posX        int
	posY        int
	owner       string
	territory   [][2]int
	expectedErr string
}{
	{
		description: "Black corner territory on 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        0,
		posY:        1,
		owner:       "BLACK",
		territory:   [][2]int{[2]int{0, 0}, [2]int{0, 1}, [2]int{1, 0}},
		expectedErr: "",
	},
	{
		description: "White center territory on 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        2,
		posY:        3,
		owner:       "WHITE",
		territory:   [][2]int{[2]int{2, 3}},
		expectedErr: "",
	},
	{
		description: "Open corner territory on 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        1,
		posY:        4,
		owner:       "NONE",
		territory:   [][2]int{[2]int{0, 3}, [2]int{0, 4}, [2]int{1, 4}},
		expectedErr: "",
	},
	{
		description: "A stone and not a territory on 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        1,
		posY:        1,
		owner:       "NONE",
		territory:   [][2]int{},
		expectedErr: "",
	},
	{
		description: "Invalid because X is too low for 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        -1,
		posY:        1,
		owner:       "",
		territory:   [][2]int(nil),
		expectedErr: "invalid coordinate",
	},
	{
		description: "Invalid because X is too high for 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        5,
		posY:        1,
		owner:       "",
		territory:   [][2]int(nil),
		expectedErr: "invalid coordinate",
	},
	{
		description: "Invalid because Y is too low for 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        1,
		posY:        -1,
		owner:       "",
		territory:   [][2]int(nil),
		expectedErr: "invalid coordinate",
	},
	{
		description: "Invalid because Y is too high for 5x5 board",
		board:       []string{"  B  ", " B B ", "B W B", " W W ", "  W  "},
		posX:        1,
		posY:        5,
		owner:       "",
		territory:   [][2]int(nil),
		expectedErr: "invalid coordinate",
	},
}

var allTerritoriesTestCases = []struct {
	description string
	input       []string
	expected    AllTerritories
}{
	{
		description: "One territory is the whole board",
		input:       []string{" "},
		expected: AllTerritories{
			Black: [][2]int{},
			White: [][2]int{},
			None:  [][2]int{[2]int{0, 0}},
		},
	},
	{
		description: "Two territory rectangular board",
		input:       []string{" BW ", " BW "},
		expected: AllTerritories{
			Black: [][2]int{[2]int{0, 0}, [2]int{0, 1}},
			White: [][2]int{[2]int{3, 0}, [2]int{3, 1}},
			None:  [][2]int{},
		},
	},
	{
		description: "Two region rectangular board",
		input:       []string{" B "},
		expected: AllTerritories{
			Black: [][2]int{[2]int{0, 0}, [2]int{2, 0}},
			White: [][2]int{},
			None:  [][2]int{},
		},
	},
}
