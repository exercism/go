# Instructions

In this exercise you will establish some constants to be used in the operation of a bank.

You have five tasks:

## 1. Represent the fixed interest rate on a loan

Create the `FixedInterestRate` untyped numeric constant to hold the value of 5% (5/100), then implement the `GetFixedInterestRate` function to return it:

```go
GetFixedInterestRate()
// Output: 0.05
```

## 2. Represent the number of days in a year

Create the `DaysPerYear` constant with type `int` to hold the value 365, then implement the `GetDaysPerYear` function to return it:

```go
GetDaysPerYear()
// Output: 365
```

## 3. Represent the months of the year

In a block, use the Go enumerator to create twelve untyped numeric constants to hold the values 1 through 12 for the months of the year. Name them `Jan`, `Feb`, `Mar`, etc., then implement the `GetMonth` function to return the value for a given month:

```go
GetMonth(Jan)
// Output: 1
```

## 4. Represent an account number

Create the `AccountNo` untyped string constant to hold the value XF348IJ, then implement the `GetAccountNumber` function to return it:

```go
GetAccountNumber()
// Output: "XF348IJ"
```
