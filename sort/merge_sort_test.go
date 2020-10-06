package sort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{1, 9, 4, 7, 0, 5, 3},
			},
			want: []int{0, 1, 3, 4, 5, 7, 9},
		},
		{
			name: "test 2",
			args: args{
				arr: []int{1, 0},
			},
			want: []int{0, 1},
		},
		{
			name: "test 3",
			args: args{
				arr: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test sub 1",
			args: args{
				left:  []int{1, 2},
				right: []int{5},
			},
			want: []int{1, 2, 5},
		},
		{
			name: "test sub 2",
			args: args{
				left:  []int{1},
				right: []int{5},
			},
			want: []int{1, 5},
		},
		{
			name: "test sub 3",
			args: args{
				left:  []int{6, 10},
				right: []int{2, 5},
			},
			want: []int{2, 5, 6, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
