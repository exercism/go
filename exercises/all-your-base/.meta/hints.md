## Implementation

Assumptions:
1. Zero is always represented in outputs as [0] instead of [].
2. In no other instances are leading zeroes present in any outputs.
3. Leading zeroes are accepted in inputs.
4. An empty sequence of input digits is considered zero, rather than an error.

Handle problem cases by returning an error whose Error() method
returns one of following messages:
* input base must be >= 2
* output base must be >= 2
* all digits must satisfy 0 <= d < input base
