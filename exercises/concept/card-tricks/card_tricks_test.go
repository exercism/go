package cards

import (
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestFavoriteCards(t *testing.T) {
	got := FavoriteCards()
	want := []int{2, 6, 9}
	if !slices.Equal(got, want) {
		t.Errorf("FavoriteCards() got = %#v, want %#v", got, want)
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
			got := GetItem(slices.Clone(tt.args.slice), tt.args.index)
			if got != tt.want {
				t.Errorf("GetItem(%#v, %d) got = %#v, want %#v", tt.args.slice, tt.args.index, got, tt.want)
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
		{
			name: "GetItem not valid bounds check",
			args: args{
				slice: []int{-2, -1, 0, 1, 2},
				index: 1,
				value: 5,
			},
			want: []int{-2, 5, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetItem(tt.args.slice, tt.args.index, tt.args.value)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SetItem(%#v, %d, %d) = %#v, want %#v",
					tt.args.slice, tt.args.index, tt.args.value, got, tt.want)
			} else if len(tt.args.slice) == len(got) {
				if reflect.ValueOf(got).Pointer() != reflect.ValueOf(tt.args.slice).Pointer() {
					t.Errorf("SetItem(%#v, %d, %d) does not return the modified input slice)", tt.args.slice,
						tt.args.index, tt.args.value)
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
			got := PrependItems(slices.Clone(tt.args.slice), tt.args.value...)
			args := make([]string, len(tt.args.value))
			for i, v := range tt.args.value {
				args[i] = strconv.Itoa(v)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("PrependItems(%#v, %s) = %#v, want %#v",
					tt.args.slice, strings.Join(args, ", "), got, tt.want)
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
			if got := RemoveItem(slices.Clone(tt.args.slice), tt.args.index); !slices.Equal(got, tt.want) {
				t.Errorf("RemoveItem(%#v, %d) = %#v, want %#v", tt.args.slice, tt.args.index, got, tt.want)
			}
		})
	}
}
