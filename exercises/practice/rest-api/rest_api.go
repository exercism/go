package restapi

// Define the Rest API interface. You should not modify the code in this block.

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

// Your code goes below here. Implement the RestApi interface.

type Api struct {}

func NewApi(database []User) RestApi {
	panic("Please implement the NewApi() function")
}

func (a *Api) GetUsers(req GetUsersRequest) GetUsersResponse {
	panic("Please implement the GetUsers() function")
}

func (a *Api) AddUser(req AddUserRequest) AddUserResponse {
	panic("Please implement the AddUser() function")
}

func (a *Api) AddIou(req AddIouRequest) AddIouResponse {
	panic("Please implement the AddIou() function")
}
