# Instructions

The day you waited so long finally came and you are now the proud owner of a beautiful farm in the Alps.

You still do not like to wake up too early in the morning to feed your cows and because you are an excellent engineer, you build a food dispenser, the `FEED-M-ALL`.

The last thing required in order to finish your project, is a piece of code that, given the number of cows and the amount of fodder for the day, does a division so each cow has the same quantity: you need to avoid conflicts, cows are very sensitive.

Depending on the day, some cows prefer to eat fresh grass instead of fodder, sometime no cows at all want to eat fodder.
While this is good for your pocket, you want to catch the division by zero returning an error.

Also, your silly nephew (who has just learned about negative numbers) sometimes will say that there are a negative number of cows.
You love your nephew so you want to return a helpful error when he does that.

## 1. Get the amount of fodder from the `FodderAmount` method

You will be passed a value that fulfills the `WeightFodder` interface.
`WeightFodder` includes a method called `FodderAmount` that returns the amount of fodder available and possibly an error.

```go
// twentyFodderNoError says there are 20.0 fodder
fodder, err := DivideFood(twentyFodderNoError, 10)
// fodder == 2.0
// err == nil
```

If `ErrScaleMalfunction` is returned by `FodderAmount` and the fodder amount is positive, double the fodder amount returned by `FodderAmount` before dividing it equally between the cows.
For any other error besides `ErrScaleMalfunction`, return 0 and the error.

```go
// twentyFodderWithErrScaleMalfunction says there are 20.0 fodder and a ErrScaleMalfunction
fodder, err := DivideFood(twentyFodderWithErrScaleMalfunction, 10)
// fodder == 4.0
// err == nil
```

## 2. Return an error for negative fodder

If the scale is broken and returning negative amounts of fodder, return an error saying "negative fodder" only if `FodderAmount` returned `ErrScaleMalfunction` or nil :

```go
// negativeFiveFodder says there are -5.0 fodder
fodder, err := DivideFood(negativeFiveFodder, 10)
// fodder == 0.0
// err.Error() == "negative fodder"
```

## 3. Prevent division by zero

After getting the fodder amount from `weightFodder`, prevent a division by zero when there are no cows at all by returning an error saying "division by zero":

```go
// twentyFodderNoError says there are 20.0 fodder
fodder, err := DivideFood(twentyFodderNoError, 0)
// fodder == 0.0
// err.Error() == "division by zero"
```

## 4. Handle negative cows

Define a custom error type called `SillyNephewError`.
It should be returned in case the number of cows is negative.

The error message should include the number of cows that was passed as argument.
You can see the format of the error message in the example below.

Note that if both a "negative fodder" error and a `SillyNephewError` could be returned, the "negative fodder" error takes precedence.

```go
// twentyFodderNoError says there are 20.0 fodder
fodder, err := DivideFood(twentyFodderNoError, -5)
// fodder == 0.0
// err.Error() == "silly nephew, there cannot be -5 cows"
```
