package highscores

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 38aef6b High Score: drop implementation details from the problem spec (#2295)

var intTestCases = []struct {
	description string
	input       []int
	call        func(h *HighScores) int
	callMsg     string
	expected    int
}{
	{
		description: "Latest score",
		input:       []int{100, 0, 90, 30},
		call:        func(h *HighScores) int { return h.Latest() },
		callMsg:     "h.Latest()",
		expected:    30,
	},
	{
		description: "Personal best",
		input:       []int{40, 100, 70},
		call:        func(h *HighScores) int { return h.PersonalBest() },
		callMsg:     "h.PersonalBest()",
		expected:    100,
	},
	{
		description: "Latest score after personal best",
		input:       []int{20, 70, 15, 25, 30},
		call:        func(h *HighScores) int { h.PersonalBest(); return h.Latest() },
		callMsg:     "h.PersonalBest(); h.Latest()",
		expected:    30,
	},
	{
		description: "Latest score after personal top scores",
		input:       []int{70, 50, 20, 30},
		call:        func(h *HighScores) int { h.TopThree(); return h.Latest() },
		callMsg:     "h.TopThree(); h.Latest()",
		expected:    30,
	},
}

var sliceTestCases = []struct {
	description string
	input       []int
	call        func(h *HighScores) []int
	callMsg     string
	expected    []int
}{
	{
		description: "List of scores",
		input:       []int{30, 50, 20, 70},
		call:        func(h *HighScores) []int { return h.Scores() },
		callMsg:     "h.Scores()",
		expected:    []int{30, 50, 20, 70},
	},
	{
		description: "Personal top three from a list of scores",
		input:       []int{10, 30, 90, 30, 100, 20, 10, 0, 30, 40, 40, 70, 70},
		call:        func(h *HighScores) []int { return h.TopThree() },
		callMsg:     "h.TopThree()",
		expected:    []int{100, 90, 70},
	},
	{
		description: "Personal top highest to lowest",
		input:       []int{20, 10, 30},
		call:        func(h *HighScores) []int { return h.TopThree() },
		callMsg:     "h.TopThree()",
		expected:    []int{30, 20, 10},
	},
	{
		description: "Personal top when there is a tie",
		input:       []int{40, 20, 40, 30},
		call:        func(h *HighScores) []int { return h.TopThree() },
		callMsg:     "h.TopThree()",
		expected:    []int{40, 40, 30},
	},
	{
		description: "Personal top when there are less than 3",
		input:       []int{30, 70},
		call:        func(h *HighScores) []int { return h.TopThree() },
		callMsg:     "h.TopThree()",
		expected:    []int{70, 30},
	},
	{
		description: "Personal top when there is only one",
		input:       []int{40},
		call:        func(h *HighScores) []int { return h.TopThree() },
		callMsg:     "h.TopThree()",
		expected:    []int{40},
	},
	{
		description: "Scores after personal best",
		input:       []int{20, 70, 15, 25, 30},
		call:        func(h *HighScores) []int { h.PersonalBest(); return h.Scores() },
		callMsg:     "h.PersonalBest(); h.Scores()",
		expected:    []int{20, 70, 15, 25, 30},
	},
	{
		description: "Scores after personal top scores",
		input:       []int{30, 50, 20, 70},
		call:        func(h *HighScores) []int { h.TopThree(); return h.Scores() },
		callMsg:     "h.TopThree(); h.Scores()",
		expected:    []int{30, 50, 20, 70},
	},
}
