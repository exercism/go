package bottlesong

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 472204b bottle-song: Reimplement test cases checking for "One green bottles" (#2102)

type bottleSongInput struct {
	startBottles int
	takeDown     int
}

var testCases = []struct {
	description string
	input       bottleSongInput
	expected    []string
}{

	{
		description: "first generic verse",
		input: bottleSongInput{
			startBottles: 10,
			takeDown:     1,
		},
		expected: []string{"Ten green bottles hanging on the wall,", "Ten green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be nine green bottles hanging on the wall."},
	},
	{
		description: "last generic verse",
		input: bottleSongInput{
			startBottles: 3,
			takeDown:     1,
		},
		expected: []string{"Three green bottles hanging on the wall,", "Three green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be two green bottles hanging on the wall."},
	},
	{
		description: "verse with 2 bottles",
		input: bottleSongInput{
			startBottles: 2,
			takeDown:     1,
		},
		expected: []string{"Two green bottles hanging on the wall,", "Two green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be one green bottle hanging on the wall."},
	},
	{
		description: "verse with 1 bottle",
		input: bottleSongInput{
			startBottles: 1,
			takeDown:     1,
		},
		expected: []string{"One green bottle hanging on the wall,", "One green bottle hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be no green bottles hanging on the wall."},
	},
	{
		description: "first two verses",
		input: bottleSongInput{
			startBottles: 10,
			takeDown:     2,
		},
		expected: []string{"Ten green bottles hanging on the wall,", "Ten green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be nine green bottles hanging on the wall.", "", "Nine green bottles hanging on the wall,", "Nine green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be eight green bottles hanging on the wall."},
	},
	{
		description: "last three verses",
		input: bottleSongInput{
			startBottles: 3,
			takeDown:     3,
		},
		expected: []string{"Three green bottles hanging on the wall,", "Three green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be two green bottles hanging on the wall.", "", "Two green bottles hanging on the wall,", "Two green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be one green bottle hanging on the wall.", "", "One green bottle hanging on the wall,", "One green bottle hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be no green bottles hanging on the wall."},
	},
	{
		description: "all verses",
		input: bottleSongInput{
			startBottles: 10,
			takeDown:     10,
		},
		expected: []string{"Ten green bottles hanging on the wall,", "Ten green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be nine green bottles hanging on the wall.", "", "Nine green bottles hanging on the wall,", "Nine green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be eight green bottles hanging on the wall.", "", "Eight green bottles hanging on the wall,", "Eight green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be seven green bottles hanging on the wall.", "", "Seven green bottles hanging on the wall,", "Seven green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be six green bottles hanging on the wall.", "", "Six green bottles hanging on the wall,", "Six green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be five green bottles hanging on the wall.", "", "Five green bottles hanging on the wall,", "Five green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be four green bottles hanging on the wall.", "", "Four green bottles hanging on the wall,", "Four green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be three green bottles hanging on the wall.", "", "Three green bottles hanging on the wall,", "Three green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be two green bottles hanging on the wall.", "", "Two green bottles hanging on the wall,", "Two green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be one green bottle hanging on the wall.", "", "One green bottle hanging on the wall,", "One green bottle hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be no green bottles hanging on the wall."},
	},
}
