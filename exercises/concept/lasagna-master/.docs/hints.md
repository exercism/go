# Hints

## 1. Estimate the preparation time

- Use the [`len()` built-in function][len-builtin] to determine the number of layers (length of the layers slice).

## 2. Compute the amounts of noodles and sauce needed

- First, set up two variables to track the amount of noodles and sauce.
- Use a for loop to iterate through the layers.
- If you encounter a `'noodles'` or `'sauce'` layer in your loop, increase the amount stored in the respective variable accordingly.

## 3. Add the secret ingredient

- Revisit [slices][concept-slices] to find out how to retrieve an element from an slice and how to replace an element in a slice.
- The index of the last element in a slice `a` is `len(a) - 1`.

## 4. Scale the recipe

- First make a new slice of the same size as the input slice
- Use a [for loop][for-loop] to iterate through the input slice and generate the output slice

[len-builtin]: https://pkg.go.dev/builtin#len
[concept-slices]: /tracks/go/concepts/slices
[for-loop]: https://tour.golang.org/flowcontrol/1
