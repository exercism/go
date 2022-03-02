## Tree logic POV exercise 

### Major Design Issues 

1. The exercise defines de-constructors Value() and Children(). These methods
   are mainly used for testing FromPov. Alternatively, we could: 
   1. ask learners to implement a String() method, or 
   2. ask students to implement other "inspection" methods (such as listing all
      arcs) 
    and then use those methods for testing. 

2. In the exercise, the order of children in trees does not matter. Currently
   this is dealt with by function treeSliceEqual. Alternatively, we could: 
   1. ask learners to sort the Children slice by value
   2. change the result type of the Children method from a slice to a map. 
      
3. There is a predefined String() method in tree.go. Its purpose is to make test
  outputs more readable. It makes use of methods Value() and Children() as
  defined by learners. Alternatively, we could: 
   1. ask learners to implement such a method.  
  
## Minor Design Issues 

1. Tree nodes (= sub-trees) are identified via their unique "value" field.
   Because of this, is "id" maybe a better field name than "value"? 
   
2. There are currently no methods that checks uniqueness of values in a tree.
   There is also no requirement to check that the first argument of New is
   nonempty. Would it be useful to add such methods to the exercise?
