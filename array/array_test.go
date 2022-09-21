package array

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInArray(t *testing.T) {
	type args[T comparable] struct {
		item T
		data []T
	}
	tests := []struct {
		name      string
		args      args[int]
		wantExist bool
		wantIndex int
	}{
		{
			name: "exists",
			args: args[int]{
				item: 15,
				data: []int{14, 12, 16, 25, 46, 15},
			},
			wantExist: true,
			wantIndex: 5,
		},
		{
			name: "not exists",
			args: args[int]{
				item: 15,
				data: []int{14, 12, 16, 25, 46},
			},
			wantExist: false,
			wantIndex: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExist, gotIndex := InArray(tt.args.item, tt.args.data)
			if gotExist != tt.wantExist {
				t.Errorf("InArray() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("InArray() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestArrayDiff(t *testing.T) {
	type args[T comparable] struct {
		src    []T
		others [][]T
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "not others",
			args: args[int]{
				src:    []int{12, 123, 354, 465},
				others: [][]int{},
			},
			want: []int{12, 123, 354, 465},
		},
		{
			name: "base",
			args: args[int]{
				src: []int{12, 123, 354, 465},
				others: [][]int{
					{1241, 1251, 2325, 46, 12, 123},
					{124, 465, 342},
				},
			},
			want: []int{354},
		},
		{
			name: "nil",
			args: args[int]{
				src: []int{12, 123, 354, 465},
				others: [][]int{
					{1241, 1251, 2325, 46, 12, 123},
					{124, 465, 342, 24, 354},
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.src, tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayIntersect(t *testing.T) {
	type args[T comparable] struct {
		arrs [][]T
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "empty",
			args: args[int]{
				arrs: nil,
			},
		},
		{
			name: "only one",
			args: args[int]{
				arrs: [][]int{{1, 2, 3}},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "multi",
			args: args[int]{
				arrs: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			},
			want: []int{3},
		},
		{
			name: "multi_nil",
			args: args[int]{
				arrs: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.arrs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombine(t *testing.T) {
	type args[T comparable] struct {
		base   []T
		others [][]T
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "empty",
			args: args[int]{
				base:   nil,
				others: nil,
			},
			want: nil,
		},
		{
			name: "base empty",
			args: args[int]{
				base:   nil,
				others: [][]int{{1, 2, 3}, {4, 5}},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "normal",
			args: args[int]{
				base:   []int{10, 9, 8},
				others: [][]int{{1, 2, 3}, {4, 5}, {4, 5}},
			},
			want: []int{10, 9, 8, 1, 2, 3, 4, 5, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.base, tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	old := []int{1, 2, 3}
	Reverse(old)
	assert.Equal(t, []int{3, 2, 1}, old)

	data := []string{"hello", "world", "good", "morning"}
	Reverse(data)
	assert.Equal(t, []string{"morning", "good", "world", "hello"}, data)
}

func TestReverseNew(t *testing.T) {
	old := []int{1, 2, 3}
	got := ReverseNew(old)
	assert.Equal(t, []int{3, 2, 1}, got)

	data := []string{"hello", "world", "good", "morning"}
	got2 := ReverseNew(data)
	assert.Equal(t, []string{"morning", "good", "world", "hello"}, got2)
}

func TestRemove(t *testing.T) {
	type args[T comparable] struct {
		data  []T
		index int
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "index<0",
			args: args[int]{
				data:  nil,
				index: -1,
			},
			want: nil,
		},
		{
			name: "index<0",
			args: args[int]{
				data:  []int{},
				index: -1,
			},
			want: []int{},
		},
		{
			name: "index<0",
			args: args[int]{
				data:  []int{1, 2, 3},
				index: -1,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "index>=len(data)",
			args: args[int]{
				data:  []int{1, 2, 3},
				index: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "index>=len(data)",
			args: args[int]{
				data:  []int{1, 2, 3},
				index: 4,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "len(data)=1,index=0",
			args: args[int]{
				data:  []int{1},
				index: 0,
			},
			want: []int{},
		},
		{
			name: "len(data)=2,index=0",
			args: args[int]{
				data:  []int{1, 2},
				index: 0,
			},
			want: []int{2},
		},
		{
			name: "len(data)=2,index=1",
			args: args[int]{
				data:  []int{1, 2},
				index: 1,
			},
			want: []int{1},
		},
		{
			name: "normal",
			args: args[int]{
				data:  []int{1, 2, 3},
				index: 1,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Remove(tt.args.data, tt.args.index), "Remove(%v, %v)", tt.args.data, tt.args.index)
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type args[T comparable] struct {
		data    []T
		element T
		count   int
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "empty input",
			args: args[int]{
				data:    nil,
				element: 1,
				count:   1,
			},
			want: nil,
		},
		{
			name: "empty input2",
			args: args[int]{
				data:    []int{},
				element: 1,
				count:   1,
			},
			want: []int{},
		},
		{
			name: "del 1",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4},
				element: 1,
				count:   1,
			},
			want: []int{2, 3, 4, 5, 1, 3, 4},
		},
		{
			name: "del all",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4},
				element: 1,
				count:   -1,
			},
			want: []int{2, 3, 4, 5, 3, 4},
		},
		{
			name: "del 2",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4},
				element: 1,
				count:   2,
			},
			want: []int{2, 3, 4, 5, 3, 4},
		},
		{
			name: "del 4-1",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4},
				element: 4,
				count:   1,
			},
			want: []int{1, 2, 3, 5, 1, 3, 4},
		},
		{
			name: "del 4-2",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4},
				element: 4,
				count:   2,
			},
			want: []int{1, 2, 3, 5, 1, 3},
		},
		{
			name: "del 4-2",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4},
				element: 4,
				count:   -2,
			},
			want: []int{1, 2, 3, 5, 1, 3},
		},
		{
			name: "del 0",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4},
				element: 4,
				count:   0,
			},
			want: []int{1, 2, 3, 4, 5, 1, 3, 4},
		},
		{
			name: "del 99-1",
			args: args[int]{
				data:    []int{1, 2, 3, 4, 5, 1, 3, 4, 99},
				element: 99,
				count:   1,
			},
			want: []int{1, 2, 3, 4, 5, 1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveElement(tt.args.data, tt.args.element, tt.args.count), "RemoveElement(%v, %v, %v)", tt.args.data, tt.args.element, tt.args.count)
		})
	}
}
