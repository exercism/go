package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"deliveryDate": &[]testCase{},
	}
	if err := gen.Gen("swift-scheduling", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		MeetingStart string `json:"meetingStart"`
		Description string `json:"description"`
	} `json:"input"`
	Expected string `json:"expected"`
}

var tmpl = `package swiftscheduling

{{.Header}}

var testCases = []struct {
	description string
	start       string
	delivery    string
	expected    string
}{
{{range .J.deliveryDate}}{
	description: 	{{printf "%q" .Description}},
	start: 		{{printf "%q" .Input.MeetingStart}},
	delivery: 	{{printf "%q" .Input.Description}},
	expected:    	{{printf "%q" .Expected}},
},
{{end}}
}`
