# Instructions

The day you waited so long finally came and you are now the proud owner of a beautiful farm in the Alps.

You still do not like waking up too early in the morning to feed your cows. 
Because you are an excellent engineer, you build a food dispenser, the `FEED-M-ALL`.

The last thing required in order to finish your project, is a piece of code that calculates the amount of fodder that each cow should get.
It is important that each cow receives the same amount, you need to avoid conflicts.
Cows are very sensitive.

Luckily, you don't need to work out all the formulas for calculating fodder amounts yourself.
You use some mysterious external library that you found on the internet.
It is supposed to result in the happiest cows.
The library exposes a type that fulfils the following interface.
You will rely on this in the code you write yourself.

```go
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}
```

As you work on your code, you will improve the error handling to make it more robust and easier to debug later on when you use it in your daily farm live.

## 1. Divide the food evenly

First of all, you focus on writing the code that is needed to calculate the amount of fodder per cow.

Implement a function `DivideFood` that accepts a `FodderCalculator` and a number of cows as an integer as arguments.
*For this task, you assume the number of cows passed in is always greater than zero.*
The function should return the amount of food per cow as a `float64` or an error if one occurred.

To make the calculation, you first need to retrieve the total amount of fodder for all the cows.
This is done by calling the `FodderAmount` method and passing the number of cows.
Additionally, you need a factor that this amount needs to be multiplied with.
You get this factor via calling the `FatteningFactor` method.
With these two values and the number of cows, you can now calculate the amount of food per cow (as a `float64`).
That is what should be returned from the `DivideFood` function.

If one of the methods you call returns an error, the execution should stop and that error should be returned (as is) from the `DivideFood` function.

```go
// For this example, we assume FodderAmount returns 50
// and FatteningFactor returns 1.5.
DivideFood(fodderCalculator, 5)
// => 15 <nil>

// Now assuming FodderAmount returns an error with message "something went wrong".
DivideFood(fodderCalculator, 5)
// => 0 "something went wrong"
```

## 2. Check the number of cows

While working on the first task above, you realized that the external library you use is not as high-quality as you thought it would be.
For example, it cannot properly handle invalid inputs.
You want to work around this limitation by adding a check for the input value in your own code.

Write a function `ValidateInputAndDivideFood` that has the same signature as `DivideFood` above.

- If the number of cows passed in is greater than 0, the function should call `DivideFood` and return the results of that call.
- If the number of cows is 0 or less, the function should return an error with message `"invalid number of cows"`.

```go
ValidateInputAndDivideFood(fodderCalculator, 5)
// => 15 <nil>

ValidateInputAndDivideFood(fodderCalculator, -2)
// => 0 "invalid number of cows"
```

## 3. Improve the error handling

Checking the number of cows before passing it along was a good move but you are not quite happy with the unspecific error message.
You decide to do better by creating a custom error type called `InvalidCowsError`.

The custom error should hold the number of cows (`int`) and a custom message (`string`) and the `Error` method should serialize the data in the following format:
```txt
{number of cows} cows are invalid: {custom message}
```

Equipped with your custom error, implement a function `ValidateNumberOfCows` that accepts the number of cows as an integer and returns an error (or nil).

- If the number of cows is less than 0, the function returns an `InvalidCowsError` with the custom message set to `"there are no negative cows"`.
- If the number of cows is 0, the function returns an `InvalidCowsError` with the custom message set to `"no cows don't need food"`.
- Otherwise, the function returns `nil` to indicate that the validation was successful.

```go
err := ValidateNumberOfCows(-5)
err.Error()
// => "-5 cows are invalid: there are no negative cows"
```

After the hard work of setting up this validation function, you notice it is already evening and you leave your desk to enjoy the sunset over the mountains.
You leave the task of actually adding the new validation function to your code for another day.