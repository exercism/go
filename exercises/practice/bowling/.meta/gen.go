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
	var j = map[string]interface{}{
		"roll":  &[]Case{},
		"score": &[]Case{},
	}
	if err := gen.Gen("bowling", j, t); err != nil {
		log.Fatal(err)
	}
}

type Case struct {
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		PreviousRolls []int `json:"previousRolls"`
		Roll          int   `json:"roll"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t Case) Score() int {
	score, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return int(score)
}

func (t Case) Valid() bool {
	_, ok := t.Expected.(float64)
	return ok
}

func (t Case) ExplainText() string {
	if !t.Valid() {
		m, ok := t.Expected.(map[string]interface{})
		if !ok {
			return ""
		}
		b, ok := m["error"].(string)
		if !ok {
			return ""
		}
		return b
	}
	return ""
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
}{ {{range .J.score}}
{
	{{printf "%q"  .Description}},
	{{printf "%#v" .Input.PreviousRolls}},
	{{printf "%v"  .Valid}},
	{{printf "%d"  .Score}},
	{{printf "%q"  .ExplainText}},
},{{end}}
}


var rollTestCases = []struct {
	description    string
	previousRolls  []int	// bowling rolls to do before the Roll(roll) test
	valid          bool     // true => no error, false => error expected
	roll           int	// pin count for the test roll
	explainText    string   // when .valid == false, error explanation text
}{ {{range .J.roll}}
{
	{{printf "%q"  .Description}},
	{{printf "%#v" .Input.PreviousRolls}},
	{{printf "%v"  .Valid}},
	{{printf "%d"  .Input.Roll}},
	{{printf "%q"  .ExplainText}},
},{{end}}
}
`
