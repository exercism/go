package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type options struct {
	printLineNumbers     bool
	printFilenamesOnly   bool
	caseInsensitiveMatch bool
	invertPatternMatch   bool
	matchEntireLine      bool
	prefixFilenames      bool
}

// Search looks for pattern in the files given, using optional flags.
func Search(pattern string, flags, files []string) []string {
	output := []string{}
	options := getOptions(flags, len(files))
	for _, file := range files {
		output = append(output, searchFile(pattern, file, options)...)
	}
	return output
}

// searchFile performs a search for pattern in the given file
// and returns the output according to the options.
func searchFile(pattern string, file string, options *options) (output []string) {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()
	var file_prefix string
	if options.prefixFilenames {
		file_prefix = file + ":"
	}
	scanner := bufio.NewScanner(f)
	for line_num := 1; scanner.Scan(); line_num++ {
		var match bool
		line := scanner.Text()
		if options.matchEntireLine {
			if options.caseInsensitiveMatch {
				match = strings.EqualFold(line, pattern)
			} else {
				match = line == pattern
			}
		} else {
			if options.caseInsensitiveMatch {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				match = strings.Contains(line, pattern)
			}
		}
		if match && !options.invertPatternMatch {
			if options.printFilenamesOnly {
				output = append(output, file)
				break
			} else if options.printLineNumbers {
				output = append(output, file_prefix+strconv.Itoa(line_num)+":"+line)
			} else {
				output = append(output, file_prefix+line)
			}
		} else if !match && options.invertPatternMatch {
			if options.printFilenamesOnly {
				output = append(output, file)
				break
			} else if options.printLineNumbers {
				output = append(output, file_prefix+strconv.Itoa(line_num)+":"+line)
			} else {
				output = append(output, file_prefix+line)
			}
		}
	}
	return output
}

// getOptions examines flags and nfiles, outputing options.
func getOptions(flags []string, nfiles int) (opt *options) {
	opt = &options{}
	for _, option := range flags {
		switch option {
		case "-n":
			opt.printLineNumbers = true
		case "-l":
			opt.printFilenamesOnly = true
		case "-i":
			opt.caseInsensitiveMatch = true
		case "-v":
			opt.invertPatternMatch = true
		case "-x":
			opt.matchEntireLine = true
		}
	}
	if nfiles > 1 {
		opt.prefixFilenames = true
	}
	return opt
}
