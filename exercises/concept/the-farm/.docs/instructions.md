# Instructions

The day you waited so long finally came and you are now the proud owner of a beautiful farm in the Alps.

You still do not like to wake up too early in the morning to feed your cows and because you are an excellent engineer, you build a food dispenser, the `FEED-M-ALL`.

The last thing required in order to finish your project, is a piece of code that, given the number of cows and the amount of fodder for the day, does a division so each cow has the same quantity: you need to avoid conflicts, cows are very sensitive.

Depending on the day, some cows prefer to eat fresh grass instead of fodder, sometime no cows at all want to eat fodder.
While this is good for your pocket, you want to catch the division by zero returning an error.

Also, your silly nephew (who has just learned about negative numbers) sometimes will say that there are a negative number of cows.
You love your nephew so you want to return a helpful error when he does that.

## 1. Get the amount of fodder from the `weightFodder` function

You will be passed a `weightFodder` function which returns the amount of fodder available and possibly an error.
If an `ErrWeight` error is returned by `weightFodder`, double the fodder amount returned by `weightFodder`.
For any other error, return `0` and the error.

## 2. Return an error for negative fodder

If the scale is broken and returning negative amounts of fodder, return a helpful error.

## 3. Prevent division by zero

After checking `weightFodder`, prevent a division by zero by returning an error when there are no cows at all.

## 4. Return a `SillyNephew` error for a negative number of cows

Help your nephew by returning a `SillyNephew` error if there are a negative number of cows.
