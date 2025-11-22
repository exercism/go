package main

import (
	"../../../../gen"
	"log"
	"strings"
	"text/template"
	"time"
	"unicode"
)

func main() {
	m := template.FuncMap{"week": Title}
	t, err := template.New("").Funcs(m).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"meetup": &[]testCase{},
	}
	if err := gen.Gen("meetup", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Year      int    `json:"year"`
		Month     int    `json:"month"`
		Week      string `json:"week"`
		DayOfWeek string `json:"dayofweek"`
	} `json:"input"`
	Expected string `json:"expected"`
}

func (c testCase) ExpectedDay() int {
	dateformat := "2006-01-02"
	parse, err := time.Parse(dateformat, c.Expected)
	if err != nil {
		log.Fatalf("[ERROR]: expected `expected` to be date of format %q, got: %q", dateformat, c.Expected)
	}
	return parse.Day()
}

// template applied to above data structure generates the Go test cases
var tmpl = `package meetup

{{.Header}}

import "time"

var testCases = []struct {
	description string
	year        int
	month       time.Month
	week        WeekSchedule
	weekday     time.Weekday
	expectedDay int
}{
{{range .J.meetup}}{
		description: 	{{printf "%q" .Description}},
		year:        	{{printf "%d" .Input.Year}},
		month:       	{{printf "%d" .Input.Month}},
		week:        	{{week .Input.Week}},
		weekday:     	time.{{.Input.DayOfWeek}},
		expectedDay:    {{printf "%d" .ExpectedDay}},
	},
{{end}}
}
`

// Title is a copy of strings.Title function of the stdlib.
// The copy is here because strings.Title is deprecated but we still
// want to use this function as the alternative would require us to support
// external dependencies which we don't yet (tracking issue https://github.com/exercism/go/issues/2379).
// Students should still be able to use strings.Title if they want.
// Since this exercise is currently deprecated, this shouldn't matter too much.
func Title(s string) string {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

// Copy of strings.isSeparator function of the stdlib.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
