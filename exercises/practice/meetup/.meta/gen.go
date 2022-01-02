package main

import (
	"fmt"
	"log"
	"strings"
	"text/template"

	"../../../../gen"
)

func main() {
	m := template.FuncMap{"week": strings.Title}
	t, err := template.New("").Funcs(m).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("meetup", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []OneCase
}

type OneCase struct {
	Description string
	Input       struct {
		Year      int
		Month     int
		Week      string
		Dayofweek string
	}
	DateString string `json:"expected"`
}

func (c OneCase) ExpectedDay() int {
	var y, m, d int
	fmt.Sscanf(c.DateString, "%4d-%2d-%2d", &y, &m, &d)
	return d
}

// template applied to above data structure generates the Go test cases
var tmpl = `package meetup

{{.Header}}

import "time"

var testCases = []struct {
	year    int
	month   time.Month
	week    WeekSchedule
	weekday time.Weekday
	expDay  int
}{
{{range .J.Cases}}{ {{.Input.Year}}, {{.Input.Month}}, {{week .Input.Week}}, time.{{.Input.Dayofweek}}, {{.ExpectedDay}}}, // {{.Description}}
{{end}}}
`
