package cards

import (
	"testing"
)

func TestGetItem(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name   string
		args   args
		want   int
		wantOk bool
	}{
		{
			name: "Retrieve item from slice by index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 4,
			},
			want:   8,
			wantOk: true,
		},
		{
			name: "Get first item from slice",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 0,
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "Get last item from slice",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 7,
			},
			want:   9,
			wantOk: true,
		},
		{
			name: "Index out of bounds",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 8,
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "Negative index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: -1,
			},
			want:   0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := GetItem(tt.args.slice, tt.args.index)
			if got != tt.want {
				t.Errorf("GetItem(slice:%v, index:%v) got = %v, want %v", tt.args.slice, tt.args.index, got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetItem(slice:%v, index:%v) gotOk = %v, want %v", tt.args.slice, tt.args.index, gotOk, tt.wantOk)
			}
		})
	}
}

func TestSetItem(t *testing.T) {
	type args struct {
		slice []int
		index int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Overwrite an existing item",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 4,
				value: 1,
			},
			want: []int{5, 2, 10, 6, 1, 7, 0, 9},
		},
		{
			name: "Overwrite first item",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 0,
				value: 8,
			},
			want: []int{8, 2, 10, 6, 8, 7, 0, 9},
		},
		{
			name: "Overwrite last item",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 7,
				value: 8,
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 8},
		},
		{
			name: "Index out of bounds",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 8,
				value: 8,
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 9, 8},
		},
		{
			name: "Negative index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: -1,
				value: 8,
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetItem(tt.args.slice, tt.args.index, tt.args.value); !slicesEqual(got, tt.want) {
				t.Errorf("SetItem(slice:%v, index:%v, value:%v) = %v, want %v",
					tt.args.slice, tt.args.index, tt.args.value, got, tt.want)
			}
		})
	}
}

func TestPrefilledSlice(t *testing.T) {
	type args struct {
		value  int
		length int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Create a prefilled slice with value 3",
			args: args{
				value:  3,
				length: 7,
			},
			want: []int{3, 3, 3, 3, 3, 3, 3},
		},
		{
			name: "Create a prefilled slice with value 10",
			args: args{
				value:  10,
				length: 2,
			},
			want: []int{10, 10},
		},
		{
			name: "Length zero",
			args: args{
				value:  3,
				length: 0,
			},
			want: []int{},
		},
		{
			name: "Negative length",
			args: args{
				value:  3,
				length: -3,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrefilledSlice(tt.args.value, tt.args.length); !slicesEqual(got, tt.want) {
				t.Errorf("PrefilledSlice(value:%v, length:%v) = %v, want %v", tt.args.value, tt.args.length, got, tt.want)
			}
		})
	}
}

func TestRemoveItem(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Remove an item",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 1,
			},
			want: []int{3, 5, 6},
		},
		{
			name: "Remove the first item",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 0,
			},
			want: []int{4, 5, 6},
		},
		{
			name: "Remove the last item",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 3,
			},
			want: []int{3, 4, 5},
		},
		{
			name: "Remove out of bounds index",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: 7,
			},
			want: []int{3, 4, 5, 6},
		},
		{
			name: "Remove negative index",
			args: args{
				slice: []int{3, 4, 5, 6},
				index: -7,
			},
			want: []int{3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveItem(copySlice(tt.args.slice), tt.args.index); !slicesEqual(got, tt.want) {
				t.Errorf("RemoveItem(slice:%v, index:%v) = %v, want %v", tt.args.slice, tt.args.index, got, tt.want)
			}
		})
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	size := len(a)
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func copySlice(s []int) []int {
	var slice = make([]int, len(s))
	copy(slice, s)
	return slice
}
