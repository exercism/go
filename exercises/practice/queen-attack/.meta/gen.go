package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"text/template"
)

const Columns = "abcdefgh"

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"canAttack": &[]testCase{},
		"create":    &[]interface{}{},
	}
	if err := gen.Gen("queen-attack", j, t); err != nil {
		log.Fatal(err)
	}
}

type Position struct {
	Position struct {
		Row    int `json:"row"`
		Column int `json:"column"`
	} `json:"position"`
}

func (p Position) String() string {
	return fmt.Sprintf(`"%c%d"`, Columns[p.Position.Column], p.Position.Row+1)
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		WhiteQueen Position `json:"white_queen"`
		BlackQueen Position `json:"black_queen"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

func (tc testCase) Pos(index int) string {
	if index == 1 {
		return tc.Input.WhiteQueen.String()
	}
	return tc.Input.BlackQueen.String()
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	pos1, pos2  string
	expected    bool
}

var testCases = []testCase{ {{range .J.canAttack}}
	{
		description: 	{{printf "%q" .Description}},
		pos1:           {{.Pos 1}},
		pos2:           {{.Pos 2}},
		expected:    	{{.Expected}},
	},{{end}}
}`
