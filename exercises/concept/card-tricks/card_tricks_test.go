package cards

import (
	"reflect"
	"testing"
)

func TestGetItem(t *testing.T) {
	type args struct {
		slice []uint8
		index int
	}
	tests := []struct {
		name   string
		args   args
		want   uint8
		wantOk bool
	}{
		{
			name: "Retrieve item from slice by index",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 4,
			},
			want:   8,
			wantOk: true,
		},
		{
			name: "Get first item from slice",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 0,
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "Get last item from slice",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 7,
			},
			want:   9,
			wantOk: true,
		},
		{
			name: "Index out of bounds",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 8,
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "Negative index",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: -1,
			},
			want:   0,
			wantOk: false,
		},
		{
			name: "Slice is nill",
			args: args{
				slice: nil,
				index: 0,
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
		slice []uint8
		index int
		value uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "Overwrite an existing item",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 4,
				value: 1,
			},
			want: []uint8{5, 2, 10, 6, 1, 7, 0, 9},
		},
		{
			name: "Overwrite first item",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 0,
				value: 8,
			},
			want: []uint8{8, 2, 10, 6, 8, 7, 0, 9},
		},
		{
			name: "Overwrite last item",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 7,
				value: 8,
			},
			want: []uint8{5, 2, 10, 6, 8, 7, 0, 8},
		},
		{
			name: "Index out of bounds",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: 8,
				value: 8,
			},
			want: []uint8{5, 2, 10, 6, 8, 7, 0, 9, 8},
		},
		{
			name: "Negative index",
			args: args{
				slice: []uint8{5, 2, 10, 6, 8, 7, 0, 9},
				index: -1,
				value: 8,
			},
			want: []uint8{5, 2, 10, 6, 8, 7, 0, 9, 8},
		},
		{
			name: "Slice is nill",
			args: args{
				slice: nil,
				index: 7,
				value: 8,
			},
			want: []uint8{8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetItem(tt.args.slice, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
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
			want: nil,
		},
		{
			name: "Negative length",
			args: args{
				value:  3,
				length: -3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrefilledSlice(tt.args.value, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				if tt.want == nil {
					t.Errorf("PrefilledSlice(value:%v, length:%v) = %v, want nil", tt.args.value, tt.args.length, got)
					return
				}
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
			name: "Remove an item from a nil slice",
			args: args{
				slice: nil,
				index: 1,
			},
			want: nil,
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
			if got := RemoveItem(copySlice(tt.args.slice), tt.args.index); !reflect.DeepEqual(got, tt.want) {
				if tt.want == nil {
					t.Errorf("RemoveItem(slice:%v, index:%v) = %v, want nil", tt.args.slice, tt.args.index, got)
					return
				}
				t.Errorf("RemoveItem(slice:%v, index:%v) = %v, want %v", tt.args.slice, tt.args.index, got, tt.want)
			}
		})
	}
}

func copySlice(s []int) []int {
	if s == nil {
		return s
	}
	var slice = make([]int, len(s))
	copy(slice, s)
	return slice
}
