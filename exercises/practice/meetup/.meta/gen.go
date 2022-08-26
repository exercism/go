package main

import (
	"log"
	"strings"
	"text/template"
	"time"

	"../../../../gen"
)

func main() {
	m := template.FuncMap{"week": strings.Title}
	t, err := template.New("").Funcs(m).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
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
