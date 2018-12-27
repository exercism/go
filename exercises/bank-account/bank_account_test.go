// API:
//
// Open(initialDeposit int64) *Account
// (*Account) Close() (payout int64, ok bool)
// (*Account) Balance() (balance int64, ok bool)
// (*Account) Deposit(amount int64) (newBalance int64, ok bool)
//
// If Open is given a negative initial deposit, it must return nil.
// Deposit must handle a negative amount as a withdrawal. Withdrawals must
// not succeed if they result in a negative balance.
// If any Account method is called on an closed account, it must not modify
// the account and must return ok = false.

// The tests will execute some operations concurrently. You should strive
// to ensure that operations on the Account leave it in a consistent state.
// For example: multiple goroutines may be depositing and withdrawing money
// simultaneously, two withdrawals occurring concurrently should not be able
// to bring the balance into the negative.

// If you are new to concurrent operations in Go it will be worth looking
// at the sync package, specifically Mutexes:
//
// https://golang.org/pkg/sync/
// https://tour.golang.org/concurrency/9
// https://gobyexample.com/mutexes

package account

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSeqOpenBalanceClose(t *testing.T) {
	// open account
	const amt = 10
	a := Open(amt)
	if a == nil {
		t.Fatalf("Open(%d) = nil, want non-nil *Account.", amt)
	}
	t.Logf("Account 'a' opened with initial balance of %d.", amt)

	// verify balance after open
	switch b, ok := a.Balance(); {
	case !ok:
		t.Fatal("a.Balance() returned !ok, want ok.")
	case b != amt:
		t.Fatalf("a.Balance() = %d, want %d", b, amt)
	}

	// close account
	switch p, ok := a.Close(); {
	case !ok:
		t.Fatalf("a.Close() returned !ok, want ok.")
	case p != amt:
		t.Fatalf("a.Close() returned payout = %d, want %d.", p, amt)
	}
	t.Log("Account 'a' closed.")

	// verify balance no longer accessible
	if b, ok := a.Balance(); ok {
		t.Log("Balance still available on closed account.")
		t.Fatalf("a.Balance() = %d, %t.  Want ok == false", b, ok)
	}

	// verify closing balance is 0
	if b, _ := a.Balance(); b != 0 {
		t.Log("Balance after close is non-zero.")
		t.Fatalf("After a.Close() balance is %d and not 0", b)
	}
}

func TestSeqOpenDepositClose(t *testing.T) {
	// open account
	const openAmt = 10
	a := Open(openAmt)
	if a == nil {
		t.Fatalf("Open(%d) = nil, want non-nil *Account.", openAmt)
	}
	t.Logf("Account 'a' opened with initial balance of %d.", openAmt)

	// deposit
	const depAmt = 20
	const newAmt = openAmt + depAmt
	switch b, ok := a.Deposit(depAmt); {
	case !ok:
		t.Fatalf("a.Deposit(%d) returned !ok, want ok.", depAmt)
	case b != openAmt+depAmt:
		t.Fatalf("a.Deposit(%d) = %d, want new balance = %d", depAmt, b, newAmt)
	}
	t.Logf("Deposit of %d accepted to account 'a'", depAmt)

	// close account
	switch p, ok := a.Close(); {
	case !ok:
		t.Fatalf("a.Close() returned !ok, want ok.")
	case p != newAmt:
		t.Fatalf("a.Close() returned payout = %d, want %d.", p, newAmt)
	}
	t.Log("Account 'a' closed.")

	// verify deposits no longer accepted
	if b, ok := a.Deposit(1); ok {
		t.Log("Deposit accepted on closed account.")
		t.Fatalf("a.Deposit(1) = %d, %t.  Want ok == false", b, ok)
	}
}

func TestMoreSeqCases(t *testing.T) {
	// open account 'a' as before
	const openAmt = 10
	a := Open(openAmt)
	if a == nil {
		t.Fatalf("Open(%d) = nil, want non-nil *Account.", openAmt)
	}
	t.Logf("Account 'a' opened with initial balance of %d.", openAmt)

	// open account 'z' with zero balance
	z := Open(0)
	if z == nil {
		t.Fatal("Open(0) = nil, want non-nil *Account.")
	}
	t.Log("Account 'z' opened with initial balance of 0.")

	// attempt to open account with negative opening balance
	if Open(-10) != nil {
		t.Fatal("Open(-10) seemed to work, " +
			"want nil result for negative opening balance.")
	}

	// verify both balances a and z still there
	switch b, ok := a.Balance(); {
	case !ok:
		t.Fatal("a.Balance() returned !ok, want ok.")
	case b != openAmt:
		t.Fatalf("a.Balance() = %d, want %d", b, openAmt)
	}
	switch b, ok := z.Balance(); {
	case !ok:
		t.Fatal("z.Balance() returned !ok, want ok.")
	case b != 0:
		t.Fatalf("z.Balance() = %d, want 0", b)
	}

	// withdrawals
	const wAmt = 3
	const newAmt = openAmt - wAmt
	switch b, ok := a.Deposit(-wAmt); {
	case !ok:
		t.Fatalf("a.Deposit(%d) returned !ok, want ok.", -wAmt)
	case b != newAmt:
		t.Fatalf("a.Deposit(%d) = %d, want new balance = %d", -wAmt, b, newAmt)
	}
	t.Logf("Withdrawal of %d accepted from account 'a'", wAmt)
	if _, ok := z.Deposit(-1); ok {
		t.Fatal("z.Deposit(-1) returned ok, want !ok.")
	}

	// verify both balances
	switch b, ok := a.Balance(); {
	case !ok:
		t.Fatal("a.Balance() returned !ok, want ok.")
	case b != newAmt:
		t.Fatalf("a.Balance() = %d, want %d", b, newAmt)
	}
	switch b, ok := z.Balance(); {
	case !ok:
		t.Fatal("z.Balance() returned !ok, want ok.")
	case b != 0:
		t.Fatalf("z.Balance() = %d, want 0", b)
	}

	// close just z
	switch p, ok := z.Close(); {
	case !ok:
		t.Fatalf("z.Close() returned !ok, want ok.")
	case p != 0:
		t.Fatalf("z.Close() returned payout = %d, want 0.", p)
	}
	t.Log("Account 'z' closed.")

	// verify 'a' balance one more time
	switch b, ok := a.Balance(); {
	case !ok:
		t.Fatal("a.Balance() returned !ok, want ok.")
	case b != newAmt:
		t.Fatalf("a.Balance() = %d, want %d", b, newAmt)
	}
}

func TestConcClose(t *testing.T) {
	if runtime.NumCPU() < 2 {
		t.Skip("Multiple CPU cores required for concurrency tests.")
	}
	if runtime.GOMAXPROCS(0) < 2 {
		runtime.GOMAXPROCS(2)
	}

	// test competing close attempts
	for rep := 0; rep < 1000; rep++ {
		const openAmt = 10
		a := Open(openAmt)
		if a == nil {
			t.Fatalf("Open(%d) = nil, want non-nil *Account.", openAmt)
		}
		var start sync.WaitGroup
		start.Add(1)
		const closeAttempts = 10
		res := make(chan string)
		for i := 0; i < closeAttempts; i++ {
			go func() { // on your mark,
				start.Wait() // get set...
				switch p, ok := a.Close(); {
				case !ok:
					if p != 0 {
						t.Errorf("a.Close() = %d, %t.  "+
							"Want payout = 0 for unsuccessful close", p, ok)
						res <- "fail"
					} else {
						res <- "already closed"
					}
				case p != openAmt:
					t.Errorf("a.Close() = %d, %t.  "+
						"Want payout = %d for successful close", p, ok, openAmt)
					res <- "fail"
				default:
					res <- "close" // exactly one goroutine should reach here
				}
			}()
		}
		start.Done() // ...go
		var closes, fails int
		for i := 0; i < closeAttempts; i++ {
			switch <-res {
			case "close":
				closes++
			case "fail":
				fails++
			}
		}
		switch {
		case fails > 0:
			t.FailNow() // error already logged by other goroutine
		case closes == 0:
			t.Fatal("Concurrent a.Close() attempts all failed.  " +
				"Want one to succeed.")
		case closes > 1:
			t.Fatalf("%d concurrent a.Close() attempts succeeded, "+
				"each paying out %d!.  Want just one to succeed.",
				closes, openAmt)
		}
	}
}

func TestConcDeposit(t *testing.T) {
	if runtime.NumCPU() < 2 {
		t.Skip("Multiple CPU cores required for concurrency tests.")
	}
	if runtime.GOMAXPROCS(0) < 2 {
		runtime.GOMAXPROCS(2)
	}
	a := Open(0)
	if a == nil {
		t.Fatal("Open(0) = nil, want non-nil *Account.")
	}
	const amt = 10
	const c = 1000
	var negBal int32
	var start, g sync.WaitGroup
	start.Add(1)
	g.Add(3 * c)
	for i := 0; i < c; i++ {
		go func() { // deposit
			start.Wait()
			a.Deposit(amt) // ignore return values
			g.Done()
		}()
		go func() { // withdraw
			start.Wait()
			for {
				if _, ok := a.Deposit(-amt); ok {
					break
				}
				time.Sleep(time.Microsecond) // retry
			}
			g.Done()
		}()
		go func() { // watch that balance stays >= 0
			start.Wait()
			if p, _ := a.Balance(); p < 0 {
				atomic.StoreInt32(&negBal, 1)
			}
			g.Done()
		}()
	}
	start.Done()
	g.Wait()
	if negBal == 1 {
		t.Fatal("Balance went negative with concurrent deposits and " +
			"withdrawals.  Want balance always >= 0.")
	}
	if p, ok := a.Balance(); !ok || p != 0 {
		t.Fatalf("After equal concurrent deposits and withdrawals, "+
			"a.Balance = %d, %t.  Want 0, true", p, ok)
	}
}

// The benchmark operations are here to encourage you to try different
// implementations to see which ones perform better. These are worth
// exploring after the tests pass.
//
// There is a basic benchmark and a parallelized version of the same
// benchmark. You run the benchmark using:
// go test --bench=.
//
// The output will look something like this:
// goos: linux
// goarch: amd64
// BenchmarkAccountOperations-8             10000000        130 ns/op
// BenchmarkAccountOperationsParallel-8     3000000         488 ns/op
// PASS
//
// You will notice that parallelism does not increase speed in this case, in
// fact it makes things slower! This is because none of the operations in our
// Account benefit from parallel processing. We are specifically protecting
// the account balance internals from being accessed by multiple processes
// simultaneously. Your protections will make the parallel processing slower
// because there is some overhead in managing the processes and protections.
//
// The interesting thing to try here is to experiment with the protections
// and see how their implementation changes the results of the parallel
// benchmark.
func BenchmarkAccountOperations(b *testing.B) {
	a := Open(0)
	defer a.Close()
	for n := 0; n < b.N; n++ {
		a.Deposit(10)
		a.Deposit(-10)
	}
}

func BenchmarkAccountOperationsParallel(b *testing.B) {
	a := Open(0)
	defer a.Close()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			a.Deposit(10)
			a.Deposit(-10)
		}
	})
}
