# Hints

## 1. Implement a general records filter

- Create a new `[]Record` to hold the filtered results
- Iterate over the records and call the predicate function on each one
- Add to the results slice only the records for which calling the predicate function on them returns `true`

## 2. Filter records within a period of time

- Make the function return an anonymous function that receives a `Record` and returns a `bool` indicating if the record is within the period of time that `ByDaysPeriod` receives as an argument.
- Inside your anonymous function, you have access to the `DaysPeriod` argument that `ByDaysPeriod` receives. Together with the `Record` argument of the anonymous function, make a boolean expression that checks if the day in the record is inside the `DaysPeriod` that `ByDaysPeriod` receives. Return that boolean expression from the anonymous function.

## 3. Filter records by category

- This is very similar to the previous task, but now you are dealing with a category instead of a period of time.
- Make the function return an anonymous function that receives a `Record` and returns a `bool` indicating if the record belongs to the category that `ByCategory` receives as an argument.
- Inside your anonymous function, you have access to the category argument that `ByCategory` receives. Together with the `Record` argument of the anonymous function, make a boolean expression that checks if the category of the record is the same as the category that `ByCategory` receives as an argument. Return that boolean expression from the anonymous function.

## 4. Calculate the total amount of expenses in a period

- Use the `ByDaysPeriod` function you made in step 2 to make a period filter function.
- Pass that filter function as the second argument of the `Filter` function you made in step 1. This allows you to filter the results by a period of time.
- Iterate over the filtered records and sum up all their amounts.

## 5. Calculate the total expenses for records of a category in a period

- Use the `ByCategory` function you made in step 3 to make a category filter function.
- Pass that filter function as the second argument of the `Filter` function you made in step 1. This allows you to filter the results by a particular category.
- After filtering the records by category, check if you have any records for thatcategory. If you don't, you must return an error.
- If you have records belonging to that category, compute the total expenses for the given period of time using the `TotalByPeriod` function you made in step 4.
