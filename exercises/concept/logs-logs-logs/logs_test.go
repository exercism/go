package logs

import (
	"testing"
)

func TestApplication(t *testing.T) {
	tests := []struct {
		name string
		log  string
		want string
	}{
		{
			name: "single character recommendation",
			log:  "❗ recommended product",
			want: "recommendation",
		},
		{
			name: "single character search",
			log:  "executed search 🔍",
			want: "search",
		},
		{
			name: "single character weather",
			log:  "forecast: ☀ sunny",
			want: "weather",
		},
		{
			name: "no characters default",
			log:  "error: could not proceed",
			want: "default",
		},
		{
			name: "multiple characters recommendation",
			log:  "❗ recommended search product 🔍",
			want: "recommendation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Application(tt.log); got != tt.want {
				t.Errorf("Application(\"%s\") = \"%s\", want \"%s\"", tt.log, got, tt.want)
			}
		})
	}
}

func TestRedact(t *testing.T) {
	tests := []struct {
		name       string
		log        string
		redactions []rune
		want       string
	}{
		{
			name:       "single occurance single invalid character",
			log:        "❗ recommended product",
			redactions: []rune{'❗'},
			want:       " recommended product",
		},
		{
			name:       "multiple occurances single redactions character",
			log:        "❗❗ recommended product",
			redactions: []rune{'❗'},
			want:       " recommended product",
		},
		{
			name:       "single occurance multiple redactions characters",
			log:        "❗ recommended product",
			redactions: []rune{'❗', 'e'},
			want:       " rcommndd product",
		},
		{
			name:       "multiple occurances multiple redactions characters",
			log:        "❗❗ recommended product",
			redactions: []rune{'❗', 'e'},
			want:       " rcommndd product",
		},
		{
			name:       "no occurances no redactions characters",
			log:        "❗❗ recommended product",
			redactions: []rune{},
			want:       "❗❗ recommended product",
		},
		{
			name:       "no occurances single redactions characters",
			log:        "recommended product",
			redactions: []rune{'❗'},
			want:       "recommended product",
		},
		{
			name:       "no occurances multiple redactions characters",
			log:        "recommended product",
			redactions: []rune{'❗', '!'},
			want:       "recommended product",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Redact(tt.log, tt.redactions); got != tt.want {
				t.Errorf("Redact(\"%s\", %c) = \"%s\", want \"%s\"", tt.log, tt.redactions, got, tt.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		name    string
		log     string
		oldChar rune
		newChar rune
		want    string
	}{
		{
			name:    "single occurance of replacement",
			log:     "❗ recommended product",
			oldChar: '❗',
			newChar: '?',
			want:    "? recommended product",
		},
		{
			name:    "multiple occurances of replacement",
			log:     "❗ recommended product ❗",
			oldChar: '❗',
			newChar: '?',
			want:    "? recommended product ?",
		},
		{
			name:    "no occurances of replacement",
			log:     "❗ recommended product ❗",
			oldChar: '?',
			newChar: '?',
			want:    "❗ recommended product ❗",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.log, tt.oldChar, tt.newChar); got != tt.want {
				t.Errorf("Replace(\"%s\", '%c', '%c') = \"%s\", want \"%s\"", tt.log, tt.oldChar, tt.newChar, got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name string
		log  string
		want int
	}{
		{
			name: "single byte characters",
			log:  "exercism",
			want: 8,
		},
		{
			name: "multiple byte characters",
			log:  "🧠exercism❗",
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.log); got != tt.want {
				t.Errorf("Count(\"%s\") = %d, want %d", tt.log, got, tt.want)
			}
		})
	}
}
