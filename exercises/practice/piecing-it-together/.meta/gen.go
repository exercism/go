package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"jigsawData": &[]testCase{},
	}
	if err := gen.Gen("piecing-it-together", j, t); err != nil {
		log.Fatal(err)
	}
}

type PuzzleDetails struct {
	Pieces      int     `json:"pieces"`
	Border      int     `json:"border"`
	Inside      int     `json:"inside"`
	Rows        int     `json:"rows"`
	Columns     int     `json:"columns"`
	AspectRatio float64 `json:"aspectRatio"`
	Format      string  `json:"format"`
	Error       string  `json:"error"`
}

type testCase struct {
	Description string        `json:"description"`
	Input       PuzzleDetails `json:"input"`
	Expected    PuzzleDetails `json:"expected"`
}

func (tc testCase) ToPuzzleDetails(p PuzzleDetails) string {
	var sb strings.Builder
	sb.WriteString("\t\tPuzzleDetails {\n")
	if p.Pieces != 0 {
		sb.WriteString(fmt.Sprintf("\t\t\tPieces: %d,\n", p.Pieces))
	}
	if p.Border != 0 {
		sb.WriteString(fmt.Sprintf("\t\t\tBorder: %d,\n", p.Border))
	}
	if p.Inside != 0 {
		sb.WriteString(fmt.Sprintf("\t\t\tInside: %d,\n", p.Inside))
	}
	if p.Rows != 0 {
		sb.WriteString(fmt.Sprintf("\t\t\tRows: %d,\n", p.Rows))
	}
	if p.Columns != 0 {
		sb.WriteString(fmt.Sprintf("\t\t\tColumns: %d,\n", p.Columns))
	}
	if p.AspectRatio != 0 {
		sb.WriteString(fmt.Sprintf("\t\t\tAspectRatio: %v,\n", p.AspectRatio))
	}
	if p.Format != "" {
		sb.WriteString(fmt.Sprintf("\t\t\tFormat: %q,\n", p.Format))
	}
	sb.WriteString("\t\t}")
	return sb.String()
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	input       PuzzleDetails
	expected    PuzzleDetails
	err         string
	wantError   bool
}{ {{range .J.jigsawData}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{.ToPuzzleDetails .Input}},
		expected:    	{{.ToPuzzleDetails .Expected}},
		err:            {{printf "%#v" .Expected.Error}},
		wantError:    	{{ne .Expected.Error ""}},
	},{{end}}
}`
