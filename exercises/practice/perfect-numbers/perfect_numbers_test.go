package perfect

import "testing"

func TestZeroGivesPositiveRequiredError(t *testing.T) {
	t.Run("GivesPositiveRequiredError", func(t *testing.T) {
		if _, err := Classify(0); err != ErrOnlyPositive {
			t.Fatalf("Classify(0) expected error %q, got: %q", ErrOnlyPositive, err)
		}
	})
}

func TestClassifiesCorrectly(t *testing.T) {
	for _, tc := range classificationTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Classify(tc.input)
			switch {
			case !tc.ok:
				// expect error
				if err == nil {
					t.Fatalf("Classify(%d) expected error, got: %q", tc.input, actual)
				}
			case err != nil:
				t.Fatalf("Classify(%d) returned error: %q, want: %q", tc.input, err, tc.expected)
			case actual != tc.expected:
				t.Fatalf("Classify(%d) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range classificationTestCases {
			Classify(c.input)
		}
	}
}
