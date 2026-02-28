package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"reflect"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Funcs(template.FuncMap{"title": Title}).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"time": &[]testCase{},
	}
	if err := gen.Gen("split-second-stopwatch", j, t); err != nil {
		log.Fatal(err)
	}
}

type operation struct {
	Command string `json:"command"`
	By string `json:"by"`
	Expected any `json:"expected"`
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Commands []operation `json:"commands"`
	} `json:"input"`
}

func (o operation) ExpectedString() string {
	if s, ok := (o.Expected).(string); ok {
		return s
	}
	return ""
}

func (o operation) ExpectedPrevLaps() string {
	if reflect.ValueOf(o.Expected).Kind() == reflect.Slice {
		m := reflect.ValueOf(o.Expected).Convert(reflect.TypeOf([]any{})).Interface().([]any)
		var vals []string
		for _, val := range m {
			vals = append(vals, fmt.Sprintf("%q", val.(string)))
		}
		return fmt.Sprintf("[]string{%s}", strings.Join(vals, ", "))

	}
	return ""
}

func (o operation) ExpectedErr() string {
	if reflect.ValueOf(o.Expected).Kind() == reflect.Map {
		m := reflect.ValueOf(o.Expected).Convert(reflect.TypeOf(map[string]any{}))
		return m.MapIndex(reflect.ValueOf("error")).Interface().(string)

	}
	return ""
}

func Title(command string) string {
	return fmt.Sprintf("%q", strings.ToUpper(command[0:1]) + command[1:])
}

var tmpl = `package splitsecondstopwatch

{{.Header}}

type operation struct{
	command string
	by string
	expected string
	expectedPrevLaps []string
	expectedErr string
}

var testCases = []struct {
	description string
	commands []operation
}{
{{range .J.time}}{
	description: 	{{printf "%q" .Description}},
	commands:       []operation {
		{{range .Input.Commands}}
		{
			command: {{title .Command}},{{if .By}}
			by: {{printf "%q" .By}}, {{end}} {{if .ExpectedString}}
			expected: {{printf "%q" .ExpectedString}}, {{end}} {{if .ExpectedPrevLaps}}
			expectedPrevLaps: {{.ExpectedPrevLaps}}, {{end}} {{if .ExpectedErr}}
			expectedErr: {{printf "%q" .ExpectedErr}}, {{end}}
		},{{end}}
	},
},
{{end}}
}`
