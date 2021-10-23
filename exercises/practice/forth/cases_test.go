package forth

// Source: exercism/problem-specifications
// Commit: 75f4c0a Corrected minor typos in the error msg expectation (doesn't match other similar error patterns and so breaks auto generated tests)
// Problem Specifications Version: 1.7.1

type testGroup struct {
	group string
	tests []testCase
}

type testCase struct {
	description string
	input       []string
	expected    []int // nil slice indicates error expected.
}

var testGroups = []testGroup{
	{
		group: "parsing and numbers",
		tests: []testCase{
			{
				"numbers just get pushed onto the stack",
				[]string{"1 2 3 4 5"},
				[]int{1, 2, 3, 4, 5},
			},
		},
	},
	{
		group: "addition",
		tests: []testCase{
			{
				"can add two numbers",
				[]string{"1 2 +"},
				[]int{3},
			},
			{
				"errors if there is nothing on the stack",
				[]string{"+"},
				[]int(nil),
			},
			{
				"errors if there is only one value on the stack",
				[]string{"1 +"},
				[]int(nil),
			},
		},
	},
	{
		group: "subtraction",
		tests: []testCase{
			{
				"can subtract two numbers",
				[]string{"3 4 -"},
				[]int{-1},
			},
			{
				"errors if there is nothing on the stack",
				[]string{"-"},
				[]int(nil),
			},
			{
				"errors if there is only one value on the stack",
				[]string{"1 -"},
				[]int(nil),
			},
		},
	},
	{
		group: "multiplication",
		tests: []testCase{
			{
				"can multiply two numbers",
				[]string{"2 4 *"},
				[]int{8},
			},
			{
				"errors if there is nothing on the stack",
				[]string{"*"},
				[]int(nil),
			},
			{
				"errors if there is only one value on the stack",
				[]string{"1 *"},
				[]int(nil),
			},
		},
	},
	{
		group: "division",
		tests: []testCase{
			{
				"can divide two numbers",
				[]string{"12 3 /"},
				[]int{4},
			},
			{
				"performs integer division",
				[]string{"8 3 /"},
				[]int{2},
			},
			{
				"errors if dividing by zero",
				[]string{"4 0 /"},
				[]int(nil),
			},
			{
				"errors if there is nothing on the stack",
				[]string{"/"},
				[]int(nil),
			},
			{
				"errors if there is only one value on the stack",
				[]string{"1 /"},
				[]int(nil),
			},
		},
	},
	{
		group: "combined arithmetic",
		tests: []testCase{
			{
				"addition and subtraction",
				[]string{"1 2 + 4 -"},
				[]int{-1},
			},
			{
				"multiplication and division",
				[]string{"2 4 * 3 /"},
				[]int{2},
			},
		},
	},
	{
		group: "dup",
		tests: []testCase{
			{
				"copies a value on the stack",
				[]string{"1 dup"},
				[]int{1, 1},
			},
			{
				"copies the top value on the stack",
				[]string{"1 2 dup"},
				[]int{1, 2, 2},
			},
			{
				"errors if there is nothing on the stack",
				[]string{"dup"},
				[]int(nil),
			},
		},
	},
	{
		group: "drop",
		tests: []testCase{
			{
				"removes the top value on the stack if it is the only one",
				[]string{"1 drop"},
				[]int{},
			},
			{
				"removes the top value on the stack if it is not the only one",
				[]string{"1 2 drop"},
				[]int{1},
			},
			{
				"errors if there is nothing on the stack",
				[]string{"drop"},
				[]int(nil),
			},
		},
	},
	{
		group: "swap",
		tests: []testCase{
			{
				"swaps the top two values on the stack if they are the only ones",
				[]string{"1 2 swap"},
				[]int{2, 1},
			},
			{
				"swaps the top two values on the stack if they are not the only ones",
				[]string{"1 2 3 swap"},
				[]int{1, 3, 2},
			},
			{
				"errors if there is nothing on the stack",
				[]string{"swap"},
				[]int(nil),
			},
			{
				"errors if there is only one value on the stack",
				[]string{"1 swap"},
				[]int(nil),
			},
		},
	},
	{
		group: "over",
		tests: []testCase{
			{
				"copies the second element if there are only two",
				[]string{"1 2 over"},
				[]int{1, 2, 1},
			},
			{
				"copies the second element if there are more than two",
				[]string{"1 2 3 over"},
				[]int{1, 2, 3, 2},
			},
			{
				"errors if there is nothing on the stack",
				[]string{"over"},
				[]int(nil),
			},
			{
				"errors if there is only one value on the stack",
				[]string{"1 over"},
				[]int(nil),
			},
		},
	},
	{
		group: "user-defined words",
		tests: []testCase{
			{
				"can consist of built-in words",
				[]string{": dup-twice dup dup ;", "1 dup-twice"},
				[]int{1, 1, 1},
			},
			{
				"execute in the right order",
				[]string{": countup 1 2 3 ;", "countup"},
				[]int{1, 2, 3},
			},
			{
				"can override other user-defined words",
				[]string{": foo dup ;", ": foo dup dup ;", "1 foo"},
				[]int{1, 1, 1},
			},
			{
				"can override built-in words",
				[]string{": swap dup ;", "1 swap"},
				[]int{1, 1},
			},
			{
				"can override built-in operators",
				[]string{": + * ;", "3 4 +"},
				[]int{12},
			},
			{
				"can use different words with the same name",
				[]string{": foo 5 ;", ": bar foo ;", ": foo 6 ;", "bar foo"},
				[]int{5, 6},
			},
			{
				"can define word that uses word with the same name",
				[]string{": foo 10 ;", ": foo foo 1 + ;", "foo"},
				[]int{11},
			},
			{
				"cannot redefine numbers",
				[]string{": 1 2 ;"},
				[]int(nil),
			},
			{
				"errors if executing a non-existent word",
				[]string{"foo"},
				[]int(nil),
			},
		},
	},
	{
		group: "case-insensitivity",
		tests: []testCase{
			{
				"DUP is case-insensitive",
				[]string{"1 DUP Dup dup"},
				[]int{1, 1, 1, 1},
			},
			{
				"DROP is case-insensitive",
				[]string{"1 2 3 4 DROP Drop drop"},
				[]int{1},
			},
			{
				"SWAP is case-insensitive",
				[]string{"1 2 SWAP 3 Swap 4 swap"},
				[]int{2, 3, 4, 1},
			},
			{
				"OVER is case-insensitive",
				[]string{"1 2 OVER Over over"},
				[]int{1, 2, 1, 2, 1},
			},
			{
				"user-defined words are case-insensitive",
				[]string{": foo dup ;", "1 FOO Foo foo"},
				[]int{1, 1, 1, 1},
			},
			{
				"definitions are case-insensitive",
				[]string{": SWAP DUP Dup dup ;", "1 swap"},
				[]int{1, 1, 1, 1},
			},
		},
	},
}
