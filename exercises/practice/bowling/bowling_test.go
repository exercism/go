package bowling

import "testing"

const previousRollErrorMessage = `
	Unexpected error occurred: %v
	while applying the previous rolls for the
	test case: %v
	The error was returned from Roll(%d) for previousRolls[%d].`

func applyPreviousRolls(g *Game, rolls []int) (index, pins int, err error) {
	for index, pins := range rolls {
		if err := g.Roll(pins); err != nil {
			return index, pins, err
		}
	}
	return 0, 0, nil
}

func TestRoll(t *testing.T) {
	for _, tc := range rollTestCases {
		t.Run(tc.description, func(t *testing.T) {
			g := NewGame()
			index, pins, err := applyPreviousRolls(g, tc.previousRolls)
			if err != nil {
				t.Fatalf(previousRollErrorMessage, err, tc.previousRolls, pins, index)
			}
			err = g.Roll(tc.roll)
			if tc.valid && err != nil {
				t.Fatalf("Roll(%d) after Previous Rolls: %#v returned unexpected error: %v", tc.roll, tc.previousRolls, err)
			} else if !tc.valid && err == nil {
				t.Fatalf("Roll(%d) after Previous Rolls: %#v expected an error, got nil\n\tExplanation: %s", tc.roll, tc.previousRolls, tc.explainText)
			}
		})
	}
}

func TestScore(t *testing.T) {
	for _, tc := range scoreTestCases {
		t.Run(tc.description, func(t *testing.T) {
			g := NewGame()
			index, pins, err := applyPreviousRolls(g, tc.previousRolls)
			if err != nil {
				t.Fatalf(previousRollErrorMessage, err, tc.previousRolls, pins, index)
			}
			score, err := g.Score()
			switch {
			case !tc.valid:
				if err == nil {
					t.Fatalf("Score() after Previous Rolls: %#v expected an error, got score %d\n\tExplanation: %s", tc.previousRolls, score, tc.explainText)
				}
			case err != nil:
				t.Fatalf("Score() after Previous Rolls: %#v returned error: %v, want: %d", tc.previousRolls, err, tc.score)
			case score != tc.score:
				t.Fatalf("Score() after Previous Rolls: %#v = %d, want: %d", tc.previousRolls, score, tc.score)
			}
		})
	}
}
