# Instructions

You manage an exclusive movie theater.
In order to attend a movie, people must purchase annual passes, then can attend movies for free.

You want to know which members shows up this month.
You have a log of people that came into the threater this month.
Some members never came in this month, and others came in multiple times.

You also want to track what age demographics attended movies.
In order to do that, you will need to collect what age groups, calculated as `age/10`, attended.

## Sets

A set is a perfect data type to store this type of data.

A set is a data type used to store distinct, unique values in no particular order.
Distinct and unique means an element is either in the set or not in the set; unlike with an array, it cannot be in the set multiple times.
Sets are typically unorderd; if you list the elements in a set, the order you get may be different from the order they were added -- or even different from the order you got on the prior list!

A fundamental property of sets is membership checks: the ability to check if a set contains an element.
Membership checks are core to almost all uses of a set.

Go does not provide a built in set data type.
It is common to use a `map[<type>]struct{}` or `map[<type>]bool` in place of a set; similar to a set, a `map` allows adding distinct, unique elements and supports membership checks.

## 1. A Place to Store Attendee Information

To track the attendance and age demographics, you want to use a set.

You don't want to write separate sets for names and ages; instead, you can use a generic set for both.
Define a `Set` type as a `map[T comparable]struct{}` to store attendees.
Write a `NewSet[T comparable]()` function that returns an empty `Set[T]`.

```go
names := NewSet[string]()
```

## 2. Tracking Attendees

Now you're ready to record attendee information.

Write an `Add[T]()` method adds an element to the set.

```go
names := NewSet[string]()
names.Add("Alice")
```

## 3. Checking Attendees

Now that you have a set of attendees, you can check if specific members and age demographics showed up.
Write a `Contains[T](element)` method that returns if a member or age demographic is in the `Set`.

```
names := NewSet[string]()
names.Add("Alice")
names.Contains("Alice")
// => true
names.Contains("Bob")
// => false
```

## 4. Comparing Attendees

Another month has passed and you have a new `Set` of age demographics.
Your marketing manager wants to know if the new releases this month have changed what age demographics are showing up.
Write an `Equals[T](other)` method that determines if two Sets have the same elements.

```go
agesOne := NewSet[int]()
names.Add(25)

agesTwo := NewSet[int]()
names.Add(25)

agesOne.Equals(agesTwo)
// => true

agesThree := NewSet[int]()
names.Add(35)

agesOne.Equals(agesThree)
// => false
```
