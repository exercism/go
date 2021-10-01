# Instructions

In this exercise you'll be working with savings accounts. Each year, the balance of your savings account is updated based on its interest rate. The interest rate your bank gives you depends on the amount of money in your account (its balance):

- -3.213% for a negative balance.
- 0.5% for a positive balance less than `1000` dollars.
- 1.621% for a positive balance greater or equal than `1000` dollars and less than `5000` dollars.
- 2.475% for a positive balance greater or equal than `5000` dollars.

You have four tasks, each of which will deal your balance and its interest rate.

## 1. Calculate the interest rate

Implement the `InterestRate()` function to calculate the interest rate based on the specified balance:

```go
InterestRate(balance: 200.75)
// Output: 0.5
```

Note that the value returned is a `float32`.

## 2. Calculate the interest

Implement the `Interest()` function to calculate the interest based on the specified balance:

```go
Interest(balance: 200.75)
// Output: 9.0375
```

Note that the value returned is a `float64`.

## 3. Calculate the annual balance update

Implement the `AnnualBalanceUpdate()` function to calculate the annual balance update, taking into account the interest rate:

```go
AnnualBalanceUpdate(balance: 200.75)
// Output: 201.75375
```

Note that the value returned is a `float64`.

## 4. Calculate the years before reaching the desired balance

Implement the `YearsBeforeDesiredBalance()` function to calculate the minimum number of years required to reach the desired balance:

```go
YearsBeforeDesiredBalance(balance: 200.75, targetBalance: 214.88)
// Output: 14
```

Note that the value returned is an `int`.
