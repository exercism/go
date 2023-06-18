# Instructions

As a chess enthusiast, you would like to write your own version of the game. Yes, there may be plenty of implementations of chess available online already, but yours will be unique!

Each square of the chessboard is identified by a letter-number pair:
 - The horizontal rows of squares, called ranks, are numbered 1 through 8.
 - The vertical columns of squares, called files, are labeled A through H.

```
   A B C D E F G H
 8 # _ _ _ # _ _ # 8
 7 _ _ _ _ _ _ _ _ 7
 6 _ _ _ _ # _ _ # 6
 5 _ # _ _ _ _ _ # 5
 4 _ _ _ _ _ _ # # 4
 3 # _ # _ _ _ _ # 3
 2 _ _ _ _ _ _ _ # 2
 1 # _ _ _ _ _ _ # 1
   A B C D E F G H
```

## 1. Declare Chessboard and File types

To implement your chessboard game you will need to define two non-struct types, one to represent the Chessboard and one to represent a File.

A Chessboard is a map of Files with string keys, in other words, the Chessboard contains keys from "A" to "H", each mapping to a File.

A File is a slice of boolean values. Every true boolean value represents an occupied square.
Example Files:
```
// This File has 3 occupied squares
[true, false, true, false, false, false, false, true]
```

## 2. Given a Chessboard and a File, count how many squares are occupied

Implement the `CountInFile(board Chessboard, file string) int` function.
It should count the total number of occupied squares by ranging over a map. Return an integer.
Return a count of zero (`0`) if the given file cannot be found in the map.

```go
CountInFile(board, "A")
// => 3
```

## 3. Given a Chessboard and a Rank, count how many squares are occupied

Implement the `CountInRank(board Chessboard, rank int) int` function.
It should count the total number of occupied squares by ranging over the given rank. Return an integer.
Return a count of zero (`0`) if the given rank is not a valid one (not between `1` and `8`, inclusive).

```go
CountInRank(board, 2)
// => 1
```

## 4. Count how many squares are present in the given chessboard

Implement the `CountAll(board Chessboard) int` function.
It should count how many squares are present in the chessboard and returns
an integer. Since you don't need to check the content of the squares,
consider using range omitting both `index` and `value`.

```go
CountAll(board)
// => 64
```

## 5. Count how many squares are occupied in the given chessboard

Implement the `CountOccupied(board Chessboard) int` function.
It should count how many squares are occupied in the chessboard.
Return an integer.

```go
CountOccupied(board)
// => 15
```
