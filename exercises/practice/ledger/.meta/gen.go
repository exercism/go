package main

import (
	"../../../../gen"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Funcs(template.FuncMap{"join": strings.Join, "replace": strings.ReplaceAll}).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"formatEntries": &[]testCase{},
	}
	if err := gen.Gen("ledger", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Currency string `json:"currency"`
		Locale   string `json:"locale"`
		Entries  []struct {
			Date        string `json:"date"`
			Description string `json:"description"`
			Change      int    `json:"amountInCents"`
		} `json:"entries"`
	} `json:"input"`
	Expected []string `json:"expected"`
	Error    bool
}

var tmpl = `{{.Header}}

import "strings"

type testCase = struct{
	description     string
	currency string
	locale   string
	entries  []Entry
	expected string
}

var testCases = []testCase{ {{range .J.formatEntries}}
	{
		description: 	{{printf "%q" .Description}},
		currency: 	{{printf "%q" .Input.Currency}},
		locale: 	{{replace .Input.Locale "_" "-" | printf "%q"}},
		entries:	[]Entry{ {{range .Input.Entries}}
			{
				Date:        {{printf "%q" .Date}},
				Description: {{printf "%q" .Description}},
				Change:      {{.Change}},
			}, {{end}}
		},
		expected:    	strings.Join([]string{ {{range .Expected}}
			{{printf "%q" .}},{{end}}
		}, "\n") + "\n",
	},{{end}}
}`
