# Instructions

A local school near you has a very active students' association.
The student's association is managed by a president and once every 2 years,
elections are run to elect a new president.

In this year's election, it was decided that a new digital system to
count the votes was needed. The school needs your help building this new system.

## 1. Create a vote counter

One of the first things that the new voting system needs is a vote counter.
This counter is a way to keep track of the votes a particular candidate has.

Create a function `NewVoteCounter` that accepts the number of initial votes for a candidate and returns a pointer refering to an `int`, initialized with the given number of intial votes.

```go
var initialVotes int
initialVotes = 2

var counter *int
counter = NewVoteCounter(initialVotes)
*counter == initialVotes // true
```

## 2. Get number of votes from a counter

You now have a way to create new counters! But now you realize the new system will also need a way to get the number of votes from a counter.

Create a function `VoteCount` that will take a counter (`*int`) as an argument and will return the number of votes in the counter. If the counter is `nil` you should assume the counter has no votes:

```go
votes := 3
voteCounter := &votes

VoteCount(&voteCounter)
// Output: 3

var nilVoteCounter *int
VoteCount(nilVoteCounter)
// Output: 0
```

## 3. Increment the votes of a counter

It's finally the time to count the votes! Now you need a way to increment the votes in a counter.

Create a function `IncrementVoteCount` that will take a counter (`*int`) as an argument and a number of votes, and will increment the counter by that number of votes. You can assume the pointer passed will never be `nil`.

```go
var votes int
votes = 3

var voteCounter *int
voteCounter = &votes

IncrementVoteCount(voteCounter, 2)
*voteCounter == 5 // true
```

## 4. Create the election results

With all the votes now counted, it's time to prepare the result announcement to the whole school.
For this, you notice that having only counters for the votes is insuficient.
There needs to be a way to associate the number of votes with a particular candidate.

Create a function `NewElectionResult` that receives the name of a candidate and their number of votes and 
returns a new election result.

```go 
var result *ElectionResult
result = NewElectionResult("Peter", 3)

result.Name == "Peter"  // true
result.Votes == 3       // true
```

The election result struct is already created for you and it's defined as:

```go
type ElectionResult struct {
    // Name of the candidate
    Name    string
    // Number of votes the candidate had
    Votes   int
}
```

## 5. Announce the results

It's time to announce the new president to the school!
The president will be announced in the little digital message boards that the school has.
The message should show the name of the new president and the votes it had, in the following format: `<candidate_name> (<votes>)`. This is an example of such message: `"Peter (51)"`.

Create a function `DisplayResult` that will receive an `*ElectionResult` as an argument and will return a string with the message to display.


```go
var result *ElectionResult
result = &ElectionResult{
    Name: "John",
    Votes: 32
}

DisplayResult(result)
// Output: John (32)
```

## 6. Vote recounting

To make sure the final results were accurate, the votes were recounted. In the recount, it was found that the number votes for some of the candidates was off by one.

Create a function `DecrementVotesOfCandidate` that receives the final results and the name of a candidate for which you should decrement its vote count. The final results are given in the form of a `map[string]int`, where the keys are the names of the candidates and the values are its total votes.

```go
var finalResults = map[string]int{
    "Mary":  10,
    "John":  51,
}

DecrementVotesOfCandidate(finalResults, "Mary")

finalResults["Mary"]
// Output: 9
```