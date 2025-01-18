package main

import (
	"log"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"answer": &[]testCase{},
	}
	if err := gen.Gen("wordy", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Question string `json:"question"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectError() bool {
	v, ok := t.Expected.(map[string]interface{})
	if ok {
		_, ok := v["error"].(string)
		return ok
	}
	return false
}

func (t testCase) ExpectedValue() int {
	v, ok := t.Expected.(float64)
	if ok {
		return int(v)
	}
	return 0
}

// template applied to above data structure generates the Go test cases
var tmpl = `package wordy

{{.Header}}

type wordyTest struct {
	description string
	question    string
	expectError bool
	expected    int
}

var tests = []wordyTest {
{{range .J.answer}}{
	description: {{printf "%q"  .Description}},
	question: {{printf "%q" .Input.Question}},
	expectError: {{printf "%t" .ExpectError }},
	expected: {{printf "%d"  .ExpectedValue}},
},
{{end}}}
`
