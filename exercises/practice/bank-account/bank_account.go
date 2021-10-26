package account

// Define the Account type here.

func Open(amt int64) *Account {
	panic("Please implement the Open function")
}

func (account *Account) Balance() (bal int64, ok bool) {
	panic("Please implement the Balance function")
}

func (account *Account) Deposit(amt int64) (newBal int64, ok bool) {
	panic("Please implement the Deposit function")
}

func (account *Account) Close() (pay int64, ok bool) {
	panic("Please implement the Close function")
}
