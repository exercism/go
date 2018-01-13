package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("bowling", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Comments []string
	Cases    []OneCase
}

// template applied to above data structure generates the Go test cases

type OneCase struct {
	Description string
	Property    string
	Input       struct {
		PreviousRolls []int
		Roll          int
	}
	Expected interface{}
}

// ScoreTest and RollTest help determine which type of test case
// to generate in the template.
func (c OneCase) ScoreTest() bool { return c.Property == "score" }
func (c OneCase) RollTest() bool  { return c.Property == "roll" }

func (c OneCase) Valid() bool {
	valid, _, _ := determineExpected(c.Expected)
	return valid
}

func (c OneCase) Score() int {
	_, score, _ := determineExpected(c.Expected)
	return score
}

func (c OneCase) ExplainText() string {
	_, _, explainText := determineExpected(c.Expected)
	return explainText
}

// determineExpected examines an .Expected interface{} object and determines
// whether a test case is valid(bool), has a score field, and/or has an expected error,
// returning valid, score, and error explanation text.
func determineExpected(expected interface{}) (bool, int, string) {
	score, ok := expected.(float64)
	if ok {
		return ok, int(score), ""
	}
	m, ok := expected.(map[string]interface{})
	if !ok {
		return false, 0, ""
	}
	iError, ok := m["error"].(interface{})
	if !ok {
		return false, 0, ""
	}
	explainText, ok := iError.(string)
	return false, 0, explainText
}

// Template to generate two sets of test cases, one for Score tests and one for Roll tests.
var tmpl = `package bowling

{{.Header}}

var scoreTestCases = []struct {
	description    string
	previousRolls  []int	// bowling rolls to do before the Score() test
	valid          bool     // true => no error, false => error expected
	score          int	// when .valid == true, the expected score value
	explainText    string   // when .valid == false, error explanation text
}{ {{range .J.Cases}}
{{if .ScoreTest}}{
	{{printf "%q"  .Description}},
	{{printf "%#v" .Input.PreviousRolls}},
	{{printf "%v"  .Valid}},
	{{printf "%d"  .Score}},
	{{printf "%q"  .ExplainText}},
},{{- end}}{{end}}
}

var rollTestCases = []struct {
	description    string
	previousRolls  []int	// bowling rolls to do before the Roll(roll) test
	valid          bool     // true => no error, false => error expected
	roll           int	// pin count for the test roll
	explainText    string   // when .valid == false, error explanation text
}{ {{range .J.Cases}}
{{if .RollTest}}{
	{{printf "%q"  .Description}},
	{{printf "%#v" .Input.PreviousRolls}},
	{{printf "%v"  .Valid}},
	{{printf "%d"  .Input.Roll}},
	{{printf "%q"  .ExplainText}},
},{{- end}}{{end}}
}
`
