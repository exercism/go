# Instructions

The new airport in Berlin hired developers for their robots lab and you are starting your job there.
They have clunky, somewhat humanoid looking robots that they are trying to use to improve customer service.

Your first task on the job is to write a program so that the robot can greet people in their native language after they scanned their passport at the self-check-in counter.

The robot is proud of its abilities so it will always say which language it can speak first and then greet the person.
For example, if someone scans a German passport the robot would say:

```txt
I can speak German: Hallo Dietrich!
```

## 1. Create the abstract greeting functionality

You will not write the code for the different languages yourself so you need to structure your code for the robot so that other developers can easily add more languages later.

As a first step, define an interface `Greeter` with two methods.

- `LanguageName` which returns the name of the language (a `string`) that the robot is supposed to greet the visitor in.
- `Greet` which accepts a visitors name (a `string`) and returns a `string` with the greeting message in a specific language.

Next, implement a function `SayHello` that accepts the name of the visitor and anything that implements the `Greeter` interface as arguments and returns the desired greeting string.
For example, imagine a German `Greeter` implementation for which `LanguageName` returns `"German"` and `Greet` returns `"Hallo {name}!"` :

```go
SayHello("Dietrich", germanGreeter)
// => "I can speak German: Hallo Dietrich!"
```

## 2. Implement Italian

Now your job is to make the robot work for people that scan Italien passports.

For that, create a struct `Italian` and implement the two methods that are needed for the struct to fulfill the `Greeter` interface you set up in task 1.
You can greet someone in Italian with `"Ciao {name}!"` .

## 3. Implement Portuguese

Before you call it a day, you are also supposed to finish the functionality to greet people in Portuguese.

For that, create a struct `Portuguese` and implement the two methods that are needed for the struct to fulfill the `Greeter` interface here as well.
You can greet someone in Portuguese with `"Olá {name}!"`.
