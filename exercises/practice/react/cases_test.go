package react

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

type CreateCell struct {
	cType    string
	name     string
	value    int
	inputs   []string
	fn1      func(int) int
	fn2      func(int, int) int
	fnString string
}

type Operation struct {
	oType           string
	cell            string
	name            string
	value           int
	wantCallbacks   map[string]int
	wantNoCallbacks []string
}

type testCase struct {
	description string
	cells       []CreateCell
	operations  []Operation
}

var testCases = []testCase{
	{
		description: "input cells have a value",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 10},
		},
		operations: []Operation{
			Operation{oType: "Value", cell: "input", value: 10},
		},
	},
	{
		description: "an input cell's value can be set",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 4},
		},
		operations: []Operation{
			Operation{oType: "SetValue", cell: "input", value: 20},
			Operation{oType: "Value", cell: "input", value: 20},
		},
	},
	{
		description: "compute cells calculate initial value",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "output", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
		},
		operations: []Operation{
			Operation{oType: "Value", cell: "output", value: 2},
		},
	},
	{
		description: "compute cells take inputs in the right order",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "one", value: 1},
			CreateCell{cType: "input", name: "two", value: 2},
			CreateCell{cType: "compute2", name: "output", inputs: []string{"one", "two"}, fn2: func(x, y int) int { return x + y*10 }, fnString: "func(x, y int) int {return x + y * 10}"},
		},
		operations: []Operation{
			Operation{oType: "Value", cell: "output", value: 21},
		},
	},
	{
		description: "compute cells update value when dependencies are changed",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "output", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
		},
		operations: []Operation{
			Operation{oType: "SetValue", cell: "input", value: 3},
			Operation{oType: "Value", cell: "output", value: 4},
		},
	},
	{
		description: "compute cells can depend on other compute cells",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "times_two", inputs: []string{"input"}, fn1: func(x int) int { return x * 2 }, fnString: "func(x int) int {return x * 2}"},
			CreateCell{cType: "compute1", name: "times_thirty", inputs: []string{"input"}, fn1: func(x int) int { return x * 30 }, fnString: "func(x int) int {return x * 30}"},
			CreateCell{cType: "compute2", name: "output", inputs: []string{"times_two", "times_thirty"}, fn2: func(x, y int) int { return x + y }, fnString: "func(x, y int) int {return x + y}"},
		},
		operations: []Operation{
			Operation{oType: "Value", cell: "output", value: 32},
			Operation{oType: "SetValue", cell: "input", value: 3},
			Operation{oType: "Value", cell: "output", value: 96},
		},
	},
	{
		description: "compute cells fire callbacks",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "output", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "output", name: "callback1"},
			Operation{oType: "SetValue", cell: "input", value: 3, wantCallbacks: map[string]int{"callback1": 4}},
		},
	},
	{
		description: "callback cells only fire on change",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "output", inputs: []string{"input"}, fn1: func(x int) int {
				if x < 3 {
					return 111
				}
				return 222
			}, fnString: "func(x int) int {if x < 3 { return 111 }; return 222}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "output", name: "callback1"},
			Operation{oType: "SetValue", cell: "input", value: 2, wantNoCallbacks: []string{"callback1"}},
			Operation{oType: "SetValue", cell: "input", value: 4, wantCallbacks: map[string]int{"callback1": 222}},
		},
	},
	{
		description: "callbacks do not report already reported values",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "output", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "output", name: "callback1"},
			Operation{oType: "SetValue", cell: "input", value: 2, wantCallbacks: map[string]int{"callback1": 3}},
			Operation{oType: "SetValue", cell: "input", value: 3, wantCallbacks: map[string]int{"callback1": 4}},
		},
	},
	{
		description: "callbacks can fire from multiple cells",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "plus_one", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
			CreateCell{cType: "compute1", name: "minus_one", inputs: []string{"input"}, fn1: func(x int) int { return x - 1 }, fnString: "func(x int) int {return x - 1}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "plus_one", name: "callback1"},
			Operation{oType: "AddCallback", cell: "minus_one", name: "callback2"},
			Operation{oType: "SetValue", cell: "input", value: 10, wantCallbacks: map[string]int{"callback1": 11, "callback2": 9}},
		},
	},
	{
		description: "callbacks can be added and removed",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 11},
			CreateCell{cType: "compute1", name: "output", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "output", name: "callback1"},
			Operation{oType: "AddCallback", cell: "output", name: "callback2"},
			Operation{oType: "SetValue", cell: "input", value: 31, wantCallbacks: map[string]int{"callback1": 32, "callback2": 32}},
			Operation{oType: "Cancel", cell: "output", name: "callback1"},
			Operation{oType: "AddCallback", cell: "output", name: "callback3"},
			Operation{oType: "SetValue", cell: "input", value: 41, wantCallbacks: map[string]int{"callback2": 42, "callback3": 42}, wantNoCallbacks: []string{"callback1"}},
		},
	},
	{
		description: "removing a callback multiple times doesn't interfere with other callbacks",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "output", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "output", name: "callback1"},
			Operation{oType: "AddCallback", cell: "output", name: "callback2"},
			Operation{oType: "Cancel", cell: "output", name: "callback1"},
			Operation{oType: "Cancel", cell: "output", name: "callback1"},
			Operation{oType: "Cancel", cell: "output", name: "callback1"},
			Operation{oType: "SetValue", cell: "input", value: 2, wantCallbacks: map[string]int{"callback2": 3}, wantNoCallbacks: []string{"callback1"}},
		},
	},
	{
		description: "callbacks should only be called once even if multiple dependencies change",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "plus_one", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
			CreateCell{cType: "compute1", name: "minus_one1", inputs: []string{"input"}, fn1: func(x int) int { return x - 1 }, fnString: "func(x int) int {return x - 1}"},
			CreateCell{cType: "compute1", name: "minus_one2", inputs: []string{"minus_one1"}, fn1: func(x int) int { return x - 1 }, fnString: "func(x int) int {return x - 1}"},
			CreateCell{cType: "compute2", name: "output", inputs: []string{"plus_one", "minus_one2"}, fn2: func(x, y int) int { return x * y }, fnString: "func(x, y int) int {return x * y}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "output", name: "callback1"},
			Operation{oType: "SetValue", cell: "input", value: 4, wantCallbacks: map[string]int{"callback1": 10}},
		},
	},
	{
		description: "callbacks should not be called if dependencies change but output value doesn't change",
		cells: []CreateCell{
			CreateCell{cType: "input", name: "input", value: 1},
			CreateCell{cType: "compute1", name: "plus_one", inputs: []string{"input"}, fn1: func(x int) int { return x + 1 }, fnString: "func(x int) int {return x + 1}"},
			CreateCell{cType: "compute1", name: "minus_one", inputs: []string{"input"}, fn1: func(x int) int { return x - 1 }, fnString: "func(x int) int {return x - 1}"},
			CreateCell{cType: "compute2", name: "always_two", inputs: []string{"plus_one", "minus_one"}, fn2: func(x, y int) int { return x - y }, fnString: "func(x, y int) int {return x - y}"},
		},
		operations: []Operation{
			Operation{oType: "AddCallback", cell: "always_two", name: "callback1"},
			Operation{oType: "SetValue", cell: "input", value: 2, wantNoCallbacks: []string{"callback1"}},
			Operation{oType: "SetValue", cell: "input", value: 3, wantNoCallbacks: []string{"callback1"}},
			Operation{oType: "SetValue", cell: "input", value: 4, wantNoCallbacks: []string{"callback1"}},
			Operation{oType: "SetValue", cell: "input", value: 5, wantNoCallbacks: []string{"callback1"}},
		},
	},
}
