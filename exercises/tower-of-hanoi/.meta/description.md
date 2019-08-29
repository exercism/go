Solve the [Tower of Hanoi](https://en.wikipedia.org/wiki/Tower_of_Hanoi) puzzle.

Tower of Hanoi consists of three rods and some disks of different sizes. Disks can slide onto any rod. The puzzle starts with the disks stacked in ascending order of size on one rod, the smallest at the top, thus making a conical shape.

The objective of the puzzle is to move the entire stack to another rod, obeying the following simple rules:

- Only one disk can be moved at a time.
- Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack or on an empty rod.
- No larger disk may be placed on top of a smaller disk.

Your implementation will accept the amount of disks, IDs for the three rods, then issue a list of moves for solving the puzzle.

Example:

Input:
- number of disks: 2
- source rod: 1
- target rod: 2
- helper rod: 3

Output: list of three moves: {1,3}, {1,2}, {3,2} 
