package markdown

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b820099 Allow prettier to format more files (#1966)

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "parses normal text as a paragraph",
		input:       "This will be a paragraph",
		expected:    "<p>This will be a paragraph</p>",
	},
	{
		description: "parsing italics",
		input:       "_This will be italic_",
		expected:    "<p><em>This will be italic</em></p>",
	},
	{
		description: "parsing bold text",
		input:       "__This will be bold__",
		expected:    "<p><strong>This will be bold</strong></p>",
	},
	{
		description: "mixed normal, italics and bold text",
		input:       "This will _be_ __mixed__",
		expected:    "<p>This will <em>be</em> <strong>mixed</strong></p>",
	},
	{
		description: "with h1 header level",
		input:       "# This will be an h1",
		expected:    "<h1>This will be an h1</h1>",
	},
	{
		description: "with h2 header level",
		input:       "## This will be an h2",
		expected:    "<h2>This will be an h2</h2>",
	},
	{
		description: "with h3 header level",
		input:       "### This will be an h3",
		expected:    "<h3>This will be an h3</h3>",
	},
	{
		description: "with h4 header level",
		input:       "#### This will be an h4",
		expected:    "<h4>This will be an h4</h4>",
	},
	{
		description: "with h5 header level",
		input:       "##### This will be an h5",
		expected:    "<h5>This will be an h5</h5>",
	},
	{
		description: "with h6 header level",
		input:       "###### This will be an h6",
		expected:    "<h6>This will be an h6</h6>",
	},
	{
		description: "h7 header level is a paragraph",
		input:       "####### This will not be an h7",
		expected:    "<p>####### This will not be an h7</p>",
	},
	{
		description: "unordered lists",
		input:       "* Item 1\n* Item 2",
		expected:    "<ul><li>Item 1</li><li>Item 2</li></ul>",
	},
	{
		description: "With a little bit of everything",
		input:       "# Header!\n* __Bold Item__\n* _Italic Item_",
		expected:    "<h1>Header!</h1><ul><li><strong>Bold Item</strong></li><li><em>Italic Item</em></li></ul>",
	},
	{
		description: "with markdown symbols in the header text that should not be interpreted",
		input:       "# This is a header with # and * in the text",
		expected:    "<h1>This is a header with # and * in the text</h1>",
	},
	{
		description: "with markdown symbols in the list item text that should not be interpreted",
		input:       "* Item 1 with a # in the text\n* Item 2 with * in the text",
		expected:    "<ul><li>Item 1 with a # in the text</li><li>Item 2 with * in the text</li></ul>",
	},
	{
		description: "with markdown symbols in the paragraph text that should not be interpreted",
		input:       "This is a paragraph with # and * in the text",
		expected:    "<p>This is a paragraph with # and * in the text</p>",
	},
	{
		description: "unordered lists close properly with preceding and following lines",
		input:       "# Start a list\n* Item 1\n* Item 2\nEnd a list",
		expected:    "<h1>Start a list</h1><ul><li>Item 1</li><li>Item 2</li></ul><p>End a list</p>",
	},
}
