package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var total int
	for _,num := range birdsPerDay{
        total += num
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var total int
    j := (week-1)*7 
	for i:=j ; i< j+7 ;i++{
        total += birdsPerDay[i]
    }

    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:=0 ; i< len(birdsPerDay) ;i++{
        if i %2 ==0{
            birdsPerDay[i]++
        }
    }

    return birdsPerDay
}
