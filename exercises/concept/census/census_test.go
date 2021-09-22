// Package census_test tests the census package.
package census_test

import (
	"census"
	"reflect"
	"testing"
)

// TestNewResident tests the census.NewResident function.
func TestNewResident(t *testing.T) {
	tests := []struct {
		name     string
		resident *census.Resident
	}{
		{
			name:     "no data collected",
			resident: &census.Resident{},
		},
		{
			name: "all data collected",
			resident: &census.Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			resident := census.NewResident(test.resident.Name, test.resident.Age, test.resident.Address)

			if !reflect.DeepEqual(resident, test.resident) {
				t.Errorf("NewResident() = %v, want %v", resident, test.resident)
			}
		})
	}
}

// TestHasRequiredInfo tests the census.HasRequiredInfo method.
func TestHasRequiredInfo(t *testing.T) {
	tests := []struct {
		name     string
		resident *census.Resident
		want     bool
	}{
		{
			name:     "no data collected",
			resident: &census.Resident{},
			want:     false,
		},
		{
			name: "all data collected",
			resident: &census.Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: true,
		},
		{
			name: "missing street",
			resident: &census.Resident{
				Name:    "Rob Pike",
				Age:     0,
				Address: map[string]string{},
			},
			want: false,
		},
		{
			name: "missing name",
			resident: &census.Resident{
				Name: "",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if got := test.resident.HasRequiredInfo(); got != test.want {
				t.Errorf("resident.HasRequiredInfo() = %t, want %t", got, test.want)
			}
		})
	}
}

// TestDelete tests the census.Delete method.
func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		resident *census.Resident
		want     *census.Resident
	}{
		{
			name:     "no data collected",
			resident: &census.Resident{},
			want:     &census.Resident{},
		},
		{
			name: "all data collected",
			resident: &census.Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: &census.Resident{},
		},
		{
			name: "some data collected",
			resident: &census.Resident{
				Name:    "Rob Pike",
				Age:     0,
				Address: map[string]string{},
			},
			want: &census.Resident{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			test.resident.Delete()

			if test.resident.Name != "" || test.resident.Age != 0 || test.resident.Address != nil {
				t.Errorf("resident.Delete() = %v, want %v", test.resident, test.want)
			}
		})
	}
}

// TestCount tests the census.Count function.
func TestCount(t *testing.T) {
	tests := []struct {
		name      string
		residents []*census.Resident
		want      int
	}{
		{
			name: "no data collected",
			residents: []*census.Resident{
				{},
			},
			want: 0,
		},
		{
			name: "all data collected",
			residents: []*census.Resident{
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
		{
			name: "some data collected",
			residents: []*census.Resident{
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if got := census.Count(test.residents); got != test.want {
				t.Errorf("census.Count() = %d, want %d", got, test.want)
			}
		})
	}
}
