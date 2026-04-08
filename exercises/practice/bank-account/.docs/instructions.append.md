# Instructions append

## Testing

The tests will execute some operations concurrently.
You should strive to ensure that operations on the Account leave it in a consistent state.
For example: multiple goroutines may be depositing and withdrawing money simultaneously, two withdrawals occurring concurrently should not be able to bring the balance into the negative.

When working locally, you can test that your code handles concurrency properly and does not introduce data races, run tests with `-race` flag on.

```bash
cd exercism/project/directory/go/bank-account
go test -race
```

## Resources

If you are new to concurrent operations in Go it will be worth looking at the sync package, specifically Mutexes:

- [Official documentation for package sync](https://golang.org/pkg/sync/)
- [Tour Of Go: sync.Mutex](https://tour.golang.org/concurrency/9)
- [Go By Example: sync.Mutex](https://gobyexample.com/mutexes)
