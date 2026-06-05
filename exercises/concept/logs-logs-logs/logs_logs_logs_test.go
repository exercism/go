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
			name: "multiple characters recommendation(1/3)",
			log:  "❗ recommended search product 🔍",
			want: "recommendation",
		},
		{
			name: "multiple characters recommendation(2/3)",
			log:  "🔍 search recommended product ❗",
			want: "search",
		},
		{
			name: "multiple characters recommendation(3/3)",
			log:  "☀ weather is sunny ❗",
			want: "weather",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Application(tc.log); got != tc.want {
				t.Errorf("Application(\"%s\") = \"%s\", want \"%s\"", tc.log, got, tc.want)
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
			name:    "single occurrence of replacement",
			log:     "❗ recommended product",
			oldChar: '❗',
			newChar: '?',
			want:    "? recommended product",
		},
		{
			name:    "multiple occurrences of replacement",
			log:     "❗ recommended product ❗",
			oldChar: '❗',
			newChar: '?',
			want:    "? recommended product ?",
		},
		{
			name:    "no occurrences of replacement",
			log:     "❗ recommended product ❗",
			oldChar: '?',
			newChar: '?',
			want:    "❗ recommended product ❗",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Replace(tc.log, tc.oldChar, tc.newChar); got != tc.want {
				t.Errorf("Replace(\"%s\", '%c', '%c') = \"%s\", want \"%s\"", tc.log, tc.oldChar, tc.newChar, got, tc.want)
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
			name:  "exact limit 1",
			log:   "exercism❗",
			limit: 9,
			want:  true,
		},
		{
			name:  "under limit 1",
			log:   "exercism❗",
			limit: 10,
			want:  true,
		},
		{
			name:  "over limit 1",
			log:   "exercism❗",
			limit: 8,
			want:  false,
		},
		{
			name:  "exact limit 2",
			log:   "exercism🔍",
			limit: 9,
			want:  true,
		},
		{
			name:  "under limit 2",
			log:   "exercism🔍",
			limit: 10,
			want:  true,
		},
		{
			name:  "over limit 2",
			log:   "exercism🔍",
			limit: 8,
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := WithinLimit(tc.log, tc.limit); got != tc.want {
				t.Errorf("WithinLimit(\"%s\", %d) = %t, want %t", tc.log, tc.limit, got, tc.want)
			}
		})
	}
}
