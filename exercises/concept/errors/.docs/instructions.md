# Errors

In this exercise, we will be implementing the programming logic in a package `deepthought`,
which checks whether an input equals to the absolute answer to life, the universe and everything.

We have a constant `answer` which represents the answer to life, and a function `AnswerToLife` which checks the user input
and returns a boolean and an error type.

A requirement is also to have a general error `ErrCalculation` that we can at the same time use in our `AnswerToLife` function,
and export to be used in other programs of ours.
