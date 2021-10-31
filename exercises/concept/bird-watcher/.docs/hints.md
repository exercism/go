# Hints

## 1. Determine the total number of birds that you counted so far

- Refer to the exercise introduction for an example how to use a for loop to iterate over a slice.
- Use a helper variable to store the total count and increase that variable as you go through the slice.
- Think about the correct initial value for that helper variable.

## 2. Calculate the number of visiting birds in a specific week

- This task is similar to the first one, you can copy your code as a starting point.
- Think about which indexes in the slice you would need to take into account for week number 1 and 2, respectively.
- Now find a general way to calculate the first and the last index that should be considered.
- With that you can set up the for loop to only iterate over the relevant section of the slice.

## 3. Fix a counting mistake

- Again you need to set up a for loop to iterate over the whole bird count slice.
- This time you only need to visit every second entry in the slice.
- Change the step so the counter variable is increased accordingly after each iteration.
- In the body of the for loop you can use the increment operator to change the value of an element in a slice in-place.
