package perfect

import "testing"

var _ error = ErrOnlyPositive

func TestGivesPositiveRequiredError(t *testing.T) {
	if _, err := Classify(0); err != ErrOnlyPositive {
		t.Errorf("Expected error %q but got %q", ErrOnlyPositive, err)
	}
}

func TestClassifiesCorrectly(t *testing.T) {
	cases := []struct {
		input    uint64
		expected Classification
	}{
		{1, ClassificationDeficient},
		{13, ClassificationDeficient},
		{12, ClassificationAbundant},
		{6, ClassificationPerfect},
		{28, ClassificationPerfect},
		{496, ClassificationPerfect},
		{8128, ClassificationPerfect},
	}
	for _, c := range cases {
		if cat, err := Classify(c.input); err != nil {
			t.Errorf("%d: Expected no error but got %s", c.input, err)
		} else if cat != c.expected {
			t.Errorf("%d: Expected %q, got %q", c.input, c.expected, cat)
		}
	}
}

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}
