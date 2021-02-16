# Instructions append

With this exercise, you are introduced to the use of
[multiple results](https://tour.golang.org/basics/6) feature and here there
is an example with [errors](https://tour.golang.org/methods/19):

    func Foo(a, b string) (string, error) {
        return "", errors.New("An error occurred!")
    }

You may be wondering about the `cases_test.go` file. We explain it in the
[leap exercise][leap-exercise-readme].

[leap-exercise-readme]: https://github.com/exercism/go/blob/master/exercises/leap/README.md
