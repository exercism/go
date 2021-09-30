# Hints

## 1. Implement `DescribeNumber`

- use `fmt.Sprintf` to format

## 3. Implement `DescribeNumberBox`

- get the result from the `Number()` method and use `fmt.Sprintf` to format

## 5. Implement `DescribeAnything` which uses them all

- either use type assertions (eg. `i.(someType)`) or a type switch (`switch v := i.(type)`)
- remember to convert `int` to a `float64`
