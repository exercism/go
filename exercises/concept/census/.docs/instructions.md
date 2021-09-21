# Instructions

It is your job to prepare the city's IT system for an upcoming census. Specifically, you are responsible for the program that will process the data from the census workers.

The program must be able to create a new resident in the system when given a resident's information. Additionally, you will create functions that ensure the required information is present in the resident's data and delete a resident's data. Lastly, you will count the residents to provide an accurate census count.

## 1. Create a new resident

When a census worker collects a resident's information, they need to register that resident by entering their name, age, and address into the system.

Implement the function `NewResident` that accepts three arguments:

- The name of the resident.
- The age of the resident.
- The address of the resident.

The function should return a pointer to a `Resident` struct that holds this information.

```go
name := "Matthew Sanabria"
age := 29
address := map[string]string{"street": "Main St.",}

NewResident(name, age, address)
// => &{Matthew Sanabria 29 map[street:Main St.]}
```

## 2. Validate required information

Residents may be reluctant to provide personal data to census workers. In those cases it's necessary to determine if the resident provided the required information to be counted in the census.

In order to be counted, a resident must provide a non-zero value for their name and an address that contains a non-zero value for the `street` key. All other information, such as the resident's age, is optional. Implement the `HasRequiredInfo` method that returns a boolean indicating if the resident has provided the required information.

```go
name := "Matthew Sanabria"
age := 0
address := make(map[string]string, 0)

resident := NewResident(name, age, address)

resident.HasRequiredInfo()
// => false
```

## 3. Delete resident information

Life moves fast and mistakes happen. A resident can move out of the city. A census worker can make mistakes when collecting data. In those cases, it's necessary to have the ability to delete a resident's data so they will not be counted.

Implement the `Delete` method that sets all of the fields the resident to their zero value.

```go
name := "Matthew Sanabria"
age := 29
address := map[string]string{"street": "Main St.",}

resident := NewResident(name, age, address)

fmt.Println(resident)
// => &{Matthew Sanabria 29 map[street:Main St.]}

resident.Delete()

fmt.Println(resident)
// => &{ 0 map[]}
```

## 4. Count the residents

Now that the system supports census data, it's time to perform the census and count the residents!

Implement the function `Count` that accepts one argument:

- A slice of pointers to `Resident` structs.

The function should return an integer indicating the number of residents that were counted in the census. A resident can only be counted if they provided the required information to the census worker.

```go
name1 := "Matthew Sanabria"
age1 := 29
address1 := map[string]string{"street": "Main St.",}

resident1 := NewResident(name1, age1, address1)

name2 := "Rob Pike"
age2 := 0
address2 := make(map[string]string, 0)

resident2 := NewResident(name2, age2, address2)

residents := []*Resident{resident1, resident2,}

Count(residents)
// => 1
}

```
