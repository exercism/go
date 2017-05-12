package perfect

import "testing"

var _ error = ErrOnlyPositive

var classificationTestCases = []struct {
	input    uint64
	expected Classification
}{
	// converted from x-common by hand
	// missing: 33550336, 30, 33550335, 2, 4, 32, 33550337, 0
	// -1 won't be tested
	// x-common version: 0.2.0
	{1, ClassificationDeficient},
	{13, ClassificationDeficient},
	{12, ClassificationAbundant},
	{6, ClassificationPerfect},
	{28, ClassificationPerfect},
	{496, ClassificationPerfect},
	{8128, ClassificationPerfect},
}

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestGivesPositiveRequiredError(t *testing.T) {
	if _, err := Classify(0); err != ErrOnlyPositive {
		t.Errorf("Expected error %q but got %q", ErrOnlyPositive, err)
	}
}

func TestClassifiesCorrectly(t *testing.T) {
	for _, c := range classificationTestCases {
		if cat, err := Classify(c.input); err != nil {
			t.Errorf("%d: Expected no error but got %s", c.input, err)
		} else if cat != c.expected {
			t.Errorf("%d: Expected %q, got %q", c.input, c.expected, cat)
		}
	}
}
