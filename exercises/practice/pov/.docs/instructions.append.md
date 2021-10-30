# Instructions append

## Implementation Notes

POV / reparent / change root of a tree

The type name is Graph because you'll probably be implementing a general
directed graph representation, although the test program will only use
it to create a tree.  The term "arc" is used here to mean a directed edge.

The test program will create a graph with New, then use AddNode to add
leaf nodes.  After that it will use AddArc to construct the rest of the tree
from the bottom up.  That is, the `to` argument will always specify a node
that has already been added.

ArcList is a dump method to let the test program see your graph.  It must
return a list of all arcs in the graph.  Format each arc as a single string
like "from -> to". The test program can then easily sort the list and
compare it to an expected result.  You do not need to bother with sorting
the list yourself.

All this graph construction and dumping must be working before you start
on the interesting part of the exercise, so it is tested separately as
a first test.

API function ChangeRoot does the interesting part of the exercise.
OldRoot is passed (as a convenience) and you must return a graph with
newRoot as the root.  You can modify the original graph in place and
return it or create a new graph and return that.  If you return a new
graph you are free to consume or destroy the original graph.  Of course
it's nice to leave it unmodified.
