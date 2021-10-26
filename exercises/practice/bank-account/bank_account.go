package account

// Define the Account type here.

func Open(amt int64) *Account {
	panic("Please implement the Open function")
}

func (account *Account) Balance() (int64, bool) {
	panic("Please implement the Balance function")
}

func (account *Account) Deposit(amt int64) (int64, bool) {
	panic("Please implement the Deposit function")
}

func (account *Account) Close() (int64, bool) {
	panic("Please implement the Close function")
}
