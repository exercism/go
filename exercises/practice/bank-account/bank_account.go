package bankaccount

type Account struct {}

func Open(amt int64) *Account {
	panic("Please implement the Open() function")
}

func (a *Account) Balance() (bal int64, ok bool) {
	panic("Please implement the Balance() function")
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	panic("Please implement the Deposit() function")
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
	panic("Please implement the Withdraw() function")
}

func (a *Account) Close() (pay int64, ok bool) {
	panic("Please implement the Close() function")
}
