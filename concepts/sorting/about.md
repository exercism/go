# About Sorting

## The Modern Way: `slices` Package

Since Go 1.21, the `slices` package is the preferred way to sort. It uses **generics**, which means it is more type-safe and generally faster than the older `sort` package methods.

To sort a slice of ordered types (like integers or strings) in ascending order, use `slices.Sort()`:

```go
import "slices"

nums := []int{4, 2, 7, 1}
slices.Sort(nums)
// nums is now [1 2 4 7]
```
## Custom Sorting with `SortFunc`
When you need to sort based on a specific logic—such as sorting a list of structs by a particular field, you use `slices.SortFunc`.
This allows you to provide a custom comparison function.

```
type Planet struct {
    Name string
    Distance int
}

planets := []Planet{
    {"Mars", 227},
    {"Venus", 108},
    {"Earth", 150},
}

slices.SortFunc(planets, func(a, b Planet) int {
    return a.Distance - b.Distance
})
```

The comparison function should return:
- A negative number if `a < b`
- Zero if `a == b`
- A positive number if `a > b`

## The Classic Way: `sort` Package
Before Go 1.21, the `sort` package was the only way to handle these tasks. You will still see this heavily in legacy codebases. It provides specific functions for basic types:
```
import "sort"

strs := []string{"c", "a", "b"}
sort.Strings(strs)
```
For custom logic in the old style, `sort.Slice` is commonly used, which relies on a "less" function:
```
sort.Slice(planets, func(i, j int) bool {
    return planets[i].Distance < planets[j].Distance
})
```
### Stability
A sort is considered stable if it preserves the original order of elements that are equal.
- In the slices package, use slices.SortStableFunc.
- In the sort package, use sort.Stable.
Standard sorting functions like slices.Sort do not guarantee stability but are often slightly faster.