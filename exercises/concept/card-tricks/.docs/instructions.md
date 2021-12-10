# Instructions

As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.

## 1. Retrieve a card from a stack

Return the card at position `index` from the given stack.

```go
card, ok := GetItem([]int{1, 2, 4, 1}, 2)
fmt.Println(card)
// Output: 4
fmt.Println(ok)
// Output: true
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return `0` and `false`:

```go
card, ok := GetItem([]int{1, 2, 4, 1}, 10)
fmt.Println(card)
// Output: 0
fmt.Println(ok)
// Output: false
```

## 2. Exchange a card in the stack

Exchange the card at position `index` with the new card provided and return the adjusted stack.
Note that this will also change the input slice which is ok.

```go
index := 2
newCard := 6
SetItem([]int{1, 2, 4, 1}, index, newCard)
// Output: []int{1, 2, 6, 1}
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack:

```go
index := -1
newCard := 6
SetItem([]int{1, 2, 4, 1}, index, newCard)
// Output: []int{1, 2, 4, 1, 6}
```

## 3. Create a stack of cards

Create a stack of given `length` and fill it with cards of the given `value`. If the given `length` is negative or zero, return an empty stack.

```go
PrefilledSlice(8, 3)
// Output: []int{8, 8, 8}
```

## 4. Remove a card from the stack

Remove the card at position `index` from the stack and return the stack.

```go
RemoveItem([]int{3, 2, 6, 4, 8}, 2)
// Output: []int{3, 2, 4, 8}
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:

```go
RemoveItem([]int{3, 2, 6, 4, 8}, 11)
// Output: []int{3, 2, 6, 4, 8}
```
