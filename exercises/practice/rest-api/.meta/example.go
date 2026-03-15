package restapi

import (
	"cmp"
	"slices"
)

type User struct {
	Name    string
	Owes    map[string]float64
	OwedBy  map[string]float64
	Balance float64
}

type GetUsersRequest struct {
	Users []string
}

type GetUsersResponse struct {
	Users []User
}

type AddUserRequest struct {
	User string
}

type AddUserResponse struct {
	User User
}

type AddIouRequest struct {
	Lender   string
	Borrower string
	Amount   float64
}

type AddIouResponse struct {
	Users []User
}

type RestApi interface {
	GetUsers(GetUsersRequest) GetUsersResponse
	AddUser(AddUserRequest) AddUserResponse
	AddIou(AddIouRequest) AddIouResponse
}

type Api struct {
	userByName map[string]*User
	users      []User
}

func NewApi(database []User) RestApi {
	api := &Api{userByName: make(map[string]*User), users: database}
	for _, user := range api.users {
		api.userByName[user.Name] = &user
	}
	slices.SortFunc(api.users, func(a, b User) int { return cmp.Compare(a.Name, b.Name) })
	return api
}

func (a *Api) GetUsers(req GetUsersRequest) GetUsersResponse {
	if req.Users == nil {
		return GetUsersResponse{a.users}
	}
	var resp []User
	for _, name := range req.Users {
		resp = append(resp, *a.userByName[name])
	}
	return GetUsersResponse{resp}
}

func (a *Api) AddUser(req AddUserRequest) AddUserResponse {
	user := User{
		Name:    req.User,
		Owes:    make(map[string]float64),
		OwedBy:  make(map[string]float64),
		Balance: 0,
	}
	a.userByName[user.Name] = &user
	a.users = append(a.users, user)
	slices.SortFunc(a.users, func(a, b User) int { return cmp.Compare(a.Name, b.Name) })
	return AddUserResponse{user}
}

func (a *Api) AddIou(req AddIouRequest) AddIouResponse {
	lender, borrower, amount := *a.userByName[req.Lender], *a.userByName[req.Borrower], req.Amount
	if amount < 0 {
		amount = -amount
		lender, borrower = borrower, lender
	}

	lender.Owes[req.Borrower] -= amount
	lender.Balance += amount
	if lender.Owes[req.Borrower] < 0 {
		lender.OwedBy[req.Borrower] = -lender.Owes[req.Borrower]
	}
	if lender.Owes[req.Borrower] <= 0 {
		delete(lender.Owes, req.Borrower)
	}

	borrower.OwedBy[req.Lender] -= amount
	borrower.Balance -= amount
	if borrower.OwedBy[req.Lender] < 0 {
		borrower.Owes[req.Lender] = -borrower.OwedBy[req.Lender]
	}
	if borrower.OwedBy[req.Lender] <= 0 {
		delete(borrower.OwedBy, req.Lender)
	}
	users := []User{lender, borrower}
	slices.SortFunc(users, func(a, b User) int { return cmp.Compare(a.Name, b.Name) })
	return AddIouResponse{users}
}
