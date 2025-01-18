# Hints

## 1. Describe simple numbers

- use `fmt.Sprintf` to format

## 2. Describe a number box

- get the result from the `Number()` method
- use `fmt.Sprintf` to format

## 3. Implement a method extracting the number from a fancy number box

- use a type assertion to check if this is a `FancyNumber`
- get the `string` from the `Value()` method
- use `strconv.Atoi` to convert the `string` to an `int` (we can throw away the error since we want 0 if it cannot be converted to an `int`)

## 4. Describe a fancy number box

- use `ExtractFancyNumber` to get the `int` value
- convert the `int` to a `float64`
- use `fmt.Sprintf` to format

## 5. Implement `DescribeAnything` which uses them all

- either use type assertions (e.g. `i.(someType)`) or a type switch (`switch v := i.(type)`)
- remember to convert `int` to a `float64`
