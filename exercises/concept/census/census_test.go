package census_test

import (
	"census"
	"testing"
)

// rawCensusData contains the information collected by census workers.
type rawCensusData struct {
	name    string
	age     int
	address map[string]string
	deleted bool
}

// tests holds the various test cases to test the census package.
var tests = []struct {
	name          string
	rawCensusData []rawCensusData
	want          int
}{
	{
		name:          "no_data",
		rawCensusData: []rawCensusData{},
		want:          0,
	},
	{
		name: "all_valid_data",
		rawCensusData: []rawCensusData{
			{
				name: "Matthew Sanabria",
				age:  29,
				address: map[string]string{
					"street": "Main St.",
				},
			},
			{
				name: "Rob Pike",
				age:  64,
				address: map[string]string{
					"street": "Gopher Ave.",
				},
			},
		},
		want: 2,
	},
	{
		name: "some_valid_data",
		rawCensusData: []rawCensusData{
			{
				name:    "Matthew Sanabria",
				age:     29,
				address: map[string]string{},
			},
			{
				name: "Rob Pike",
				age:  0,
				address: map[string]string{
					"street": "Gopher Ave.",
				},
			},
		},
		want: 1,
	},
	{
		name: "all_invalid_data",
		rawCensusData: []rawCensusData{
			{
				name:    "",
				age:     0,
				address: nil,
			},
		},
		want: 0,
	},
	{
		name: "all_deleted",
		rawCensusData: []rawCensusData{
			{
				name: "Matthew Sanabria",
				age:  29,
				address: map[string]string{
					"street": "Main St.",
				},
				deleted: true,
			},
			{
				name: "Rob Pike",
				age:  64,
				address: map[string]string{
					"street": "Gopher Ave.",
				},
				deleted: true,
			},
		},
		want: 0,
	},
	{
		name: "some_deleted",
		rawCensusData: []rawCensusData{
			{
				name: "Matthew Sanabria",
				age:  29,
				address: map[string]string{
					"street": "Main St.",
				},
			},
			{
				name: "Rob Pike",
				age:  64,
				address: map[string]string{
					"street": "Gopher Ave.",
				},
				deleted: true,
			},
		},
		want: 1,
	},
}

// TestCensus tests the census package.
func TestCensus(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			residents := make([]*census.Resident, 0, len(test.rawCensusData))

			for _, censusData := range test.rawCensusData {
				resident := census.NewResident(censusData.name, censusData.age, censusData.address)

				if censusData.deleted {
					resident.Delete()
				}

				residents = append(residents, resident)
			}

			if got := census.Count(residents); got != test.want {
				t.Errorf("Count() = %d, want %d", got, test.want)
			}
		})
	}
}
