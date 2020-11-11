package interview

import (
	"reflect"
	"testing"
)

func TestMatchNumber(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{0, 1},
				n:   1,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				arr: []int{1, 1},
				n:   2,
			},
			want: true,
		},
		{
			name: "test 3",
			args: args{
				arr: []int{1, 2},
				n:   3,
			},
			want: true,
		},
		{
			name: "test 5",
			args: args{
				arr: []int{2, 3},
				n:   5,
			},
			want: true,
		},
		{
			name: "test 4",
			args: args{
				arr: []int{2, 3},
				n:   4,
			},
			want: false,
		},
		{
			name: "test 7",
			args: args{
				arr: []int{3, 5},
				n:   7,
			},
			want: false,
		},
		{
			name: "test 8",
			args: args{
				arr: []int{3, 5},
				n:   8,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchNumber(tt.args.arr, tt.args.n); got != tt.want {
				t.Errorf("MatchNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberToThreeJinZhi(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 5",
			args: args{
				n: 5,
			},
			want: []int{2, 1},
		},
		{
			name: "test 14",
			args: args{
				n: 14,
			},
			want: []int{2, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberToThreeJinZhi(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberToThreeJinZhi() = %v, want %v", got, tt.want)
			}
		})
	}
}
