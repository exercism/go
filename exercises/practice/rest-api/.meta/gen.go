package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"get":  &[]getTestCase{},
		"post": &[]postTestCase{},
	}
	if err := gen.Gen("rest-api", j, t); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Name    string             `json:"name"`
	Owes    map[string]float64 `json:"owes"`
	OwedBy  map[string]float64 `json:"owed_by"`
	Balance float64            `json:"balance"`
}

type GetUserInput struct {
	Database struct {
		Users []User `json:"users"`
	} `json:"database"`
	Url     string `json:"url"`
	Payload struct {
		Users []string `json:"users"`
	} `json:"payload"`
}

type getTestCase struct {
	Description string       `json:"description"`
	Input       GetUserInput `json:"input"`
	Expected    struct {
		Users []User `json:"users"`
	} `json:"expected"`
}

type postTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Database struct {
			Users []User `json:"users"`
		} `json:"database"`
		Url     string `json:"url"`
		Payload struct {
			User     string  `json:"user"`
			Lender   string  `json:"lender"`
			Borrower string  `json:"borrower"`
			Amount   float64 `json:"amount"`
		} `json:"payload"`
	} `json:"input"`
	Expected struct {
		User
		Users []User `json:"users"`
	} `json:"expected"`
}

var tmpl = `{{.Header}}

type testCaseGetUsers struct {
	description string
	database    []User
	payload     GetUsersRequest
	response    GetUsersResponse
}

var getUsersTestCases = []testCaseGetUsers { {{range .J.get}}
	{
		description: {{printf "%q" .Description}},
		database:    []User { {{range .Input.Database.Users }}
			User{
				Name:    {{printf "%q" .Name}},
				Owes:    {{printf "%#v" .Owes}},
				OwedBy:  {{printf "%#v" .OwedBy}},
				Balance: {{.Balance}},
			},{{end}}
		}, {{if ne .Input.Payload nil}}
		payload:     GetUsersRequest {Users: {{printf "%#v" .Input.Payload.Users}} }, {{end}}
		response:    GetUsersResponse { 
			Users: []User{ {{range .Expected.Users}}
				User{
					Name:    {{printf "%q" .Name}},
					Owes:    {{printf "%#v" .Owes}},
					OwedBy:  {{printf "%#v" .OwedBy}},
					Balance: {{.Balance}},
				},{{end}}
			},
		},
	},{{end}}
}

type testCaseAddUser struct {
	description string
	database    []User
	payload     AddUserRequest
	response    AddUserResponse
}

var addUserTestCases = []testCaseAddUser { {{range .J.post}} {{if eq .Input.Url "/add"}}
	{
		description: {{printf "%q" .Description}},
		database:    []User { {{range .Input.Database.Users }}
			User{
				Name:    {{printf "%q" .Name}},
				Owes:    {{printf "%#v" .Owes}},
				OwedBy:  {{printf "%#v" .OwedBy}},
				Balance: {{.Balance}},
			},{{end}}
		},
		payload:     AddUserRequest{User: {{printf "%q" .Input.Payload.User}}},
		response:    AddUserResponse{
			User: User{
				Name:    {{printf "%q" .Expected.Name}},
				Owes:    {{printf "%#v" .Expected.Owes}},
				OwedBy:  {{printf "%#v" .Expected.OwedBy}},
				Balance: {{.Expected.Balance}},
			},
		},
	},{{end}}{{end}}
}

type testCaseAddIOU struct {
	description string
	database    []User
	payload     AddIouRequest
	response    AddIouResponse
}

var addIouTestCases = []testCaseAddIOU { {{range .J.post}} {{if eq .Input.Url "/iou"}}
	{
		description: {{printf "%q" .Description}},
		database:    []User { {{range .Input.Database.Users }}
			User{
				Name:    {{printf "%q" .Name}},
				Owes:    {{printf "%#v" .Owes}},
				OwedBy:  {{printf "%#v" .OwedBy}},
				Balance: {{.Balance}},
			},{{end}}
		},
		payload:     AddIouRequest{
			Lender:   {{printf "%q" .Input.Payload.Lender}},
			Borrower: {{printf "%q" .Input.Payload.Borrower}},
			Amount:   {{.Input.Payload.Amount}},
		},
		response:    AddIouResponse { 
			Users: []User{ {{range .Expected.Users}}
				User{
					Name:    {{printf "%q" .Name}},
					Owes:    {{printf "%#v" .Owes}},
					OwedBy:  {{printf "%#v" .OwedBy}},
					Balance: {{.Balance}},
				},{{end}}
			},
		},
	},{{end}}{{end}}
}
`
