// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"text/template"
)

// Most code at the beginning here not specific to the exercise.
// It could be used for other problems.

func getPath() (jPath, jOri, jCommit string) {
	// Ideally draw from a .json which is pulled from the official x-common
	// repository.  For development however, accept a file in current directory
	// if there is no .json in source control.  Also allow an override in any
	// case by environment variable.
	if jPath = os.Getenv("EXTEST"); jPath > "" {
		return jPath, "local file", "" // override
	}
	c := exec.Command("git", "show", "--oneline", "HEAD")
	c.Dir = "../../../exercism/x-common"
	ori, err := c.Output()
	if err != nil {
		return "", "local file", "" // no source control
	}
	if _, err = os.Stat(filepath.Join(c.Dir, jFile)); err != nil {
		return "", "local file", "" // not in source control
	}
	// good.  return source control dir and commit.
	return c.Dir, "exercism/x-common", string(bytes.TrimSpace(ori))
}

func main() {
	jPath, jOri, jCommit := getPath()
	d, err := ioutil.ReadFile(filepath.Join(jPath, jFile))
	if err != nil {
		log.Fatal(err)
	}
	j := &js{}
	if err = json.Unmarshal(d, j); err != nil {
		// This error message is usually enough if the problem is a wrong
		// data structure defined here. Sadly it doesn't locate the error well
		// in the case of invalid JSON.  Use a real validator tool if you can't
		// spot the problem right away.
		log.Print(`unexpected data structure`)
		log.Fatal(err)
	}
	gen(j, jPath, jOri, jCommit)
}

func gen(j *js, jPath, jOri, jCommit string) {
	t := template.New("").Funcs(template.FuncMap{
		"str":  str,
		"strs": strSlice,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("cases_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	d := struct {
		Ori    string
		Commit string
		J      *js
	}{jOri, jCommit, j}
	if err = t.Execute(f, &d); err != nil {
		log.Print(err)
		return
	}
	if exec.Command("go", "fmt", "cases_test.go").Run(); err != nil {
		log.Print(err)
		return
	}
}

// Code from here on is custom-set specific.
//
// There's some extra complexity because json test cases use ints but it's
// all converted to strings here.  It just seemed like strings would make
// a better and more practical example.

func str(n int) string {
	if n >= 1 && n <= 26 {
		return string('a' - 1 + n)
	}
	return strconv.Itoa(n)
}

func strSlice(ns []int) []string {
	s := make([]string, len(ns))
	for i, n := range ns {
		s[i] = str(n)
	}
	return s
}

const jFile = "custom-set.json"

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Equal               binaryBool // binary function, boolean result
	Add                 eleOp      // set-element operator
	Delete              eleOp      // ...
	IsEmpty             unaryBool  `json:"is-empty"`
	Size                unaryInt   // Set.Len
	Element             eleBool    // Set.Has
	Subset              binaryBool
	Disjoint            binaryBool
	Union               binaryOp // set-set operator
	Intersection        binaryOp
	Difference          binaryOp
	SymmetricDifference binaryOp `json:"symmetric-difference"`
}

type binaryBool struct {
	Description []string
	Cases       []struct {
		Description string
		Set1        []int
		Set2        []int
		Expected    bool
	}
}

type eleOp struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Element     int
		Expected    []int
	}
}

type unaryBool struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Expected    bool
	}
}

type unaryInt struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Expected    int
	}
}

type eleBool struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Element     int
		Expected    bool
	}
}

type binaryOp struct {
	Description []string
	Cases       []struct {
		Description string
		Set1        []int
		Set2        []int
		Expected    []int
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `
{{/* nested templates for repeated stuff */}}

{{define "cmts"}}{{range .}}// {{.}}
{{end}}{{end}}

{{define "binaryBool"}} = []binBoolCase{
{{range .}}{ // {{.Description}}
	{{strs .Set1 | printf "%#v"}},
	{{strs .Set2 | printf "%#v"}},
	{{.Expected}},
},
{{end}}}{{end}}

{{define "binaryOp"}} = []binOpCase{
{{range .}}{ // {{.Description}}
	{{strs .Set1 | printf "%#v"}},
	{{strs .Set2 | printf "%#v"}},
	{{strs .Expected | printf "%#v"}},
},
{{end}}}{{end}}

{{define "eleOp"}} = []eleOpCase{
{{range .}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{str .Element | printf "%q"}},
	{{strs .Expected | printf "%#v"}},
},
{{end}}}{{end}}

{{/* begin template body */}}
package stringset

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

{{template "cmts" .J.Equal.Description}}var eqCases{{template "binaryBool" .J.Equal.Cases}}

{{template "cmts" .J.Add.Description}}var addCases{{template "eleOp" .J.Add.Cases}}

{{template "cmts" .J.Delete.Description}}var delCases{{template "eleOp" .J.Delete.Cases}}

{{range .J.IsEmpty.Description}}// {{.}}
{{end}}var emptyCases = []unaryBoolCase{
{{range .J.IsEmpty.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{.Expected}},
},
{{end}}}

{{range .J.Size.Description}}// {{.}}
{{end}}var lenCases = []unaryIntCase{
{{range .J.Size.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{.Expected}},
},
{{end}}}

{{range .J.Element.Description}}// {{.}}
{{end}}var hasCases = []eleBoolCase{
{{range .J.Element.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{str .Element | printf "%q"}},
	{{.Expected}},
},
{{end}}}

{{template "cmts" .J.Subset.Description}}var subsetCases{{template "binaryBool" .J.Subset.Cases}}

{{template "cmts" .J.Disjoint.Description}}var disjointCases{{template "binaryBool" .J.Disjoint.Cases}}

{{template "cmts" .J.Union.Description}}var unionCases{{template "binaryOp" .J.Union.Cases}}

{{template "cmts" .J.Intersection.Description}}var intersectionCases{{template "binaryOp" .J.Intersection.Cases}}

{{template "cmts" .J.Difference.Description}}var differenceCases{{template "binaryOp" .J.Difference.Cases}}

{{template "cmts" .J.SymmetricDifference.Description}}var symmetricDifferenceCases{{template "binaryOp" .J.SymmetricDifference.Cases}}
`
