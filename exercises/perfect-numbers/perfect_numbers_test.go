package perfect

import "testing"

var _ error = ErrOnlyPositive

var classificationTestCases = []struct {
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

// Test that the classifications are not equal to each other.
// If they are equal, then the tests will return false positives.
func TestClassificationsNotEqual(t *testing.T) {
	classifications := []struct {
		class Classification
		name  string
	}{
		{ClassificationAbundant, "ClassificationAbundant"},
		{ClassificationDeficient, "ClassificationDeficient"},
		{ClassificationPerfect, "ClassificationPerfect"},
	}

	for i, pair1 := range classifications {
		for j := i + 1; j < len(classifications); j++ {
			pair2 := classifications[j]
			if pair1.class == pair2.class {
				t.Fatalf("%s should not be equal to %s", pair1.name, pair2.name)
			}
		}
	}
}

func BenchmarkClassify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range classificationTestCases {
			Classify(c.input)
		}
	}
}
