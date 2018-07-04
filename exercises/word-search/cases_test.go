package wordsearch

// Source: exercism/problem-specifications
// Commit: 2b02251 word-search: Remove unneeded exception case comment
// Problem Specifications Version: 1.2.1

var testCases = []struct {
	description string
	puzzle      []string             // puzzle strings
	words       []string             // words to search for
	expected    map[string][2][2]int // expected coordinates
	expectError bool
}{
	{
		"Should accept an initial game grid and a target search word",
		[]string{"jefblpepre"},
		[]string{"clojure"},
		map[string][2][2]int{},
		true,
	},
	{
		"Should locate one word written left to right",
		[]string{"clojurermt"},
		[]string{"clojure"},
		map[string][2][2]int{"clojure": {{0, 0}, {6, 0}}},
		false,
	},
	{
		"Should locate the same word written left to right in a different position",
		[]string{"mtclojurer"},
		[]string{"clojure"},
		map[string][2][2]int{"clojure": {{2, 0}, {8, 0}}},
		false,
	},
	{
		"Should locate a different left to right word",
		[]string{"coffeelplx"},
		[]string{"coffee"},
		map[string][2][2]int{"coffee": {{0, 0}, {5, 0}}},
		false,
	},
	{
		"Should locate that different left to right word in a different position",
		[]string{"xcoffeezlp"},
		[]string{"coffee"},
		map[string][2][2]int{"coffee": {{1, 0}, {6, 0}}},
		false,
	},
	{
		"Should locate a left to right word in two line grid",
		[]string{"jefblpepre", "tclojurerm"},
		[]string{"clojure"},
		map[string][2][2]int{"clojure": {{1, 1}, {7, 1}}},
		false,
	},
	{
		"Should locate a left to right word in three line grid",
		[]string{"camdcimgtc", "jefblpepre", "clojurermt"},
		[]string{"clojure"},
		map[string][2][2]int{"clojure": {{0, 2}, {6, 2}}},
		false,
	},
	{
		"Should locate a left to right word in ten line grid",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}},
		false,
	},
	{
		"Should locate that left to right word in a different position in a ten line grid",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "clojurermt", "jalaycalmp"},
		[]string{"clojure"},
		map[string][2][2]int{"clojure": {{0, 8}, {6, 8}}},
		false,
	},
	{
		"Should locate a different left to right word in a ten line grid",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "fortranftw", "alxhpburyi", "clojurermt", "jalaycalmp"},
		[]string{"fortran"},
		map[string][2][2]int{"fortran": {{0, 6}, {6, 6}}},
		false,
	},
	{
		"Should locate multiple words",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "fortranftw", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"fortran", "clojure"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "fortran": {{0, 6}, {6, 6}}},
		false,
	},
	{
		"Should locate a single word written right to left",
		[]string{"rixilelhrs"},
		[]string{"elixir"},
		map[string][2][2]int{"elixir": {{5, 0}, {0, 0}}},
		false,
	},
	{
		"Should locate multiple words written in different horizontal directions",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"elixir", "clojure"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "elixir": {{5, 4}, {0, 4}}},
		false,
	},
	{
		"Should locate words written top to bottom",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure", "elixir", "ecmascript"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}},
		false,
	},
	{
		"Should locate words written bottom to top",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure", "elixir", "ecmascript", "rust"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "rust": {{8, 4}, {8, 1}}},
		false,
	},
	{
		"Should locate words written top left to bottom right",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure", "elixir", "ecmascript", "rust", "java"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "rust": {{8, 4}, {8, 1}}},
		false,
	},
	{
		"Should locate words written bottom right to top left",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure", "elixir", "ecmascript", "rust", "java", "lua"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "lua": {{7, 8}, {5, 6}}, "rust": {{8, 4}, {8, 1}}},
		false,
	},
	{
		"Should locate words written bottom left to top right",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "lisp": {{2, 5}, {5, 2}}, "lua": {{7, 8}, {5, 6}}, "rust": {{8, 4}, {8, 1}}},
		false,
	},
	{
		"Should locate words written top right to bottom left",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp", "ruby"},
		map[string][2][2]int{"clojure": {{0, 9}, {6, 9}}, "ecmascript": {{9, 0}, {9, 9}}, "elixir": {{5, 4}, {0, 4}}, "java": {{0, 0}, {3, 3}}, "lisp": {{2, 5}, {5, 2}}, "lua": {{7, 8}, {5, 6}}, "ruby": {{7, 5}, {4, 8}}, "rust": {{8, 4}, {8, 1}}},
		false,
	},
	{
		"Should fail to locate a word that is not in the puzzle",
		[]string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"},
		[]string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp", "ruby", "haskell"},
		map[string][2][2]int{},
		true,
	},
}
