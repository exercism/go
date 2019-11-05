# Khan Go Contest

This repository is literally and liberally stolen from the [Exercism.io](http://exercism.io) track.

Fork this repo, solve the exercises, and submit a pull request for judging with the description containing your personal Tally. No more commits after the time limit, but you can fuss with your Pull Request more.

Please notify the judges which if any your fully 100% passing solutions should be considered for Special Awards:
+ **Greased Lightning Performance Award** - for exercises with benchmarks showing that you made yours go so dang fast
+ **Elegant Étude Award** - for exceptional cleverness, radical simplicity, and amazing aesthetics to your solutions.

# Contest Scoring rubric:

Each exercise in this repository is worth 0 to 5 points. A solution that will pass all the tests will get 5 points, and -1 point for every test or test case, or example test that fails. The minimum score for each exercise is 0.  

Some exercises are easier and quicker to solve than others (Hint: Start with Darts). A good developer skill to practice is recognizing a tar pit and avoiding it!

If you are reading past this before you tried it and got stuck, you are wasting precious time!

## Development setup

You will need a github account and you will need to fork StevenACoffman/go to your account.
See [GitHub Help](https://help.github.com/articles/fork-a-repo/) if you are unfamiliar with the process.
Clone your fork with the command: `git clone https://github.com/<you>/go`.
Test your clone by `cd`ing to the go directory and typing
`bin/test-without-stubs`. You should see tests pass for all exercises.

This repo only imports from the standard library and isn't expected to be imported by other packages. Please keep your solutions that way.

Your Go code should be formatted using the [gofmt](https://golang.org/cmd/gofmt/) tool. For the other file types in the repository you may want to copy the setup used in the [.editorconfig](http://editorconfig.org/) file from [the exercism.io repository](https://github.com/exercism/exercism.io/blob/master/.editorconfig).


## Exercism Go style

Let's walk through an example, non-existent, exercise, which we'll call
`fizzbuzz` to see what could be included in its implementation.

### Exercise files: Overview

For any exercise you may see a number of files present in a directory under each exercise directory:

```sh
~/go/fizzbuzz
$ tree -a
.
├── cases_test.go
├── example.go
├── fizzbuzz.go
├── fizzbuzz_test.go
├── .meta
│   └── description.md
│   └── gen.go
│   └── hints.md
│   └── metadata.yml
└── README.md
```

This list of files *can vary* across exercises. Not all exercises use
all of these files. Exercises originate their test data and README
text from the Exercism [problem-specification repository](https://github.com/exercism/problem-specifications/tree/master/exercises). This repository collects common information for all exercises across all
tracks.  However, should track-specific documentation need to be
included with the exercise, files in an exercise's `.meta/` directory
can be used to override or augment the exercise's README.

Let's briefly describe each file:

* **cases_test.go** - Contains [generated test cases](#generating-test-cases),
  using test data sourced from the problem-specifications repository.
  Only in some exercises.

* **example.go** - An [example solution](#example-solutions) for
  the exercise used to verify the test suite. Ignored
  by the `exercism fetch` command. See also [ignored files](#ignored-files).

* **fizzbuzz.go** - A *stub file*, in some early exercises
  to give users a starting point.

* **fizzbuzz_test.go** - The main test file for the exercise.

* **.meta/description.md** - Use to generate a [track specific description](https://github.com/exercism/docs/blob/master/language-tracks/exercises/anatomy/readmes.md) of the exercise in the exercise's README.

* **.meta/hints.md** - Use to add track specific information in
  addition to the generic exercise's problem-specification
  description in the README.

### Stub files

Stub files, such as *leap.go*, are a starting point for solutions. Not all exercises
need to have a stub file, only exercises early in the syllabus.

### Tests

The test file is fetched for the solver and deserves
attention for consistency and appearance.

The `leap` exercise makes use of data-driven tests. Test cases are defined as
data, then a test function iterates over the data. In this exercise, as they are
generated, the test cases are defined in the *cases_test.go* file. The test function
that iterates over this data is defined in the *leap_test.go* file.

Identifiers within the test function appear in actual-expected order as described
at [Useful Test Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures).
Here the identifier `observed` is used instead of actual. That's fine. More
common are words `got` and `want`. They are clear and short. Note [Useful Test
Failures](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures)
is part of [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
Really we like most of the advice on that page.

In Go we generally have all tests enabled and do not ask the solver to edit the
test program, to enable progressive tests for example. `t.Fatalf()`, as seen
in the *leap_test.go* file, will stop tests at the first failure encountered,
so the solver is not faced with too many failures at once.

### Testable examples

Some exercises can contain [Example tests](https://blog.golang.org/examples)
that document the exercise API. These examples are run alongside the standard
exercise tests and will verify that the exercise API is working as expected.
They are not required by all exercises and are not intended to replace the
data-driven tests. They are most useful for providing examples of how an
exercise's API is used. Have a look at the example tests in the [clock exercise](https://github.com/exercism/go/blob/master/exercises/clock/example_clock_test.go)
to see them in action.

### Errors

We like errors in Go. It's not idiomatic Go to ignore invalid data or have undefined
behavior. Sometimes our Go tests require an error return where other language
tracks don't.

### Benchmarks

In most test files there will also be benchmark tests, as can be seen at the end
of the *leap_test.go* file. In Go, benchmarking is a first-class citizen of the
testing package. We throw in benchmarks because they're interesting, and because
it is idiomatic in Go to think about performance. There is no critical use for
these though. Usually they will just bench the combined time to run over all
the test data rather than attempt precise timings on single function calls. They
are useful if they let the solver try a change and see a performance effect.
