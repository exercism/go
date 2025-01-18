package cards

import (
	"reflect"
	"testing"
)

func TestFavoriteCards(t *testing.T) {
	got := FavoriteCards()
	want := []int{2, 6, 9}
	if !slicesEqual(got, want) {
		t.Errorf("NewCards() got = %v, want %v", got, want)
	}
}

func TestGetItem(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Retrieve item from slice by index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 4,
			},
			want: 8,
		},
		{
			name: "Get first item from slice",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 0,
			},
			want: 5,
		},
		{
			name: "Get last item from slice",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 7,
			},
			want: 9,
		},
		{
			name: "Index out of bounds",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: 9,
			},
			want: -1,
		},
		{
			name: "Negative index",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				index: -1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetItem(tt.args.slice, tt.args.index)
			if got != tt.want {
				t.Errorf("GetItem(slice:%v, index:%v) got = %v, want %v", tt.args.slice, tt.args.index, got, tt.want)
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
			got := SetItem(tt.args.slice, tt.args.index, tt.args.value)
			if !slicesEqual(got, tt.want) {
				t.Errorf("SetItem(slice:%v, index:%v, value:%v) = %v, want %v",
					tt.args.slice, tt.args.index, tt.args.value, got, tt.want)
			}
			if len(tt.args.slice) == len(got) {
				for i := range got {
					got[i] = -1
				}
				if reflect.ValueOf(got).Pointer() != reflect.ValueOf(tt.args.slice).Pointer() {
					t.Errorf("SetItem(slice:%v, index:%v) does not return the modified input slice)", tt.args.slice,
						tt.args.value)
				}
			}
		})
	}
}

func TestPrependItems(t *testing.T) {
	type args struct {
		slice []int
		value []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Prepend one item",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				value: []int{1},
			},
			want: []int{1, 5, 2, 10, 6, 8, 7, 0, 9},
		},
		{
			name: "Prepend two items",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				value: []int{0, 6},
			},
			want: []int{0, 6, 5, 2, 10, 6, 8, 7, 0, 9},
		},
		{
			name: "prepend nil",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				value: nil,
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 9},
		},
		{
			name: "prepend zero items",
			args: args{
				slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
				value: []int{},
			},
			want: []int{5, 2, 10, 6, 8, 7, 0, 9},
		},
		{
			name: "Prepend slice to itself",
			args: args{
				slice: []int{5, 2, 10, 6},
				value: []int{5, 2, 10, 6},
			},
			want: []int{5, 2, 10, 6, 5, 2, 10, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrependItems(tt.args.slice, tt.args.value...)
			if !slicesEqual(got, tt.want) {
				t.Errorf("PrependItems(slice:%v, value:%v) = %v, want %v",
					tt.args.slice, tt.args.value, got, tt.want)
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
