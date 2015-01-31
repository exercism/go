// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

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
		return "", "local file", "" // not in source control
	}
	// good.  return source control dir and commit.
	return c.Dir, "exercism/x-common", string(bytes.TrimSpace(ori))
}

func main() {
	jPath, jOri, jCommit := getPath()
	d, err := ioutil.ReadFile(filepath.Join(jPath, "clock.json"))
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

// The JSON structure we expect to be able to umarshal into
type js struct {
	Create struct {
		Description []string
		Cases       []timeCase
	}
	Add struct {
		Description []string
		Cases       []addCase
	}
	Equal struct {
		Description []string
		Cases       []eqCase
	}
}

// Handle the three tests similarly

type timeCase struct {
	Description  string
	Hour, Minute int
	Expected     string
}

func genTimeTests(f *os.File, j *js) {
	// json.Unmarshal will happily skip parts of the data structure that
	// don't match.  Check here that we really got some data.
	if len(j.Create.Cases) == 0 {
		log.Fatal(`Missing "Create" test cases`)
	}
	genCmts(f, j.Create.Description)
	fmt.Fprintln(f, `var timeTests = []struct {
	h, m int
	want string
}{`)
	for i := range j.Create.Cases {
		tc := &j.Create.Cases[i]
		fmt.Fprintf(f, "{%d, %d, %q}, // %s\n",
			tc.Hour, tc.Minute, tc.Expected, tc.Description)
	}
	fmt.Fprint(f, "}\n\n")
}

type addCase struct {
	Description       string
	Hour, Minute, Add int
	Expected          string
}

func genAddTests(f *os.File, j *js) {
	// generate addTests
	if len(j.Add.Cases) == 0 {
		log.Fatal(`Missing "Add" test cases`)
	}
	genCmts(f, j.Add.Description)
	fmt.Fprintln(f, `var addTests = []struct {
	h, m, a int
	want    string
}{`)
	for i := range j.Add.Cases {
		tc := &j.Add.Cases[i]
		fmt.Fprintf(f, "{%d, %d, %d, %q}, // %s\n",
			tc.Hour, tc.Minute, tc.Add, tc.Expected, tc.Description)
	}
	fmt.Fprint(f, "}\n\n")
}

type eqCase struct {
	Description    string
	Clock1, Clock2 struct{ Hour, Minute int }
	Expected       bool
}

func genEqTests(f *os.File, j *js) {
	if len(j.Equal.Cases) == 0 {
		log.Fatal(`Missing "Equal" test cases`)
	}
	genCmts(f, j.Equal.Description)
	fmt.Fprintln(f, `type hm struct{ h, m int }

var eqTests = []struct {
	c1, c2 hm
	want   bool
}{`)
	for i := range j.Equal.Cases {
		tc := &j.Equal.Cases[i]
		fmt.Fprintf(f, `// %s
{
	hm{%d, %d},
	hm{%d, %d},
	%t,
},
`, tc.Description,
			tc.Clock1.Hour, tc.Clock1.Minute,
			tc.Clock2.Hour, tc.Clock2.Minute,
			tc.Expected)
	}
	fmt.Fprintln(f, "}")
}

func gen(j *js, jPath, jOri, jCommit string) {
	f, err := os.Create("cases_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package clock\n")

	fmt.Fprintln(f, "// Source:", jOri)
	if jCommit > "" {
		fmt.Fprintf(f, "// Commit: %s\n\n", jCommit)
	}

	genTimeTests(f, j)
	genAddTests(f, j)
	genEqTests(f, j)

	exec.Command("go", "fmt", "cases_test.go").Run()
}

func genCmts(f *os.File, cmts []string) {
	for _, c := range cmts {
		fmt.Fprintln(f, "//", c)
	}
}
