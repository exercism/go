package main

import (
	"log"
	"strings"
	"text/template"

	"../../../gen"
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
	Cases []struct {
		Description string
		Year        int
		Month       int
		Week        string
		Dayofweek   string
		Dayofmonth  int
	}
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
{{range .J.Cases}}{ {{.Year}}, {{.Month}}, {{week .Week}}, time.{{.Dayofweek}}, {{.Dayofmonth}}}, // {{.Description}}
{{end}}}
`
