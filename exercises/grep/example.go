package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type grepOptions struct {
	n_line_numbers      bool
	l_file_only         bool
	i_any_case          bool
	v_invert            bool
	x_match_entire_line bool
	prefix_filenames    bool
}

// GrepFiles searches for pattern in the files given, using optional flags.
func GrepFiles(pattern string, flags []string, files []string) []string {
	output := []string{}
	options := getOptions(flags, len(files))
	for _, file := range files {
		output = append(output, grepOneFile(pattern, file, options)...)
	}
	return output
}

// grepOneFile performs a search for pattern in the given file
// and returns the proper output according to the options.
func grepOneFile(pattern string, file string, options *grepOptions) (output []string) {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()
	var file_prefix string
	if options.prefix_filenames {
		file_prefix = file + ":"
	}
	scanner := bufio.NewScanner(f)
	for line_num := 1; scanner.Scan(); line_num++ {
		var match bool
		line := scanner.Text()
		if options.x_match_entire_line {
			if options.i_any_case {
				match = strings.EqualFold(line, pattern)
			} else {
				match = line == pattern
			}
		} else {
			if options.i_any_case {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				match = strings.Contains(line, pattern)
			}
		}
		if match && !options.v_invert {
			if options.l_file_only {
				output = append(output, file)
				break
			} else if options.n_line_numbers {
				output = append(output, file_prefix+strconv.Itoa(line_num)+":"+line)
			} else {
				output = append(output, file_prefix+line)
			}
		} else if !match && options.v_invert {
			if options.l_file_only {
				output = append(output, file)
				break
			} else if options.n_line_numbers {
				output = append(output, file_prefix+strconv.Itoa(line_num)+":"+line)
			} else {
				output = append(output, file_prefix+line)
			}
		}
	}
	return output
}

// getOptions examines flags and nfiles, outputing options.
func getOptions(flags []string, nfiles int) (options *grepOptions) {
	options = &grepOptions{}
	for _, option := range flags {
		switch option {
		case "-n":
			options.n_line_numbers = true
		case "-l":
			options.l_file_only = true
		case "-i":
			options.i_any_case = true
		case "-v":
			options.v_invert = true
		case "-x":
			options.x_match_entire_line = true
		}
	}
	if nfiles > 1 {
		options.prefix_filenames = true
	}
	return options
}
