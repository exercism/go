# Append

## Zipper Behavior

The Zipper provides a functional way to nagivate a `Node` tree.

All Zipper calls should return a _new_ Zipper object.
If the Zipper does not change values (i.e. no `Set*` method was called), the Zipper should use the same underlying tree.
When a `Set*` method is called, the new Zipper should have a new tree, leaving the original tree unchanged.

Zipper calls that change the focus to an invalid `Node` should return an error.
