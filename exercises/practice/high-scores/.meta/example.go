package highscores

import "slices"

type HighScores struct {
	scores []int
}

func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores}
}

func (s *HighScores) Scores() []int {
	return s.scores
}

func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

func (s *HighScores) PersonalBest() int {
	scores := slices.Clone(s.scores)
	slices.Sort(scores)
	return scores[len(scores)-1]
}

func (s *HighScores) TopThree() []int {
	scores := slices.Clone(s.scores)
	slices.Sort(scores)
	slices.Reverse(scores)
	return scores[:min(3, len(scores))]
}
