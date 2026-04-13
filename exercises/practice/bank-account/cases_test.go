package bankaccount

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 0df8c72 [Bank account]: Add canonical data (#2192)

type Operation struct {
	Name   string
	Amount int64
}

type testCase struct {
	description string
	operations  []Operation
	expected    int64
	expectedErr string
}

var testCases = []testCase{
	{
		description: "Newly opened account has zero balance",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "balance"},
		},
		expected: 0,
	},
	{
		description: "Single deposit",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 100},
			Operation{Name: "balance"},
		},
		expected: 100,
	},
	{
		description: "Multiple deposits",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 100},
			Operation{Name: "deposit", Amount: 50},
			Operation{Name: "balance"},
		},
		expected: 150,
	},
	{
		description: "Withdraw once",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 100},
			Operation{Name: "withdraw", Amount: 75},
			Operation{Name: "balance"},
		},
		expected: 25,
	},
	{
		description: "Withdraw twice",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 100},
			Operation{Name: "withdraw", Amount: 80},
			Operation{Name: "withdraw", Amount: 20},
			Operation{Name: "balance"},
		},
		expected: 0,
	},
	{
		description: "Can do multiple operations sequentially",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 100},
			Operation{Name: "deposit", Amount: 110},
			Operation{Name: "withdraw", Amount: 200},
			Operation{Name: "deposit", Amount: 60},
			Operation{Name: "withdraw", Amount: 50},
			Operation{Name: "balance"},
		},
		expected: 20,
	},
	{
		description: "Cannot check balance of closed account",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "close"},
			Operation{Name: "balance"},
		},
		expectedErr: "account not open",
	},
	{
		description: "Cannot deposit into closed account",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "close"},
			Operation{Name: "deposit", Amount: 50},
		},
		expectedErr: "account not open",
	},
	{
		description: "Cannot deposit into unopened account",
		operations: []Operation{
			Operation{Name: "deposit", Amount: 50},
		},
		expectedErr: "account not open",
	},
	{
		description: "Cannot withdraw from closed account",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "close"},
			Operation{Name: "withdraw", Amount: 50},
		},
		expectedErr: "account not open",
	},
	{
		description: "Cannot close an account that was not opened",
		operations: []Operation{
			Operation{Name: "close"},
		},
		expectedErr: "account not open",
	},
	{
		description: "Cannot open an already opened account",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "open"},
		},
		expectedErr: "account already open",
	},
	{
		description: "Reopened account does not retain balance",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 50},
			Operation{Name: "close"},
			Operation{Name: "open"},
			Operation{Name: "balance"},
		},
		expected: 0,
	},
	{
		description: "Cannot withdraw more than deposited",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 25},
			Operation{Name: "withdraw", Amount: 50},
		},
		expectedErr: "amount must be less than balance",
	},
	{
		description: "Cannot withdraw negative",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: 100},
			Operation{Name: "withdraw", Amount: -50},
		},
		expectedErr: "amount must be greater than 0",
	},
	{
		description: "Cannot deposit negative",
		operations: []Operation{
			Operation{Name: "open"},
			Operation{Name: "deposit", Amount: -50},
		},
		expectedErr: "amount must be greater than 0",
	},
}
