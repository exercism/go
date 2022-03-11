// Package census tests the census package.
package census

import (
	"reflect"
	"testing"
)

// TestNewResident tests the NewResident function.
func TestNewResident(t *testing.T) {
	testCases := map[string]struct {
		resident *Resident
	}{
		"no data collected": {
			resident: &Resident{},
		},
		"all data collected": {
			resident: &Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {

			resident := NewResident(testCase.resident.Name, testCase.resident.Age, testCase.resident.Address)

			if !reflect.DeepEqual(resident, testCase.resident) {
				t.Errorf("NewResident() = %#v, want %#v", resident, testCase.resident)
			}
		})
	}
}

// TestHasRequiredInfo tests the Resident.HasRequiredInfo method.
func TestHasRequiredInfo(t *testing.T) {
	testCases := map[string]struct {
		resident *Resident
		want     bool
	}{
		"no data collected": {
			resident: &Resident{},
			want:     false,
		},
		"all data collected": {
			resident: &Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: true,
		},
		"missing name": {
			resident: &Resident{
				Name: "",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: false,
		},
		"empty map as address": {
			resident: &Resident{
				Name:    "Rob Pike",
				Age:     0,
				Address: map[string]string{},
			},
			want: false,
		},
		"missing street": {
			resident: &Resident{
				Name: "Hossein",
				Age:  30,
				Address: map[string]string{
					"street": "",
				},
			},
			want: false,
		},
		"age is optional": {
			resident: &Resident{
				Name: "Rob Pike",
				Age:  0,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: true,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {

			if got := testCase.resident.HasRequiredInfo(); got != testCase.want {
				t.Errorf("%#v.HasRequiredInfo() = %t, want %t", testCase.resident, got, testCase.want)
			}
		})
	}
}

// TestDelete tests the Resident.Delete method.
func TestDelete(t *testing.T) {
	testCases := map[string]struct {
		resident *Resident
		want     *Resident
	}{
		"no data collected": {
			resident: &Resident{},
			want:     &Resident{},
		},
		"all data collected": {
			resident: &Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: &Resident{},
		},
		"some data collected": {
			resident: &Resident{
				Name:    "Rob Pike",
				Age:     0,
				Address: map[string]string{},
			},
			want: &Resident{},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {

			testCase.resident.Delete()

			if testCase.resident.Name != "" || testCase.resident.Age != 0 || testCase.resident.Address != nil {
				t.Errorf("resident.Delete() = %#v, want %#v", testCase.resident, testCase.want)
			}
		})
	}
}

// TestCount tests the Count function.
func TestCount(t *testing.T) {
	testCases := map[string]struct {
		name      string
		residents []*Resident
		want      int
	}{
		"no data collected": {
			residents: []*Resident{
				{},
			},
			want: 0,
		},
		"all data collected": {
			residents: []*Resident{
				{
					Name: "Matthew Sanabria",
					Age:  29,
					Address: map[string]string{
						"street": "Main St.",
					},
				},
			},
			want: 1,
		},
		"some data collected": {
			residents: []*Resident{
				{
					Name: "Matthew Sanabria",
					Age:  29,
					Address: map[string]string{
						"street": "Main St.",
					},
				},
				{
					Name:    "Rob Pike",
					Age:     0,
					Address: map[string]string{},
				},
				{
					Name:    "",
					Age:     0,
					Address: map[string]string{},
				},
			},
			want: 1,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {

			if got := Count(testCase.residents); got != testCase.want {
				t.Errorf("Count() = %d, want %d", got, testCase.want)
			}
		})
	}
}
