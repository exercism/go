// Package census tests the census package.
package census

import (
	"reflect"
	"testing"
)

// TestNewResident tests the NewResident function.
func TestNewResident(t *testing.T) {
	tests := []struct {
		name     string
		resident *Resident
	}{
		{
			name:     "no data collected",
			resident: &Resident{},
		},
		{
			name: "all data collected",
			resident: &Resident{
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

			resident := NewResident(test.resident.Name, test.resident.Age, test.resident.Address)

			if !reflect.DeepEqual(resident, test.resident) {
				t.Errorf("NewResident() = %#v, want %#v", resident, test.resident)
			}
		})
	}
}

// TestHasRequiredInfo tests the Resident.HasRequiredInfo method.
func TestHasRequiredInfo(t *testing.T) {
	tests := []struct {
		name     string
		resident *Resident
		want     bool
	}{
		{
			name:     "no data collected",
			resident: &Resident{},
			want:     false,
		},
		{
			name: "all data collected",
			resident: &Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: true,
		},
		{
			name: "missing name",
			resident: &Resident{
				Name: "",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: false,
		},
		{
			name: "empty map as address",
			resident: &Resident{
				Name:    "Rob Pike",
				Age:     0,
				Address: map[string]string{},
			},
			want: false,
		},
		{
			name: "missing street",
			resident: &Resident{
				Name: "Hossein",
				Age:  30,
				Address: map[string]string{
					"street": "",
				},
			},
			want: false,
		},
		{
			name: "age is optional",
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if got := test.resident.HasRequiredInfo(); got != test.want {
				t.Errorf("%#v.HasRequiredInfo() = %t, want %t", test.resident, got, test.want)
			}
		})
	}
}

// TestDelete tests the Resident.Delete method.
func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		resident *Resident
		want     *Resident
	}{
		{
			name:     "no data collected",
			resident: &Resident{},
			want:     &Resident{},
		},
		{
			name: "all data collected",
			resident: &Resident{
				Name: "Matthew Sanabria",
				Age:  29,
				Address: map[string]string{
					"street": "Main St.",
				},
			},
			want: &Resident{},
		},
		{
			name: "some data collected",
			resident: &Resident{
				Name:    "Rob Pike",
				Age:     0,
				Address: map[string]string{},
			},
			want: &Resident{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			test.resident.Delete()

			if test.resident.Name != "" || test.resident.Age != 0 || test.resident.Address != nil {
				t.Errorf("resident.Delete() = %#v, want %#v", test.resident, test.want)
			}
		})
	}
}

// TestCount tests the Count function.
func TestCount(t *testing.T) {
	tests := []struct {
		name      string
		residents []*Resident
		want      int
	}{
		{
			name: "no data collected",
			residents: []*Resident{
				{},
			},
			want: 0,
		},
		{
			name: "all data collected",
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
		{
			name: "some data collected",
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if got := Count(test.residents); got != test.want {
				t.Errorf("Count() = %d, want %d", got, test.want)
			}
		})
	}
}
