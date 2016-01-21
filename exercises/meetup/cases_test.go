package meetup

// Source: exercism/x-common
// Commit: b237b7b Merge pull request #124 from soniakeys/meetup-common-tests

import "time"

var testCases = []struct {
	year    int
	month   time.Month
	week    WeekSchedule
	weekday time.Weekday
	expDay  int
}{
	{2013, 5, Teenth, time.Monday, 13},    // monteenth of May 2013
	{2013, 8, Teenth, time.Monday, 19},    // monteenth of August 2013
	{2013, 9, Teenth, time.Monday, 16},    // monteenth of September 2013
	{2013, 3, Teenth, time.Tuesday, 19},   // tuesteenth of March 2013
	{2013, 4, Teenth, time.Tuesday, 16},   // tuesteenth of April 2013
	{2013, 8, Teenth, time.Tuesday, 13},   // tuesteenth of August 2013
	{2013, 1, Teenth, time.Wednesday, 16}, // wednesteenth of January 2013
	{2013, 2, Teenth, time.Wednesday, 13}, // wednesteenth of February 2013
	{2013, 6, Teenth, time.Wednesday, 19}, // wednesteenth of June 2013
	{2013, 5, Teenth, time.Thursday, 16},  // thursteenth of May 2013
	{2013, 6, Teenth, time.Thursday, 13},  // thursteenth of June 2013
	{2013, 9, Teenth, time.Thursday, 19},  // thursteenth of September 2013
	{2013, 4, Teenth, time.Friday, 19},    // friteenth of April 2013
	{2013, 8, Teenth, time.Friday, 16},    // friteenth of August 2013
	{2013, 9, Teenth, time.Friday, 13},    // friteenth of September 2013
	{2013, 2, Teenth, time.Saturday, 16},  // saturteenth of February 2013
	{2013, 4, Teenth, time.Saturday, 13},  // saturteenth of April 2013
	{2013, 10, Teenth, time.Saturday, 19}, // saturteenth of October 2013
	{2013, 5, Teenth, time.Sunday, 19},    // sunteenth of May 2013
	{2013, 6, Teenth, time.Sunday, 16},    // sunteenth of June 2013
	{2013, 10, Teenth, time.Sunday, 13},   // sunteenth of October 2013
	{2013, 3, First, time.Monday, 4},      // first Monday of March 2013
	{2013, 4, First, time.Monday, 1},      // first Monday of April 2013
	{2013, 5, First, time.Tuesday, 7},     // first Tuesday of May 2013
	{2013, 6, First, time.Tuesday, 4},     // first Tuesday of June 2013
	{2013, 7, First, time.Wednesday, 3},   // first Wednesday of July 2013
	{2013, 8, First, time.Wednesday, 7},   // first Wednesday of August 2013
	{2013, 9, First, time.Thursday, 5},    // first Thursday of September 2013
	{2013, 10, First, time.Thursday, 3},   // first Thursday of October 2013
	{2013, 11, First, time.Friday, 1},     // first Friday of November 2013
	{2013, 12, First, time.Friday, 6},     // first Friday of December 2013
	{2013, 1, First, time.Saturday, 5},    // first Saturday of January 2013
	{2013, 2, First, time.Saturday, 2},    // first Saturday of February 2013
	{2013, 3, First, time.Sunday, 3},      // first Sunday of March 2013
	{2013, 4, First, time.Sunday, 7},      // first Sunday of April 2013
	{2013, 3, Second, time.Monday, 11},    // second Monday of March 2013
	{2013, 4, Second, time.Monday, 8},     // second Monday of April 2013
	{2013, 5, Second, time.Tuesday, 14},   // second Tuesday of May 2013
	{2013, 6, Second, time.Tuesday, 11},   // second Tuesday of June 2013
	{2013, 7, Second, time.Wednesday, 10}, // second Wednesday of July 2013
	{2013, 8, Second, time.Wednesday, 14}, // second Wednesday of August 2013
	{2013, 9, Second, time.Thursday, 12},  // second Thursday of September 2013
	{2013, 10, Second, time.Thursday, 10}, // second Thursday of October 2013
	{2013, 11, Second, time.Friday, 8},    // second Friday of November 2013
	{2013, 12, Second, time.Friday, 13},   // second Friday of December 2013
	{2013, 1, Second, time.Saturday, 12},  // second Saturday of January 2013
	{2013, 2, Second, time.Saturday, 9},   // second Saturday of February 2013
	{2013, 3, Second, time.Sunday, 10},    // second Sunday of March 2013
	{2013, 4, Second, time.Sunday, 14},    // second Sunday of April 2013
	{2013, 3, Third, time.Monday, 18},     // third Monday of March 2013
	{2013, 4, Third, time.Monday, 15},     // third Monday of April 2013
	{2013, 5, Third, time.Tuesday, 21},    // third Tuesday of May 2013
	{2013, 6, Third, time.Tuesday, 18},    // third Tuesday of June 2013
	{2013, 7, Third, time.Wednesday, 17},  // third Wednesday of July 2013
	{2013, 8, Third, time.Wednesday, 21},  // third Wednesday of August 2013
	{2013, 9, Third, time.Thursday, 19},   // third Thursday of September 2013
	{2013, 10, Third, time.Thursday, 17},  // third Thursday of October 2013
	{2013, 11, Third, time.Friday, 15},    // third Friday of November 2013
	{2013, 12, Third, time.Friday, 20},    // third Friday of December 2013
	{2013, 1, Third, time.Saturday, 19},   // third Saturday of January 2013
	{2013, 2, Third, time.Saturday, 16},   // third Saturday of February 2013
	{2013, 3, Third, time.Sunday, 17},     // third Sunday of March 2013
	{2013, 4, Third, time.Sunday, 21},     // third Sunday of April 2013
	{2013, 3, Fourth, time.Monday, 25},    // fourth Monday of March 2013
	{2013, 4, Fourth, time.Monday, 22},    // fourth Monday of April 2013
	{2013, 5, Fourth, time.Tuesday, 28},   // fourth Tuesday of May 2013
	{2013, 6, Fourth, time.Tuesday, 25},   // fourth Tuesday of June 2013
	{2013, 7, Fourth, time.Wednesday, 24}, // fourth Wednesday of July 2013
	{2013, 8, Fourth, time.Wednesday, 28}, // fourth Wednesday of August 2013
	{2013, 9, Fourth, time.Thursday, 26},  // fourth Thursday of September 2013
	{2013, 10, Fourth, time.Thursday, 24}, // fourth Thursday of October 2013
	{2013, 11, Fourth, time.Friday, 22},   // fourth Friday of November 2013
	{2013, 12, Fourth, time.Friday, 27},   // fourth Friday of December 2013
	{2013, 1, Fourth, time.Saturday, 26},  // fourth Saturday of January 2013
	{2013, 2, Fourth, time.Saturday, 23},  // fourth Saturday of February 2013
	{2013, 3, Fourth, time.Sunday, 24},    // fourth Sunday of March 2013
	{2013, 4, Fourth, time.Sunday, 28},    // fourth Sunday of April 2013
	{2013, 3, Last, time.Monday, 25},      // last Monday of March 2013
	{2013, 4, Last, time.Monday, 29},      // last Monday of April 2013
	{2013, 5, Last, time.Tuesday, 28},     // last Tuesday of May 2013
	{2013, 6, Last, time.Tuesday, 25},     // last Tuesday of June 2013
	{2013, 7, Last, time.Wednesday, 31},   // last Wednesday of July 2013
	{2013, 8, Last, time.Wednesday, 28},   // last Wednesday of August 2013
	{2013, 9, Last, time.Thursday, 26},    // last Thursday of September 2013
	{2013, 10, Last, time.Thursday, 31},   // last Thursday of October 2013
	{2013, 11, Last, time.Friday, 29},     // last Friday of November 2013
	{2013, 12, Last, time.Friday, 27},     // last Friday of December 2013
	{2013, 1, Last, time.Saturday, 26},    // last Saturday of January 2013
	{2013, 2, Last, time.Saturday, 23},    // last Saturday of February 2013
	{2013, 3, Last, time.Sunday, 31},      // last Sunday of March 2013
	{2013, 4, Last, time.Sunday, 28},      // last Sunday of April 2013
	{2012, 2, Last, time.Wednesday, 29},   // last Wednesday of February 2012
	{2014, 12, Last, time.Wednesday, 31},  // last Wednesday of December 2014
	{2015, 2, Last, time.Sunday, 22},      // last Sunday of February 2015
	{2012, 12, First, time.Friday, 7},     // first Friday of December 2012
}
