package hangman

import (
	"fmt"
	"testing"
)

func TestGuess(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			var err error
			setup := fmt.Sprintf("Word: %s. Guesses: %c.", tc.word, tc.guesses)
			final := ""
			if len(tc.guesses) > 0 {
				final = fmt.Sprintf("Final Guess('%c')", tc.guesses[len(tc.guesses)-1])
			}
			g := NewGame(tc.word)
			for _, c := range tc.guesses {
				err = g.Guess(c)
			}
			if tc.error != "" {
				if err == nil {
					t.Fatalf("%s\n%s did not return an error; expected %v", setup, final, tc.error)
				} else if err.Error() != tc.error {
					t.Fatalf("%s\n%s = %v; expected %v", setup, final, err.Error(), tc.error)
				}
			} else if err != nil {
				t.Fatalf("%s\n%s returned unexpected %v", setup, final, err.Error())
			} else if got := g.State(); got != tc.state {
				t.Fatalf("%s\nState() = %q, want %q", setup, got, tc.state)
			} else if got := g.MaskedWord(); got != tc.maskedWord {
				t.Fatalf("%s\nMaskedWord() = %q, want %q", setup, got, tc.maskedWord)
			} else if got := g.RemainingGuesses(); got != tc.remainingFailures {
				t.Fatalf("%s\nRemainingGuesses() = %d, want %d", setup, got, tc.remainingFailures)
			}
		})
	}
}

func BenchmarkGuess(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			g := NewGame(tc.word)
			for _, c := range tc.guesses {
				_ = g.Guess(c)
			}
		}
	}
}
