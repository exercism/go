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
func searchFile(pattern, file string, options *options) (output []string) {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()
	var filePrefix string
	if options.prefixFilenames {
		filePrefix = file + ":"
	}
	scanner := bufio.NewScanner(f)
	for lineNum := 1; scanner.Scan(); lineNum++ {
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
			switch {
			case options.printFilenamesOnly:
				return append(output, file)
			case options.printLineNumbers:
				output = append(output, filePrefix+strconv.Itoa(lineNum)+":"+line)
			default:
				output = append(output, filePrefix+line)
			}
		} else if !match && options.invertPatternMatch {
			switch {
			case options.printFilenamesOnly:
				return append(output, file)
			case options.printLineNumbers:
				output = append(output, filePrefix+strconv.Itoa(lineNum)+":"+line)
			default:
				output = append(output, filePrefix+line)
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
