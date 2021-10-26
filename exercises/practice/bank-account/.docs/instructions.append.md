# Instructions append

## Implementation Notes

If Open is given a negative initial deposit, it must return nil.
Deposit must handle a negative amount as a withdrawal. Withdrawals must
not succeed if they result in a negative balance.
If any Account method is called on an closed account, it must not modify
the account and must return ok = false.

The tests will execute some operations concurrently. You should strive
to ensure that operations on the Account leave it in a consistent state.
For example: multiple goroutines may be depositing and withdrawing money
simultaneously, two withdrawals occurring concurrently should not be able
to bring the balance into the negative.

If you are new to concurrent operations in Go it will be worth looking
at the sync package, specifically Mutexes:

https://golang.org/pkg/sync/
https://tour.golang.org/concurrency/9
https://gobyexample.com/mutexes
