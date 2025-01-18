# Hints

## 1. Determine if you will need a driver's license

- Use the [equals comparison operator][comparison-operators] to check whether your input equals a certain string.
- Use one of the two [logical operators][logical-operators] you learned about in the boolean concept to combine the two requirements.
- You do **not** need an if-statement to solve this task. You can return the boolean expression you build directly.

## 2. Choose between two potential vehicles to buy

- Use a [comparison operator][comparison-operators] to determine which option comes first in lexicographical order.
- Then set the value of a helper variable depending of the outcome of that comparison with the help of an an [if-else statement][if-statement].
- Finally construct the recommendation sentence. For that you can use the `+` operator to concatenate the two strings.

## 3. Calculate an estimation for the price of a used vehicle

- Start with determining the percentage of the original price based on the age of the vehicle. Save it in a helper variable. Use an [if-else if-else statement][if-statement] as mentioned in the instructions.
- In the two if conditions use [comparison operators][comparison-operators] to compare the age of the car to the threshold values.
- To calculate the result, apply the percentage to the original price. For example `30% of x` can be calculated by dividing `30` by `100` and multiplying with `x`.

[comparison-operators]: https://golang.org/ref/spec#Comparison_operators
[logical-operators]: https://golang.org/ref/spec#Logical_operators
[if-statement]: https://golang.org/ref/spec#If_statements
