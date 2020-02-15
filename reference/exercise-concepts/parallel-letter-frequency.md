# Concepts of Parallel Letter Frequency

## Concurrency

The goal of this exercise is to learn to write concurrent code to count the letters inside different texts.

## Goroutines

In order to count concurrent the apppearance of different letters in separate texts it is important to be able to create goroutines with the `go` keyword, one for each text where the letters are counted.

## Channel

In this exercise the student can learn to use a channel to make the main goroutine communicate with the goroutines that count letters.

## Functions

The function that has to be called concurrently to count needs to be created in order to get the total count at the end.

### Anonymous functions

However, since it is a simple function it is not necessary to define an external function but it can be coded inside the `go` statement.

## For range

To iterate on the texts and on the frequency map a plain `for` clause is no more needed than a for range statement as there is no specific manipulation to do on the index of the `for` loop.

## Map

The type `FreqMap` is itself a map so students have to familarize with it and with operations such as accessing via `[]` ,looping and its zero-value when the key requested does not exists
