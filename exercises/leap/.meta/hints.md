You will see a `cases_test.go` file in this exercise. This holds the
test cases used in the `leap_test.go`. You can mostly ignore this file.

However, if you are interested... we sometimes generate the test data
from a [cross language repository](https://github.com/exercism/problem-specifications/tree/master/exercises/leap).  
They may have a [.json file](https://github.com/exercism/problem-specifications/blob/master/exercises/leap/canonical-data.json)   
that contains common test data. We have an   
[intermediary Go program](https://github.com/exercism/go/blob/master/exercises/leap/.meta/gen.go) (per exercise)  
that takes the JSON and turns in into Go structs that are fed into the  
`<exercise>_test.go` file. That is where the `cases_test.go` file comes from. Don't worry about it.
