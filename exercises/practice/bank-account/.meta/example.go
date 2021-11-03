package account

import "sync"

type Account struct {
	mu      *sync.RWMutex
	open    bool
	balance int64
}

func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{
		open:    true,
		balance: amt,
		mu:      new(sync.RWMutex),
	}
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.mu.RLock()
	bal, ok = a.balance, a.open
	a.mu.RUnlock()
	return
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open {
		return a.balance, false
	}
	if amt < 0 && a.balance+amt < 0 {
		return a.balance, false
	}
	a.balance += amt
	return a.balance, true
}

func (a *Account) Close() (pay int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open {
		return 0, false
	}
	a.open = false
	if a.balance < 0 {
		return 0, true
	}
	pay = a.balance
	a.balance = 0
	return pay, true
}
