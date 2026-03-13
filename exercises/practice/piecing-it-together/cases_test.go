package piecingittogether

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 087e5e3 Add 'piecing-it-together' exercise (#2554)

var testCases = []struct {
	description string
	input       PuzzleDetails
	expected    PuzzleDetails
	err         string
	wantError   bool
}{
	{
		description: "1000 pieces puzzle with 1.6 aspect ratio",
		input: PuzzleDetails{
			Pieces:      1000,
			AspectRatio: 1.6,
		},
		expected: PuzzleDetails{
			Pieces:      1000,
			Border:      126,
			Inside:      874,
			Rows:        25,
			Columns:     40,
			AspectRatio: 1.6,
			Format:      "landscape",
		},
		err:       "",
		wantError: false,
	},
	{
		description: "square puzzle with 32 rows",
		input: PuzzleDetails{
			Rows:   32,
			Format: "square",
		},
		expected: PuzzleDetails{
			Pieces:      1024,
			Border:      124,
			Inside:      900,
			Rows:        32,
			Columns:     32,
			AspectRatio: 1,
			Format:      "square",
		},
		err:       "",
		wantError: false,
	},
	{
		description: "400 pieces square puzzle with only inside pieces and aspect ratio",
		input: PuzzleDetails{
			Inside:      324,
			AspectRatio: 1,
		},
		expected: PuzzleDetails{
			Pieces:      400,
			Border:      76,
			Inside:      324,
			Rows:        20,
			Columns:     20,
			AspectRatio: 1,
			Format:      "square",
		},
		err:       "",
		wantError: false,
	},
	{
		description: "1500 pieces landscape puzzle with 30 rows and 1.6 aspect ratio",
		input: PuzzleDetails{
			Rows:        30,
			AspectRatio: 1.6666666666666667,
		},
		expected: PuzzleDetails{
			Pieces:      1500,
			Border:      156,
			Inside:      1344,
			Rows:        30,
			Columns:     50,
			AspectRatio: 1.6666666666666667,
			Format:      "landscape",
		},
		err:       "",
		wantError: false,
	},
	{
		description: "300 pieces portrait puzzle with 70 border pieces",
		input: PuzzleDetails{
			Pieces: 300,
			Border: 70,
			Format: "portrait",
		},
		expected: PuzzleDetails{
			Pieces:      300,
			Border:      70,
			Inside:      230,
			Rows:        25,
			Columns:     12,
			AspectRatio: 0.48,
			Format:      "portrait",
		},
		err:       "",
		wantError: false,
	},
	{
		description: "puzzle with insufficient data",
		input: PuzzleDetails{
			Pieces: 1500,
			Format: "landscape",
		},
		expected:  PuzzleDetails{},
		err:       "Insufficient data",
		wantError: true,
	},
	{
		description: "puzzle with contradictory data",
		input: PuzzleDetails{
			Rows:    100,
			Columns: 1000,
			Format:  "square",
		},
		expected:  PuzzleDetails{},
		err:       "Contradictory data",
		wantError: true,
	},
}
