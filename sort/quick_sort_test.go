package sort

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
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
			if got := quickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
