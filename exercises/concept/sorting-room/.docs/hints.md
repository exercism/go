# Hints

## 1. Implement `DescribeNumber`

- use `fmt.Sprintf` to format

## 2. Implement `DescribeBoolean`

- use an if statement (or `fmt.Sprintf`)

## 3. Implement `DescribeNumberBox`

- get the result from the `Number()` method and use `fmt.Sprintf` to format

## 4. Implement `DescribeBooleanBox`

- use an if statement on the result from the `Boolean()` method (or `fmt.Sprintf`)

## 5. Implement `DescribeAnything` which uses them all

- either use type assertions (eg. `i.(someType)`) or a type switch (`switch v := i.(type)`)
- remember to convert `int` to a `float64`
