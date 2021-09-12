# Instructions

As a chess enthusiast, you would like to write your own version of the game. Yes, there maybe plenty of implementations of chess available online already, but yours will be unique!

Each square of the chessboard is identified by a letter-number pair. The vertical columns of squares, called files, are labeled A through H. The horizontal rows of squares, called ranks, are numbered 1 to 8.

## 1. Given a Chessboard and a Rank, count how many squares are occupied

Decorate the `Chessboard` type with the `CountInRank(rank byte) int` function.
It should count occupied squares ranging over a map. Return an integer.

```go
CountInRank('A')
// => 6
```

## 2. Given a Chessboard and a File, count how many squares are occupied

Decorate the `Chessboard` type with the `CountInFile(file int) int` function.
It should count occupied squares ranging over the given file. Return an integer.

```go
CountInFile(2)
// => 5
```

## 3. Count how many squares are present in the given chessboard

Decorate the `Chessboard` type with the `CountAll() int` function.
It should count how many squares are present in the chessboard and returns
an integer. Since you don't need to check the content of the squares,
consider using range omitting both `index` and `value`.

```go
CountAll()
// => 64
```

## 4. Count how many squares are occupied in the given chessboard

Decorate the `Chessboard` type with the `CountOccupied() int` function.
It should count how many squares are occupied in the chessboard.
Return an integer.

```go
CountOccupied()
// => 15
```
