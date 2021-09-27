package logs

import (
	"strings"
	"testing"
)

func TestLogApplication(t *testing.T) {
	tests := []struct {
		name string
		log  string
		want string
	}{
		{
			name: "single sequence recommendation",
			log:  "❗❗ recommended product",
			want: "recommendation",
		},
		{
			name: "single sequence search",
			log:  "executed search 🔍🔎",
			want: "search",
		},
		{
			name: "single sequence weather",
			log:  "forecast: ☁☀ rain",
			want: "weather",
		},
		{
			name: "no sequence default",
			log:  "error: could not proceed",
			want: "default",
		},
		{
			name: "multiple sequences recommendation",
			log:  "❗❗ recommended search product 🔍🔎",
			want: "recommendation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogApplication(tt.log); got != tt.want {
				t.Errorf("LogApplication(\"%s\") = \"%s\", want \"%s\"", escapeWhiteSpace(tt.log), got, tt.want)
			}
		})
	}
}

func TestLogCLean(t *testing.T) {
	tests := []struct {
		name    string
		log     string
		invalid []rune
		want    string
	}{
		{
			name:    "single invalid character",
			log:     "❗❗ recommended product",
			invalid: []rune{'❗'},
			want:    " recommended product",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogClean(tt.log, tt.invalid); got != tt.want {
				t.Errorf("LogClean(\"%s\", \"%v\") = \"%s\", want \"%s\"", escapeWhiteSpace(tt.log), tt.invalid, got, tt.want)
			}
		})
	}
}

func escapeWhiteSpace(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	return strings.ReplaceAll(s, "\t", "\\t")
}
