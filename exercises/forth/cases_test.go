package forth

// Source: exercism/x-common
// Commit: 8804a76 Remove non-ASCII characters tests from forth exercise

var parsingGroup = []testCase{
	{
		"empty input results in empty stack",
		[]string{},
		[]int{},
	},
	{
		"numbers just get pushed onto the stack",
		[]string{"1 2 3 4 5"},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"all non-word characters are separators",
		[]string{"1\x002\x133\n4\r5 6\t7"},
		[]int{1, 2, 3, 4, 5, 6, 7},
	},
}

var additionGroup = []testCase{
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
}

var subtractionGroup = []testCase{
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
}

var multiplicationGroup = []testCase{
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
}

var divisionGroup = []testCase{
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
}

var arithmeticGroup = []testCase{
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
}

var dupGroup = []testCase{
	{
		"copies the top value on the stack",
		[]string{"1 DUP"},
		[]int{1, 1},
	},
	{
		"is case-insensitive",
		[]string{"1 2 Dup"},
		[]int{1, 2, 2},
	},
	{
		"errors if there is nothing on the stack",
		[]string{"dup"},
		[]int(nil),
	},
}

var dropGroup = []testCase{
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
}

var swapGroup = []testCase{
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
}

var overGroup = []testCase{
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
}

var userdefinedGroup = []testCase{
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
		"cannot redefine numbers",
		[]string{": 1 2 ;"},
		[]int(nil),
	},
	{
		"errors if executing a non-existent word",
		[]string{"foo"},
		[]int(nil),
	},
}

var testSections = []testcaseSection{
	{"parsing", parsingGroup},
	{"addition(+)", additionGroup},
	{"subtraction(-)", subtractionGroup},
	{"multiplication(*)", multiplicationGroup},
	{"division(/)", divisionGroup},
	{"arithmetic", arithmeticGroup},
	{"dup", dupGroup},
	{"drop", dropGroup},
	{"swap", swapGroup},
	{"over", overGroup},
	{"user-defined", userdefinedGroup},
}
