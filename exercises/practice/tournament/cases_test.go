package tournament

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b820099 Allow prettier to format more files (#1966)

type TestCase = struct {
	description string
	input       string
	expected    string
}

var testCases = []TestCase{
	{
		description: "just the header if no input",
		input: `

`,
		expected: `
Team                           | MP |  W |  D |  L |  P
`,
	},
	{
		description: "a win is three points, a loss is zero points",
		input: `
Allegoric Alaskans;Blithering Badgers;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  1 |  1 |  0 |  0 |  3
Blithering Badgers             |  1 |  0 |  0 |  1 |  0
`,
	},
	{
		description: "a win can also be expressed as a loss",
		input: `
Blithering Badgers;Allegoric Alaskans;loss
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  1 |  1 |  0 |  0 |  3
Blithering Badgers             |  1 |  0 |  0 |  1 |  0
`,
	},
	{
		description: "a different team can win",
		input: `
Blithering Badgers;Allegoric Alaskans;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Blithering Badgers             |  1 |  1 |  0 |  0 |  3
Allegoric Alaskans             |  1 |  0 |  0 |  1 |  0
`,
	},
	{
		description: "a draw is one point each",
		input: `
Allegoric Alaskans;Blithering Badgers;draw
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  1 |  0 |  1 |  0 |  1
Blithering Badgers             |  1 |  0 |  1 |  0 |  1
`,
	},
	{
		description: "There can be more than one match",
		input: `
Allegoric Alaskans;Blithering Badgers;win
Allegoric Alaskans;Blithering Badgers;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  2 |  2 |  0 |  0 |  6
Blithering Badgers             |  2 |  0 |  0 |  2 |  0
`,
	},
	{
		description: "There can be more than one winner",
		input: `
Allegoric Alaskans;Blithering Badgers;loss
Allegoric Alaskans;Blithering Badgers;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  2 |  1 |  0 |  1 |  3
Blithering Badgers             |  2 |  1 |  0 |  1 |  3
`,
	},
	{
		description: "There can be more than two teams",
		input: `
Allegoric Alaskans;Blithering Badgers;win
Blithering Badgers;Courageous Californians;win
Courageous Californians;Allegoric Alaskans;loss
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  2 |  2 |  0 |  0 |  6
Blithering Badgers             |  2 |  1 |  0 |  1 |  3
Courageous Californians        |  2 |  0 |  0 |  2 |  0
`,
	},
	{
		description: "typical input",
		input: `
Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskans;Courageous Californians;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
`,
	},
	{
		description: "incomplete competition (not all pairs have played)",
		input: `
Allegoric Alaskans;Blithering Badgers;loss
Devastating Donkeys;Allegoric Alaskans;loss
Courageous Californians;Blithering Badgers;draw
Allegoric Alaskans;Courageous Californians;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  2 |  1 |  1 |  0 |  4
Courageous Californians        |  2 |  0 |  1 |  1 |  1
Devastating Donkeys            |  1 |  0 |  0 |  1 |  0
`,
	},
	{
		description: "ties broken alphabetically",
		input: `
Courageous Californians;Devastating Donkeys;win
Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskans;loss
Courageous Californians;Blithering Badgers;win
Blithering Badgers;Devastating Donkeys;draw
Allegoric Alaskans;Courageous Californians;draw
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Allegoric Alaskans             |  3 |  2 |  1 |  0 |  7
Courageous Californians        |  3 |  2 |  1 |  0 |  7
Blithering Badgers             |  3 |  0 |  1 |  2 |  1
Devastating Donkeys            |  3 |  0 |  1 |  2 |  1
`,
	},
	{
		description: "ensure points sorted numerically",
		input: `
Devastating Donkeys;Blithering Badgers;win
Devastating Donkeys;Blithering Badgers;win
Devastating Donkeys;Blithering Badgers;win
Devastating Donkeys;Blithering Badgers;win
Blithering Badgers;Devastating Donkeys;win
`,
		expected: `
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  5 |  4 |  0 |  1 | 12
Blithering Badgers             |  5 |  1 |  0 |  4 |  3
`,
	},
}
