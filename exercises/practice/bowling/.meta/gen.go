package main

import (
	"exercism-go/gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
		"roll":  &TestRoll{},
		"score": &TestScore{},
	}
	if err := gen.Gen("bowling", &j, t); err != nil {
		log.Fatal(err)
	}
}

type TestScore []struct {
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Previousrolls []int `json:"previousRolls"`
	} `json:"input"`
	Expected int `json:"expected"`
}

type TestRoll []struct {
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Previousrolls []interface{} `json:"previousRolls"`
		Roll          int           `json:"roll"`
	} `json:"input"`
	Expected struct {
		Error string `json:"error"`
	} `json:"expected"`
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
