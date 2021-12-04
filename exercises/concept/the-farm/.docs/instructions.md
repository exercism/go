# Instructions

The day you waited so long finally came and you are now the proud owner of a beautiful farm in the Alps.

You still do not like to wake up too early in the morning to feed your cows and because you are an excellent engineer, you build a food dispenser, the `FEED-M-ALL`.

The last thing required in order to finish your project, is a piece of code that, given the number of cows and the amount of fodder for the day, does a division so each cow has the same quantity: you need to avoid conflicts, cows are very sensitive.

Depending on the day, some cows prefer to eat fresh grass instead of fodder, sometime no cows at all want to eat fodder.
While this is good for your pocket, you want to catch the division by zero returning an error.

Also, your silly nephew (who has just learned about negative numbers) sometimes will say that there are a negative number of cows.
You love your nephew so you want to return a helpful error when he does that.

## 1. Get the amount of fodder from the `weightFodder` function

You will be passed a `WeightFodder` which has a function called `FodderAmount` which returns the amount of fodder available and possibly an error.

```go
// twentyFodderNoError says there are 20.0 fodder
fodder, err := DivideFood(twentyFodderNoError, 10)
// fodder == 2.0
// err == nil
```

If a `ScaleError` is returned by `FodderAmount`, double the fodder amount returned by `FodderAmount`.
For any other error, return `0` and the error.

```go
// twentyFodderNoError says there are 20.0 fodder and a ScaleError
fodder, err := DivideFood(twentyFodderWithScaleError, 10)
// fodder == 4.0
// err == nil
```

## 2. Return an error for negative fodder

If the scale is broken and returning negative amounts of fodder, return an error saying "Negative fodder":

```go
// twentyFodderNoError says there are -5.0 fodder
fodder, err := DivideFood(negativeFiveFodder, 10)
// fodder == 0.0
// err.Error() == "Negative fodder"
```

## 3. Prevent division by zero

After getting the fodder amount from `weightFodder`, prevent a division by zero when there are no cows at all by returning an error saying "Division by zero":

```go
// twentyFodderNoError says there are 20.0 fodder
fodder, err := DivideFood(twentyFodderNoError, 0)
// fodder == 0.0
// err.Error() == "Division by zero"
```

## 4. Return a `SillyNephewError` for a negative number of cows

Help your nephew by returning a `SillyNephewError` if the number of cows is negative.

```go
// twentyFodderNoError says there are 20.0 fodder
fodder, err := DivideFood(twentyFodderNoError, -5)
// fodder == 0.0
// err == SillyNephewError{-5}
```
