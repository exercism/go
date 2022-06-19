# Instructions append

### 1. Open an account

An account can be opened. On opening, its initial balance should not be negative. In case of the negative balance, `nil` is returned.

### 2. Close an account

An account can be closed. Closed account has balance = 0. But on closing, an actual balance should be returned - `(balance,true)`. Other close requests should finish returning `(0,false)`.

### 3. Get account balance 

Account balance is accessible. If an account is closed, `(0,false)` is returned, otherwise - `(balance,true)`

### 4. Deposit/withdraw an account

Deposit/withdraw operations are supported. They are handled by a single function and the exact operation depends on an amonth value provided - it is withdraw if the value is negative, deposit otherwise.
In case withdraw is requested but not enough money present, the request canceled, returning `false` with a balance keeping the same. Otherwise - the balance is changed according to the requested amount value. 

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

https://golang.org/pkg/sync/
https://tour.golang.org/concurrency/9
https://gobyexample.com/mutexes
