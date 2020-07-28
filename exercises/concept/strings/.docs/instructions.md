Once there was an eccentric programmer living in a strange house with barred windows.
One day he accepted a job from an online job board to build a party robot. The
robot is supposed to greet people and help them to their seats. The first edition
was very technical and showed the programmer's lack of human interaction. Some of
which also made it into the next edition.

## 1. Welcome a new guest to the party

Implement the `Welcome` function to return a welcome message using the given name:

```go
Welcome("Christiane")
// Output:
// Welcome to my party, Christiane!
```

## 2. Welcome a new guest to the party whose birthday is today

Implement the `HappyBirthday` function to return a birthday message using the given name and age of the person.
Unfortunately the programmer is a bit of a show-off, so the robot has to demonstrate its
knowledge of every guest's birthday.

```go
HappyBirthday("Frank", 58)
// Output:
// Happy birthday Frank! You are now 58 years old!
```

## 3. Give directions

Implement the `AssignTable` function to give directions. The robot provides the table number in hex (uppercase)
due to a misalignment of its creator and the rest of the world. Fortunately the precision on the distance
was limited to 1 digit.

```go
AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
// Output:
// Welcome to my party, Christiane!
// You have been assigned to table 1B. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.
```
