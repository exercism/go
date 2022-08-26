package wordsearch

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 76c0ba7 word-search: Add cases checking for concatenation and wrapping

var testCases = []struct {
	description string
	puzzle      []string
	words       []string
	expectError bool
	expected    map[string][2][2]int
}{
	{
		description: "Should accept an initial game grid and a target search word",
		puzzle:      []string{"jefblpepre"},
		words:       []string{"clojure"},
		expectError: true,
		expected:    map[string][2][2]int{"clojure": {{-1, -1}, {-1, -1}}},
	},
	{
		description: "Should locate one word written left to right",
		puzzle:      []string{"clojurermt"},
		words:       []string{"clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 0}, {6, 0}}},
	},
	{
		description: "Should locate the same word written left to right in a different position",
		puzzle:      []string{"mtclojurer"},
		words:       []string{"clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{2, 0}, {8, 0}}},
	},
	{
		description: "Should locate a different left to right word",
		puzzle:      []string{"coffeelplx"},
		words:       []string{"coffee"},
		expectError: false,
		expected:    map[string][2][2]int{"coffee": {{0, 0}, {5, 0}}},
	},
	{
		description: "Should locate that different left to right word in a different position",
		puzzle:      []string{"xcoffeezlp"},
		words:       []string{"coffee"},
		expectError: false,
		expected:    map[string][2][2]int{"coffee": {{1, 0}, {6, 0}}},
	},
	{
		description: "Should locate a left to right word in two line grid",
		puzzle:      []string{"jefblpepre", "tclojurerm"},
		words:       []string{"clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{1, 1}, {7, 1}}},
	},
	{
		description: "Should locate a left to right word in three line grid",
		puzzle:      []string{"camdcimgtc", "jefblpepre", "clojurermt"},
		words:       []string{"clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 2}, {6, 2}}},
	},
	{
		description: "Should locate a left to right word in ten line grid",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}},
	},
	{
		description: "Should locate that left to right word in a different position in a ten line grid",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "clojurermt", "jalaycalmp"},
		words:       []string{"clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 8}, {6, 8}}},
	},
	{
		description: "Should locate a different left to right word in a ten line grid",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "fortranftw", "alxhpburyi", "clojurermt", "jalaycalmp"},
		words:       []string{"fortran"},
		expectError: false,
		expected:    map[string][2][2]int{"fortran": {{0, 6}, {6, 6}}},
	},
	{
		description: "Should locate multiple words",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "fortranftw", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"fortran", "clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "fortran": {{0, 6}, {6, 6}}},
	},
	{
		description: "Should locate a single word written right to left",
		puzzle:      []string{"rixilelhrs"},
		words:       []string{"elixir"},
		expectError: false,
		expected:    map[string][2][2]int{"elixir": {{5, 0}, {0, 0}}},
	},
	{
		description: "Should locate multiple words written in different horizontal directions",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"elixir", "clojure"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "elixir": {{5, 4}, {0, 4}}},
	},
	{
		description: "Should locate words written top to bottom",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure", "elixir", "ecmascript"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}},
	},
	{
		description: "Should locate words written bottom to top",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure", "elixir", "ecmascript", "rust"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "rust": {{8, 4}, {8, 1}}},
	},
	{
		description: "Should locate words written top left to bottom right",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure", "elixir", "ecmascript", "rust", "java"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "rust": {{8, 4}, {8, 1}}},
	},
	{
		description: "Should locate words written bottom right to top left",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure", "elixir", "ecmascript", "rust", "java", "lua"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "lua": {{7, 8}, {5, 6}}, "rust": {{8, 4}, {8, 1}}},
	},
	{
		description: "Should locate words written bottom left to top right",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "lisp": {{2, 5}, {5, 2}}, "lua": {{7, 8}, {5, 6}}, "rust": {{8, 4}, {8, 1}}},
	},
	{
		description: "Should locate words written top right to bottom left",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp", "ruby"},
		expectError: false,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "lisp": {{2, 5}, {5, 2}}, "lua": {{7, 8}, {5, 6}}, "ruby": {{7, 5}, {4, 8}}, "rust": {{8, 4}, {8, 1}}},
	},
	{
		description: "Should fail to locate a word that is not in the puzzle",
		puzzle:      []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		words:       []string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp", "ruby", "haskell"},
		expectError: true,
		expected:    map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "haskell": {{-1, -1}, {-1, -1}}, "java": {{0, 0}, {3, 3}}, "lisp": {{2, 5}, {5, 2}}, "lua": {{7, 8}, {5, 6}}, "ruby": {{7, 5}, {4, 8}}, "rust": {{8, 4}, {8, 1}}},
	},
	{
		description: "Should fail to locate words that are not on horizontal, vertical, or diagonal lines",
		puzzle:      []string{"abc", "def"},
		words:       []string{"aef", "ced", "abf", "cbd"},
		expectError: true,
		expected:    map[string][2][2]int{"abf": {{-1, -1}, {-1, -1}}, "aef": {{-1, -1}, {-1, -1}}, "cbd": {{-1, -1}, {-1, -1}}, "ced": {{-1, -1}, {-1, -1}}},
	},
	{
		description: "Should not concatenate different lines to find a horizontal word",
		puzzle:      []string{"abceli", "xirdfg"},
		words:       []string{"elixir"},
		expectError: true,
		expected:    map[string][2][2]int{"elixir": {{-1, -1}, {-1, -1}}},
	},
	{
		description: "Should not wrap around horizontally to find a word",
		puzzle:      []string{"silabcdefp"},
		words:       []string{"lisp"},
		expectError: true,
		expected:    map[string][2][2]int{"lisp": {{-1, -1}, {-1, -1}}},
	},
	{
		description: "Should not wrap around vertically to find a word",
		puzzle:      []string{"s", "u", "r", "a", "b", "c", "t"},
		words:       []string{"rust"},
		expectError: true,
		expected:    map[string][2][2]int{"rust": {{-1, -1}, {-1, -1}}},
	},
}
