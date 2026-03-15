package main

import (
	"../../../../gen"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Funcs(template.FuncMap{"lower": strings.ToLower}).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"territory":   &[]territoryTestCase{},
		"territories": &[]territoriesTestCase{},
	}
	if err := gen.Gen("go-counting", j, t); err != nil {
		log.Fatal(err)
	}
}

type territoryTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Board []string `json:"board"`
		PosX  int      `json:"x"`
		PosY  int      `json:"y"`
	} `json:"input"`
	Expected struct {
		Owner     string   `json:"owner"`
		Territory [][2]int `json:"territory"`
		Error     string   `json:"error"`
	} `json:"expected"`
}

type territoriesTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Board []string `json:"board"`
	} `json:"input"`
	Expected struct {
		TerritoryBlack [][2]int `json:"territoryBlack"`
		TerritoryWhite [][2]int `json:"territoryWhite"`
		TerritoryNone  [][2]int `json:"territoryNone"`
	} `json:"expected"`
}

var tmpl = `{{.Header}}

var oneTerritoryTestCases = []struct {
	description string
	board       []string
	posX        int
	posY        int
	owner       string
	territory   [][2]int
	expectedErr string
}{ {{range .J.territory}}
	{
		description: {{printf "%q" .Description}},
		board:       {{printf "%#v" .Input.Board}},
		posX:        {{.Input.PosX}},
		posY:        {{.Input.PosY}},
		owner:       {{printf "%q" .Expected.Owner}},
		territory:   {{printf "%#v" .Expected.Territory}},
		expectedErr: {{printf "%q" .Expected.Error | lower}},
	},{{end}}
}

var allTerritoriesTestCases = []struct {
	description string
	input       []string
	expected    AllTerritories
}{ {{range .J.territories}}
	{
		description: {{printf "%q" .Description}},
		input: 	     {{printf "%#v" .Input.Board}},
		expected:    AllTerritories{
			Black: {{printf "%#v" .Expected.TerritoryBlack}},
			White: {{printf "%#v" .Expected.TerritoryWhite}},
			None: {{printf "%#v" .Expected.TerritoryNone}},
		},
	},{{end}}
}`
