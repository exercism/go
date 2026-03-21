package circularbuffer

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

type Operation struct {
	name    string
	value   byte
	wantErr bool
}

type testCase struct {
	description string
	capacity    int
	operations  []Operation
}

var testCases = []testCase{
	{
		description: "reading empty buffer should fail",
		capacity:    1,
		operations: []Operation{
			Operation{name: "ReadByte", wantErr: true},
		},
	},
	{
		description: "can read an item just written",
		capacity:    1,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "ReadByte", value: '1', wantErr: false},
		},
	},
	{
		description: "each item may only be read once",
		capacity:    1,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "ReadByte", value: '1', wantErr: false},
			Operation{name: "ReadByte", wantErr: true},
		},
	},
	{
		description: "items are read in the order they are written",
		capacity:    2,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '2', wantErr: false},
			Operation{name: "ReadByte", value: '1', wantErr: false},
			Operation{name: "ReadByte", value: '2', wantErr: false},
		},
	},
	{
		description: "full buffer can't be written to",
		capacity:    1,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '2', wantErr: true},
		},
	},
	{
		description: "a read frees up capacity for another write",
		capacity:    1,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "ReadByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '2', wantErr: false},
			Operation{name: "ReadByte", value: '2', wantErr: false},
		},
	},
	{
		description: "read position is maintained even across multiple writes",
		capacity:    3,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '2', wantErr: false},
			Operation{name: "ReadByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '3', wantErr: false},
			Operation{name: "ReadByte", value: '2', wantErr: false},
			Operation{name: "ReadByte", value: '3', wantErr: false},
		},
	},
	{
		description: "items cleared out of buffer can't be read",
		capacity:    1,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "Reset"},
			Operation{name: "ReadByte", wantErr: true},
		},
	},
	{
		description: "clear frees up capacity for another write",
		capacity:    1,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "Reset"},
			Operation{name: "WriteByte", value: '2', wantErr: false},
			Operation{name: "ReadByte", value: '2', wantErr: false},
		},
	},
	{
		description: "clear does nothing on empty buffer",
		capacity:    1,
		operations: []Operation{
			Operation{name: "Reset"},
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "ReadByte", value: '1', wantErr: false},
		},
	},
	{
		description: "overwrite acts like write on non-full buffer",
		capacity:    2,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "Overwrite", value: '2'},
			Operation{name: "ReadByte", value: '1', wantErr: false},
			Operation{name: "ReadByte", value: '2', wantErr: false},
		},
	},
	{
		description: "overwrite replaces the oldest item on full buffer",
		capacity:    2,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '2', wantErr: false},
			Operation{name: "Overwrite", value: '3'},
			Operation{name: "ReadByte", value: '2', wantErr: false},
			Operation{name: "ReadByte", value: '3', wantErr: false},
		},
	},
	{
		description: "overwrite replaces the oldest item remaining in buffer following a read",
		capacity:    3,
		operations: []Operation{
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '2', wantErr: false},
			Operation{name: "WriteByte", value: '3', wantErr: false},
			Operation{name: "ReadByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '4', wantErr: false},
			Operation{name: "Overwrite", value: '5'},
			Operation{name: "ReadByte", value: '3', wantErr: false},
			Operation{name: "ReadByte", value: '4', wantErr: false},
			Operation{name: "ReadByte", value: '5', wantErr: false},
		},
	},
	{
		description: "initial clear does not affect wrapping around",
		capacity:    2,
		operations: []Operation{
			Operation{name: "Reset"},
			Operation{name: "WriteByte", value: '1', wantErr: false},
			Operation{name: "WriteByte", value: '2', wantErr: false},
			Operation{name: "Overwrite", value: '3'},
			Operation{name: "Overwrite", value: '4'},
			Operation{name: "ReadByte", value: '3', wantErr: false},
			Operation{name: "ReadByte", value: '4', wantErr: false},
			Operation{name: "ReadByte", wantErr: true},
		},
	},
}
