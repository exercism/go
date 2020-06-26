## Goal

The goal of this exercise is to teach the student the rules about comments in Go.

## Learning objectives

- Exported identifiers should have a comment
- Comments should be whole sentences
- Function comments start with the function name (no colon after function name but complete sentence)
- Package comments start with package XY
- golint to check for missing and wrongly formatted comments

## Out of scope

- gofmt and formatting code

## Concepts

- `comments`

## Prerequisites

None

## Representer

## Analyzer

The following rule could be added to the analyzer.

Verify that there are no comments written about the local variables in function Log().
If there is, a reply could be:

`Variables inside of functions are private and do not require comments for the user of the package. Adding a comment here may still be useful for the maintainers of this package.`
