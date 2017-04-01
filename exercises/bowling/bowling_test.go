package bowling

import "testing"

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func applyPreviousRolls(g *Game, rolls []int) error {
	for _, pins := range rolls {
		if err := g.Roll(pins); err != nil {
			return err
		}
	}
	return nil
}

func TestScore(t *testing.T) {
	for _, tc := range scoreTestCases {
		g := NewGame()
		err := applyPreviousRolls(g, tc.previousRolls)
		if err != nil {
			t.Fatalf("%s :\n	unexpected error \"%s\" applying Previous Rolls: %#v before Score()",
				tc.description, err, tc.previousRolls)
		}
		score, err := g.Score()
		if tc.valid {
			if err != nil {
				t.Fatalf("%s : Score() after Previous Rolls: %#v expected %d, got error %s",
					tc.description, tc.previousRolls, tc.score, err)
			} else {
				if score != tc.score {
					t.Fatalf("%s : Score() after Previous Rolls: %#v expected %d, got %d",
						tc.description, tc.previousRolls, tc.score, score)
				}
			}
		} else if err == nil {
			t.Fatalf("%s : Score() after Previous Rolls: %#v expected error \"%s\", got score %d",
				tc.description, tc.previousRolls, tc.errorString, score)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestRoll(t *testing.T) {
	for _, tc := range rollTestCases {
		g := NewGame()
		err := applyPreviousRolls(g, tc.previousRolls)
		if err != nil {
			t.Fatalf("%s :\n	unexpected error \"%s\" applying Previous Rolls: %#v before Roll(%d)",
				tc.description, err, tc.previousRolls, tc.roll)
		}
		err = g.Roll(tc.roll)
		if tc.valid && err != nil {
			t.Fatalf("%s : Roll(%d) after Previous Rolls: %#v got unexpected error \"%s\".",
				tc.description, tc.roll, tc.previousRolls, err)

		} else if !tc.valid && err == nil {
			t.Fatalf("%s : Roll(%d) after Previous Rolls: %#v did not get expected error \"%s\".",
				tc.description, tc.roll, tc.previousRolls, tc.errorString)
		}
		t.Logf("PASS: %s", tc.description)
	}
}
