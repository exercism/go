package wordy

import "testing"

func TestAnswer(t *testing.T) {
	for _, test := range tests {
		switch answer, ok := Answer(test.question); {
		case !ok:
			if test.ok {
				t.Fatalf("FAIL: %s\nAnswer(%q)\nreturned ok = false, expecting true.", test.description, test.question)
			}
		case !test.ok:
			t.Errorf("FAIL: %s\nAnswer(%q)\nreturned %d, %t, expecting ok = false.", test.description, test.question, answer, ok)
		case answer != test.answer:
			t.Errorf("FAIL: %s\nAnswer(%q)\nreturned %d, expected %d.", test.description, test.question, answer, test.answer)
		}
		t.Logf("PASS: %s", test.description)
	}
}

// Benchmark combined time to answer all questions.
func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Answer(test.question)
		}
	}
}
