# CLI (Command Line Interface)

You can run the tests by opening a command prompt in the exercise's directory, and then running the [`go test` command][docs-go-test].
Alternatively, most IDE's have built-in support for running tests, including [Visual Studio Code][docs-run-unit-tests-visual-studio-code] and [JetBrains GoLand][docs-run-unit-tests-goland].

Initially, only the first test will be enabled. This is to encourage you to solve the exercise one step at a time. Once you get the first test passing, remove the `Skip` property from the next test and work on getting that test passing.

Once none of the tests are skipped and they are all passing, you can submit your solution by opening a command prompt in the exercise's directory and running the using [`exercism submit` command][docs-exercism-cli].

For further reading there is an introduction to [testing in Go][testing-in-go] including on how to write tests yourself.

[docs-go-test]: https://golang.org/cmd/go/#hdr-Test_packages
[docs-exercism-cli]: https://exercism.io/cli
[docs-run-unit-tests-visual-studio-code]: https://code.visualstudio.com/docs/languages/go#_test
[docs-run-unit-tests-goland]: https://www.jetbrains.com/help/go/performing-tests.html
[testing-in-go]: ../../../reference/testing.md
