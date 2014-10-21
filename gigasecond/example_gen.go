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
	"text/template"
)

// Code at the beginning here is not specific to gigasecond.
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
	t, err := template.New("").Parse(tmpl)
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

// Code from here on is gigasecond-specific definitions.

const jFile = "gigasecond.json"

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Add struct {
		Description []string
		Cases       []struct {
			Input    string
			Expected string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package gigasecond

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}
{{range .J.Add.Description}}// {{.}}
{{end}}var addCases = []struct {
	in   string
	want string
}{
{{range .J.Add.Cases}}{
	{{printf "%q" .Input}},
	{{printf "%q" .Expected}},
},
{{end}}}
`
