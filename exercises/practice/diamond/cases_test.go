package diamond

// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description   string
	input         string
	expected      []string
	expectedError error
}{
	{
		description:   "Degenerate case with a single 'A' row",
		input:         "A",
		expectedError: nil,
		expected: []string{
			"A",
		},
	},
	{
		description:   "Degenerate case with no row containing 3 distinct groups of spaces",
		input:         "B",
		expectedError: nil,
		expected: []string{
			" A ",
			"B B",
			" A ",
		},
	},
	{
		description:   "Smallest non-degenerate case with odd diamond side length",
		input:         "C",
		expectedError: nil,
		expected: []string{
			"  A  ",
			" B B ",
			"C   C",
			" B B ",
			"  A  ",
		},
	},
	{
		description:   "Smallest non-degenerate case with even diamond side length",
		input:         "D",
		expectedError: nil,
		expected: []string{
			"   A   ",
			"  B B  ",
			" C   C ",
			"D     D",
			" C   C ",
			"  B B  ",
			"   A   ",
		},
	},
	{
		description:   "Largest possible diamond",
		input:         "Z",
		expectedError: nil,
		expected: []string{
			"                         A                         ",
			"                        B B                        ",
			"                       C   C                       ",
			"                      D     D                      ",
			"                     E       E                     ",
			"                    F         F                    ",
			"                   G           G                   ",
			"                  H             H                  ",
			"                 I               I                 ",
			"                J                 J                ",
			"               K                   K               ",
			"              L                     L              ",
			"             M                       M             ",
			"            N                         N            ",
			"           O                           O           ",
			"          P                             P          ",
			"         Q                               Q         ",
			"        R                                 R        ",
			"       S                                   S       ",
			"      T                                     T      ",
			"     U                                       U     ",
			"    V                                         V    ",
			"   W                                           W   ",
			"  X                                             X  ",
			" Y                                               Y ",
			"Z                                                 Z",
			" Y                                               Y ",
			"  X                                             X  ",
			"   W                                           W   ",
			"    V                                         V    ",
			"     U                                       U     ",
			"      T                                     T      ",
			"       S                                   S       ",
			"        R                                 R        ",
			"         Q                               Q         ",
			"          P                             P          ",
			"           O                           O           ",
			"            N                         N            ",
			"             M                       M             ",
			"              L                     L              ",
			"               K                   K               ",
			"                J                 J                ",
			"                 I               I                 ",
			"                  H             H                  ",
			"                   G           G                   ",
			"                    F         F                    ",
			"                     E       E                     ",
			"                      D     D                      ",
			"                       C   C                       ",
			"                        B B                        ",
			"                         A                         ",
		},
	},
}
