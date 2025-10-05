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
		if t == "" {
			if filename == "" {
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
	return filenames
}

func deleteFiles(filenames []string) {
	for _, file := range filenames {
		os.Remove(file)
	}
}

func TestSearch(t *testing.T) {
	files := createFiles(fileContentData)
	defer deleteFiles(files)

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Search(tc.pattern, tc.flags, tc.files)

			// We do not care whether the result is nil or an empty slice.
			if len(tc.expected) == 0 && len(actual) == 0 {
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Search(%q,%v,%v)\ngot: %v\nwant: %v", tc.pattern, tc.flags, tc.files, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	files := createFiles(fileContentData)
	defer deleteFiles(files)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Search(tc.pattern, tc.flags, tc.files)
		}
	}
}
