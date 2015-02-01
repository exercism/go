package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func Gen(jFile string, j interface{}, t *template.Template) error {
	// find and read the json source file
	jPath, jOri, jCommit := getPath(jFile)
	jSrc, err := ioutil.ReadFile(filepath.Join(jPath, jFile))
	if err != nil {
		return err
	}

	// unmarshal the json source to a Go structure
	if err = json.Unmarshal(jSrc, j); err != nil {
		// This error message is usually enough if the problem is a wrong
		// data structure defined here. Sadly it doesn't locate the error well
		// in the case of invalid JSON.  Use a real validator tool if you can't
		// spot the problem right away.
		return fmt.Errorf(`unexpected data structure: %v`, err)
	}
	// package up a little meta data
	d := struct {
		Ori    string
		Commit string
		J      interface{}
	}{jOri, jCommit, j}

	// create an output file for the Go test cases.
	f, err := os.Create("cases_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// render the Go test cases into the output file
	if err = t.Execute(f, &d); err != nil {
		return err
	}
	// clean it up
	if exec.Command("go", "fmt", "cases_test.go").Run(); err != nil {
		return err
	}
	return nil
}

func getPath(jFile string) (jPath, jOri, jCommit string) {
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
