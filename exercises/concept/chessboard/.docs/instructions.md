# Instructions

As a chess enthusiast, you would like to write your own version of the game. Yes, there may be plenty of implementations of chess available online already, but yours will be unique!

Each square of the chessboard is identified by a letter-number pair:
 - The vertical columns of squares, called files, are numbered 1 through 8.
 - The horizontal rows of squares, called ranks, are labelled A through H.

```
  1 2 3 4 5 6 7 8
A # _ # _ _ _ _ # A
B _ _ _ _ # _ _ _ B
C _ _ # _ _ _ _ _ C
D _ _ _ _ _ _ _ _ D
E _ _ _ _ _ # _ # E
F _ _ _ _ _ _ _ _ F
G _ _ _ # _ _ _ _ G
H # # # # # # _ # H
  1 2 3 4 5 6 7 8
```

## 1. Given a Chessboard and a Rank, count how many squares are occupied

Implement the `CountInRank(board Chessboard, rank string) int` function.
It should count the total number of occupied squares by ranging over a map. Return an integer.
Return a count of zero (`0`) if the given rank cannot be found in the map.

```go
CountInRank(board, "A")
// => 6
```

## 2. Given a Chessboard and a File, count how many squares are occupied

Implement the `CountInFile(board Chessboard, file int) int` function.
It should count the total number of occupied squares by ranging over the given file. Return an integer.
Return a count of zero (`0`) if the given file is not a valid one (not between `1` and `8`, inclusive).

```go
CountInFile(board, 2)
// => 5
```

## 3. Count how many squares are present in the given chessboard

Implement the `CountAll(board Chessboard) int` function.
It should count how many squares are present in the chessboard and returns
an integer. Since you don't need to check the content of the squares,
consider using range omitting both `index` and `value`.

```go
CountAll(board)
// => 64
```

## 4. Count how many squares are occupied in the given chessboard

Implement the `CountOccupied(board Chessboard) int` function.
It should count how many squares are occupied in the chessboard.
Return an integer.

```go
CountOccupied(board)
// => 15
```
