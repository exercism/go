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

func TestWithinLimit(t *testing.T) {
	tests := []struct {
		name  string
		log   string
		limit int
		want  bool
	}{
		{
			name:  "exact limit",
			log:   "exercism❗",
			limit: 9,
			want:  true,
		},
		{
			name:  "under limit",
			log:   "exercism❗",
			limit: 10,
			want:  true,
		},
		{
			name:  "over limit",
			log:   "exercism❗",
			limit: 8,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithinLimit(tt.log, tt.limit); got != tt.want {
				t.Errorf("WithinLimit(\"%s\", %d) = %t, want %t", tt.log, tt.limit, got, tt.want)
			}
		})
	}
}
