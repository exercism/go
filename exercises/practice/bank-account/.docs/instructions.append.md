# Instructions append

## Opening an account

An account can be opened. On opening, its initial balance should not be negative. If a negative balance is given to the `Open` function, `nil` must be returned, otherwise the newly created account must be returned.

## Closing an account

An account can be closed. When closing an account, the `Close` method must return the balance the account has and a boolean `true` indicating the account was closed successfully. Closing an account does not succeed if the account is already closed. When an account is closed, its balance must be set to `0`.

## Getting the account balance

The `Balance` method allows a user to check the current balance of an account. It must return the balance of the account and a boolean indicating if the operation succeeded. Checking the balance only doesn't succeed if the account is closed.

## Deposits / Withrawals

Deposits and withdrawals are both handled by the `Deposit` method. If the argument given to the method it's a positive amount, then that amount of money must be deposited in the account. If the amount given is negative, that amount of money must be withdrawn from the account.
The return value of `Deposit` indicates if the operation succeeded. A deposit always succeeds, however a withdrawal might fail if there is not enough money in the account.

## Implementation Notes

The tests will execute some operations concurrently. You should strive
to ensure that operations on the Account leave it in a consistent state.
For example: multiple goroutines may be depositing and withdrawing money
simultaneously, two withdrawals occurring concurrently should not be able
to bring the balance into the negative.

When working locally, you can test that your code handles concurrency properly and does not introduce
data races, run tests with `-race` flag on.

```bash
$ cd exercism/project/directory/go/bank-account
$ go test -race
```

If you are new to concurrent operations in Go it will be worth looking
at the sync package, specifically Mutexes:

- [Official documentation for package sync](https://golang.org/pkg/sync/)
- [Golang Tour: sync.Mutex](https://tour.golang.org/concurrency/9)
- [Go By Example: sync.Mutex](https://gobyexample.com/mutexes)
