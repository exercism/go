package restapi

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var getUsersTestCases = []struct {
	description string
	database    []User
	payload     GetUsersRequest
	response    GetUsersResponse
}{
	{
		description: "no users",
		database:    []User{},
		payload:     GetUsersRequest{Users: []string(nil)},
		response: GetUsersResponse{
			Users: []User{},
		},
	},
	{
		description: "get single user",
		database: []User{
			User{
				Name:    "Adam",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{},
				Balance: 0,
			},
			User{
				Name:    "Bob",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{},
				Balance: 0,
			},
		},
		payload: GetUsersRequest{Users: []string{"Bob"}},
		response: GetUsersResponse{
			Users: []User{
				User{
					Name:    "Bob",
					Owes:    map[string]float64{},
					OwedBy:  map[string]float64{},
					Balance: 0,
				},
			},
		},
	},
}

var addUserTestCases = []struct {
	description string
	database    []User
	payload     AddUserRequest
	response    AddUserResponse
}{
	{
		description: "add user",
		database:    []User{},
		payload:     AddUserRequest{User: "Adam"},
		response: AddUserResponse{
			User: User{
				Name:    "Adam",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{},
				Balance: 0,
			},
		},
	},
}

var addIouTestCases = []struct {
	description string
	database    []User
	payload     AddIouRequest
	response    AddIouResponse
}{
	{
		description: "both users have 0 balance",
		database: []User{
			User{
				Name:    "Adam",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{},
				Balance: 0,
			},
			User{
				Name:    "Bob",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{},
				Balance: 0,
			},
		},
		payload: AddIouRequest{
			Lender:   "Adam",
			Borrower: "Bob",
			Amount:   3,
		},
		response: AddIouResponse{
			Users: []User{
				User{
					Name:    "Adam",
					Owes:    map[string]float64{},
					OwedBy:  map[string]float64{"Bob": 3},
					Balance: 3,
				},
				User{
					Name:    "Bob",
					Owes:    map[string]float64{"Adam": 3},
					OwedBy:  map[string]float64{},
					Balance: -3,
				},
			},
		},
	},
	{
		description: "borrower has negative balance",
		database: []User{
			User{
				Name:    "Adam",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{},
				Balance: 0,
			},
			User{
				Name:    "Bob",
				Owes:    map[string]float64{"Chuck": 3},
				OwedBy:  map[string]float64{},
				Balance: -3,
			},
			User{
				Name:    "Chuck",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{"Bob": 3},
				Balance: 3,
			},
		},
		payload: AddIouRequest{
			Lender:   "Adam",
			Borrower: "Bob",
			Amount:   3,
		},
		response: AddIouResponse{
			Users: []User{
				User{
					Name:    "Adam",
					Owes:    map[string]float64{},
					OwedBy:  map[string]float64{"Bob": 3},
					Balance: 3,
				},
				User{
					Name:    "Bob",
					Owes:    map[string]float64{"Adam": 3, "Chuck": 3},
					OwedBy:  map[string]float64{},
					Balance: -6,
				},
			},
		},
	},
	{
		description: "lender has negative balance",
		database: []User{
			User{
				Name:    "Adam",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{},
				Balance: 0,
			},
			User{
				Name:    "Bob",
				Owes:    map[string]float64{"Chuck": 3},
				OwedBy:  map[string]float64{},
				Balance: -3,
			},
			User{
				Name:    "Chuck",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{"Bob": 3},
				Balance: 3,
			},
		},
		payload: AddIouRequest{
			Lender:   "Bob",
			Borrower: "Adam",
			Amount:   3,
		},
		response: AddIouResponse{
			Users: []User{
				User{
					Name:    "Adam",
					Owes:    map[string]float64{"Bob": 3},
					OwedBy:  map[string]float64{},
					Balance: -3,
				},
				User{
					Name:    "Bob",
					Owes:    map[string]float64{"Chuck": 3},
					OwedBy:  map[string]float64{"Adam": 3},
					Balance: 0,
				},
			},
		},
	},
	{
		description: "lender owes borrower",
		database: []User{
			User{
				Name:    "Adam",
				Owes:    map[string]float64{"Bob": 3},
				OwedBy:  map[string]float64{},
				Balance: -3,
			},
			User{
				Name:    "Bob",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{"Adam": 3},
				Balance: 3,
			},
		},
		payload: AddIouRequest{
			Lender:   "Adam",
			Borrower: "Bob",
			Amount:   2,
		},
		response: AddIouResponse{
			Users: []User{
				User{
					Name:    "Adam",
					Owes:    map[string]float64{"Bob": 1},
					OwedBy:  map[string]float64{},
					Balance: -1,
				},
				User{
					Name:    "Bob",
					Owes:    map[string]float64{},
					OwedBy:  map[string]float64{"Adam": 1},
					Balance: 1,
				},
			},
		},
	},
	{
		description: "lender owes borrower less than new loan",
		database: []User{
			User{
				Name:    "Adam",
				Owes:    map[string]float64{"Bob": 3},
				OwedBy:  map[string]float64{},
				Balance: -3,
			},
			User{
				Name:    "Bob",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{"Adam": 3},
				Balance: 3,
			},
		},
		payload: AddIouRequest{
			Lender:   "Adam",
			Borrower: "Bob",
			Amount:   4,
		},
		response: AddIouResponse{
			Users: []User{
				User{
					Name:    "Adam",
					Owes:    map[string]float64{},
					OwedBy:  map[string]float64{"Bob": 1},
					Balance: 1,
				},
				User{
					Name:    "Bob",
					Owes:    map[string]float64{"Adam": 1},
					OwedBy:  map[string]float64{},
					Balance: -1,
				},
			},
		},
	},
	{
		description: "lender owes borrower same as new loan",
		database: []User{
			User{
				Name:    "Adam",
				Owes:    map[string]float64{"Bob": 3},
				OwedBy:  map[string]float64{},
				Balance: -3,
			},
			User{
				Name:    "Bob",
				Owes:    map[string]float64{},
				OwedBy:  map[string]float64{"Adam": 3},
				Balance: 3,
			},
		},
		payload: AddIouRequest{
			Lender:   "Adam",
			Borrower: "Bob",
			Amount:   3,
		},
		response: AddIouResponse{
			Users: []User{
				User{
					Name:    "Adam",
					Owes:    map[string]float64{},
					OwedBy:  map[string]float64{},
					Balance: 0,
				},
				User{
					Name:    "Bob",
					Owes:    map[string]float64{},
					OwedBy:  map[string]float64{},
					Balance: 0,
				},
			},
		},
	},
}
