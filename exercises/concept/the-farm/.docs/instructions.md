# Instructions

The day you waited so long finally came over and you are now the proud owner
of a beautiful farm in the Alps.

You still don't like to wake up too early in the morning to feed your cows and
because you are an excellent engineer, you build a food dispenser, the
`FEED-M-ALL`.

The last thing required in order to finish your project, is a piece of code
that, given the number of cows and the amount of fodder for the day, does a
division so each cow has the same quantity: you need to avoid conflicts,
cows are very sensitive.

Depending on the day, some cows prefer to eat fresh grass instead of fodder,
sometime no cows at all want to eat fodder. While this is good for your
pocket, you want to catch the division by zero returning an error.

## Implement the required method

Implement the `DivideFood` function.

You will be passed a `weightFodder` function which returns the amount of fodder
available but also an error, sometimes. You should not handle that error,
return it to the `FEED-M-ALL`, instead.

Return an error when there are no cows at all, preventing a division by zero.

Return a different error if cows are a negative number.
