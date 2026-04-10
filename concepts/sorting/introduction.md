# Sorting basics

Sorting is the process of arranging data in a specific order (typically ascending or descending). In Go, sorting is performed <b>in-place</b>, meaning the original slice is modified rather than returning a new one.

## Efficiency
Go uses an algorithm called <b>pdqsort</b> (Pattern-defeating Quicksort). It is highly efficient because it adapts its strategy based on the data:
- It uses <b>Quick Sort</b> for large, random data.
- It switches to <b>Heap Sort</b> if the recursion depth gets too high (to avoid worst-case O(n<sup>2</sup>) performance).
- It uses <b>Insertion Sort</b> for very small segments of data.