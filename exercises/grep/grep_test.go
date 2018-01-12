package grep

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func createFiles(content []string) (filenames []string) {
	// Parse fileContentData, making the list of filenames
	// with their content.
	var filename string
	var f *os.File
	for _, d := range content {
		t := strings.TrimSpace(d)
		if len(t) == 0 {
			if len(filename) == 0 {
				continue
			}
			// Close file
			f.Close()
			filenames = append(filenames, filename)
			filename = ""
			f = nil
			continue
		}
		if strings.Contains(t, ".txt") {
			filename = t
			// Open file
			var err error
			f, err = os.Create(filename)
			if err != nil {
				panic(err)
			}
			continue
		}
		fields := strings.Split(t, "|")
		if len(fields) == 3 {
			// Write string into file with newline.
			_, err := f.WriteString(strings.TrimRight(fields[1], " ") + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
	if f != nil {
		f.Close()
		filenames = append(filenames, filename)
	}
	return
}

func deleteFiles(filenames []string) {
	for _, file := range filenames {
		os.Remove(file)
	}
}

func TestGrepFiles(t *testing.T) {
	files := createFiles(fileContentData)
	defer deleteFiles(files)

	for _, tc := range testCases {
		actual := GrepFiles(tc.pattern, tc.flags, tc.files)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("FAIL: %s\nGrepFiles for pattern %q\nexpected %v\nactual %v.",
				tc.description, tc.pattern, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkResultOf(b *testing.B) {
	files := createFiles(fileContentData)
	defer deleteFiles(files)

	b.StopTimer()
	for i := 0; i < b.N; i++ {

		b.StartTimer()

		for _, tc := range testCases {
			GrepFiles(tc.pattern, tc.flags, tc.files)
		}

		b.StopTimer()
	}

}
