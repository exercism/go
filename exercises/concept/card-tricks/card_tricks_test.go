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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetItem(slices.Clone(tc.args.slice), tc.args.index)
			if got != tc.want {
				t.Errorf("GetItem(%#v, %d) got = %#v, want %#v", tc.args.slice, tc.args.index, got, tc.want)
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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SetItem(tc.args.slice, tc.args.index, tc.args.value)
			if !slices.Equal(got, tc.want) {
				t.Errorf("SetItem(%#v, %d, %d) = %#v, want %#v",
					tc.args.slice, tc.args.index, tc.args.value, got, tc.want)
			} else if len(tc.args.slice) == len(got) {
				if reflect.ValueOf(got).Pointer() != reflect.ValueOf(tc.args.slice).Pointer() {
					t.Errorf("SetItem(%#v, %d, %d) does not return the modified input slice)", tc.args.slice,
						tc.args.index, tc.args.value)
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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PrependItems(slices.Clone(tc.args.slice), tc.args.value...)
			args := make([]string, len(tc.args.value))
			for i, v := range tc.args.value {
				args[i] = strconv.Itoa(v)
			}

			if !slices.Equal(got, tc.want) {
				t.Errorf("PrependItems(%#v, %s) = %#v, want %#v",
					tc.args.slice, strings.Join(args, ", "), got, tc.want)
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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := RemoveItem(slices.Clone(tc.args.slice), tc.args.index); !slices.Equal(got, tc.want) {
				t.Errorf("RemoveItem(%#v, %d) = %#v, want %#v", tc.args.slice, tc.args.index, got, tc.want)
			}
		})
	}
}
