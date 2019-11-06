# Khan Go Contest

This repository is literally and liberally stolen from the [Exercism.io](http://exercism.io) track.

Fork this repo, solve the exercises, and submit a pull request __*TO THIS REPO, NOT EXERCISM/GO*__ for judging with the description containing your personal Tally. No more commits after the time limit, but you can fuss with your Pull Request more.

Please notify the judges in Slack the `#go-contest` channel which one, if any of your fully 100% passing solutions, should be considered for Special Awards:
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


## Exercise files.

Let's walk through an example, non-existent, exercise, which we'll call `fizzbuzz` to see what could be included in its implementation.

```sh
~/go/fizzbuzz$ ls
cases_test.go
fizzbuzz.go
fizzbuzz_test.go
README.md
```

The `README.md` describes the problem; your code goes in `fizzbuzz.go`, which contains a stub of the functions you need to define.
You won't need to modify the test files, although you are encouraged to look at them for usage examples.
As usual, you can run the tests with `go test`.

Many of the test files use "table-driven tests".
Such tests will include their test cases in `cases_test.go`, and the code that runs them is in `fizzbuzz_test.go`.
Some exercises can contain [Example tests](https://blog.golang.org/examples) that document the exercise API.

Most of the exercises also include benchmark tests.
You can run these with e.g. `go test -bench .`.
They are not required, but take a look at them if you want to compete for the **Greased Lightning Award**.
