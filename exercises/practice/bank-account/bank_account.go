package account

// Define the Account type here.

func Open(amount int64) *Account {
	panic("Please implement the Open function")
}

func (a *Account) Balance() (int64, bool) {
	panic("Please implement the Balance function")
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	panic("Please implement the Deposit function")
}

func (a *Account) Close() (int64, bool) {
	panic("Please implement the Close function")
}
