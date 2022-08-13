package forth

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b230e1e forth: Add local-scope test

var testCases = []struct {
	description string
	input       []string
	expected    []int  // nil slice indicates error expected.
	explainText string // error explanation text
}{
	{
		description: "numbers just get pushed onto the stack",
		input:       []string{"1 2 3 4 5"},
		expected:    []int{1, 2, 3, 4, 5},
		explainText: "",
	},
	{
		description: "pushes negative numbers onto the stack",
		input:       []string{"-1 -2 -3 -4 -5"},
		expected:    []int{-1, -2, -3, -4, -5},
		explainText: "",
	},
	{
		description: "can add two numbers",
		input:       []string{"1 2 +"},
		expected:    []int{3},
		explainText: "",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"+"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "errors if there is only one value on the stack",
		input:       []string{"1 +"},
		expected:    []int(nil),
		explainText: "only one value on the stack",
	},
	{
		description: "can subtract two numbers",
		input:       []string{"3 4 -"},
		expected:    []int{-1},
		explainText: "",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"-"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "errors if there is only one value on the stack",
		input:       []string{"1 -"},
		expected:    []int(nil),
		explainText: "only one value on the stack",
	},
	{
		description: "can multiply two numbers",
		input:       []string{"2 4 *"},
		expected:    []int{8},
		explainText: "",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"*"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "errors if there is only one value on the stack",
		input:       []string{"1 *"},
		expected:    []int(nil),
		explainText: "only one value on the stack",
	},
	{
		description: "can divide two numbers",
		input:       []string{"12 3 /"},
		expected:    []int{4},
		explainText: "",
	},
	{
		description: "performs integer division",
		input:       []string{"8 3 /"},
		expected:    []int{2},
		explainText: "",
	},
	{
		description: "errors if dividing by zero",
		input:       []string{"4 0 /"},
		expected:    []int(nil),
		explainText: "divide by zero",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"/"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "errors if there is only one value on the stack",
		input:       []string{"1 /"},
		expected:    []int(nil),
		explainText: "only one value on the stack",
	},
	{
		description: "addition and subtraction",
		input:       []string{"1 2 + 4 -"},
		expected:    []int{-1},
		explainText: "",
	},
	{
		description: "multiplication and division",
		input:       []string{"2 4 * 3 /"},
		expected:    []int{2},
		explainText: "",
	},
	{
		description: "copies a value on the stack",
		input:       []string{"1 dup"},
		expected:    []int{1, 1},
		explainText: "",
	},
	{
		description: "copies the top value on the stack",
		input:       []string{"1 2 dup"},
		expected:    []int{1, 2, 2},
		explainText: "",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"dup"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "removes the top value on the stack if it is the only one",
		input:       []string{"1 drop"},
		expected:    []int{},
		explainText: "",
	},
	{
		description: "removes the top value on the stack if it is not the only one",
		input:       []string{"1 2 drop"},
		expected:    []int{1},
		explainText: "",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"drop"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "swaps the top two values on the stack if they are the only ones",
		input:       []string{"1 2 swap"},
		expected:    []int{2, 1},
		explainText: "",
	},
	{
		description: "swaps the top two values on the stack if they are not the only ones",
		input:       []string{"1 2 3 swap"},
		expected:    []int{1, 3, 2},
		explainText: "",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"swap"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "errors if there is only one value on the stack",
		input:       []string{"1 swap"},
		expected:    []int(nil),
		explainText: "only one value on the stack",
	},
	{
		description: "copies the second element if there are only two",
		input:       []string{"1 2 over"},
		expected:    []int{1, 2, 1},
		explainText: "",
	},
	{
		description: "copies the second element if there are more than two",
		input:       []string{"1 2 3 over"},
		expected:    []int{1, 2, 3, 2},
		explainText: "",
	},
	{
		description: "errors if there is nothing on the stack",
		input:       []string{"over"},
		expected:    []int(nil),
		explainText: "empty stack",
	},
	{
		description: "errors if there is only one value on the stack",
		input:       []string{"1 over"},
		expected:    []int(nil),
		explainText: "only one value on the stack",
	},
	{
		description: "can consist of built-in words",
		input:       []string{": dup-twice dup dup ;", "1 dup-twice"},
		expected:    []int{1, 1, 1},
		explainText: "",
	},
	{
		description: "execute in the right order",
		input:       []string{": countup 1 2 3 ;", "countup"},
		expected:    []int{1, 2, 3},
		explainText: "",
	},
	{
		description: "can override other user-defined words",
		input:       []string{": foo dup ;", ": foo dup dup ;", "1 foo"},
		expected:    []int{1, 1, 1},
		explainText: "",
	},
	{
		description: "can override built-in words",
		input:       []string{": swap dup ;", "1 swap"},
		expected:    []int{1, 1},
		explainText: "",
	},
	{
		description: "can override built-in operators",
		input:       []string{": + * ;", "3 4 +"},
		expected:    []int{12},
		explainText: "",
	},
	{
		description: "can use different words with the same name",
		input:       []string{": foo 5 ;", ": bar foo ;", ": foo 6 ;", "bar foo"},
		expected:    []int{5, 6},
		explainText: "",
	},
	{
		description: "can define word that uses word with the same name",
		input:       []string{": foo 10 ;", ": foo foo 1 + ;", "foo"},
		expected:    []int{11},
		explainText: "",
	},
	{
		description: "cannot redefine non-negative numbers",
		input:       []string{": 1 2 ;"},
		expected:    []int(nil),
		explainText: "illegal operation",
	},
	{
		description: "cannot redefine negative numbers",
		input:       []string{": -1 2 ;"},
		expected:    []int(nil),
		explainText: "illegal operation",
	},
	{
		description: "errors if executing a non-existent word",
		input:       []string{"foo"},
		expected:    []int(nil),
		explainText: "undefined operation",
	},
	{
		description: "DUP is case-insensitive",
		input:       []string{"1 DUP Dup dup"},
		expected:    []int{1, 1, 1, 1},
		explainText: "",
	},
	{
		description: "DROP is case-insensitive",
		input:       []string{"1 2 3 4 DROP Drop drop"},
		expected:    []int{1},
		explainText: "",
	},
	{
		description: "SWAP is case-insensitive",
		input:       []string{"1 2 SWAP 3 Swap 4 swap"},
		expected:    []int{2, 3, 4, 1},
		explainText: "",
	},
	{
		description: "OVER is case-insensitive",
		input:       []string{"1 2 OVER Over over"},
		expected:    []int{1, 2, 1, 2, 1},
		explainText: "",
	},
	{
		description: "user-defined words are case-insensitive",
		input:       []string{": foo dup ;", "1 FOO Foo foo"},
		expected:    []int{1, 1, 1, 1},
		explainText: "",
	},
	{
		description: "definitions are case-insensitive",
		input:       []string{": SWAP DUP Dup dup ;", "1 swap"},
		expected:    []int{1, 1, 1, 1},
		explainText: "",
	},
}
