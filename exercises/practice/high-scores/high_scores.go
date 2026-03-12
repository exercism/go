package highscores

type HighScores struct{}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	panic("Please implement the NewHighScores function")
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	panic("Please implement the Scores function")
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	panic("Please implement the Latest function")
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	panic("Please implement the PersonalBest function")
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	panic("Please implement the TopThree function")
}
