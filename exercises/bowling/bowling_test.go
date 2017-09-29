package bowling

import "testing"

const previousRollErrorMessage = `FAIL: %s
	Unexpected error occurred: %q
	while applying the previous rolls for the
	test case: %v
	The error was returned from Roll(%d) for previousRolls[%d].`

func applyPreviousRolls(g *Game, rolls []int) (int, int, error) {
	for index, pins := range rolls {
		var err error
		if err = g.Roll(pins); err != nil {
			return index, pins, err
		}
	}
	return 0, 0, nil
}

func TestScore(t *testing.T) {
	for _, tc := range scoreTestCases {
		g := NewGame()
		index, pins, err := applyPreviousRolls(g, tc.previousRolls)
		if err != nil {
			t.Fatalf(previousRollErrorMessage,
				tc.description, err, tc.previousRolls, pins, index)
		}
		score, err := g.Score()
		if tc.valid {
			var _ error = err
			if err != nil {
				t.Fatalf("FAIL: %s : Score() after Previous Rolls: %#v expected %d, got error %s",
					tc.description, tc.previousRolls, tc.score, err)
			} else {
				if score != tc.score {
					t.Fatalf("%s : Score() after Previous Rolls: %#v expected %d, got %d",
						tc.description, tc.previousRolls, tc.score, score)
				}
			}
		} else if err == nil {
			t.Fatalf("FAIL: %s : Score() after Previous Rolls: %#v expected an error, got score %d\n\tExplanation: %s",
				tc.description, tc.previousRolls, score, tc.explainText)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestRoll(t *testing.T) {
	for _, tc := range rollTestCases {
		g := NewGame()
		index, pins, err := applyPreviousRolls(g, tc.previousRolls)
		if err != nil {
			t.Fatalf(previousRollErrorMessage,
				tc.description, err, tc.previousRolls, pins, index)
		}
		err = g.Roll(tc.roll)
		if tc.valid && err != nil {
			var _ error = err
			t.Fatalf("FAIL: %s : Roll(%d) after Previous Rolls: %#v got unexpected error \"%s\".",
				tc.description, tc.roll, tc.previousRolls, err)

		} else if !tc.valid && err == nil {
			t.Fatalf("FAIL: %s : Roll(%d) after Previous Rolls: %#v expected an error.\n\tExplanation: %s",
				tc.description, tc.roll, tc.previousRolls, tc.explainText)
		}
		t.Logf("PASS: %s", tc.description)
	}
}
